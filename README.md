## pingxx log module

- Usage:

        logger:=pingpp-log.New(pingpp-log.NewConfig(pingpp-log.ToConsole | pingpp-log.ToFile, pingxx.log))
        info:=pingpp-log.NewLogInfo().SetAcctId(xx).SetAppId(xx).SetChannel(xx).SetMode(x).SetAgent(x).SetRefer(xx).SetUrl()
        info.SetRequestInfo(request) 或者 info.SetResponseInfo(response)
        logger.SetModule("module_name").SetLogInfo(info).Debug("xxxxxxx")
    