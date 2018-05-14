package main

import (
	"fmt"
	"log"
	"os"

	"github.com/o-kos/cls-micro/pkg"
	"github.com/youpy/go-wav"
)

type sampleMaker func([]byte, int) (float32, float32)

func from16s(data []byte, offset int) (r, i float32) {
	r = float32(data[offset+0])
	i = float32(data[offset+1])
	return
}

func from32s(data []byte, offset int) (r, i float32) {
	r = float32(data[offset+0])
	i = float32(data[offset+1])
	// soffset := offset + (j * bitsPerSample / 8)
	// bits := uint32(
	// 	(int(bytes[soffset+3]) << 24) +
	// 	(int(bytes[soffset+2]) << 16) +
	// 	(int(bytes[soffset+1]) <<  8) +
	// 	(int(bytes[soffset+0]) <<  0)
	// )
	// samples[i * 2 + 0] = math.Float32frombits(bits)
	// samples[i * 2 + 1] = math.Float32frombits(bits)
	return
}

func main() {
	cfg, err := config.New("cls")
	if err != nil {
		log.Panic("Unable to parse config file", err)
		return
	}

	fmt.Printf("Processing file %s ", cfg.FileName)
	file, err := os.Open(cfg.FileName)
	defer file.Close()
	if err != nil {
		log.Panic("Unable to open file", err)
		return
	}
	reader := wav.NewReader(file)
	format, err := reader.Format()
	if err != nil {
		log.Panic("Unable to read wav header", err)
		return
	}
	if format.NumChannels != 2 {
		log.Fatal("Unable to process non-iq files")
		return
	}
	af, ok := map[uint16]string{wav.AudioFormatPCM: "i", wav.AudioFormatIEEEFloat: "f"}[format.AudioFormat]
	if !ok {
		log.Fatal("Unsupported format type", format.AudioFormat)
		return
	}
	fmt.Printf("(%d kHz, %d%s)\n", format.SampleRate/1000, format.BitsPerSample, af)

	var ms uint16
	for _, val := range cfg.Data.Duration {
		if val > ms {
			ms = val
		}
	}
	if ms == 0 {
		log.Fatal("Data durations is too small")
		return
	}

	count := int(ms) * int(format.SampleRate) / 1000
	fmt.Printf("Try to read %d samples... ", count)
	bytes := make([]byte, count*int(format.BlockAlign))
	n, err := reader.Read(bytes)
	if err != nil {
		log.Panic("Unable to read source file samples", err)
		return
	}
	if n < len(bytes) {
		log.Fatalf("Source file is too short (%d bytes from %d)", n, len(bytes))
		return
	}

	samples := make([]float32, count*2)
	var fn sampleMaker
	if format.AudioFormat == wav.AudioFormatIEEEFloat {
		fn = from16s
	} else {
		fn = from32s
	}
	for i, offset := 0, 0; i < count; i++ {
		samples[i*2], samples[i*2+1] = fn(bytes, offset)
		offset += int(format.BlockAlign)
	}
	fmt.Println("Ok")
}
