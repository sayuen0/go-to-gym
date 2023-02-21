#!/bin/bash

SERVER_CN=localhost

openssl genrsa -out server.key 4096
openssl req -new -key server.key -subj "/CN=$SERVER_CN" -out server.csr
openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650 -sha256
