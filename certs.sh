#!/bin/bash

DIR_NAME=`dirname "$BASH_SOURCE"`
CERTS_ROOT=${DIR_NAME}/certs

SERVICE_KEY=${CERTS_ROOT}/service.key
SERVICE_PEM=${CERTS_ROOT}/service.pem
SERVICE_CSR=${CERTS_ROOT}/tls.csr

mkdir -p $CERTS_ROOT

openssl req -nodes -new -newkey rsa:2048 -keyout ${SERVICE_KEY} -out ${SERVICE_CSR} -subj "/CN=localhost"
openssl x509 -req -days 365 -in ${SERVICE_CSR} -signkey ${SERVICE_KEY} -out ${SERVICE_PEM}

CLIENT_KEY=${CERTS_ROOT}/client.key
CLIENT_PEM=${CERTS_ROOT}/client.pem

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ${CLIENT_KEY} -out ${CLIENT_PEM} -subj "/CN=localhost"