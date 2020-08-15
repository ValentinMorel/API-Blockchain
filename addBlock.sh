#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '{"Data": "Second Block"}' localhost:8080/api/AddBlock
