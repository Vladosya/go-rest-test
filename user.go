package todo

type User struct {
	Id       int    `json:"-"` // json:".." чтобы корректо принимать и выводить в http запросах
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
