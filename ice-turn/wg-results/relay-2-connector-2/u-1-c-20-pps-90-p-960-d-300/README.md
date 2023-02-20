
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 1 |
| `Per user threads` | 20 |
| `PPS` | 90 |
| `Packet Size` | 960 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 100% |
| `400 ms > Response Time < 1s` | 0% | 
| `Packet Loss` | 0.03% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.48 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 31.7 Mbps | 30.8 Mbps |
| `CPU Usage` | 4.1% | 10% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |
| Network Out | ![](coturn/relay-network-out.png) |  ![](coturn/connector-network-out.png) |