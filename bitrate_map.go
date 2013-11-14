package main

type BitrateMap struct {
	m map[int]int
}

func NewBitrateMap() BitrateMap {
	b := BitrateMap{
		m: make(map[int]int),
	}

	return b
}

func (b *BitrateMap) Add(ver MP3Version, lay MP3Layer, bits byte, bitrate int) {

	i := 0

	i += int(ver) << 6
	i += int(lay) << 4
	i += int(bits)

	b.m[i] = bitrate
}

func (b BitrateMap) Get(ver MP3Version, lay MP3Layer, bits byte) int {
	i := 0

	i += int(ver) << 6
	i += int(lay) << 4
	i += int(bits)

	if bitrate, exists := b.m[i]; exists {
		return bitrate
	} else {
		return 0
	}
}

func init() {
	bitrates = NewBitrateMap()

	bitrates.Add(Version1, LayerI, 1, 32)
	bitrates.Add(Version1, LayerI, 2, 64)
	bitrates.Add(Version1, LayerI, 3, 96)
	bitrates.Add(Version1, LayerI, 4, 128)
	bitrates.Add(Version1, LayerI, 5, 160)
	bitrates.Add(Version1, LayerI, 6, 192)
	bitrates.Add(Version1, LayerI, 7, 224)
	bitrates.Add(Version1, LayerI, 8, 256)
	bitrates.Add(Version1, LayerI, 9, 288)
	bitrates.Add(Version1, LayerI, 10, 320)
	bitrates.Add(Version1, LayerI, 11, 352)
	bitrates.Add(Version1, LayerI, 12, 384)
	bitrates.Add(Version1, LayerI, 13, 416)
	bitrates.Add(Version1, LayerI, 14, 448)

	bitrates.Add(Version1, LayerII, 1, 32)
	bitrates.Add(Version1, LayerII, 2, 48)
	bitrates.Add(Version1, LayerII, 3, 56)
	bitrates.Add(Version1, LayerII, 4, 64)
	bitrates.Add(Version1, LayerII, 5, 80)
	bitrates.Add(Version1, LayerII, 6, 96)
	bitrates.Add(Version1, LayerII, 7, 112)
	bitrates.Add(Version1, LayerII, 8, 128)
	bitrates.Add(Version1, LayerII, 9, 160)
	bitrates.Add(Version1, LayerII, 10, 192)
	bitrates.Add(Version1, LayerII, 11, 224)
	bitrates.Add(Version1, LayerII, 12, 256)
	bitrates.Add(Version1, LayerII, 13, 320)
	bitrates.Add(Version1, LayerII, 14, 384)

	bitrates.Add(Version1, LayerIII, 1, 32)
	bitrates.Add(Version1, LayerIII, 2, 40)
	bitrates.Add(Version1, LayerIII, 3, 48)
	bitrates.Add(Version1, LayerIII, 4, 56)
	bitrates.Add(Version1, LayerIII, 5, 64)
	bitrates.Add(Version1, LayerIII, 6, 80)
	bitrates.Add(Version1, LayerIII, 7, 96)
	bitrates.Add(Version1, LayerIII, 8, 112)
	bitrates.Add(Version1, LayerIII, 9, 128)
	bitrates.Add(Version1, LayerIII, 10, 160)
	bitrates.Add(Version1, LayerIII, 11, 192)
	bitrates.Add(Version1, LayerIII, 12, 224)
	bitrates.Add(Version1, LayerIII, 13, 256)
	bitrates.Add(Version1, LayerIII, 14, 320)

	bitrates.Add(Version2, LayerI, 1, 32)
	bitrates.Add(Version2, LayerI, 2, 48)
	bitrates.Add(Version2, LayerI, 3, 56)
	bitrates.Add(Version2, LayerI, 4, 64)
	bitrates.Add(Version2, LayerI, 5, 80)
	bitrates.Add(Version2, LayerI, 6, 96)
	bitrates.Add(Version2, LayerI, 7, 112)
	bitrates.Add(Version2, LayerI, 8, 128)
	bitrates.Add(Version2, LayerI, 9, 144)
	bitrates.Add(Version2, LayerI, 10, 160)
	bitrates.Add(Version2, LayerI, 11, 176)
	bitrates.Add(Version2, LayerI, 12, 192)
	bitrates.Add(Version2, LayerI, 13, 224)
	bitrates.Add(Version2, LayerI, 14, 256)

	bitrates.Add(Version2, LayerII, 1, 8)
	bitrates.Add(Version2, LayerII, 2, 16)
	bitrates.Add(Version2, LayerII, 3, 24)
	bitrates.Add(Version2, LayerII, 4, 32)
	bitrates.Add(Version2, LayerII, 5, 40)
	bitrates.Add(Version2, LayerII, 6, 48)
	bitrates.Add(Version2, LayerII, 7, 56)
	bitrates.Add(Version2, LayerII, 8, 64)
	bitrates.Add(Version2, LayerII, 9, 80)
	bitrates.Add(Version2, LayerII, 10, 96)
	bitrates.Add(Version2, LayerII, 11, 112)
	bitrates.Add(Version2, LayerII, 12, 128)
	bitrates.Add(Version2, LayerII, 13, 144)
	bitrates.Add(Version2, LayerII, 14, 160)

	bitrates.Add(Version2, LayerIII, 1, 8)
	bitrates.Add(Version2, LayerIII, 2, 16)
	bitrates.Add(Version2, LayerIII, 3, 24)
	bitrates.Add(Version2, LayerIII, 4, 32)
	bitrates.Add(Version2, LayerIII, 5, 40)
	bitrates.Add(Version2, LayerIII, 6, 48)
	bitrates.Add(Version2, LayerIII, 7, 56)
	bitrates.Add(Version2, LayerIII, 8, 64)
	bitrates.Add(Version2, LayerIII, 9, 80)
	bitrates.Add(Version2, LayerIII, 10, 96)
	bitrates.Add(Version2, LayerIII, 11, 112)
	bitrates.Add(Version2, LayerIII, 12, 128)
	bitrates.Add(Version2, LayerIII, 13, 144)
	bitrates.Add(Version2, LayerIII, 14, 160)
}
