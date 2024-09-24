#!/bin/bash

for i in {1..9}
do
  ./bin/consumer --conf config/service.conf &
done
