package cfgmaker

import (
	"github.com/yukiouma/cfg-maker/internal/cfgmaker"
)

func New(data any) cfgmaker.CfgMaker {
	return cfgmaker.NewMaker(data)
}
