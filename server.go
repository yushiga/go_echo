package main

import (
	"github.com/hmhr/go_sample/config"
	"github.com/hmhr/go_sample/db"
	"github.com/hmhr/go_sample/logger"
	"github.com/hmhr/go_sample/route"

	"flag"
)

/*
main関数
*/
func main() {
	// ログ設定
	logger.Setup()
	logger.ZapLog.Info("aaaaaaaaa")

	// 実行時引数から環境設定
	flag.Parse()
	config.SetEnvironment(flag.Args()[0])

	// DB接続設定
	db.Init()
	// router設定
	e := route.Init()
	// アプリ実行
	e.Logger.Fatal(e.Start(":1323"))
}
