#  RE:TERNAL
-------------

<img src="https://i.postimg.cc/7hwhx4Dp/reternal.png" alt="Drawing" style="width: 300px;"/>

![Go](https://img.shields.io/badge/Go-1.11.4-green.svg)
![Python](https://img.shields.io/badge/Python-3.6-green.svg)
![version](https://img.shields.io/badge/Version-Alpha_0.0.1-orange.svg)

---------------------

RE:TERNAL is a centralised purple team simulation platform. Reternal uses agents installed on a simulation network to execute various known
red-teaming techniques in order to test blue-teaming capabilities. The simulations are mapped to the MITRE ATT&CK framework. This repo contains the agent code (written in Golang)
and Celery task to build and download the compiled agent via the payload API. Since the project is still in development, please do not run this in production or on an exposed interface yet.

#### Reternal components
- **API:** https://github.com/d3vzer0/reternal-backend.git
- **UI:** https://github.com/d3vzer0/reternal-ui.git
- **Agent:** https://github.com/d3vzer0/reternal-agent.git
- **C2:** https://github.com/d3vzer0/reternal-c2.git
- **Quickstart:** https://github.com/d3vzer0/reternal-quickstart.git
- **Mitre/Command Mapping:** https://github.com/d3vzer0/reternal-mitre.git

<img src="https://i.postimg.cc/15nGCgws/Untitled-Diagram-3.png" alt="Drawing" style="width: 600px;"/>


#### Component installation
Reternal components are primarily aimed to be run as docker containers since the component configuration depends on environment variables set by docker-compose or the dockerfile. A docker-compose with all the default options can be found on the reternal-quickstart repository. If you don't want to run the service within containers, adjust the config.py files with your own custom values.

#### Pre-Requirements
  - Go 1.11
  - Redis (for task queuing)
  - Celery (for task management)
  - Python v3.6
  - +- 3 PIP components (found in requirements.txt)

#### Feature Requests & Bugs
We use the Github to manage Feature requests and Bug reports.

#### Developers and Contact

Joey Dreijer < joeydreijer@gmail.com >  
Yaleesa Borgman < yaleesa@gmail.com >

#### Whats up with the name?

This project has been re-developed so many times, it will probably never really finish. Hence RE (Redo) and Ternal (Eternal).

#### Special Thanks
  - MITRE ATT&CK - Framework used for mapping simulations: https://attack.mitre.org/wiki/Main_Page
  - Uber Metta -  Using Metta's templates for MITRE techniques with small (optional) adjustments to the purple_action format: https://github.com/uber-common/metta


