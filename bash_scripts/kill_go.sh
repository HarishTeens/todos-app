#!/bin/bash

# Define the range of port numbers
start_port=8080
end_port=8094

# Loop over the port range
for (( port=start_port; port<=end_port; port++ )); do
    # Run lsof to find the process listening on the port
    lsof_output=$(sudo lsof -n -i :$port | grep LISTEN)
    if [ -n "$lsof_output" ]; then
        # Extract the process ID (2nd column) from lsof output
        pid=$(echo "$lsof_output" | awk '{print $2}')
        if [ -n "$pid" ]; then
            # Kill the process
            echo "Killing process $pid listening on port $port"
            sudo kill -9 $pid
        else
            echo "No process found listening on port $port"
        fi
    else
        echo "No process found listening on port $port"
    fi
done
