#!/bin/bash

curl -X  POST -H "Content-Type: application/json" -d '{"Data": "genesis"}' localhost:8080/api/CreateBlockchain
