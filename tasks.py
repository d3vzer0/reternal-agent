from config import config
from celery import Celery
from pathlib import Path
from logger import stdlog

import os
import shlex
import json
import hashlib
import subprocess
import base64

build_app = Celery('reternal-agent', broker=config['celery']['broker'], backend=config['celery']['results'])
build_app.conf.task_routes = config['celery']['routes']
build_app.conf.broker_transport_options = {'fanout_prefix': True}
build_app.conf.broker_transport_options = {'fanout_patterns': True}
stdlog.info("Started agent building daemon")

supported_platforms = {
    'windows': ['386', 'amd64'],
    'linux': ['386', 'amd64'],
    'darwin': ['386', 'amd64']
}

class InvalidPlatform(Exception):
    pass

@build_app.task(name="agent.build")
def build_agent(platform, arch, base_url, build_id):
    if platform not in supported_platforms or \
        arch not in supported_platforms[platform]: raise InvalidPlatform('Unsupported platform')

    stdlog.info('Building: {0}-{1} {2}'.format(platform, arch, base_url)) 
    os.environ['GOOS'] = shlex.quote(platform)
    os.environ['GOARCH'] = shlex.quote(arch)
    custom_build_url = '-X main.base_url={0}'.format(shlex.quote(base_url))
    build_path = '{0}/{1}'.format(config['golang']['dst'],
        hashlib.md5('{0}-{1}'.format(platform, arch).encode()).hexdigest())
    src_code = '{0}/corebeacon.go'.format(config['golang']['src'])

    try:
        build_output = subprocess.check_output(['go', 'build', '-ldflags',
            custom_build_url,'-o', build_path, src_code])
        with open(build_path, 'rb') as build_file:
            build_data = build_file.read()
            build_encoded = base64.b64encode(build_data).decode('utf-8')
            stdlog.info('Finished building: {0}-{1} {2}'.format(platform, arch, base_url)) 
            result = {"result":"success", "data":"Succesfully built agent", "file":build_encoded}

    except InvalidPlatform:
        stdlog.info('Platform not supported: {0}-{1} {2}'.format(platform, arch, base_url))
        result = {"result":"failed", "data":"Unable to run subprocess"}

    except Exception as err:
        stdlog.info('Failed to build: {0}-{1} {2}'.format(platform, arch, base_url)) 
        result = {"result":"failed", "data":"Unable to run subprocess"}

    return result