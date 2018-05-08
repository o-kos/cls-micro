package config

import (
	"testing"

	"github.com/franela/goblin"
)

func TestNew(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#New", func() {
		def := Config{
			Clients: Clients{
				Count: 2,
				Fiber: 4,
			},
			Data: Data{
				Duration:   []int{750, 1000, 3000},
				PacketSize: 100,
			},
		}

		g.It("should create correct default values cfg", func() {
			res, err := New("config_test_empty")
			g.Assert(err == nil).IsTrue()
			g.Assert(*res).Equal(def)
		})

		g.It("should load correct cfg from yaml file", func() {
			res, err := New("config_test_def")
			g.Assert(err == nil).IsTrue()
			g.Assert(*res).Equal(def)
		})
	})
}
