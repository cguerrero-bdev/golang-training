package dao

import (
	"context"

	"github.com/cguerrero-bdev/golang-training/final-project/api/util"
	"github.com/jackc/pgx/v4"
)

type QuestionEntity struct {
	Id         int
	Statement  string
	UserId     int
	Answere    string
	AnsweredBy int
}

type QuestionDao struct {
	Connection *pgx.Conn
}

const questionSelect = "select id, statement, created_by, answer, answered_by from question "

func (questionDao *QuestionDao) GetQuestions() ([]QuestionEntity, *util.ApplicationError) {

	rows, err := questionDao.Connection.Query(context.Background(), questionSelect)

	if err != nil {

		return nil, util.GenerateApplicationErrorFromError(err)
	}

	result := make([]QuestionEntity, 0)

	for rows.Next() {

		questionEntity, _ := questionRowsToEntity(rows)

		result = append(result, questionEntity)
	}

	return result, nil
}

func (questionDao *QuestionDao) GetQuestionById(id int) (*QuestionEntity, error) {

	row := questionDao.Connection.QueryRow(context.Background(),
		questionSelect+"where id=$1",
		id)

	return questionRowToEntity(row)

}

func (questionDao *QuestionDao) GetQuestionsByUserId(id int) ([]QuestionEntity, error) {

	rows, err := questionDao.Connection.Query(context.Background(), questionSelect+"where created_by=$1", id)

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

func (questionDao *QuestionDao) CreateQuestion(q QuestionEntity) (QuestionEntity, error) {

	s := "insert into question (id,statement,created_by) values($1,$2,$3)"

	_, err := questionDao.Connection.Exec(context.Background(), s, q.Id, q.Statement, q.UserId)

	return q, err

}

func (questionDao *QuestionDao) UpdateQuestion(q *QuestionEntity) (*QuestionEntity, *util.ApplicationError) {

	s := "update question set statement=$1, answer = $2, answered_by = $3 where id = $4"

	answeredBy := &q.AnsweredBy

	if *answeredBy == 0 {
		answeredBy = nil
	}

	_, err := questionDao.Connection.Exec(context.Background(), s, q.Statement, q.Answere, answeredBy, q.Id)

	if err != nil {

		return nil, util.GenerateApplicationErrorFromError(err)
	}

	return q, nil

}

func (questionDao *QuestionDao) DeleteQuestion(id int) error {

	s := "delete from question where id = $1"

	_, err := questionDao.Connection.Exec(context.Background(), s, id)

	return err

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

func questionRowToEntity(row pgx.Row) (*QuestionEntity, error) {

	questionEntity := &QuestionEntity{}

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
