package main

import (
	"github.com/wyvern-go/pingxx_log"
	"time"
)

var PingxxLog *pingxx_log.Logger

func main() {
	var i int
	for {
		i++
		PingxxLog.Debug("test for log %d", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func init() {
	cfg := pingxx_log.NewConfig(pingxx_log.ToConsole)
	cfg.SetCacheSize(2048) //可设置,也可以不设置,默认1024
	cfg.SetCententType(pingxx_log.ToConsole, pingxx_log.FormatJson)
	cfg.SetLevel(pingxx_log.Debug)
	PingxxLog = pingxx_log.New(cfg)
	PingxxLog.SetModule("example")
	PingxxLog.Cache.CacheMonitor()//监控缓存 可开可不开 不开注释即可
}