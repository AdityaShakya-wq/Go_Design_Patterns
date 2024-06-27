package main

import "fmt"

type quizAnswer struct {
	answerId string
	question string
	answer   string
	position int
}

type quizAnswerBuilder struct {
	answerId string
	question string
	answer   string
	position int
}

func newBuilder() quizAnswerBuilder {
	return quizAnswerBuilder{}
}

func (q *quizAnswerBuilder) setAnswerId(answerId string) *quizAnswerBuilder {
	q.answerId = answerId
	return q
}

func (q *quizAnswerBuilder) setQuestion(question string) *quizAnswerBuilder {
	q.question = question
	return q
}

func (q *quizAnswerBuilder) setAnswer(answer string) *quizAnswerBuilder {
	q.answer = answer
	return q
}

func (q *quizAnswerBuilder) setPosition(position int) *quizAnswerBuilder {
	q.position = position
	return q
}

func (q *quizAnswerBuilder) build() quizAnswer {
	return quizAnswer{
		answerId: q.answerId,
		question: q.question,
		answer:   q.answer,
		position: q.position,
	}
}

type director struct {
	builder quizAnswerBuilder
}

func newDirector(q quizAnswerBuilder) *director {
	return &director{
		builder: q,
	}
}

func (d *director) buildQuizAnswer(answerId string, question string, answer string, position int) quizAnswer {
	d.builder.setAnswerId(answerId)
	d.builder.setQuestion(question)
	d.builder.setAnswer(answer)
	d.builder.setPosition(position)
	return d.builder.build()
}

func main() {
	fmt.Println("Builder Pattern \n")
	quiz := newBuilder()
	quiz.setAnswerId("Id1").setQuestion("First question").setAnswer("answer 1").setPosition(1).build()
	fmt.Printf("Built quiz %v", quiz)
	fmt.Printf("Built quiz answer id %v", quiz.answerId)

	quizBuilderV2 := newBuilder()
	quizDirector := newDirector(quizBuilderV2)
	quizV2 := quizDirector.buildQuizAnswer("Id2", "Second question", "answer 2", 2)

	fmt.Printf("Built quiz with director %v", quizV2)
}
