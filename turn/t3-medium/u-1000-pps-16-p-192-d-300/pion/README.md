
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 1000 |
| `duration` | 300s |
| `packets per second` | 16 |
| `packet size` | 192 bytes |

## Results

### Pion/Turn
| Parameter | Value                |
| :-------- |:------------------------- |
| `Throughput` | 60.8 Mbps |
| `CPU Usage` | 65.8% |
| `Response Time < 400ms` | 67.89% |
| `400 ms > Response Time < 1s` | 31.86% |
| `Packet Loss` | 0.23% |
| `Bad Packet Loss` | 0% |
| `Score` | 9.18 |

![CPU](cpu.png)
![Network In (Bytes)](network-in.png)
![Network Out (Bytes)](network-out.png)
