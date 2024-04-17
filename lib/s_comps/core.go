package s_comps

import (
	"github.com/zmloong/lotus/core"
	"github.com/zmloong/lotus/sys/proto"
)

type (
	ISC_GateRouteComp interface {
		core.IServiceComp
		RegisterRoute(comId uint16, f func(s core.IUserSession, msg proto.IMessage) (code core.ErrorCode, err string)) (err error)
	}
)

func NewGateRouteComp() ISC_GateRouteComp {
	comp := new(SComp_GateRouteComp)
	return comp
}
