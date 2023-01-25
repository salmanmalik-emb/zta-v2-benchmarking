
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 3 |
| `Per user threads` | 1 |
| `PPS` | 90 |
| `Packet Size` | 780 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 100% |
| `400 ms > Response Time < 1s` | 0% | 
| `Packet Loss` | 0.2% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.49 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 3.94 Mbps | 3.81 Mbps |
| `CPU Usage` | 0.18% | 0.4% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |