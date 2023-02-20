
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 5 |
| `Per user threads` | 1 |
| `PPS` | 90 |
| `Packet Size` | 1000 |

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
| `Throughput` | 8.22 Mbps | 7.98 Mbps |
| `CPU Usage` | 0.25% | 0.66% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |