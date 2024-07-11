# .NET Core vs Golang performance

This repository holds code to test http api performance between .NET Core and Golang HTTP.
Each service has `/test` endpoint which calls another api using http client and returns that api response as JSON.

## Start containers

`docker-compose up --build`

docker-compose should start 3 containers
1) Golang api with GET `http://localhost:5002/data` endpoint
2) Golang api with GET `http://localhost:5001/test` endpoint which calls 1 endpoint
3) .NET Core api with GET `http://localhost:5000/test` endpoint which calls 1 endpoint

## Run load tests

```
brew install wrk
cd wrk
// .net core
URL=http://localhost:5000 make run

// golang
URL=http://localhost:5001 make run
```

## Check for file descriptors leaks

Connect to docker container while wrk is running
`docker exec -it <CONTAINER_ID> /bin/bash`

Count TIME_WAIT state
`ss -t4 state time-wait | wc -l`

## Results

### .net core api (http://localhost:5000)

```
Running 2m test @ http://localhost:5000
  8 threads and 256 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.06ms    4.86ms 180.39ms   96.50%
    Req/Sec     5.58k   843.84     7.24k    77.55%
  Latency Distribution
     50%    5.37ms
     75%    6.59ms
     90%    8.31ms
     99%   18.00ms
  4442164 requests in 1.67m, 0.94GB read
  Socket errors: connect 0, read 125, write 0, timeout 0
Requests/sec:  44359.35
Transfer/sec:      9.59MB
```

Resources used
```
CPU: 100%
MEMORY: 82MB
TIME_WAIT file descriptors: ~3
```

### golang api (http://localhost:5001)

```
Running 2m test @ http://localhost:5001
  8 threads and 256 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.29ms    8.54ms 248.88ms   91.18%
    Req/Sec     3.38k   522.76     5.13k    83.79%
  Latency Distribution
     50%    8.49ms
     75%   12.56ms
     90%   17.41ms
     99%   39.32ms
  2684929 requests in 1.67m, 491.01MB read
  Socket errors: connect 0, read 150, write 0, timeout 0
Requests/sec:  26821.91
```

Resources used
```
CPU: 100%
MEMORY: 25.57MB
TIME_WAIT file descriptors: ~10
```

## My machine spec

* MacBook Pro (15-inch, 2017)
* Processor 2,9 GHz Intel Core i7
* Memory 16 GB 2133 MHz LPDDR3
* Docker version 18.06.0-ce (4 CPUs, 2 GiB memory)
* Golang 1.11.2
* Dotnet 2.1.0
