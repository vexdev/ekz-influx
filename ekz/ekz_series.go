package ekz

import (
	"encoding/json"
	"slices"
)

type EkzData struct {
	Series       EkzSeries `json:"series"`
	SeriesHt     EkzSeries `json:"seriesHt"`
	SeriesNt     EkzSeries `json:"seriesNt"`
	SeriesNetz   EkzSeries `json:"seriesNetz"`
	SeriesNetzHt EkzSeries `json:"seriesNetzHt"`
	SeriesNetzNt EkzSeries `json:"seriesNetzNt"`
}

type EkzSeries struct {
	Level      string            `json:"level"`
	EnergyType string            `json:"energyType"`
	SourceType string            `json:"sourceType"`
	TariffType string            `json:"tariffType"`
	Ab         string            `json:"ab"`
	Bis        string            `json:"bis"`
	Values     []EkzSeriesValues `json:"values"`
}

type EkzSeriesValues struct {
	Value     float64 `json:"value"`
	Timestamp int     `json:"timestamp"`
	Date      string  `json:"date"`
	Time      string  `json:"time"`
	Status    string  `json:"status"`
}

func EkzDataFromJson(data []byte) (EkzData, error) {
	var ekzData EkzData
	err := json.Unmarshal(data, &ekzData)
	if err != nil {
		return EkzData{}, err
	}
	return ekzData, nil
}

func (e *EkzData) GetValidValuesSorted() []EkzSeriesValues {
	var sortedSeries []EkzSeriesValues
	for _, series := range []EkzSeries{e.Series, e.SeriesHt, e.SeriesNt, e.SeriesNetz, e.SeriesNetzHt, e.SeriesNetzNt} {
		for _, value := range series.Values {
			if value.Status != "VALID" {
				continue
			}
			sortedSeries = append(sortedSeries, value)
		}
	}
	slices.SortFunc(sortedSeries, func(a, b EkzSeriesValues) int {
		return a.Timestamp - b.Timestamp
	})
	return sortedSeries
}
