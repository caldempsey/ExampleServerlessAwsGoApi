package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	formatBadRequest = "Bad Request. Value %q for %s is not allowed. Allowable values: %v"
	formatNotFound   = "Not Found. Could not find the requested %s: %q"
)

func convertTimeToRFC3339(hoursData map[string]interface{}) (rfc3339Data map[string]interface{}, err error) {
	rfc3339Data = make(map[string]interface{})
	var floatHours float64
	var intSeconds int64
	for k, v := range hoursData {

		k = strings.Replace(k, ",", ".", 1)

		floatHours, err = strconv.ParseFloat(k, 32)
		if err != nil {
			return
		}

		floatSeconds := floatHours * 3600
		stringSeconds := fmt.Sprintf("%.0f", floatSeconds)
		intSeconds, err = strconv.ParseInt(stringSeconds, 10, 64)
		if err != nil {
			return
		}

		rfc3339Data[time.Unix(intSeconds, 0).Format("2006-01-02T15:04:05.000Z")] = v
	}

	return
}
