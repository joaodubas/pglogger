package main

// Final data structure that is serialized and sent to Graphite
// Every field populates a new metric on Graphite side
// The API we are using sends everything as string
type LogMinute struct {
	Timestamp   int64
	Connections string
	Sessions    string
	Selects     string
	Inserts     string
	Updates     string
	Deletes     string
	Max         string
	Min         string
	Duration    string
}

// Creating a Container type for our Logs that
// can be sorted according timestamp (older, smaller timestamp, first)
type Logs []LogMinute

func (m Logs) Len() int {
	return len(m)
}
func (m Logs) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m Logs) Less(i, j int) bool {
	return m[i].Timestamp < m[j].Timestamp
}

/*
 * JSON format generated by pgbadger:
 * {
 *   "per_minute_info": {
 *     "20150731": {
 *       "10": {
 *         "SELECT": {
 *           "count": 42,
 *           (...)
 *         },
 *         "INSERT": {
 *           (...)
 *         }
 *       },
 *       "11": {
 *         (...)
 *       }
 *     }
 *   }
 *  (...)
 * }
 */
type LogFile struct {
	PerMinuteInfo map[string]map[string]map[string]Counter `json:"per_minute_info"`
}

type HasCount struct {
	Count int `json:"count"`
}

type Counter struct {
	Others     HasCount `json:"OTHERS"`
	Select     HasCount `json:"SELECT"`
	Insert     HasCount `json:"INSERT"`
	Update     HasCount `json:"UPDATE"`
	Delete     HasCount `json:"DELETE"`
	Connection HasCount `json:"connection"`
	Session    HasCount `json:"session"`
	Query      struct {
		Duration float64 `json:"duration"`
		Max      string  `json:"max"`
		Min      string  `json:"min"`
	} `json:"query"`
}
