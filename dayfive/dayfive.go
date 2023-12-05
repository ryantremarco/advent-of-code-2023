package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

const newline = "\n"

func main() {
	parsedInput := parseInput(input)
	fmt.Println("Part One:", partOne(parsedInput))
	fmt.Println("Part Two:", partTwo(parsedInput))
}

type MappingRange struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

type Mapping []MappingRange

func (m Mapping) Get(i int) int {
	v := -1
	found := false

	for _, r := range m {
		if i >= r.SourceStart && i < r.SourceStart+r.Length {
			v = i - r.SourceStart + r.DestinationStart
			found = true
			break
		}
	}

	if !found {
		return i
	}

	return v
}

type SeedRange struct {
	Start  int
	Length int
}
type SeedIterator struct {
	Incrementor int
	Ranges      []SeedRange
	Current     int
}

func (s *SeedIterator) Next() bool {
	totalSeeds := 0
	for _, seedRange := range s.Ranges {
		if s.Incrementor < totalSeeds+seedRange.Length {
			s.Current = seedRange.Start + (s.Incrementor - totalSeeds)
			s.Incrementor++
			return true
		}
		totalSeeds += seedRange.Length
	}

	return false
}

type PlantingIterator struct {
	Incrementor  int
	SeedIterator SeedIterator
	Almanac      Almanac
	Current      Planting
}

func (p *PlantingIterator) Next() bool {

	if p.SeedIterator.Next() {
		seed := p.SeedIterator.Current
		soil := p.Almanac.seedToSoil.Get(seed)
		fertilizer := p.Almanac.soilToFertilizer.Get(soil)
		water := p.Almanac.fertilizerToWater.Get(fertilizer)
		light := p.Almanac.waterToLight.Get(water)
		temperature := p.Almanac.lightToTemperature.Get(light)
		humidity := p.Almanac.temperatureToHumidity.Get(temperature)
		location := p.Almanac.humidityToLocation.Get(humidity)

		p.Current = Planting{
			Seed:        seed,
			Soil:        soil,
			Fertilizer:  fertilizer,
			Water:       water,
			Light:       light,
			Temperature: temperature,
			Humidity:    humidity,
			Location:    location,
		}

		return true
	}

	return false
}

type Almanac struct {
	seeds                 []int
	seedToSoil            Mapping
	soilToFertilizer      Mapping
	fertilizerToWater     Mapping
	waterToLight          Mapping
	lightToTemperature    Mapping
	temperatureToHumidity Mapping
	humidityToLocation    Mapping
}

func (a Almanac) ToPlantingIterator(seedRanges []SeedRange) PlantingIterator {
	seedIterator := SeedIterator{Ranges: seedRanges}
	return PlantingIterator{SeedIterator: seedIterator, Almanac: a}
}

type Planting struct {
	Seed        int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

func parseInput(input string) Almanac {
	sections := strings.Split(input, newline+newline)

	var seeds []int
	seedStrs := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	for _, seedStr := range seedStrs {
		seeds = append(seeds, toInt(seedStr))
	}

	parseMapping := func(mappings string) Mapping {
		m := Mapping{}
		mappingLines := strings.Split(mappings, newline)
		for _, mappingLine := range mappingLines[1:] {
			if mappingLine == "" {
				continue
			}

			intStrs := strings.Split(mappingLine, " ")
			destinationStart := toInt(intStrs[0])
			sourceStart := toInt(intStrs[1])
			length := toInt(intStrs[2])

			m = append(m, MappingRange{
				SourceStart:      sourceStart,
				DestinationStart: destinationStart,
				Length:           length,
			})
		}

		return m
	}

	return Almanac{
		seeds:                 seeds,
		seedToSoil:            parseMapping(sections[1]),
		soilToFertilizer:      parseMapping(sections[2]),
		fertilizerToWater:     parseMapping(sections[3]),
		waterToLight:          parseMapping(sections[4]),
		lightToTemperature:    parseMapping(sections[5]),
		temperatureToHumidity: parseMapping(sections[6]),
		humidityToLocation:    parseMapping(sections[7]),
	}
}

func toInt[T ~string | ~rune | ~byte](t T) int {
	i, err := strconv.Atoi(string(t))
	if err != nil {
		panic(err)
	}
	return i
}
func contains[T comparable](haystack []T, needle T) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}

	return false
}

func partOne(almanac Almanac) int {
	var lowestLoc *int

	var seedRanges []SeedRange
	for _, seed := range almanac.seeds {
		seedRanges = append(seedRanges, SeedRange{
			Start:  seed,
			Length: 1,
		})
	}

	plantings := almanac.ToPlantingIterator(seedRanges)
	for plantings.Next() {
		planting := plantings.Current
		if lowestLoc == nil || planting.Location < *lowestLoc {
			lowestLoc = &planting.Location
		}
	}

	if lowestLoc == nil {
		return -1
	}

	return *lowestLoc
}

// This is a very slow solution to the puzzle.
// I had just gone through the seed and planting arrays one by one previously and just updated that to use iterators instead to save on (coding) time.
func partTwo(almanac Almanac) int {
	var lowestLoc *int

	var seedRanges []SeedRange
	for i := 0; i < len(almanac.seeds); i += 2 {
		seed := almanac.seeds[i]
		length := almanac.seeds[i+1]
		seedRanges = append(seedRanges, SeedRange{
			Start:  seed,
			Length: length,
		})
	}

	plantings := almanac.ToPlantingIterator(seedRanges)
	for plantings.Next() {
		planting := plantings.Current
		if lowestLoc == nil || planting.Location < *lowestLoc {
			lowestLoc = &planting.Location
		}
	}

	if lowestLoc == nil {
		return -1
	}

	return *lowestLoc
}
