package utils

import (
	"testing"

	"github.com/franela/goblin"
)

func TestNewSampleReader(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#NewSampleReader", func() {
		g.It("Should process error if wav not found", func() {
			_, err := NewSamplesReader("./dummy.wav")
			g.Assert(err == nil).IsFalse()
			g.Assert(err.Error()).Equal("Unable to open ./dummy.wav: The system cannot find the file specified.")
		})

		g.It("Should process error if wav file has bad header", func() {
			_, err := NewSamplesReader("./reader_test.go")
			g.Assert(err == nil).IsFalse()
			g.Assert(err.Error()).Equal("Unable to read wav header: Given bytes is not a RIFF format")
		})

		g.It("Should process error if wav has non-iq format", func() {
			_, err := NewSamplesReader("./data/110B_8k_16s.wav")
			g.Assert(err == nil).IsFalse()
			g.Assert(err.Error()).Equal("Unable to process non-iq file")
		})

		g.It("Should process error if wav has unsupported format type", func() {
			_, err := NewSamplesReader("./data/110B_8k_24c.wav")
			g.Assert(err == nil).IsFalse()
			g.Assert(err.Error()).Equal("Unsupported format type: 2")
		})

		g.It("Should correct read 16 bit iq file", func() {
			r, err := NewSamplesReader("./data/110B_8k_16c.wav")
			g.Assert(err == nil).IsTrue()
			g.Assert(r == nil).IsFalse()
			g.Assert(r.reader == nil).IsFalse()
			g.Assert(int(r.SampleRate)).Equal(8000)
			g.Assert(r.SampleType).Equal("16s")
		})

		g.It("Should correct read 32 bit iq file", func() {
			r, err := NewSamplesReader("./data/110B_8k_32c.wav")
			g.Assert(err == nil).IsTrue()
			g.Assert(r == nil).IsFalse()
			g.Assert(r.reader == nil).IsFalse()
			g.Assert(int(r.SampleRate)).Equal(8000)
			g.Assert(r.SampleType).Equal("32f")
		})
	})
}
