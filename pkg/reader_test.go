package utils

import (
	"strings"
	"testing"

	"github.com/franela/goblin"
)

func TestReader(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#NewSampleReader", func() {
		g.It("Should process error if wav not found", func() {
			_, err := NewSamplesReader("./dummy.wav")
			g.Assert(err == nil).IsFalse()
			g.Assert(strings.HasPrefix(err.Error(), "Unable to open ./dummy.wav: ")).IsTrue()
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

		g.It("Should correct open 16 bit iq file", func() {
			r, err := NewSamplesReader("./data/110B_8k_16c.wav")
			g.Assert(err == nil).IsTrue()
			g.Assert(r == nil).IsFalse()
			g.Assert(r.reader == nil).IsFalse()
			g.Assert(int(r.SampleRate)).Equal(8000)
			g.Assert(r.SampleType).Equal("16s")
		})

		g.It("Should correct open 32 bit iq file", func() {
			r, err := NewSamplesReader("./data/110B_8k_32c.wav")
			g.Assert(err == nil).IsTrue()
			g.Assert(r == nil).IsFalse()
			g.Assert(r.reader == nil).IsFalse()
			g.Assert(int(r.SampleRate)).Equal(8000)
			g.Assert(r.SampleType).Equal("32f")
		})
	})

	g.Describe("#from16s", func() {
		g.It("Should make 32c sample from 16s pairs in byte array", func() {
			//                 0r    0r    0i    0i    1r    1r    1i    1i    2r    2r    2i    2i    3r    3r    3i    3i
			bytes := [...]byte{0xB5, 0x04, 0x74, 0xF5, 0xEC, 0x0E, 0xCA, 0xF1, 0xD6, 0x06, 0x0E, 0xF4, 0xD9, 0xF3, 0xDE, 0xFA}
			//                      0r    0i    1r     1i     2r    2i     3r     3i
			samples := [...]float32{1205, -2700, 3820, -3638, 1750, -3058, -3111, -1314}
			for i := 0; i < 4; i++ {
				re, im := from16s(bytes[:], i*4)
				g.Assert(re).Equal(samples[i*2+0])
				g.Assert(im).Equal(samples[i*2+1])
			}
		})
	})

	g.Describe("#from32f", func() {
		g.It("Should make 32c sample from 32f pairs in byte array", func() {
			bytes := [...]byte{
				//0r                    0i                      1r                      1i
				0xC0, 0xAD, 0x16, 0x3D, 0x80, 0xBE, 0xA8, 0xBD, 0x80, 0xC3, 0xEE, 0x3D, 0x80, 0x60, 0xE3, 0xBD,
				//2r                    2i                      3r                      3i
				0x80, 0xBF, 0x5A, 0x3D, 0x80, 0x1D, 0xBF, 0xBD, 0x00, 0x78, 0xC2, 0xBD, 0x00, 0x3D, 0x24, 0xBD,
			}
			//                      0r           0i          1r           1i          2r           2i            3r            3i
			samples := [...]float32{0.036786795, -0.0823946, 0.116583824, -0.1110239, 0.053405285, -0.093317986, -0.094955444, -0.040097237}
			// samples := [...]float32{1024896448, -1113014656, 1039057792, -1109172096, 1029357440, -1111548544, -1111328768, -1121698560}
			for i := 0; i < 4; i++ {
				re, im := from32f(bytes[:], i*8)
				g.Assert(re).Equal(samples[i*2+0])
				g.Assert(im).Equal(samples[i*2+1])
			}
		})
	})

	g.Describe("#Read", func() {
		g.It("Should check correct reader pointer")
		g.It("Should check too short file")
		g.It("Should correct read 16s file")
		g.It("Should correct read 32f file")
	})
}
