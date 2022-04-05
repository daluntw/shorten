package db

import (
	"github.com/linxGnu/grocksdb"
	"go.uber.org/zap"
)

var globalConn *Conn

type Conn struct {
	dbConn *grocksdb.DB
}

func NewDB(filename string) (*Conn, error) {

	bbt := grocksdb.NewDefaultBlockBasedTableOptions()
	bbt.SetBlockCache(grocksdb.NewLRUCache(1 << 30))
	// set 1g lru cache
	// TODO: modify to changeable

	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbt)
	opts.SetCreateIfMissing(true)

	conn, err := grocksdb.OpenDb(opts, filename)
	if err != nil {
		return nil, err
	}

	return &Conn{dbConn: conn}, nil
}

func GetGlobalConn() *Conn {
	if globalConn == nil {
		zap.S().Panic("global db conn seems not init yet.")
	}
	return globalConn
}

func SetGlobalConn(c *Conn) {
	globalConn = c
}

func (c *Conn) Close() {
	c.dbConn.Close()
}

func (c *Conn) Get(key []byte) (*Record, error) {

	v, err := c.dbConn.Get(grocksdb.NewDefaultReadOptions(), key)
	if err != nil || v.Exists() == false || v.Size() == 0 {
		return nil, err
	}

	defer v.Free()

	r := &Record{}
	if _, err = r.UnmarshalMsg(v.Data()); err != nil {
		return r, err
	}

	return r, nil
}

func (c *Conn) Set(key []byte, r *Record) (exist bool, err error) {

	if gr, _ := c.Get(key); gr != nil {
		return true, nil
	}

	v, err := r.MarshalMsg(nil)
	if err != nil {
		return false, err
	}

	return false, c.dbConn.Put(grocksdb.NewDefaultWriteOptions(), key, v)
}