package main

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	log "klog"
	"klog/klog"
)

func main() {
	// 初始化logger
	loggerConfig := &zap.Config{}
	l, _ := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	klog.InitLogger(l)

	// Debug、Info(with field)、Warnf、Errorw使用
	log.Debug("This is a debug message")
	log.Info("This is a info message", log.Int32("int_key", 10))
	log.Warnf("This is a formatted %s message", "WARN")
	log.Errorw("Message printed with Errorw", "X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")

	// Context使用
	ctx := context.WithValue(context.Background(), "requestID", "2a7b9f24-4ace-4b2a-9464-69238b45b953")
	log.L(ctx).Infof("fetch datasource success")
	log.L(ctx).Errorf("fetch datasource failed")

	// WithValues使用
	lv := log.WithValues("X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	lv.Infow("Info message printed with [WithValues] logger")
	lv.Infow("Debug message printed with [WithValues] logger")

	ln := lv.WithName("test")
	ln.Info("Message printed with [WithName] logger")

	// level使用
	log.V(log.InfoLevel).Info("This is a V level message")
	log.V(log.ErrorLevel).Infow("This is a V level message with fields", "X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	log.V(0).Info("This is a V level message")
	log.V(1).Info("This is a V level message")
}
