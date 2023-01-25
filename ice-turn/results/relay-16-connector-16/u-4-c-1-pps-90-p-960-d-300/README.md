
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 250 |
| `Per user threads` | 1 |
| `PPS` | 90 |
| `Packet Size` | 960 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 100% |
| `400 ms > Response Time < 1s` | 0% | 
| `Packet Loss` | 0.23% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.42 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 6.37 Mbps | 6.13 Mbps |
| `CPU Usage` | 0.2% | 0.5% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |