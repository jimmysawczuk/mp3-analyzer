package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"unicode/utf8"
)

type MP3 struct {
	id3_data map[string]string

	Title  string `json:"title"`
	Artist string `json:"artist"`

	Size     int64         `json:"size"`
	Duration time.Duration `json:"duration"`
	Bitrate  float64       `json:"bitrate"`
}

type MP3Version int16

const (
	// Version25 MP3Version = 0
	Version2 = 2
	Version1 = 3
)

type MP3Layer int16

const (
	LayerIII MP3Layer = 1
	LayerII           = 2
	LayerI            = 3
)

var bitrates BitrateMap
var sampling_rates SampleMap

type FrameHeader struct {
	raw_string string
	valid      bool

	Version MP3Version
	Layer   MP3Layer

	Bitrate      int
	SamplingRate int

	FrameSize int
}

func NewMP3() MP3 {
	i := MP3{}
	i.id3_data = make(map[string]string)

	return i
}

func (i *MP3) AddID3(tag string, value string) {
	i.id3_data[tag] = value

	switch tag {
	case "TIT2":
		i.Title = value
	case "TPE1":
		i.Artist = value
	}
}

func main() {
	mp3 := NewMP3()

	id3_seen := false

	fp := os.Stdin
	buf := bytes.NewBuffer([]byte{})

	n, err := buf.ReadFrom(fp)
	if err != nil {
		panic(err)
	}
	mp3.Size = n

	by := buf.Bytes()
	i := 0
	var seconds float64
	var bitrate float64

	for i < len(by) {
		if hdr := parseHeader(by[i : i+4]); hdr.valid {
			idx := i + hdr.FrameSize
			next_hdr := parseHeader(by[idx : idx+4])
			if next_hdr.valid {
				// fmt.Printf("%#v %#v\n", hdr, next_hdr)
				if !id3_seen && bytes.NewBuffer(by[0:3]).String() == "ID3" {
					parseID3(bytes.NewBuffer(by[0:i]), &mp3)
					id3_seen = true
				}

				seconds += float64(hdr.FrameSize*8.0) / float64(1000.0*hdr.Bitrate)
				bitrate += float64(hdr.FrameSize*8.0) / float64(1000.0*hdr.Bitrate) * float64(hdr.Bitrate)

				i = i + hdr.FrameSize
			} else {
				i++
			}
		} else {
			i++
		}
	}

	mp3.Bitrate = bitrate / seconds
	mp3.Duration, _ = time.ParseDuration(fmt.Sprintf("%.3fs", seconds))

	// fmt.Println(mp3)

	json_data, err := json.MarshalIndent(mp3, "", "   ")
	fmt.Println(bytes.NewBuffer(json_data).String())
}

func parseHeader(by []byte) (fh FrameHeader) {

	fh.raw_string = fmt.Sprintf("%08b %08b %08b %08b", by[0], by[1], by[2], by[3])

	fh.valid = by[0] == byte(255) && by[1]>>5 == byte(7)
	if !fh.valid {
		return
	}

	num := make([]int, len(by))
	for i := 0; i < len(by); i++ {
		num[i] = int(by[i])
	}

	version := num[1] >> 3 % 0x4
	layer := num[1] >> 1 % 0x4

	fh.Version = MP3Version(version)
	fh.Layer = MP3Layer(layer)

	bitrate_mask := num[2] >> 4
	bitrate := bitrates.Get(MP3Version(version), MP3Layer(layer), byte(bitrate_mask))
	fh.Bitrate = bitrate

	sample_rate := (num[2] % 0x10) >> 2
	fh.SamplingRate = sampling_rates.Get(sample_rate, MP3Version(version))

	// if we have valid bitrates and sampling rates, we're valid
	fh.valid = fh.valid && fh.Bitrate > 0 && fh.SamplingRate > 0

	// the last two bits of the header cannot be 10
	fh.valid = fh.valid && num[3]%4 != 3

	padding := num[2] >> 1 % 2

	if fh.valid {
		if fh.Layer == LayerI {
			fh.FrameSize = (12*(fh.Bitrate*1000)/(fh.SamplingRate) + padding) * 4
		} else {
			fh.FrameSize = 144*(fh.Bitrate*1000)/(fh.SamplingRate) + padding
		}
	}

	return
}

func parseID3(data *bytes.Buffer, mp3 *MP3) {

	h := data.Next(10)
	_ = h

	for data.Len() > 0 {
		frame_header := data.Next(10)

		if len(frame_header) == 10 {
			tag := frame_header[0:4]
			size := int(frame_header[4])<<24 + int(frame_header[5])<<16 + int(frame_header[6])<<8 + int(frame_header[7])
			flags := frame_header[8:10]

			_ = flags

			tag_data := data.Next(int(size))

			tag_name := bytes.NewBuffer(tag).String()
			tag_value := make([]rune, 0)

			for len(tag_data) > 0 {
				r, size := utf8.DecodeRune(tag_data)

				if int(r) != 0 {
					tag_value = append(tag_value, r)
				}

				tag_data = tag_data[size:]
			}

			mp3.AddID3(tag_name, string(tag_value))
		}
	}
}
