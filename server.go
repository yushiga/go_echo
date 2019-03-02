package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yushiga/go_echo/config"
	"github.com/yushiga/go_echo/db"
	"github.com/yushiga/go_echo/logger"
	"github.com/yushiga/go_echo/route"
)

/*
main関数
*/
func main() {
	// 実行時引数から環境設定
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		panic("No Startup Argument")
	}

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
