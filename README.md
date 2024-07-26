# Basic Todo App (for million concurrent users)


1. Basic API
2. Use Connnection Pool Manager, dont use GORM. Cause it tries to establish a connection for each API call. So inefficient.
Using connection pool size of 20
3. Index Columns on tables
4. Use Go Routines
With Go Routines all the requests are handled. without it, it hits a failure after a threshold. With Asynchronous handling using Go Routines,
even at 1Million concurrent users it was running alrighty.
5. Enhance NGINX
    - Enable worker processes, give 200(takes up RAM)
    - NGINX worker's file descriptor limit
    - max number of connections per worker
6. Increase File descriptor limit

    In Linux, everything is a file. A TCP connection is also a file, in order to support more TCP connections. The limit of max open files needs to be increased. There are two limits, hard and soft. Hard limit is already at its max. needed to increase soft limit
    ```
    ulimit -Hn // Hard
    ulimit -Sn // Soft
    ```
    Need to edit /etc/security/limits.conf
    *       nofile  soft <limit>
7. Handle TCP bottleneck

    Every TCP connection requires a TCP header frame. That includes, Src IP, Src Address, Dest IP, Dest Address.
    The TCP header needs to be unique in order for the system to uniquely identify a TCP connection.
    The port range in any machine 0-65535. OS reserves 0-500. so the rest is up for grabs. Needed to increase the port range to utilise them all.
    by default its 32k-65k. 
    `$ sudo sysctl -w net.ipv4.ip_local_port_range="1024 65535"`

    So to generate 1 million requests per second. Use all 65k ports on src machine and run the application on 20ports on the client machine. 
    That is 1.2 million unique combinations with just port number. This helped me resolve the problem.
    Another solution is to use IP aliasing but it didnt get picked up automatically.


<hr/>

1. Use Premium Tier Networking: GCP offers premium tier networking, which can provide lower latency and higher reliability compared to standard tier networking. Consider using premium tier networking for your instances.
2. Leverage Google's Network Backbone: GCP benefits from Google's extensive global network infrastructure, including high-capacity fiber optic cables. Utilize this backbone for faster data transmission between regions and zones.
3. Utilize Content Delivery Networks (CDNs): Distribute your static content and assets through a CDN like Google Cloud CDN or other third-party CDNs. CDNs cache content closer to users, reducing latency by serving content from edge locations.
4. Implement HTTP/2 or QUIC: Use modern protocols like HTTP/2 or QUIC, which are designed to reduce latency through techniques such as multiplexing, header compression, and connection reuse.
5. Tune TCP/IP Parameters: Adjust TCP/IP settings on your VM instances to optimize network performance. This may include tuning parameters like TCP window size, TCP congestion control algorithms, and TCP keepalive settings.
6. Reduce Packet Loss and Jitter: Minimize packet loss and jitter by using reliable network connections and optimizing network configurations. High packet loss and jitter can significantly impact latency-sensitive applications.

<hr/>

###  Signoz Setup

Install

```
git clone -b main https://github.com/SigNoz/signoz.git
cd signoz/deploy/
./install.sh
```
Run

```
SERVICE_NAME=goTodosApp INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 go run *.go
```

Stop

```
docker-compose -f ./docker/clickhouse-setup/docker-compose.yaml down -v
```


### Huge thanks to
- [Scaling to 12 Million](https://mrotaru.wordpress.com/2013/10/10/scaling-to-12-million-concurrent-connections-how-migratorydata-did-it/)