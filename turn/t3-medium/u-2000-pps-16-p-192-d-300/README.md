
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 2000 |
| `duration` | 300s |
| `packets per second` | 16 |
| `packet size` | 192 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 122.53 Mbps |  121.46 Mbps|
| `CPU Usage` | 81.1% |  89.5% |
| `Response Time < 400ms` | 76.127% |  46.04% |
| `400 ms > Response Time < 1s` | 23.85% |  46.06% |
| `Packet Loss` | 0% |  4.89% |
| `Bad Packet Loss` | 0% |  0% |
| `Score` | 9.31 |  7.52 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |