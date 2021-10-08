package meta

type Cricketmeta struct {
	Cricketer_name string `json:"name"`
	Cricketer_type string `json:"type"`
	Runs           int64  `json:"runs"`
	Wickets        int64  `json:"wickets"`
	Average        int64  `json:"average"`
	Matches        int64  `json:"match"`
}
