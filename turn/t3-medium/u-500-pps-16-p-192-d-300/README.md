
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 500 |
| `duration` | 300s |
| `packets per second` | 16 |
| `packet size` | 192 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 30.4 Mbps |  30.4 Mbps|
| `CPU Usage` | 28.8% |  40.4% |
| `Response Time < 400ms` | 71.22% |  71.16% |
| `400 ms > Response Time < 1s` | 28.77% |  28.81% |
| `Packet Loss` | 0% |  0.02% |
| `Bad Packet Loss` | 0% |  0% |
| `Score` | 9.31 |  9.296 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |