FROM daluntw/go-1.18-rocksdb-7.0.4-alpine

WORKDIR /srv
COPY . /srv/
RUN CGO_CFLAGS="-I /tmp/rocksdb-$ROCKSDB_VERSION/include" \
    CGO_LDFLAGS="-L /tmp/rocksdb-$ROCKSDB_VERSION -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4" \
    go build -o app .

ENTRYPOINT ["/srv/app"]
