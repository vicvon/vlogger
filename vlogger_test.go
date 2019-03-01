package vlogger

import "testing"

func TestVlogger(t *testing.T) {
    cfg := &VloggerConfig{
        LogFileName: "vlog.log",
        LogMaxAge: 1,
        LogMaxSize: 5,
        LogMaxBackUp: 5,
        LogCompress: false,
        LogLevel: "debug",
    }

    log := NewVlogger(cfg)
    for i := 0; i < 1; i++ {
        log.Infof("this info")
        log.Warnf("this info")
        log.Debugf("this info")
        log.Errorf("this info")
    }
}

