package pkg

type AuthUser struct {
	Password string `json:"password" bson:"password"`
	Username string `json:"username" bson:"username"`
}

// type RedisAuthenticateUser struct {
// 	Token    string `json:"token"`
// 	Username string `json:"username"`
// }
