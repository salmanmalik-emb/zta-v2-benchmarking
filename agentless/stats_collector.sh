#!/bin/sh

# Define Interface Name to Monitor
iface=wg1

touch stats.log

get_net_packets () {
  curr_rx_packets=$(cat /sys/class/net/"$iface"/statistics/rx_packets)
  # curr_tx_packets=$(cat /sys/class/net/"$iface"/statistics/tx_packets)
  # total_curr_packets=$((curr_rx_packets + curr_tx_packets))

  echo "$curr_rx_packets"
}

get_net_bytes () {
  curr_rx_bytes=$(cat /sys/class/net/"$iface"/statistics/rx_bytes)
  # curr_tx_bytes=$(cat /sys/class/net/"$iface"/statistics/tx_bytes)
  # total_curr_bytes=$((curr_rx_bytes + curr_tx_bytes))

  echo "$curr_rx_bytes"
}


last_packets_count=$(get_net_packets)
packets_count_per_minute=0
total_packets_count_so_far=0

last_bytes_size=$(get_net_bytes)
bytes_size_per_minute=0
total_bytes_size_so_far=0

counter=1

# Main loop to continuously calculate and log stats per minute
while true
do
  # Calculate and log the CPU utilization
  cpu_util=$(sar -u 60 1 | tail -n 1 | awk '{print 100-$NF}')

  # Calculate and log the memory usage
  mem_total=$(free -m | awk '{if (NR==2) print $2}')
  mem_used=$(free -m | awk '{if (NR==2) print $3}')
  mem_util=$(awk "BEGIN {printf \"%.2f\", $mem_used/$mem_total*100}")
  
  # Calculate and log the total packets so far
  total_curr_packets=$(get_net_packets)

  packets_count_per_minute=$((total_curr_packets - last_packets_count))
  total_packets_count_so_far=$((total_packets_count_so_far + packets_count_per_minute))
  last_packets_count=$total_curr_packets

  # Calculate and log total packets size so far
  total_curr_bytes=$(get_net_bytes)
  
  bytes_size_per_minute=$((total_curr_bytes - last_bytes_size))
  total_bytes_size_so_far=$((total_bytes_size_so_far + bytes_size_per_minute))
  
  last_bytes_size=$total_curr_bytes

  # Calculate and log the average throughput

  throughput_per_minute=$(echo "scale=4; ($bytes_size_per_minute / 1250000) * 8 / 60" | bc) # formula to calculate throughput in Mbps

  overall_throughput=$(echo "scale=4; (($total_bytes_size_so_far / 1250000) * 8) / (60 * $counter)" | bc) # calculate throughput in Mbps for overall duration of script run

  # Write stats to log file
  echo "$(date '+%Y-%m-%d %H:%M:%S'), PPM: $packets_count_per_minute, BPM: $bytes_size_per_minute, TPM: $throughput_per_minute Mbps, CPU: $cpu_util%, Mem: $mem_util%, TPs: $total_packets_count_so_far, TBs: $total_bytes_size_so_far, TT: $overall_throughput" >> stats.log

  # Wait for 1 minute before logging stats again
  # sleep 60

  # Counter to keep tracks of minutes passed

  if [[ $throughput_per_minute -gt 3 ]]
  then
  counter=$((counter+1))
  fi
done
