#!/bin/bash

if [ ! -f ./ice-turn-client ]
then
        wget -O ice-turn-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/ice-turn/build/client
        chmod +x ice-turn-client
fi

if [ ! -f /etc/ztna/client.json ]
then
        echo '{"wg_private_key":"","pre_shared_key":"","wg_iface_name":"utun151","wg_port":51821,"wg_mtu":1420,"wg_addr":"","tunnel_config":{"stuns":[{"uri":"stun:turnserver1.extremecloudztna.com:3478"}],"turns":[{"hostConfig":{"uri":"turn:turnserver1.extremecloudztna.como:3478"},"user":"salman","password":"123"}],"signal_service":{"uri":"turnserver1.extremecloudztna.com:443","protocol":"https"}},"peers":[{"public_key":"s1zvmG7exxlZTh9NM+z+2oG4ehelVBGBB9P38IVLlQg=","allowed_ips":["100.64.0.1/32"]}]}' > /etc/ztna/client.json
fi

# run client in backgroud
start /b ./ice-turn-client service start 

# delay until client is connected to connector
while ! ping -c 1 100.64.0.1; do
  sleep 5
done

SERVICE_ENDPOINT=100.64.0.1:7000
CONN=1
DURATION=300
DELAY_AFTER_CLOSE=60
PPS=90
PACKET_SIZE=960


if [ ! -f ./packet-bouncer-client ]
then
        wget -O packet-bouncer-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/packet-bouncer/client/client
        chmod +x packet-bouncer-client
fi
./packet-bouncer-client service start --endpoint $SERVICE_ENDPOINT --clients $CONN --duration $DURATION --pps $PPS --packet $PACKET_SIZE 2>&1 | tail -n1 > output.json
curl -X POST http://turnserver1.extremecloudztna.com:8888/results -H "Content-Type: application/json" -d @./output.json