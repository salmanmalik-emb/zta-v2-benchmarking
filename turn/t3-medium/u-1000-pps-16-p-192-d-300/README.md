
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 1000 |
| `duration` | 300s |
| `packets per second` | 16 |
| `packet size` | 192 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 60.8 Mbps |  60.8 Mbps|
| `CPU Usage` | 51.1% |  65.8% |
| `Response Time < 400ms` | 68.07% |  67.89% |
| `400 ms > Response Time < 1s` | 31.92% |  31.86% |
| `Packet Loss` | 0% |  0.23% |
| `Bad Packet Loss` | 0% |  0% |
| `Score` | 9.26 |  9.18 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |