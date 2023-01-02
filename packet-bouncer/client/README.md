# Packet-Bouncer-Client

Generates and measures user defined traffic and packet loss.

## Usage

```
Usage: turnhammer <server> <username> <password> [-j <parallel-connections>] [-s <pkt-size>] [--pps <pps>] [-d <duration>] [--delay-after-stopping-sender <delay-after-stopping-sender>] [--delay-between-allocations <delay-between-allocations>] [--calc] [-f] [--video] [--audio] [-J] [-C]


Options:
  --clients
                    number of simultaneous connections
  --packet    packet size
  --pps             packets per second
  -d, --duration    experiment duration, seconds
  --delay-after-stopping-sender
                    seconds to wait and receive after stopping sender
```
