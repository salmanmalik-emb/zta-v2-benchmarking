#!/bin/bash

if wg show utun151 > /dev/null 2>&1; 
then
    if [ ! -f /etc/wireguard/utun151.conf ]
    then
        touch /etc/wireguard/utun151.conf
    fi
    # Bring the interface down
    wg-quick down utun151
    rm /etc/wireguard/utun151.conf
fi
if [ ! -f /etc/ztna/wg-client.json ]
then
        echo "creating json file"
        mkdir /etc/ztna
        touch wg-client.json
fi
echo '{"wg_private_key":"","pre_shared_key":"","wg_iface_name":"utun151","wg_port":51821,"wg_mtu":1420,"wg_addr":"","tunnel_config":{"stuns":[{"uri":"stun:turnserver1.extremecloudztna.com:3479"}],"turns":[{"hostConfig":{"uri":"turn:turnserver1.extremecloudztna.com:3479"},"user":"salman1","password":"123"}],"signal_service":{"uri":"turnserver1.extremecloudztna.com:443","protocol":"https"}},"peers":[{"public_key":"8/0GtB5f2ahSBG56EFa3rGSGXwvib+IDrp4UQqNZrQg=","allowed_ips":["100.64.0.1/32"]}]}' > /etc/ztna/wg-client.json

if [ ! -f /etc/systemd/system/ice-turn-client.service ]
then
        echo "downloading ice turn client"
        wget -O ice-turn-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/ice-turn/build/client
        sudo cp ice-turn-client /usr/local/bin/ice-turn-client
        sudo chmod +x /usr/local/bin/ice-turn-client
        
        wget -O ice-turn-client.service https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/ice-turn/build/wg-client.service
        sudo cp ice-turn-client.service /etc/systemd/system
        
        sudo systemctl daemon-reload
        
fi

systemctl start ice-turn-client.service
sleep 10s

# delay until client is connected to connector
while  ! ping -c 1 100.64.0.1; do
  echo "pinging 100.64.0.1"
  systemctl restart ice-turn-client.service
  sleep 10s
done

#SERVICE_ENDPOINT=18.189.156.25:8050
SERVICE_ENDPOINT=100.64.0.1:7000
CONN=10
DURATION=60
DELAY_AFTER_CLOSE=60
PPS=90
PACKET_SIZE=960


if [ ! -f ./packet-bouncer-client ]
then
        echo "downloading packet bouncer client"
        wget -O packet-bouncer-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/packet-bouncer/client/client
        chmod +x packet-bouncer-client
fi

echo "waiting 1 min before running load test to avoid packet count mixup"
sleep 60s
echo "running load client script"
./packet-bouncer-client service start --endpoint $SERVICE_ENDPOINT --clients $CONN --duration $DURATION --pps $PPS --packet $PACKET_SIZE --delay-after-stopping-sender $DELAY_AFTER_CLOSE 2>&1 | tail -n1 > output.json
cat output.json
curl -X POST http://13.229.240.224:8888/results -H "Content-Type: application/json" -d @./output.json

systemctl stop ice-turn-client.service
sudo systemctl disable ice-turn-client
sudo rm /etc/systemd/system/ice-turn-client.service
sudo systemctl daemon-reload
rm /etc/ztna/wg-client.json
rm ./packet-bouncer-client

if wg show utun151 > /dev/null 2>&1; then
    if [ ! -f /etc/wireguard/utun151.conf ]
    then
        touch /etc/wireguard/utun151.conf
    fi
    # Bring the interface down
    wg-quick down utun151
    rm /etc/wireguard/utun151.conf
fi