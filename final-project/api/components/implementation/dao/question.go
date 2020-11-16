package dao

import (
	"context"
	"log"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/dao"
	"github.com/jackc/pgx/v4"
)

type QuestionDao struct {
	Connection *pgx.Conn

	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

const questionSelect = "select id, statement, created_by, answer, answered_by from question "

func (questionDao *QuestionDao) GetQuestions() ([]dao.QuestionEntity, error) {

	rows, err := questionDao.Connection.Query(context.Background(), questionSelect)

	if err != nil {
		questionDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	result := make([]dao.QuestionEntity, 0)

	for rows.Next() {

		questionEntity, err := questionRowsToEntity(rows)
		if err != nil {
			questionDao.ErrorLogger.Println(err.Error())
			return nil, err
		}

		result = append(result, questionEntity)
	}

	return result, nil
}

func (questionDao *QuestionDao) GetQuestionById(id int) (*dao.QuestionEntity, error) {

	row := questionDao.Connection.QueryRow(context.Background(),
		questionSelect+"where id=$1",
		id)

	result, err := questionRowToEntity(row, questionDao.ErrorLogger)

	if err != nil {
		questionDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (questionDao *QuestionDao) GetQuestionsByUserId(id int) ([]dao.QuestionEntity, error) {

	rows, err := questionDao.Connection.Query(context.Background(), questionSelect+"where created_by=$1", id)
	if err != nil {
		questionDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	result := make([]dao.QuestionEntity, 0)

	for rows.Next() {

		questionEntity, err := questionRowsToEntity(rows)
		if err != nil {
			questionDao.ErrorLogger.Println(err.Error())
			return nil, err
		}

		result = append(result, questionEntity)
	}

	return result, nil
}

func (questionDao *QuestionDao) CreateQuestion(q *dao.QuestionEntity) (*dao.QuestionEntity, error) {

	s := "insert into question (id,statement,created_by) values($1,$2,$3)"

	_, err := questionDao.Connection.Exec(context.Background(), s, q.Id, q.Statement, q.UserId)

	if err != nil {
		questionDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	return q, nil

}

func (questionDao *QuestionDao) UpdateQuestion(q *dao.QuestionEntity) (*dao.QuestionEntity, error) {

	s := "update question set statement=$1, answer = $2, answered_by = $3 where id = $4"

	answeredBy := &q.AnsweredBy

	if *answeredBy == 0 {
		answeredBy = nil
	}

	_, err := questionDao.Connection.Exec(context.Background(), s, q.Statement, q.Answer, answeredBy, q.Id)

	if err != nil {
		questionDao.ErrorLogger.Println(err.Error())
		return nil, err
	}

	return q, nil

}

func (questionDao *QuestionDao) DeleteQuestion(id int) error {

	s := "delete from question where id = $1"

	_, err := questionDao.Connection.Exec(context.Background(), s, id)

	if err != nil {
		questionDao.ErrorLogger.Println(err.Error())
		return err
	}

	return nil
}

func questionRowsToEntity(rows pgx.Rows) (dao.QuestionEntity, error) {

	questionEntity := dao.QuestionEntity{}

	var answer *string
	var answeredBy *int

	err := rows.Scan(
		&questionEntity.Id,
		&questionEntity.Statement,
		&questionEntity.UserId,
		&answer,
		&answeredBy,
	)

	if answer != nil {
		questionEntity.Answer = *answer
	}

	if answeredBy != nil {
		questionEntity.AnsweredBy = *answeredBy
	}

	return questionEntity, err
}

func questionRowToEntity(row pgx.Row, errorLogger *log.Logger) (*dao.QuestionEntity, error) {

	questionEntity := &dao.QuestionEntity{}

	var answer *string
	var answeredBy *int

	err := row.Scan(
		&questionEntity.Id,
		&questionEntity.Statement,
		&questionEntity.UserId,
		&answer,
		&answeredBy,
	)

	if answer != nil {
		questionEntity.Answer = *answer
	}

	if answeredBy != nil {
		questionEntity.AnsweredBy = *answeredBy
	}

	if err != nil {
		return nil, err
	}

	return questionEntity, nil
}
