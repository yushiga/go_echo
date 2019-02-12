package logger

import (
	"io/ioutil"

	"github.com/yushiga/go_echo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	yaml "gopkg.in/yaml.v2"
)

var ZapLog *zap.Logger

/*
ログ設定
*/
func Setup() {
	// ログをファイル出力
	zapyml, err := ioutil.ReadFile("/aaa/bbb/zap_conf.yml")
	if err != nil {
		panic(err)
	}
	var zapconf zap.Config
	if err := yaml.Unmarshal(zapyml, &zapconf); err != nil {
		panic(err)
	}
	enc := zapcore.NewJSONEncoder(zapconf.EncoderConfig)
	sink := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   zapconf.OutputPaths[1],
			MaxAge:     config.Config.Logger.MaxAge, //days
			MaxBackups: config.Config.Logger.MaxBackup,
		},
	)
	ZapLog = zap.New(
		zapcore.NewCore(enc, sink, zapconf.Level),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	defer ZapLog.Sync()
}
