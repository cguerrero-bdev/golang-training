package dao

type QuestionEntity struct {
	Id         int
	Statement  string
	UserId     int
	Answer     string
	AnsweredBy int
}

type QuestionDao interface {
	GetQuestions() ([]QuestionEntity, error)
	GetQuestionById(id int) (*QuestionEntity, error)
	CreateQuestion(q *QuestionEntity) (*QuestionEntity, error)
	GetQuestionsByUserId(id int) ([]QuestionEntity, error)
	UpdateQuestion(q *QuestionEntity) (*QuestionEntity, error)
	DeleteQuestion(id int) error
}
