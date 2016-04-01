## pingxx log module

- Usage:

        cfg := pingxx_log.NewConfig("pingxx.log", pingxx_log.ToConsole | pingxx_log.ToFile)
        cfg.SetCententType(pingxx_log.ToFile, pingxx_log.JSON)
        cfg.SetCententType(pingxx_log.ToConsole, pingxx_log.STDOUTPUT)
        cfg.SetLevel(pingxx_log.Debug)
        Glogger := pingxx_log.New(cfg)
        Glogger.SetModule("pingxx_ar")
        
        如果要设置acct_id等信息,可以通过下面方式设定
        info:=pingxx_log.NewLogInfo().SetAcctId("acct_id1111").SetAppId("app_id222").SetChannel("alipay").SetMode(1).SetAgent("xxx").SetRefer("xxx").SetUrl("xxx")
        Glogger.SetLogInfo(info)
        
        如果要设置request
        info.SetRequestInfo(req *http.Request)
        Glogger.SetLogInfo(info)
        如果要设置response
        info.SetResponseInfo(response http.Response)
        Glogger.SetLogInfo(info)
        
        最后通过Debug,Info,Error,Warn等方式写入文件或stdout
        Glogger.Debug("aaa%d",11111) // Glogger.Info("")
    