package apiserver

import (
	"github.com/juju/loggo"
	"github.com/samsung-cnct/cma-aws/pkg/util"
	"github.com/samsung-cnct/cma-aws/pkg/util/cluster"
)

var (
	logger loggo.Logger
)

type Server struct {
	ErrorMap *cluster.ErrorMap
}

func SetLogger() {
	logger = util.GetModuleLogger("internal.cluster-manager-api", loggo.INFO)
}
