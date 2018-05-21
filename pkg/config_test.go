package utils

import (
	"testing"

	"github.com/franela/goblin"
)

func TestNewCfg(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#NewCfg", func() {
		def := Config{
			Clients: Clients{
				Count: 2,
				Fiber: 4,
			},
			Data: Data{
				Durations:  []uint16{750, 1000, 3000},
				PacketSize: 100,
			},
		}

		g.It("Should create correct default values cfg", func() {
			res, err := NewCfg("config_test_empty")
			g.Assert(err == nil).IsTrue()
			g.Assert(*res).Equal(def)
		})

		g.It("Should load correct cfg from yaml file", func() {
			res, err := NewCfg("config_test_def")
			g.Assert(err == nil).IsTrue()
			g.Assert(*res).Equal(def)
		})
	})
}
