package logic

type QuestionEntity interface {
	GetId() int
	GetText() string
	GetUserName() string
}

type Question interface {
	GetId() int
	GetText() string
	GetUserName() string
}

type QuestionRepository interface {
	GetQuestions()
	GetQuestionById(id int)
	GetQuestionByUserId(id int)
	CreateQuestion(q QuestionEntity) QuestionEntity
	UpdateQuestion(q QuestionEntity) QuestionEntity
	DeleteQuestionById(id int)
}

type QuestionManager struct {
	QuestionRepository QuestionRepository
}

func (questionManager *QuestionManager) GetQuestions() {

}

func (questionManager *QuestionManager) GetQuestionById(id int) {

}

func (questionManager *QuestionManager) GetQuestionByUserId(id int) {

}

func (questionManager *QuestionManager) CreateQuestion(q Question) Question {

	return questionManager.QuestionRepository.CreateQuestion(q)

}

func (questionManager *QuestionManager) UpdateQuestion(question Question) {

}

func (questionManager *QuestionManager) DeleteQuestion(id int) {

}
