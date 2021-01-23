package spore

import "github.com/l0k18/OSaaS/pkg/util/datadir"

type Shell struct {
	dataDir string
}

func New() *Shell {
	s := &Shell{
		dataDir: datadir.Get("spore", false),
	}
	return s
}