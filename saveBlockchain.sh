#!/bin/bash
curl -X POST -H "Content-Typ: application/json" -d '{"Data": "save.json"}' localhost:8080/api/persistBlockchain

