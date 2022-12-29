
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 1500 |
| `duration` | 300s |
| `packets per second` | 90 |
| `packet size` | 960 bytes |

## Results


|  Item | Coturn            |  Pion/Turn |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 2239.14 Mbps | 1870.5 Mbps |
| `CPU Usage` | 77.9%% | 53.1%% |
| `Response Time < 400ms` | 67.96% | 5.721% |
| `400 ms > Response Time < 1s` | 32.03% | 6.99% |
| `Packet Loss` | 0.004% | 83.78%% |
| `Bad Packet Loss` | 0% | 36.42% |
| `Score` | 9.26 | -54.54 |
| CPU | ![](coturn/cpu.png) |  ![](pion/cpu.png) |
| Network In | ![](coturn/network-in.png) |  ![](pion/network-in.png) |
| Network Out | ![](coturn/network-out.png) |  ![](pion/network-out.png) |