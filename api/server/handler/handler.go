package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/takutakukatokatojapan/go_redis/api/types"
	"github.com/takutakukatokatojapan/go_redis/infrastructure/appctx"
	"net/http"
)

func SetHandler(context echo.Context) error {
	c := context.(*appctx.APPContext)
	var req types.SetRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "json can not bind")
	}
	if err := c.Redis.Set(req.ID, req.Name); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, types.SetResponse{Code: "500"})
	}
	return c.JSON(http.StatusOK, types.SetResponse{Code: "200"})
}

func GetHandler(context echo.Context) error {
	c := context.(*appctx.APPContext)
	var req types.GetRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "json can not bind")
	}
	r, err := c.Redis.Get(req.ID)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, types.GetResponse{Code: "500"})
	}
	return c.JSON(http.StatusOK, types.GetResponse{Name: r, Code: "200"})
}
