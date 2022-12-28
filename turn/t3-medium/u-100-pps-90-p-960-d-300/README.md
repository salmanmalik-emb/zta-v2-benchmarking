
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 100 |
| `duration` | 300 |
| `packets per second` | 90 |
| `packet size` | 960 |

## Results


| Coturn            |  Pion/Turn
:-------------------------:|:-------------------------:|:-------------------------:
| `Throughput` | 100 |  100 |
| `CPU Usage` | 300 |  100 |
| `Response Time < 400ms` | 90 |  100 |
| `400 ms > Response Time < 1s` | 960 |  100 |
| `Packet Loss` | 960 |  100 |
| `Bad Packet Loss` | 960 |  100 |
| `Score` | 960 |  100 |
| CPU Screenshot | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In (Bytes) | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out (Bytes) | ![](coturn/network-out.png) |  ![](pion/network-out.png) |