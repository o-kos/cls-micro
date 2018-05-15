package utils

import (
	"fmt"
	"log"
	"os"

	wav "github.com/youpy/go-wav"
)

// WavReader - wrapper for reading signal samples from wav file
type WavReader struct {
	reader wav.Reader
	SampleRate uint32
	SampleType string
	Samples map[uint16][]float32
}

// NewWavReader - constructor for WavReader
func NewWavReader(fileName string, []durations int) (*WavReader, error) {
	file, err := os.Open(FileName)
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
}

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

// ReadSamples read count pcm samples from wav file by fileName
func ReadSamples(reader *wav.Reader, count int) ([]float32, error) {
	file, err := os.Open(cfg.FileName)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %v", err)
	}
	reader := wav.NewReader(file)
	format, err := reader.Format()
	if err != nil {
		return nil, fmt.Errorf("Unable to read wav header", err)
	}
	if format.NumChannels != 2 {
		return nil, fmt.Errorf("Unable to process non-iq files")
	}
	af, ok := map[uint16]string{wav.AudioFormatPCM: "i", wav.AudioFormatIEEEFloat: "f"}[format.AudioFormat]
	if !ok {
		return nil, fmt.Errorf("Unsupported format type", format.AudioFormat)
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
	return &cfg, nil
}
