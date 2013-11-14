package main

type SampleMap struct {
	m map[int]int
}

func NewSampleMap() SampleMap {
	b := SampleMap{
		m: make(map[int]int),
	}

	return b
}

func (b *SampleMap) Add(bits int, ver MP3Version, rate int) {

	i := 0

	i += int(ver) << 2
	i += int(bits)

	b.m[i] = rate
}

func (b SampleMap) Get(bits int, ver MP3Version) int {
	i := 0

	i += int(ver) << 2
	i += int(bits)

	if sample_rate, exists := b.m[i]; exists {
		return sample_rate
	} else {
		return 0
	}
}

func init() {
	sampling_rates = NewSampleMap()

	sampling_rates.Add(0, Version1, 44100)
	sampling_rates.Add(1, Version1, 48000)
	sampling_rates.Add(2, Version1, 32000)

	sampling_rates.Add(0, Version2, 22050)
	sampling_rates.Add(1, Version2, 24000)
	sampling_rates.Add(2, Version2, 16000)
}
