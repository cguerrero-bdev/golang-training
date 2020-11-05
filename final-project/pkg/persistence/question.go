package persistence

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type QuestionEntity struct {
	Id     int
	Text   string
	UserId int
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

func (questionRepository *QuestionRepository) GetQuestionById(id int) (QuestionEntity, error) {

	result := QuestionEntity{}

	err := questionRepository.Connection.QueryRow(context.Background(),
		"select id, text, created_by from question where id=$1",
		id).Scan(&result.Id, &result.Text, &result.UserId)

	return result, err

}

func (questionRepository *QuestionRepository) GetQuestionByUserId(id int) {

}

func (questionRepository *QuestionRepository) CreateQuestion(q QuestionEntity) (QuestionEntity, error) {

	s := "insert into question (id,text,created_by) values($1,$2,$3)"

	_, err := questionRepository.Connection.Exec(context.Background(), s, q.Id, q.Text, q.UserId)

	return q, err

}

func (questionRepository *QuestionRepository) UpdateQuestion(q QuestionEntity) QuestionEntity {
	return q
}

func (questionRepository *QuestionRepository) DeleteQuestion(id int) {

}
