
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
| `Throughput` | 148.82 Mbps | 148.82 Mbps |
| `CPU Usage` | 31.3% | 45.7% |
| `Response Time < 400ms` | 85.515% | 76.254% |
| `400 ms > Response Time < 1s` | 14.484% | 23.338%  |
| `Packet Loss` | 0% | 0% |
| `Bad Packet Loss` | 0% | 0% |
| `Score` | 9.4 | |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
