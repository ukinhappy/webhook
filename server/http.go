package server

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/ukinhappy/webhook/hook"
	"github.com/ukinhappy/webhook/logger"
	"net/http"
)

func Server() {
	e := echo.New()
	e.GET("/ping", Ping)
	e.POST("/webhook_deploy", WebHookDeploy)
	logger.Fatalf("start server failed %v", e.Start(viper.GetString("http.addr")))
}

func Ping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "success")
}

func WebHookDeploy(ctx echo.Context) error {
	var param hook.WebHookRequest
	if err := ctx.Bind(&param); err != nil {
		logger.Errorf("bind param failed %v", err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	logger.Infof("%v", param)
	event := ctx.Request().Header.Get("X-GitHub-Event")
	if err := hook.Do(param.Repository.Name, "", event); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "success")
}
