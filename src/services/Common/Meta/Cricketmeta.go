package Meta

import (
	"encoding/json"
)

type Cricketmeta struct {
	Cricketer_name string `json:"name"`
	Cricketer_type string `json:"type"`
	Runs           int64  `json:"runs"`
	Wickets        int64  `json:"wickets"`
	Average        int64  `json:"average"`
	Matches        int64  `json:"match"`
	encoded        []byte
	err            error
}

func (cric *Cricketmeta) Encode() ([]byte, error) {
	if cric.encoded == nil && cric.err == nil {
		cric.encoded, cric.err = json.Marshal(cric)
	}
	return cric.encoded, cric.err
}

func (cric *Cricketmeta) Length() int {
	if cric.encoded == nil && cric.err == nil {
		cric.encoded, cric.err = json.Marshal(cric)
	}
	return len(cric.encoded)
}
