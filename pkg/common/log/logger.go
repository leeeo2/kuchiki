package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

type Config struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool

	// Level available values are debug, info, warn, error
	Level string
	// Console available values are stdout, stderr, 1, 2
	Console string

	// SQLSlowThreshold is used in gorm logger
	SQLSlowThreshold time.Duration
	// GormTraceLevel is level for gorm trace
	GormTraceLevel string
	// GormCallerBorder is the border for caller stack, caller is the direct caller to this file
	GormCallerBorder string
}

type Logger struct {
	config Config
	z      *zap.Logger
	level  zapcore.Level
}

func newLogger(config *Config, callSkip int) (*Logger, error) {

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "M",
		LevelKey:         "L",
		TimeKey:          "T",
		NameKey:          "N",
		CallerKey:        "C",
		FunctionKey:      "F",
		StacktraceKey:    "S",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "\t",
	}

	wrirteSyncers := make([]zapcore.WriteSyncer, 0)
	lumberjack := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		LocalTime:  config.LocalTime,
		Compress:   config.Compress,
	}
	wrirteSyncers = append(wrirteSyncers, zapcore.AddSync(lumberjack))
	switch strings.ToLower(config.Console) {
	case "stdout", "1":
		wrirteSyncers = append(wrirteSyncers, zapcore.AddSync(os.Stdout))
	case "stderr", "2":
		wrirteSyncers = append(wrirteSyncers, zapcore.AddSync(os.Stderr))
	}

	level := zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, err
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(wrirteSyncers...),
		level)

	options := []zap.Option{
		zap.WithCaller(true),
		zap.AddCallerSkip(callSkip),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	l := zapcore.InfoLevel
	if err := l.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, err
	}

	return &Logger{
		config: *config,
		z:      zap.New(core, options...),
		level:  l,
	}, nil
}

func (l *Logger) Close() error {
	return l.z.Sync()
}

//func (l *Logger) Enabled(level zapcore.Level) bool {
//	return l.level.Enabled(level)
//}

func zapFields(ctx context.Context, fields ...interface{}) []zap.Field {
	ret := make([]zap.Field, 0, len(fields))
	for i := 0; i < len(fields)-1; i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			continue
		}
		ret = append(ret, zap.Any(key, fields[i+1]))
	}
	return ret
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...interface{}) {
	if l.level.Enabled(zap.DebugLevel) {
		l.z.Debug(msg, zapFields(ctx, fields)...)
	}
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...interface{}) {
	if l.level.Enabled(zap.InfoLevel) {
		l.z.Info(msg, zapFields(ctx, fields)...)
	}
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...interface{}) {
	if l.level.Enabled(zap.WarnLevel) {
		l.z.Warn(msg, zapFields(ctx, fields)...)
	}
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...interface{}) {
	if l.level.Enabled(zap.ErrorLevel) {
		l.z.Warn(msg, zapFields(ctx, fields)...)
	}
}

func (l *Logger) Panic(ctx context.Context, msg string, fields ...interface{}) {
	if l.level.Enabled(zap.PanicLevel) {
		l.z.Panic(msg, zapFields(ctx, fields)...)
	}
}

func Setup(cfg *Config) (*Logger, error) {
	logger, err := newLogger(cfg, 2)
	if err != nil {
		fmt.Printf("error to init logger:%v\n", err)
		return nil, err
	}
	return logger, nil
}

////// Global
var gLogger *Logger

func SetupGlobal(cfg *Config) error {
	if gLogger != nil {
		err := gLogger.Close()
		if err != nil {
			fmt.Printf("close gLogger failed,err:%v\n", err)
		}
	}

	logger, err := newLogger(cfg, 2)
	if err != nil {
		fmt.Printf("new logger failed.err:%v\n", err)
		return err
	}
	gLogger = logger
	zap.ReplaceGlobals(gLogger.z)
	return nil
}

func Global() *Logger {
	return gLogger
}

func Debug(ctx context.Context, msg string, fields ...interface{}) {
	gLogger.Debug(ctx, msg, fields)
}

func Info(ctx context.Context, msg string, fields ...interface{}) {
	gLogger.Info(ctx, msg, fields)
}

func Warn(ctx context.Context, msg string, fields ...interface{}) {
	gLogger.Warn(ctx, msg, fields)
}

func Error(ctx context.Context, msg string, fields ...interface{}) {
	gLogger.Error(ctx, msg, fields)
}

func Panic(ctx context.Context, msg string, fields ...interface{}) {
	gLogger.Panic(ctx, msg, fields)
}
