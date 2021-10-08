package Meta

import (
	"encoding/json"
)

type Usermeta struct {
	Username string `json:"name"`
	Password string `json:"pass"`
	encoded  []byte
	err      error
}

func (user *Usermeta) Encode() ([]byte, error) {
	if user.encoded == nil && user.err == nil {
		user.encoded, user.err = json.Marshal(user)
	}
	return user.encoded, user.err
}

func (user *Usermeta) Length() int {
	if user.encoded == nil && user.err == nil {
		user.encoded, user.err = json.Marshal(user)
	}
	return len(user.encoded)
}
