package main

import (
	"github.com/astrix-network/astrixd/infrastructure/logger"
	"github.com/astrix-network/astrixd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
