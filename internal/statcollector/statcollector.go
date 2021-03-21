package statcollector

import (
	"fmt"
	"rlcache/internal/constants"
	"rlcache/internal/tsvreader"
	"strconv"
)

type StatisticsCollector struct {
	SizeMap map[string]float64
	FrequencyMap map[string]float64
	currentTime int64
	linesProcessed int
	linesSkipped int
}

func NewStatisticsCollector() *StatisticsCollector {
	return &StatisticsCollector{
		SizeMap: map[string]float64{},
		FrequencyMap: map[string]float64{},
		currentTime: 0,
		linesProcessed: 0,
		linesSkipped: 0,
	}
}

func (s *StatisticsCollector) UpdateLine(line tsvreader.TSVString) bool {
	timestamp, _ := strconv.ParseInt(line["timestamp"], 10, 64)
	size, _ := strconv.ParseFloat(line["size"], 10)
	data_id := line["data_id"]

	if timestamp < s.currentTime {
		return false
	}

	s.currentTime = timestamp

	if _, ok := s.SizeMap[data_id]; ok {
		s.SizeMap[data_id] += 1./size
		s.FrequencyMap[data_id] += 1.
	} else {
		s.SizeMap[data_id] = 1./size
		s.FrequencyMap[data_id] = 1.
	}

	return true
}

func (s *StatisticsCollector) UpdateFile(lines []tsvreader.TSVString) {
	for _, line := range lines {
		updateResult := s.UpdateLine(line)
		if updateResult {
			s.linesProcessed += 1
		} else {
			s.linesSkipped += 1
		}
		if s.linesProcessed % 10000 == 0 {
			fmt.Printf("\rProcessed %s skipped %s",
				constants.MakeBoldInt(s.linesProcessed), constants.MakeBoldInt(s.linesSkipped))
		}
	}
	fmt.Printf("\rProcessed %s skipped %s\n",
		constants.MakeBoldInt(s.linesProcessed), constants.MakeBoldInt(s.linesSkipped))
}