package model

type Question struct {
	Id         int    `json:"id"`
	Statement  string `json:"statement"`
	UserId     int    `json:"username"`
	Answer     string `json:"answer"`
	AnsweredBy int    `json:"answeredby"`
}
