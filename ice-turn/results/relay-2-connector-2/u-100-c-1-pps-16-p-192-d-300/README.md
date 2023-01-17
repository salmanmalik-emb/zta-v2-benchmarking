
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 100 |
| `Per user threads` | 1 |
| `PPS` | 90 |
| `Packet Size` | 960 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 100% |
| `400 ms > Response Time < 1s` | 0% | 
| `Packet Loss` | 0.14% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.42 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 49.73 Mbps | 46.9 Mbps |
| `CPU Usage` | 9.1% | 29% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |
| Network Out | ![](coturn/relay-network-out.png) |  ![](coturn/connector-network-out.png) |