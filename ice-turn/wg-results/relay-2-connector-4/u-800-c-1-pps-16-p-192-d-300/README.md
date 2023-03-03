
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 800 |
| `Per user threads` | 1 |
| `PPS` | 16 |
| `Packet Size` | 192 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 99.4% |
| `400 ms > Response Time < 1s` | 0.03% | 
| `Packet Loss` | 0.56% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.31 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 288 Mbps | 272 Mbps |
| `CPU Usage` | 60.5% | 78% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |
| Network Out | ![](coturn/relay-network-out.png) |  ![](coturn/connector-network-out.png) |