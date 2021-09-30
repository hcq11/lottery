package format


type Response struct {
	Code int
	Msg string
}

type User struct {
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}