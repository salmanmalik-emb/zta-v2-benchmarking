
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 300 |
| `duration` | 300s |
| `packets per second` | 90 |
| `packet size` | 960 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 446.47 Mbps | 445.09 Mbps |
| `CPU Usage` | 78.5% | 88.3% |
| `Response Time < 400ms` | 77.9% | 31.40% |
| `400 ms > Response Time < 1s` | 21.84% | 54.37% |
| `Packet Loss` | 0.2% | 14.22% |
| `Bad Packet Loss` | 0% | 0% |
| `Score` | 9.28 | 4.4 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
