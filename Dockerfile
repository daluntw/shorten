FROM golang:1.18-alpine

WORKDIR /tmp
ENV ROCKSDB_VERSION='7.0.4'
RUN apk add -U git build-base linux-headers cmake bash zlib-dev bzip2-dev snappy-dev lz4-dev librdkafka-dev zstd-dev curl perl
RUN wget https://github.com/facebook/rocksdb/archive/v${ROCKSDB_VERSION}.zip -O rocksdb-${ROCKSDB_VERSION}.zip
RUN unzip rocksdb-${ROCKSDB_VERSION}.zip -d .
WORKDIR rocksdb-${ROCKSDB_VERSION}
RUN make -j8 static_lib

WORKDIR /srv
COPY . /srv/
RUN CGO_CFLAGS="-I /tmp/rocksdb-$ROCKSDB_VERSION/include" \
    CGO_LDFLAGS="-L /tmp/rocksdb-$ROCKSDB_VERSION -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4" \
    go build -o app .

ENTRYPOINT ["/srv/app"]