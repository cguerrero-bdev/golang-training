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

func (questionRepository *QuestionRepository) GetQuestions() ([]QuestionEntity, error) {

	rows, err := questionRepository.Connection.Query(context.Background(), "select id, text, created_by from question")

	result := make([]QuestionEntity, 0)

	for rows.Next() {

		questionEntity := QuestionEntity{}

		err := rows.Scan(&questionEntity.Id, &questionEntity.Text, &questionEntity.UserId)
		if err != nil {
			return nil, err
		}

		result = append(result, questionEntity)
	}

	return result, err
}

func (questionRepository *QuestionRepository) GetQuestionById(id int) (QuestionEntity, error) {

	result := QuestionEntity{}

	err := questionRepository.Connection.QueryRow(context.Background(),
		"select id, text, created_by from question where id=$1",
		id).Scan(&result.Id, &result.Text, &result.UserId)

	return result, err

}

func (questionRepository *QuestionRepository) GetQuestionsByUserId(id int) ([]QuestionEntity, error) {

	rows, err := questionRepository.Connection.Query(context.Background(), "select id, text, created_by from question where created_by=$1", id)

	result := make([]QuestionEntity, 0)

	for rows.Next() {

		questionEntity := QuestionEntity{}

		err := rows.Scan(&questionEntity.Id, &questionEntity.Text, &questionEntity.UserId)
		if err != nil {
			return nil, err
		}

		result = append(result, questionEntity)
	}

	return result, err
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
