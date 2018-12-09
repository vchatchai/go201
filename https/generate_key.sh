#!/bin/bash 
openssl req -newkey rsa:2048 -nodes -keyout domain.key -out domain.csr -subj "/C=TH/ST=Bangkok/L=Bangsue/O=test/CN=test.com"
openssl req -key domain.key -new -x509 -days 365 -out domain.crt -subj "/C=TH/ST=Bangkok/L=Bangsue/O=test/CN=test.com"