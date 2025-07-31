# ZTA 2.0 Benchmarking

This repository contains comprehensive load testing and performance benchmarking results for Zero Trust Architecture (ZTA) 2.0 components. The benchmarks evaluate various networking and connectivity solutions across different deployment scenarios and configurations.

## Overview

Zero Trust Architecture 2.0 represents the next generation of security-first networking, where every connection must be verified and validated. This repository provides detailed performance analysis and benchmarking data for key ZTA components including TURN servers, relay systems, VPN technologies, and network access solutions.

## Components & Benchmarks

### 🔄 [TURN Server Performance](turn/README.md)
Load testing results for different open source TURN servers across various AWS EC2 instance types.
- **Servers Tested**: Coturn, Pion/Turn
- **Instance Types**: t3.medium, c5.large, c5.xlarge, c5.2xlarge
- **Metrics**: Throughput (Mbps), Connection capacity, CPU utilization, Performance scores
- **Key Findings**: Performance comparison across different hardware configurations

### 🚀 [Agentless Service Access](agentless/README.md)
Performance benchmarking for agentless network access solutions with various load patterns.
- **Test Scenarios**: Different user counts (700-1000), file sizes (5.4-30.8 MB)
- **Rate Limiting**: Tests with 400kbps and 1mbps bandwidth limits
- **Architecture**: Relay and Connector performance analysis
- **Metrics**: Throughput, CPU utilization, Packets per second (PPS)

### 🔐 [ICE-TURN with VPN Technologies](ice-turn/)
Benchmarking of Interactive Connectivity Establishment (ICE) with TURN servers using different VPN backends.
- **VPN Technologies**: 
  - IPsec results ([ipsec-results/](ice-turn/ipsec-results/))
  - WireGuard results ([wg-results/](ice-turn/wg-results/))
- **Test Configurations**: Multiple relay and connector combinations
- **Load Patterns**: Various user counts, connection types, and packet parameters

### ⚖️ [Agent-based vs Agentless Comparison](agentbased_agentless/)
Direct performance comparison between agent-based and agentless architectures.
- **Test Scenarios**: With and without relay configurations
- **Load Patterns**: 200-700 agent-based users, 200-500 agentless users
- **Rate Limiting**: 188kbps bandwidth constraints
- **Comparative Analysis**: Performance trade-offs between deployment models

### 📦 [Packet Bouncer Tools](packet-bouncer/)
Custom Go-based tools for network packet testing and bouncing.
- **Components**: Client and server implementations
- **Purpose**: Network connectivity and performance testing utilities
- **Technology**: Go modules with configurable packet handling

### 📊 [Results Capture System](capture-results/)
Automated system for collecting and aggregating benchmark results.
- **Features**: HTTP-based results collection, Docker containerization
- **Technology**: Go-based implementation with web interface
- **Purpose**: Centralized benchmarking data collection and analysis

### 🏢 [PA-5250 vs ICE-TURN](pa-vs-ice-turn/)
Performance comparison between Palo Alto PA-5250 firewall and ICE-TURN relay solutions.
- **Hardware**: PA-5250 enterprise firewall comparison
- **Software**: Coturn relay with connector architecture
- **Metrics**: CPU utilization vs throughput analysis
- **Throughput Range**: 3.8 Mbps to 34 Mbps testing scenarios

## Key Performance Insights

### TURN Server Benchmarks
- **Best Performance**: c5.2xlarge instances achieving 2.2+ Gbps throughput
- **CPU Efficiency**: Coturn generally shows better CPU utilization than Pion/Turn
- **Scalability**: Performance scales linearly with CPU cores and memory

### Agentless vs Traditional Solutions
- **High Throughput**: Agentless solutions achieve 1.3+ Gbps with proper configuration
- **CPU Efficiency**: Connector components show significantly lower CPU usage than relay
- **Rate Limiting Impact**: Bandwidth limits effectively control resource utilization

### VPN Technology Comparison
- **IPsec**: Traditional enterprise VPN performance baselines
- **WireGuard**: Modern, efficient VPN protocol benchmarks
- **Relay Architecture**: Performance impact of different relay configurations

## Repository Structure

```
├── turn/                    # TURN server benchmarks
├── agentless/              # Agentless access performance tests
├── ice-turn/               # ICE-TURN with VPN technologies
├── agentbased_agentless/   # Architecture comparison tests
├── packet-bouncer/         # Network testing tools
├── capture-results/        # Results collection system
└── pa-vs-ice-turn/         # Hardware vs software comparison
```

## Usage

Each component directory contains detailed README files with specific test configurations, results, and analysis. Navigate to individual directories for:

- Detailed test methodologies
- Raw performance data
- Configuration parameters
- Visual performance charts
- Setup and reproduction instructions

## Test Environment

- **Cloud Provider**: Amazon Web Services (AWS)
- **Instance Types**: EC2 instances ranging from t3.medium to c5.2xlarge
- **Operating System**: Linux-based testing environments
- **Network Configuration**: Various bandwidth limits and connection patterns
- **Monitoring**: CPU, memory, network throughput, and connection metrics

## Contributing

When adding new benchmarks or updating existing results:

1. Follow the existing directory structure
2. Include detailed README with test parameters
3. Provide raw data and visualizations
4. Document test environment and methodology
5. Update this main README with new component information

---

*This benchmarking suite provides comprehensive performance data to guide ZTA 2.0 implementation decisions and architecture choices.*
