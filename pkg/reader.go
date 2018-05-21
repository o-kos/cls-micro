package utils

import (
	"fmt"
	"os"

	wav "github.com/youpy/go-wav"
)

// SamplesReader - wrapper for reading signal samples from wav file
type SamplesReader struct {
	reader     *wav.Reader
	SampleRate uint32
	SampleType string
	Samples    []float32
}

// NewSamplesReader - constructor for WavReader
func NewSamplesReader(fileName string) (*SamplesReader, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("Unable to %v", err)
	}
	r := wav.NewReader(file)
	format, err := r.Format()
	if err != nil {
		return nil, fmt.Errorf("Unable to read wav header: %v", err)
	}
	if format.NumChannels != 2 {
		return nil, fmt.Errorf("Unable to process non-iq file")
	}
	af, ok := map[uint16]string{wav.AudioFormatPCM: "s", wav.AudioFormatIEEEFloat: "f"}[format.AudioFormat]
	if !ok {
		return nil, fmt.Errorf("Unsupported format type: %d", format.AudioFormat)
	}

	wr := SamplesReader{r, format.SampleRate, fmt.Sprintf("%d%s", format.BitsPerSample, af), nil}
	return &wr, nil
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

// Read read count pcm samples from wav file by fileName
func (r *SamplesReader) Read(count int) error {
	format, _ := r.reader.Format()
	bytes := make([]byte, count*int(format.BlockAlign))
	n, err := r.reader.Read(bytes)
	if err != nil {
		return fmt.Errorf("Unable to read source file samples %v", err)
	}
	if n < len(bytes) {
		return fmt.Errorf("Source file is too short (%d bytes from %d)", n, len(bytes))
	}

	r.Samples = make([]float32, count*2)
	var fn sampleMaker
	if format.AudioFormat == wav.AudioFormatIEEEFloat {
		fn = from16s
	} else {
		fn = from32s
	}
	for i, offset := 0, 0; i < count; i++ {
		r.Samples[i*2], r.Samples[i*2+1] = fn(bytes, offset)
		offset += int(format.BlockAlign)
	}
	return nil
}
