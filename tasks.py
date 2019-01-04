from config import config
from celery import Celery
from pathlib import Path
from logger import stdlog

import os
import shlex
import json
import subprocess
import base64

build_app = Celery('reternal-agent', broker=config['celery']['broker'], backend=config['celery']['results'])
build_app.conf.task_routes = config['celery']['routes']
build_app.conf.broker_transport_options = {'fanout_prefix': True}
build_app.conf.broker_transport_options = {'fanout_patterns': True}
stdlog.info("Started agent building daemon")

@build_app.task(name="agent.build")
def build_agent(platform, arch, base_url, build_id):
    stdlog.info("Building: %s-%s %s") %(platform, arch, base_url)
    set_env_platform = os.environ["GOOS"] = platform
    set_env_arch = os.environ["GOARCH"] = arch
    escaped_url = shlex.quote(base_url)

    custom_build_url = '-X main.base_url=%s' %(escaped_url)
    build_id += ".exe" if platform == "windows" else ""
    build_path = "%s/%s" %(config["golang"]["dst"], build_id)
    src_code = "%s/corebeacon.go" %(config["golang"]["src"])

    try:
        build_output = subprocess.check_output(["go", "build", "-ldflags", custom_build_url, "-o", build_path, src_code])
        build_data = open(build_path, "rb").read() 
        build_encoded = base64.b64encode(build_data).decode('utf-8')
        stdlog.info("Finished building: %s-%s %s") %(platform, arch, base_url)
        result = {"result":"success", "data":"Succesfully built agent", "file":build_encoded}


    except Exception as err:
        stdlog.info("Failed to build: %s-%s %s") %(platform, arch, base_url)
        result = {"result":"failed", "data":"Unable to run subprocess"}

    return result