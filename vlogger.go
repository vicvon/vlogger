package vlogger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

type VloggerConfig struct {
    LogFileName         string
    LogMaxSize          int
    LogMaxAge           int
    LogMaxBackUp        int
    LogCompress         bool
    LogLevel            string
}

type Vlogger struct {
    logCfg *VloggerConfig
    log *zap.SugaredLogger
}

func NewVlogger(cfg * VloggerConfig) *Vlogger {
    w := zapcore.AddSync(&lumberjack.Logger{
        Filename: cfg.LogFileName,
        MaxSize: cfg.LogMaxSize,
        MaxAge: cfg.LogMaxAge,
        MaxBackups: cfg.LogMaxBackUp,
        Compress: cfg.LogCompress,
    })

    logLvl := func () zapcore.Level {
        switch cfg.LogLevel {
        case "debug":
            return zapcore.DebugLevel
        case "info":
            return zapcore.InfoLevel
        case "warn":
            return zapcore.WarnLevel
        case "error":
            return zapcore.ErrorLevel
        default:
            return zapcore.WarnLevel
        }
    }()

    zapCfg := zap.NewProductionEncoderConfig()
    zapCfg.EncodeTime = zapcore.ISO8601TimeEncoder

    core := zapcore.NewCore(zapcore.NewJSONEncoder(zapCfg), w, logLvl)
    logger := zap.New(core, zap.AddCaller()).Sugar()

    return &Vlogger{
        logCfg: cfg,
        log: logger,
    }
}

func (l *Vlogger) Infof(template string, args ...interface{}) {
    l.log.Infof(template, args...)
}

func (l *Vlogger) Debugf(template string, args ...interface{}) {
    l.log.Debugf(template, args...)
}

func (l *Vlogger) Warnf(template string, args ...interface{}) {
    l.log.Warnf(template, args...)
}

func (l *Vlogger) Errorf(template string, args ...interface{}) {
    l.log.Errorf(template, args...)
}
