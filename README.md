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


### Huge thanks to
- [Scaling to 12 Million](https://mrotaru.wordpress.com/2013/10/10/scaling-to-12-million-concurrent-connections-how-migratorydata-did-it/)