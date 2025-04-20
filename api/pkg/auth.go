package pkg

import "time"

type AuthUser struct {
	Password string `json:"password" bson:"password"`
	Username string `json:"username" bson:"username"`
}

type CookieAuthUser struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
