#!/bin/bash
TURN_SERVER=46.137.194.93:3478
TURN_USER=salman
TURN_PASS=123
CONN=20
DURATATION=300
DELAY_AFTER_CLOSE=120
PPS=16
PACKET_SIZE=192


if [ ! -f ./turnhammer ]
then
        wget -O turnhammer https://github.com/vi/turnhammer/releases/download/v0.2.0/turnhammer_linux64
        chmod +x turnhammer
fi
./turnhammer $TURN_SERVER $TURN_USER $TURN_PASS -j $CONN -d $DURATATION --json  --pps $PPS -s $PACKET_SIZE --force --delay-after-stopping-sender $DELAY_AFTER_CLOSE 2> /dev/null > output.json
curl -X POST http://turnserver1.extremecloudztna.com:8888/results -H "Content-Type: application/json" -d @./output.json