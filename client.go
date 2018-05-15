package main

import (
	"fmt"
	"log"
	"os"

	"github.com/o-kos/cls-micro/pkg"
	"github.com/youpy/go-wav"
)

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
	wr := wav.NewReader(file)
	format, err := wr.Format()
	if err != nil {
		log.Panic("Unable to read wav header", err)
		return
	}
	fmt.Printf("(%d kHz, %d%s)\n", format.SampleRate/1000, format.BitsPerSample, af)
	fmt.Printf("Try to read %d samples... ", count := int(ms) * int(format.SampleRate) / 1000)
	samples := reader.NewSamples(&wr, cfg.Data.MaxDuration())

	// if format.NumChannels != 2 {
	// 	log.Fatal("Unable to process non-iq files")
	// 	return
	// }
	// af, ok := map[uint16]string{wav.AudioFormatPCM: "i", wav.AudioFormatIEEEFloat: "f"}[format.AudioFormat]
	// if !ok {
	// 	log.Fatal("Unsupported format type", format.AudioFormat)
	// 	return
	// }
	// fmt.Printf("(%d kHz, %d%s)\n", format.SampleRate/1000, format.BitsPerSample, af)

	// count := int(ms) * int(format.SampleRate) / 1000
	// fmt.Printf("Try to read %d samples... ", count)
	// bytes := make([]byte, count*int(format.BlockAlign))
	// n, err := reader.Read(bytes)
	// if err != nil {
	// 	log.Panic("Unable to read source file samples", err)
	// 	return
	// }
	// if n < len(bytes) {
	// 	log.Fatalf("Source file is too short (%d bytes from %d)", n, len(bytes))
	// 	return
	// }

	// samples := make([]float32, count*2)
	// var fn sampleMaker
	// if format.AudioFormat == wav.AudioFormatIEEEFloat {
	// 	fn = from16s
	// } else {
	// 	fn = from32s
	// }
	// for i, offset := 0, 0; i < count; i++ {
	// 	samples[i*2], samples[i*2+1] = fn(bytes, offset)
	// 	offset += int(format.BlockAlign)
	// }
	fmt.Println("Ok")
}
