#!/bin/sh

while true
sleep 3
do
	/data/crshow/checkfiles rsql >>/data/logs/crshow/log-$(date +\%Y-\%m-\%d).log
	sleep 2
	/data/crshow/checkfiles rfile >>/data/logs/crshow/log-$(date +\%Y-\%m-\%d).log
	sleep 2
done