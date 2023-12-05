package day5

import (
	"slices"
	"strconv"
	"strings"
)

type PosMap struct {
	dest   []int
	source []int
	r      []int
}

type Almanac struct {
	seeds             []int
	seedLoc           []int
	seedToSoil        PosMap
	soilToFertlizer   PosMap
	fertilizerToWater PosMap
	waterToLight      PosMap
	lightToTemp       PosMap
	tempToHumidity    PosMap
	humidityToLoc     PosMap
}

func Part1(input *string) string {
	data := strings.Split(*input, "\n\n")
	alamnac := &Almanac{}
	alamnac.seeds = parseSeeds(data[0])
	alamnac.seedToSoil = parseMap(data[1])
	alamnac.soilToFertlizer = parseMap(data[2])
	alamnac.fertilizerToWater = parseMap(data[3])
	alamnac.waterToLight = parseMap(data[4])
	alamnac.lightToTemp = parseMap(data[5])
	alamnac.tempToHumidity = parseMap(data[6])
	alamnac.humidityToLoc = parseMap(data[7])
	alamnac.SeedsToLocs()
	min := slices.Min(alamnac.seedLoc)
	return strconv.Itoa(min)
}

func Part2(input *string) string {
	data := strings.Split(*input, "\n\n")
	alamnac := &Almanac{}
	alamnac.seeds = parseSeeds(data[0])
	alamnac.seedToSoil = parseMap(data[1])
	alamnac.soilToFertlizer = parseMap(data[2])
	alamnac.fertilizerToWater = parseMap(data[3])
	alamnac.waterToLight = parseMap(data[4])
	alamnac.lightToTemp = parseMap(data[5])
	alamnac.tempToHumidity = parseMap(data[6])
	alamnac.humidityToLoc = parseMap(data[7])
	alamnac.SeedRangeToLocs()
	min := slices.Min(alamnac.seedLoc)
	return strconv.Itoa(min)
}

func parseSeeds(seeds string) []int {
	seeds = strings.Split(seeds, ": ")[1]
	seeds_split := strings.Split(seeds, " ")
	var seedSlice []int
	for _, seed := range seeds_split {
		s, _ := strconv.Atoi(seed)
		seedSlice = append(seedSlice, s)
	}
	return seedSlice
}

func parseMap(m string) PosMap {
	lines := strings.Split(m, "\n")
	pm := &PosMap{}
	for i := 1; i < len(lines); i++ {
		vals := strings.Split(lines[i], " ")
		dest, _ := strconv.Atoi(vals[0])
		source, _ := strconv.Atoi(vals[1])
		r, _ := strconv.Atoi(vals[2])
		pm.dest = append(pm.dest, dest)
		pm.source = append(pm.source, source)
		pm.r = append(pm.r, r)
	}
	return *pm
}

func (a *Almanac) SeedsToLocs() {
	for _, seed := range a.seeds {
		a.seedLoc = append(a.seedLoc, a.seedToLoc(seed))
	}
}

func (a *Almanac) seedToLoc(c int) int {
	c = a.seedToSoil.NextPos(c)
	c = a.soilToFertlizer.NextPos(c)
	c = a.fertilizerToWater.NextPos(c)
	c = a.waterToLight.NextPos(c)
	c = a.lightToTemp.NextPos(c)
	c = a.tempToHumidity.NextPos(c)
	c = a.humidityToLoc.NextPos(c)
	return c
}

func (p *PosMap) NextPos(c int) int {
	for i, source := range p.source {
		r := p.r[i]
		dest := p.dest[i]
		if c >= source && c <= source+r {
			return dest + (c - source)
		}
	}
	return c
}

func (a *Almanac) SeedRangeToLocs() {
	for i := 0; i < len(a.seeds); i = i + 2 {
		seed := a.seeds[i]
		totalSeeds := a.seeds[i+1]
		minLoc := int(^uint(0) >> 1)
		for j := seed; j < seed+totalSeeds; j++ {
			loc := a.seedToLoc(j)
			if loc < minLoc {
				minLoc = loc
			}
		}
		a.seedLoc = append(a.seedLoc, minLoc)
	}
}
