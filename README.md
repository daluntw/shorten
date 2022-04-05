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

> Although this is go project, but we have to compile RocksDB. So it will take all of your cpu & memory. :(
> If build failed on OSX & Windows, please adjust your Docker VM memory. 

### Benchmark

# TODO