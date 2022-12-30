package log

import (
	"context"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"time"
)

type GormLogger struct {
	l          *Logger
	traceLevel zapcore.Level
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *g
	switch level {
	case logger.Error:
		newLogger.traceLevel = zapcore.ErrorLevel
	case logger.Warn:
		newLogger.traceLevel = zapcore.WarnLevel
	case logger.Info:
		newLogger.traceLevel = zapcore.InfoLevel
	case logger.Silent:
		newLogger.traceLevel = zapcore.DebugLevel
	default:
		newLogger.traceLevel = zapcore.DebugLevel
	}
	return &newLogger
}

func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	g.l.Error(ctx, msg, data)
}

func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	g.l.Warn(ctx, msg, data)
}
func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	g.l.Info(ctx, msg, data)
}

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if !g.l.level.Enabled(g.traceLevel) {
		return
	}
	elapsed := time.Since(begin)
	sql, rows := fc()
	var rowsField interface{} = rows
	if rows == -1 {
		rowsField = "-"
	}
	fields := []interface{}{"elapsed", elapsed, "rows", rowsField, "sql", sql}

	switch {
	case err != nil:
		fields = append(fields, "err", err)
		g.Error(ctx, "gorm.trace", fields)
	case elapsed > g.l.config.SQLSlowThreshold && g.l.config.SQLSlowThreshold != 0:
		fields = append(fields, "slow_threshold", g.l.config.SQLSlowThreshold)
		g.Warn(ctx, "gorm.trace", fields)
	}
}

func NewGormLogger(config *Config) (l logger.Interface, err error) {
	if config == nil {
		config = &Config{
			Filename:         "",
			MaxSize:          0,
			MaxAge:           0,
			MaxBackups:       0,
			LocalTime:        false,
			Compress:         false,
			Level:            "info",
			Console:          "stderr",
			SQLSlowThreshold: time.Second,
			GormTraceLevel:   "debug",
			GormCallerBorder: "models/model.go",
		}
	}
	traceLevel := zapcore.InfoLevel
	err = traceLevel.Set(config.GormTraceLevel)
	if err != nil {
		return nil, err
	}
	logger, err := newLogger(config, 2)
	if err != nil {
		return nil, err
	}

	l = &GormLogger{l: logger, traceLevel: traceLevel}
	return l, nil
}
