# Packet-Bouncer-Client

Generates and measures user defined traffic and packet loss.

## Usage

```
Usage: packet-bouncer-client [--endpoint <server>] [--packet <pkt-size>] [--pps <pps>] [--duration <duration>] [--delay-after-stopping-sender <delay-after-stopping-sender>]


Options:
  --clients   number of simultaneous connections
  --packet    packet size
  --pps             packets per second
  --duration    experiment duration, seconds
  --delay-after-stopping-sender seconds to wait and receive after stopping sender
```


```
#!/bin/bash
SERVICE_ENDPOINT=100.64.0.1:7000
CONN=20
DURATION=300
DELAY_AFTER_CLOSE=120
PPS=16
PACKET_SIZE=192


if [ ! -f ./packet-bouncer-client ]
then
        wget -O packet-bouncer-client https://github.com/salmanmalik-emb/zta-v2-benchmarking/raw/main/packet-bouncer/client/client
        chmod +x packet-bouncer-client
fi
./packet-bouncer-client --endpoint $SERVICE_ENDPOINT --clients $CONN --duration $DURATION --pps $PPS --packet $PACKET_SIZE --delay-after-stopping-sender $DELAY_AFTER_CLOSE 2> /dev/null > output.json
curl -X POST http://turnserver1.extremecloudztna.com:8888/results -H "Content-Type: application/json" -d @./output.json
```
