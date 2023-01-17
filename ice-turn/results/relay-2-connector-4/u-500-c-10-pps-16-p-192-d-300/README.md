
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 800 |
| `duration` | 300s |
| `packets per second` | 54 |
| `packet size` | 960 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 709.97 Mbps | 712.704 Mbps |
| `CPU Usage` | 94.2% | 98.1% |
| `Response Time < 400ms` | 51.1744% | 7.663% |
| `400 ms > Response Time < 1s` | 48.64% | 46.924% |
| `Packet Loss` | 0.181% | 45.412% |
| `Bad Packet Loss` | 0% | 0.08% |
| `Score` | 9.0775 | -5.845 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |