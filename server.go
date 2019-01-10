package main

import (
	"github.com/hmhr/go_sample/config"
	"github.com/hmhr/go_sample/db"
	"github.com/hmhr/go_sample/logger"
	"github.com/hmhr/go_sample/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"flag"
)

/*
main関数
*/
func main() {
	// 実行時引数から環境設定
	flag.Parse()
	config.SetEnvironment(flag.Args()[0])

	// ログ設定
	logger.Setup()

	// DB接続設定
	db.Init()

	// router設定
	e := route.Init()

	// エラー時にスタックトレースを出力
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logger.ZapLog.Error(err.Error())
	}

	// アプリ実行
	e.Logger.Fatal(e.Start(":1323"))
}
