package id

import (
	"errors"

	"github.com/appleboy/gorush/notify"
)

func GetTraceId(notifications []notify.PushNotification) (interface{}, error) {
	if len(notifications) > 0 {
		if data, ok := notifications[0].Data["body"]; ok {
			if dataMap, ok := data.(map[string]interface{}); ok {
				if traceId, ok := dataMap["traceId"]; ok {
					return traceId, nil
				} else {
					return nil, errors.New("traceId not found")
				}
			} else {
				return nil, errors.New("body is not a map")
			}
		} else {
			return nil, errors.New("body not found")
		}
	} else {
		return nil, errors.New("no notifications")
	}
}
