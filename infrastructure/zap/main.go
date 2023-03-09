package zap

import "go.uber.org/zap"

func New(isDebug bool) *zap.SugaredLogger {
	zapOpt := []zap.Option{
		zap.AddCallerSkip(2),
	}

	logger, err := zap.NewProduction(zapOpt...)
	if err != nil {
		panic(err.Error())
	}

	if isDebug {
		logger, err = zap.NewDevelopment(zapOpt...)
		if err != nil {
			panic(err.Error())
		}
	}

	return logger.Sugar()
}
