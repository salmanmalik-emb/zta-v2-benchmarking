
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 800 |
| `duration` | 300s |
| `packets per second` | 90 |
| `packet size` | 960 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 1198.76 Mbps | 1160.5 Mbps |
| `CPU Usage` | 82% | 85.4% |
| `Response Time < 400ms` | 68.03% | 7.35% |
| `400 ms > Response Time < 1s` | 31.96% | 35.09% |
| `Packet Loss` | 0.004% | 55.701% |
| `Bad Packet Loss` | 0% | 1.52% |
| `Score` | 9.25 | -10.615 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |