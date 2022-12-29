
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 1500 |
| `duration` | 300s |
| `packets per second` | 16 |
| `packet size` | 192 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 91.46 Mbps |  90.8 Mbps|
| `CPU Usage` | 70.2% |  80% |
| `Response Time < 400ms` | 41.3% |  39.94% |
| `400 ms > Response Time < 1s` | 58.69% |  58.53% |
| `Packet Loss` | 0% |  1.516% |
| `Bad Packet Loss` | 0% |  0% |
| `Score` | 9.061 |  8.557 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |