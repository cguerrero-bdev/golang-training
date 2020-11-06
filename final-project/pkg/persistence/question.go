package persistence

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type QuestionEntity struct {
	Id         int
	Statement  string
	UserId     int
	Answere    string
	AnsweredBy int
}

type QuestionRepository struct {
	Connection *pgx.Conn
}

const questionSelect = "select id, statement, created_by, answer, answered_by from question "

func (questionRepository *QuestionRepository) GetQuestions() ([]QuestionEntity, error) {

	rows, err := questionRepository.Connection.Query(context.Background(), questionSelect)

	result := make([]QuestionEntity, 0)

	for rows.Next() {

		questionEntity, _ := questionRowsToEntity(rows)

		result = append(result, questionEntity)
	}

	return result, err
}

func (questionRepository *QuestionRepository) GetQuestionById(id int) (QuestionEntity, error) {

	row := questionRepository.Connection.QueryRow(context.Background(),
		questionSelect+"where id=$1",
		id)

	return questionRowToEntity(row)

}

func (questionRepository *QuestionRepository) GetQuestionsByUserId(id int) ([]QuestionEntity, error) {

	rows, err := questionRepository.Connection.Query(context.Background(), questionSelect+"where created_by=$1", id)

	result := make([]QuestionEntity, 0)

	for rows.Next() {

		questionEntity, err := questionRowsToEntity(rows)

		if err != nil {
			return nil, err
		}

		result = append(result, questionEntity)
	}

	return result, err
}

func (questionRepository *QuestionRepository) CreateQuestion(q QuestionEntity) (QuestionEntity, error) {

	s := "insert into question (id,statement,created_by) values($1,$2,$3)"

	_, err := questionRepository.Connection.Exec(context.Background(), s, q.Id, q.Statement, q.UserId)

	return q, err

}

func (questionRepository *QuestionRepository) UpdateQuestion(q QuestionEntity) (QuestionEntity, error) {

	s := "update question set statement=$1, answer = $2, answered_by = $3 where id = $4"

	_, err := questionRepository.Connection.Exec(context.Background(), s, q.Statement, q.Answere, q.AnsweredBy, q.Id)

	return q, err

}

func (questionRepository *QuestionRepository) DeleteQuestion(id int) {

}

func questionRowsToEntity(rows pgx.Rows) (QuestionEntity, error) {

	questionEntity := QuestionEntity{}

	var answere *string
	var answeredBy *int

	err := rows.Scan(
		&questionEntity.Id,
		&questionEntity.Statement,
		&questionEntity.UserId,
		&answere,
		&answeredBy,
	)

	if answere != nil {
		questionEntity.Answere = *answere
	}

	if answeredBy != nil {
		questionEntity.AnsweredBy = *answeredBy
	}

	return questionEntity, err
}

func questionRowToEntity(row pgx.Row) (QuestionEntity, error) {

	questionEntity := QuestionEntity{}

	var answere *string
	var answeredBy *int

	err := row.Scan(
		&questionEntity.Id,
		&questionEntity.Statement,
		&questionEntity.UserId,
		&answere,
		&answeredBy,
	)

	if answere != nil {
		questionEntity.Answere = *answere
	}

	if answeredBy != nil {
		questionEntity.AnsweredBy = *answeredBy
	}

	return questionEntity, err
}
