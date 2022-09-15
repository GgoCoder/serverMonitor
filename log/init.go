package log

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(fileName string, level zapcore.LevelEnabler) *zap.SugaredLogger{
    writeSyncer := getLogWriter(fileName)
    encoder := getEncoder()
    core := zapcore.NewCore(encoder, writeSyncer, level)

    logger := zap.New(core)
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
    return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil{
		log.Panic("failed to create log file", fileName)
	}
    return zapcore.AddSync(file)
}
