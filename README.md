## pingxx日志

### pingxx日志架构

![pingxx_log](http://7xs3v3.com1.z0.glb.clouddn.com/pingxx_log.png)
-------------------------------------------------------------------
pingxx日志,内置了2个缓存container,所有的日志输入时,会把日志中的数据先存入container1
中,待container1存满之后,会将container1释放并输出,并且,接下来的日志会放入container2
中,如此循环

- 打印文本日志
    - 输入格式是文本:
    
            cfg := pingxx_log.NewConfig(pingxx_log.ToConsole)
            cfg.SetCententType(pingxx_log.ToConsole, pingxx_log.FormatText)
            
            2016/05/23 15:51:19 [example.go:14] <example>  Debug: test for log 1
            2016/05/23 15:51:19 [example.go:14] <example>  Debug: test for log 2
            2016/05/23 15:51:19 [example.go:14] <example>  Debug: test for log 3
            
    - 输出格式是json:
        
            cfg := pingxx_log.NewConfig(pingxx_log.ToConsole)
            cfg.SetCententType(pingxx_log.ToConsole, pingxx_log.FormatJson)
            
            {"log_id":"","log_level":"Debug","module":"example","time":"2016/05/23 16:12:41","filename":"example.go","line":14,"remark":"test for log 1"}
            {"log_id":"","log_level":"Debug","module":"example","time":"2016/05/23 16:12:42","filename":"example.go","line":14,"remark":"test for log 2"}
            {"log_id":"","log_level":"Debug","module":"example","time":"2016/05/23 16:12:42","filename":"example.go","line":14,"remark":"test for log 3"}

- Usage:

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
            cfg.SetCententType(pingxx_log.ToConsole, pingxx_log.FormatText)
            cfg.SetLevel(pingxx_log.Debug)
            PingxxLog = pingxx_log.New(cfg)
            PingxxLog.SetModule("example")
            PingxxLog.Cache.CacheMonitor()//监控缓存 可开可不开 不开注释即可
        }
    