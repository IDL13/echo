package unmarshal

type Redis struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type Auth struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Name struct {
	Name string `json:"name"`
}

type Date struct {
	Email string `json:"email"`
	Msg   string `json:"msg"`
}
