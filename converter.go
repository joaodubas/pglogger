package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"
)

// Date layout in Go is based on a fixed date
// instead of using ambigous masks (yyyy-dd-mm, yadda-yadda)
// The fixed date is: Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
const dateLayout = "200601020304"

// For some weird reason we're getting the duration doubled
// We divide by 2 to handle that
func convDuration(counter Counter) string {
	duration := counter.Query.Duration / 2
	return fmt.Sprintf("%.2f", duration)
}

func handleEmpty(str string) string {
	if str == "" {
		return "0"
	}
	return str
}

func conv(v int) string {
	return handleEmpty(strconv.Itoa(v))
}

// Convert pgbadger generated logs to our LogMinute struct
func ConvertLogs(lines string) Logs {
	log.Info("Converting logs...")
	var logFile LogFile

	json.Unmarshal([]byte(lines), &logFile)
	fmt.Println("SELECTS:", logFile.PerMinuteInfo["20150801"]["11"]["00"])

	var logs Logs
	for date, info := range logFile.PerMinuteInfo {
		for hour, info := range info {
			for min, info := range info {
				timeStr := fmt.Sprintf("%s%s%s", date, hour, min)
				moment, err := time.Parse(dateLayout, timeStr)
				if err != nil {
					log.Panic(err)
				}
				fmt.Printf("%s:%s:%s SELECT: %v\n", date, hour, min, info.Insert.Count)
				logs = append(logs, newLogLine(moment, info))
			}
		}
	}
	sort.Sort(logs)
	return logs
}
