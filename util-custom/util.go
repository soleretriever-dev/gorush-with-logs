package custom

import (
	"encoding/json"
	"time"
)

type PerformanceData struct {
	Name        string  `json:"name"`
	PipelineId  string  `json:"pipelineId"`
	TraceId     string  `json:"traceId"`
	TimeElapsed float64 `json:"timeElapsed"`
}

func CalculatePerformanceData(pipelineId string, traceId string, name string, timeStart time.Time, timeEnd time.Time) []byte {

	var timeElapsed = timeEnd.Sub(timeStart).Seconds()

	var data = PerformanceData{
		Name:        name,
		PipelineId:  pipelineId,
		TraceId:     traceId,
		TimeElapsed: timeElapsed,
	}

	jsonData, _ := json.Marshal(data)

	return jsonData

}
