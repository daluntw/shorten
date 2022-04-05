package main

import (
    "flag"
    "github.com/daluntw/shorten/db"
    "github.com/daluntw/shorten/handler"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

var (
    debug  bool
    addr   string
    dbname string
)

func init() {

    flag.BoolVar(&debug, "debug", false, "-debug=false")
    flag.StringVar(&addr, "addr", ":9999", "-addr=:9999")
    flag.StringVar(&dbname, "db", "/srv/shorten.db", "-db=/srv/shorten.db")
    flag.Parse()

    if debug == false {
        gin.SetMode(gin.ReleaseMode)
    } else {
        gin.SetMode(gin.DebugMode)
    }

    logger, err := zap.NewProduction()
    if err != nil {
        panic(err)
    }
    zap.ReplaceGlobals(logger)
}

func main() {

    c, cErr := db.NewDB(dbname)
    if cErr != nil {
        zap.S().Panic("db can not open or create: ", cErr)
    }

    db.SetGlobalConn(c)
    defer c.Close()

    r := gin.New()
    r.POST("/api/v1/urls", handler.SetHandler)
    r.GET("/:id", handler.GetHandler)

    zap.S().Panic(r.Run(addr))
}
