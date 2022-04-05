# Shorten

A useful tool to short url with speed.

### Feature & Tech Stack

- less than 300 lines (without test and generated code)
- use Gin HTTP Framework
- checksum in id, prevent malicious id
- use MessagePack as serializer
- use RocksDB as storage, it's super-fast and easy to replace to other distributed KV for scalability

### Usage

```shell
$ docker build -t shorten .
$ docker run --rm -it -p 80:80 shorten -debug=false -addr=:80 
```

> If build failed on OSX & Windows, please adjust your Docker VM memory. It will take ~5G Mem and 5min

### Benchmark

Test on 2x 4C8G with NVMe VPS (1x server, 1x client) 

#### 1M malicious id get request
```
Elapsed 18.29s / result 54673.784 RPS
Latency Percentile:
  P50       P75     P90     P95       P99      P99.9     P99.99
  4.264ms  5.29ms  7.3ms  10.307ms  16.283ms  20.815ms  28.519ms
```

#### 1M fixed expired id get request
```
Elapsed 22.565s / result 44315.798 RPS
Latency Percentile:
  P50        P75      P90       P95       P99      P99.9     P99.99
  4.255ms  6.891ms  11.445ms  17.065ms  29.681ms  39.789ms  48.444ms
```

#### 1M write request
```
Elapsed 51.065s / result 19582.519 RPS 
Latency Percentile:
  P50        P75       P90       P95       P99       P99.9     P99.99
  9.576ms  14.778ms  24.353ms  36.428ms  68.204ms  117.723ms  189.657ms
```

# TODO