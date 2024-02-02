package custom

import (
	"encoding/json"
	"time"
)

type PerformanceData struct {
	Name        string  `json:"name"`
	TraceId     string  `json:"traceId"`
	TimeElapsed float64 `json:"timeElapsed"`
}

func CalculatePerformanceData(traceId string, name string, timeStart time.Time, timeEnd time.Time) []byte {

	var timeElapsed = timeEnd.Sub(timeStart).Seconds()

	var data = PerformanceData{
		Name:        name,
		TraceId:     traceId,
		TimeElapsed: timeElapsed,
	}

	jsonData, _ := json.Marshal(data)

	return jsonData

}
