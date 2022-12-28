
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 100 |
| `duration` | 300 |
| `packets per second` | 90 |
| `packet size` | 960 |

## Results

### Coturn
| Parameter | Value                |
| :-------- |:------------------------- |
| `Throughput` | 100 |
| `CPU Usage` | 300 |
| `Response Time < 400ms` | 90 |
| `400 ms > Response Time < 1s` | 960 |
| `Packet Loss` | 960 |
| `Bad Packet Loss` | 960 |
| `Score` | 960 |



![CPU](coturn/cpu.png)
![Network In (Bytes)](coturn/network-in.png)
![Network Out (Bytes)](coturn/network-in.png)

### Pion/Turn
| Parameter | Value                |
| :-------- |:------------------------- |
| `Throughput` | 100 |
| `CPU Usage` | 300 |
| `Response Time < 400ms` | 90 |
| `400 ms > Response Time < 1s` | 960 |
| `Packet Loss` | 960 |
| `Bad Packet Loss` | 960 |
| `Score` | 960 |

![CPU](pion/cpu.png)
![Network In (Bytes)](pion/network-in.png)
![Network Out (Bytes)](pion/network-in.png)
