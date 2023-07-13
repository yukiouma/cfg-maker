package cfgreader

import (
	cfgreader "github.com/yukiouma/cfg-maker/internal/cfgmaker"
)

func New(data any) cfgreader.CfgMaker {
	return cfgreader.NewMaker(data)
}
