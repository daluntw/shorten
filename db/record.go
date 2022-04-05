package db

import "time"

//go:generate msgp
type Record struct {
	Dest string `msg:"dest"`
	Expire time.Time `msg:"expire"`
}