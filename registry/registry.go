package registry

import (
	"github.com/CubicCommunity/BloqGo/include"
)

var Commands []*include.Command

func Register(cmd *include.Command) {
	Commands = append(Commands, cmd)
}
