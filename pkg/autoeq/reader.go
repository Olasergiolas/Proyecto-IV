package autoeq

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	bandTypePos int = 3
	freqPos     int = 5
	gainPos     int = 8
	qualityPos  int = 11
)

func CreateBandMap(bands []band) map[string]band {
	m := make(map[string]band)

	for i := 0; i < len(bands); i++ {
		m["band"+strconv.Itoa(i)] = bands[i]
	}
	return m
}

func OpenParametricData(path string) []string {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	sc := bufio.NewScanner(data)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func GenerateBands(lines []string) []band {
	var b band
	var bands []band
	var freq, gain, quality float64

	for _, l := range lines[1:] {
		l_split := strings.Split(l, " ")
		freq, _ = strconv.ParseFloat(l_split[freqPos], 32)
		gain, _ = strconv.ParseFloat(l_split[gainPos], 32)
		quality, _ = strconv.ParseFloat(l_split[qualityPos], 32)

		b = NewBand(float32(freq), float32(gain), float32(quality), easyEffectsBandTypes[l_split[bandTypePos]])
		bands = append(bands, b)
	}
	return bands
}
