package log

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLoggerConfig struct {
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
	Logger        *zap.SugaredLogger
}

func NewGormLogger(slowThreshold time.Duration, logLevel logger.LogLevel) logger.Interface {
	log := InitLogger()

	return &GormLoggerConfig{
		SlowThreshold: slowThreshold,
		LogLevel:      logLevel,
		Logger:        log,
	}
}

func (l *GormLoggerConfig) LogMode(LogLevel logger.LogLevel) logger.Interface {
	l.LogLevel = LogLevel
	l.Logger.Infof("Gorm log level set to %s", LogLevel)
	return l
}

func (l *GormLoggerConfig) Info(ctx context.Context, msg string, params ...interface{}) {
	l.Logger.Infof(msg, params...)
}

func (l *GormLoggerConfig) Warn(ctx context.Context, msg string, params ...interface{}) {
	l.Logger.Warnf(msg, params...)
}

func (l *GormLoggerConfig) Error(ctx context.Context, msg string, params ...interface{}) {
	l.Logger.Errorf(msg, params...)
}

func (l *GormLoggerConfig) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []interface{}{
		"elapsed", elapsed,
		"rows", rows,
		"sql", sql,
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		fields = append(fields, "error", err)

		l.Logger.Errorw("query", fields...)
		return
	}

	if err == gorm.ErrRecordNotFound {
		l.Logger.Warnw("query not found", fields...)
		return
	}

	if elapsed > l.SlowThreshold {
		l.Logger.Warnw("slow query", fields...)
		return
	}

	l.Logger.Infow("query", fields...)
}
