package logger

//
//import (
//	"AvitoTask/config"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
////type Logger struct {
////	cfg    *config.Config
////	logger zap.Logger
////}
//
////func NewLogger(cfg *config.Config) *Logger {
////	return &Logger{cfg: cfg,
////		logger: zap.NewProduction()}
////}
//
//func NewLogger(cfg *config.Config) *zap.Logger {
//	encoderCfg := zap.NewProductionEncoderConfig()
//	encoderCfg.TimeKey = "timestamp"
//	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
//
//	config := zap.Config{
//		Level:             GetInfoLvl(cfg),
//		Development:       cfg.Logger.Development,
//		DisableCaller:     cfg.Logger.DisableCaller,
//		DisableStacktrace: cfg.Logger.DisableStacktrace,
//		Encoding:          cfg.Logger.Encoding,
//		EncoderConfig:     encoderCfg,
//		OutputPaths: []string{
//			"stderr",
//		},
//		ErrorOutputPaths: []string{
//			"stderr",
//		},
//	}
//
//	return zap.Must(config.Build())
//}
//
//func GetInfoLvl(cfg *config.Config) zap.AtomicLevel {
//	return zap.NewAtomicLevelAt(loggerLevelMap[cfg.Logger.Level])
//}
