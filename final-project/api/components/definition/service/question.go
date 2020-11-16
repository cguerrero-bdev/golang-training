package service

type Question struct {
	Id         int
	Statement  string
	UserName   string
	Answer     string
	AnsweredBy string
}

type QuestionManager interface {
	GetQuestions() ([]Question, error)
	GetQuestionById(id int) (*Question, error)
	GetQuestionsByUserName(userName string) ([]Question, error)
	CreateQuestion(question *Question) (*Question, error)
	UpdateQuestion(question *Question) (*Question, error)
	DeleteQuestion(id int) error
}
