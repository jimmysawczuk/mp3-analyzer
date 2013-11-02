package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type MP3 struct {
	id3_data map[string]string

	Title  string
	Artist string

	Size   int64
	Length time.Duration
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

	fp := os.Stdin

	file_buffer := bytes.NewBuffer([]byte{})
	n, _ := file_buffer.ReadFrom(fp)

	header := bytes.NewBuffer([]byte{})
	data := bytes.NewBuffer([]byte{})

	id3_seen := false

	var seconds float64 = 0
	mp3 := NewMP3()
	mp3.Size = n

	for {
		a, err := file_buffer.ReadByte()
		if err != nil {
			break
		}

		data.WriteByte(a)

		if a == byte(255) && header.Len() == 0 {
			header.WriteByte(a)
		} else if a>>5 == byte(7) && header.Len() == 1 {
			header.WriteByte(a)
		} else if header.Len() >= 2 && header.Len() < 4 {
			header.WriteByte(a)
		} else if header.Len() > 0 {
			header.Reset()
		}

		if header.Len() == 4 {
			parseHeader(header)

			if !id3_seen {
				parseID3(data, &mp3)
				id3_seen = true
			} else {
				seconds += float64(data.Len()*8) / float64(1000.0*48.0)
			}

			header.Reset()
			data.Reset()
		}
	}

	mp3.Length, _ = time.ParseDuration(fmt.Sprintf("%.3fs", seconds))

	fmt.Println(mp3)
}

func parseHeader(header *bytes.Buffer) {
	// we'll need this function at some point if we want to programmatically determine the bitrate
}

func parseID3(data *bytes.Buffer, mp3 *MP3) {
	h := data.Next(10)
	_ = h

	for data.Len() > 0 {
		frame_header := data.Next(10)

		if len(frame_header) == 10 {
			tag := frame_header[0:4]
			size_buf := frame_header[4:8]
			size := uint32(0)
			binary.Read(bytes.NewBuffer(size_buf), binary.BigEndian, &size)
			tag_data := data.Next(int(size))

			mp3.AddID3(bytes.NewBuffer(tag).String(), bytes.NewBuffer(tag_data).String())
		}
	}
}
