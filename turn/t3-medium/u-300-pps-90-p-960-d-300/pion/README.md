
### Parameters

| Parameter | Value                |
| :-------- |:------------------------- |
| `concurent connections` | 100 |
| `duration` | 300s |
| `packets per second` | 90 |
| `packet size` | 960 bytes |

## Results

### Pion/Turn
| Parameter | Value                |
| :-------- |:------------------------- |
| `Throughput` | 100 Mbps |
| `CPU Usage` | 300% |
| `Response Time < 400ms` | 90% |
| `400 ms > Response Time < 1s` | 960% |
| `Packet Loss` | 0% |
| `Bad Packet Loss` | 0% |
| `Score` | 9 |

![CPU](cpu.png)
![Network In (Bytes)](network-in.png)
![Network Out (Bytes)](network-out.png)
