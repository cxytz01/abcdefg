#!/bin/bash

while true
do
  curl -X 'POST' 'http://127.0.0.1:7005/api/v1/messages'
  sleep 2
done
