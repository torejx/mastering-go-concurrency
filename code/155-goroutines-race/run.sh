#!/bin/bash
for i in {1..10}
do
   echo "Execution $i"
   go run main.go
   echo ""
done