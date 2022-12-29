
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 100 |
| `duration` | 300s |
| `packets per second` | 90 |
| `packet size` | 960 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 748.2 Mbps | 734.55 Mbps |
| `CPU Usage` | 99.2% |  97.2% |
| `Response Time < 400ms` | 51.23% | 14.88% |
| `400 ms > Response Time < 1s` | 43.13% | 41.74% |
| `Packet Loss` | 5.63% | 42.73% |
| `Bad Packet Loss` | 0% | |
| `Score` | 7.336 | -5.2225 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |