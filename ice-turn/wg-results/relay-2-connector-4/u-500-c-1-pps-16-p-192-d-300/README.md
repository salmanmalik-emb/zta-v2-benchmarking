
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `Users` | 500 |
| `Per user threads` | 1 |
| `PPS` | 16 |
| `Packet Size` | 192 |

## Results

|  Item | Value            |
| :------------------------- |:------------------------- |
| `Response Time < 400ms` | 99.99% |
| `400 ms > Response Time < 1s` | 0.002% | 
| `Packet Loss` | 0.005% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.42 |

|  Item | Relay            | Connector |
| :------------------------- |:------------------------- |:------------------------- |
| `Throughput` | 175 Mbps | 173.33 Mbps |
| `CPU Usage` | 40.7% | 58.7% |
| CPU | ![](coturn/relay-cpu.png) |  ![](coturn/connector-cpu.png) |
| Network In | ![](coturn/relay-network-in.png) |  ![](coturn/connector-network-in.png) |
| Network Out | ![](coturn/relay-network-out.png) |  ![](coturn/connector-network-out.png) |