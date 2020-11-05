package persistence

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type QuestionEntity interface {
	GetId() int
	GetText() string
	GetUserName() string
}

type QuestionRepository struct {
	Connection *pgx.Conn
}

/*
func getQuestions() error {
	rows, _ := connection.Query(context.Background(), "select * from tasks")

	for rows.Next() {
		var id int32
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%d. %s\n", id, description)
	}

	return rows.Err()
}

*/

func (questionRepository *QuestionRepository) GetQuestions() {

}

func (questionRepository *QuestionRepository) GetQuestionById(id int) {

}

func (questionRepository *QuestionRepository) GetQuestionByUserId(id int) {

}

func (questionRepository *QuestionRepository) CreateQuestion(q QuestionEntity) QuestionEntity {
	s := "insert into question (id,text,created_by) values($1,$2,$3)"
	questionRepository.Connection.Exec(context.Background(), s, q.GetId(), q.GetText(), q.GetUserName())
	return q

}

func (questionRepository *QuestionRepository) UpdateQuestion(q QuestionEntity) QuestionEntity {
	return q
}

func (questionRepository *QuestionRepository) DeleteQuestion(id int) {

}
