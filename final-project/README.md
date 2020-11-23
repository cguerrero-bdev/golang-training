# Golang Development Program - Level 6 Final Project
---
# Statement
You are to design the backend side of a system for the following business idea.
We want to build a site called QuestionsAndAnswers.com that will compete with Quora/Stackoverflow and others with 1 major difference. We will only allow 1 answer per question. If someone thinks they have a better answer, they will have to update the existing answer for that question instead of adding another answer. In essence, each question can only have 0 or 1 answer.
The backend should support the following operations:
- Get one question by its ID
- Get a list of all questions
- Get all the questions created by a given user
- Create a new question
- Update an existing question (the statement and/or the answer)
- Delete an existing question

---
## Tech stack
---
- API service: go
- Database: PostgreSQL

## Usage
---
**Get one question by its ID** 
```sh
curl --location --request GET 'localhost:8080/questions/100'
````
**Get a list of all questions**
```sh
curl --location --request GET 'localhost:8080/questions'
````
**Get all the questions created by a given user**
```sh
curl --location --request GET 'localhost:8080/questions/user/1'
````
**Create a new question**
```sh
curl --location --request POST 'localhost:8080/questions' \
--header 'Content-Type: text/plain' \
--data-raw '{"id":100,"statement":"q100","userid":1,"answer":"","answeredby":0}'
````
**Update an existing question (the statement and/or the answer)**
```sh
curl --location --request PUT 'localhost:8080/questions/100' \
--header 'Content-Type: text/plain' \
--data-raw '{"id":100,"statement":"Question 100","userid":1,"answer":"Answer Question 100","answeredby":1}'
````
**Delete an existing question**
```sh
curl --location --request DELETE 'localhost:8080/questions/100'
````