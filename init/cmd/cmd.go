package cmd

import (
	"crud-server/config"
	"crud-server/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}

	if err := c.network.ServerStart(c.config.Server.Port); err != nil {
		panic(err)
	}

	return c
}
