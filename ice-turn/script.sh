#!/bin/bash

if [ ! -f /etc/ztna/client.json ]
then
        echo "creating json file"
        mkdir /etc/ztna
        touch client.json
        echo '{"wg_private_key":"","pre_shared_key":"","wg_iface_name":"utun151","wg_port":51821,"wg_mtu":1420,"wg_addr":"","tunnel_config":{"stuns":[{"uri":"stun:turnserver1.extremecloudztna.com:3478"}],"turns":[{"hostConfig":{"uri":"turn:turnserver1.extremecloudztna.como:3478"},"user":"salman","password":"123"}],"signal_service":{"uri":"turnserver1.extremecloudztna.com:443","protocol":"https"}},"peers":[{"public_key":"s1zvmG7exxlZTh9NM+z+2oG4ehelVBGBB9P38IVLlQg=","allowed_ips":["100.64.0.1/32"]}]}' > /etc/ztna/client.json
fi

if [ ! -f ./ice-turn-client ]
then
        echo "downloading ice turn client"
        wget -O ice-turn-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/ice-turn/build/client
        sudo cp ice-turn-client /usr/local/bin/ice-turn-client

        wget -O ice-turn-client.service https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/ice-turn/build/client.service
        sudo cp ice-turn-client.service /etc/systemd/system
        
        sudo systemctl daemon-reload
        systemctl start ice-turn-client.service
fi

# delay until client is connected to connector
while ! ping -c 1 100.64.0.1; do
  echo "pinging 100.64.0.1"
  sleep 5
done

SERVICE_ENDPOINT=100.64.0.1:7000
CONN=8
DURATION=10
DELAY_AFTER_CLOSE=10
PPS=25
PACKET_SIZE=1420


if [ ! -f ./packet-bouncer-client ]
then
        echo "downloading packet bouncer client"
        wget -O packet-bouncer-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/packet-bouncer/client/client
        chmod +x packet-bouncer-client
fi

echo "running load client script"
./packet-bouncer-client service start --endpoint $SERVICE_ENDPOINT --clients $CONN --duration $DURATION --pps $PPS --packet $PACKET_SIZE 2>&1 | tail -n1 > output.json
curl -X POST http://turnserver1.extremecloudztna.com:8888/results -H "Content-Type: application/json" -d @./output.json


systemctl stop ice-turn-client.service