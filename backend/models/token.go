package models

import (
	"fmt"
	"time"
)

type Token struct {
	Expires int    `json:"expires_in,omitempty" bson:"expires_at"`
	Access  string `json:"access_token,omitempty" bson:"access_token"`
	Refresh string `json:"refresh_token,omitempty" bson:"refresh_token"`
}

func (t Token) String() string {
	return fmt.Sprintf("Token : %s\nRefresh : %s\nExpire in : %d\n", t.Access, t.Refresh, t.Expires)
}

//Convert token to timestamp
func (t *Token) ConvertToTime() {

	now := time.Now()

	future := now.Add(time.Second * 7200)

	//In timestamp format
	t.Expires = int(future.Unix())
}

//Token is valid if it's superior at now
func (t Token) IsValid() bool {

	now := time.Now().Unix()

	return t.Expires > int(now)
}

func NewToken(expires int, access, refresh string) *Token {

	return &Token{
		Expires: expires,
		Access:  access,
		Refresh: refresh,
	}

}
