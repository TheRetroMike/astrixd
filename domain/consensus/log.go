package consensus

import (
	"github.com/astrix-network/astrixd/infrastructure/logger"
	"github.com/astrix-network/astrixd/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
