
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 6 |
| `Per user threads` | 1 |
| `PPS` | 90 |
| `Packet Size` | 960 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 100% |
| `400 ms > Response Time < 1s` | 0% | 
| `Packet Loss` | 0.3% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.48 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 94.46 Mbps | 9.09 Mbps |
| `CPU Usage` | 0.28% | 0.78% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |