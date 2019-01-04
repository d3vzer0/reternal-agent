FROM golang:1.11.4-stretch

RUN apt-get update && apt-get -y upgrade && apt-get -y install python3-pip

COPY requirements.txt .
RUN pip3 install --no-cache-dir -r requirements.txt

ARG GO_SRC=/reternal-agent/agent/src
ENV GO_SRC="${GO_SRC}"

ARG GO_DST=/reternal-agent/agent/dist
ENV GO_DST="${GO_DST}"

ARG CELERY_BACKEND=redis://redis-service:6379
ENV CELERY_BACKEND="${CELERY_BACKEND}"

ARG CELERY_BROKER=redis://redis-service:6379
ENV CELERY_BROKER="${CELERY_BROKER}"

COPY . /reternal-agent
WORKDIR /reternal-agent

CMD celery -A tasks worker -Q agent
