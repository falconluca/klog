# klog

抽取了 kubernetes 日志 logger。可用于业务开发。

一个简单的示例，创建一个 `main.go` 文件，内容如下：

```go
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
	log.Errorw("Message printed with Errorw", 
		"X-Request-ID", "fbf54504-64da-4088-9b86-67824a7fb508")

	// Context使用
	ctx := context.WithValue(context.Background(), 
		"requestID", "2a7b9f24-4ace-4b2a-9464-69238b45b953")
	log.L(ctx).Infof("fetch datasource success")
	log.L(ctx).Errorf("fetch datasource failed")

	// WithValues使用
	lv := log.WithValues("X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	lv.Infow("Info message printed with [WithValues] logger")
	lv.Infow("Debug message printed with [WithValues] logger")

	ln := lv.WithName("test")
	ln.Info("Message printed with [WithName] logger")

	// Level使用
	log.V(log.InfoLevel).Info("This is a V level message")
	log.V(log.ErrorLevel).Infow("This is a V level message with fields", 
		"X-Request-ID", "7a7b9f24-4cae-4b2a-9464-69088b45b904")
	log.V(0).Info("This is a V level message")
	log.V(1).Info("This is a V level message")
}

```

执行代码：

```bash
$ go run main.go 
2023-03-28 10:06:09.843 INFO    example/main.go:19      This is a info message  {"int_key": 10}
2023-03-28 10:06:09.843 WARN    example/main.go:20      This is a formatted WARN message
2023-03-28 10:06:09.843 ERROR   example/main.go:21      Message printed with Errorw     {"X-Request-ID": "fbf54504-64da-4088-9b86-67824a7fb508"}
2023-03-28 10:06:09.843 INFO    example/main.go:25      fetch datasource success        {"requestID": "2a7b9f24-4ace-4b2a-9464-69238b45b953"}
2023-03-28 10:06:09.843 ERROR   example/main.go:26      fetch datasource failed {"requestID": "2a7b9f24-4ace-4b2a-9464-69238b45b953"}
2023-03-28 10:06:09.843 INFO    example/main.go:30      Info message printed with [WithValues] logger   {"X-Request-ID": "7a7b9f24-4cae-4b2a-9464-69088b45b904"}
2023-03-28 10:06:09.843 INFO    example/main.go:31      Debug message printed with [WithValues] logger  {"X-Request-ID": "7a7b9f24-4cae-4b2a-9464-69088b45b904"}
2023-03-28 10:06:09.843 INFO    test    example/main.go:34      Message printed with [WithName] logger  {"X-Request-ID": "7a7b9f24-4cae-4b2a-9464-69088b45b904"}
2023-03-28 10:06:09.843 INFO    example/main.go:37      This is a V level message
2023-03-28 10:06:09.843 ERROR   example/main.go:38      This is a V level message with fields   {"X-Request-ID": "7a7b9f24-4cae-4b2a-9464-69088b45b904"}
2023-03-28 10:06:09.843 INFO    example/main.go:39      This is a V level message
2023-03-28 10:06:09.843 WARN    example/main.go:40      This is a V level message
```

上述代码使用 `falconluca/klog` 包默认的全局 `logger`，分别使用了 **Context使用** 、**WithValues使用** 和 **Level使用** 级别打印了一条日志。