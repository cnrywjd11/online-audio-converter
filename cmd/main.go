package main

import (
	"github.com/cnrywjd11/online-audio-converter/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(
		middleware.Recover(),
	)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"method":"${method}","uri":"${uri}","status":${status},"error":"${error}"}` + "\n",
	}))
	e.POST("/convert", handler.ConvertAudioHandler)
	e.Logger.Fatal(e.Start(":80"))
}
