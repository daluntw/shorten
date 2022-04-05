package handler

import (
	"fmt"
	"github.com/daluntw/shorten/db"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func GetHandler(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" || IDValidate(id) == false {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"resp": "id not valid"})
		return
	}

	if r, err := db.GetGlobalConn().Get([]byte(id)); r == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"resp": "id not found"})
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		zap.S().Error("msgpack decode err: ", err)
	} else if r.Expire.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "expired"})
	} else {
		ctx.Redirect(http.StatusFound, r.Dest)
	}
}

func SetHandler(ctx *gin.Context) {

	type SetRequest struct {
		URL string `json:"url"`
		ExpireStr string `json:"expireAt"`
	}

	req := SetRequest{}
	id := IDGenerate()
	if jsonErr := ctx.BindJSON(&req); jsonErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": jsonErr.Error()})
	} else if URLValidate(req.URL) == false {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "url parse failed"})
	} else if t, timeErr := time.Parse(time.RFC3339, req.ExpireStr); timeErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": timeErr.Error()})
	} else if exist, setErr := db.GetGlobalConn().Set([]byte(id), &db.Record{
		Dest:   req.URL,
		Expire: t,
	}); exist == true {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "key collision, try again"})
	} else if setErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": setErr.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id": id,
			"shortUrl": fmt.Sprintf("https://dalun.in/%s", id),
		})
	}
}