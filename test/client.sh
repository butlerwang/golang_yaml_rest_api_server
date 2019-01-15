#!/bin/bash

echo -e "\n\n=> test 1: create valid 1\n\n"
curl  --data-binary @valid1.yaml -H "Content-type: text/x-yaml" http://localhost:31415/api/v1/app -v -H "X-token: 123456"
echo -e "\n\n=> test 2: create valid 2\n\n"
curl  --data-binary @valid1.yaml -H "Content-type: text/x-yaml" http://localhost:31415/api/v1/app -v -H "X-token: 123456"
