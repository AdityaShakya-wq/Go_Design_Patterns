package models

type QuizAnswer struct {
	answerId string
	question string
	answer   string
	position int
}

// Public getter methods for QuizAnswer struct
func (q QuizAnswer) AnswerId() string {
	return q.answerId
}

func (q QuizAnswer) Question() string {
	return q.question
}

func (q QuizAnswer) Answer() string {
	return q.answer
}

func (q QuizAnswer) Position() int {
	return q.position
}

func (q QuizAnswer) ToBuilder() *QuizAnswerBuilder {
	quizAnswerBuilder := NewBuilder()
	quizAnswerBuilder.SetAnswerId(q.answerId)
	quizAnswerBuilder.SetQuestion(q.question)
	quizAnswerBuilder.SetAnswer(q.answer)
	quizAnswerBuilder.SetPosition(q.position)
	return &quizAnswerBuilder
}

type QuizAnswerBuilder struct {
	answerId string
	question string
	answer   string
	position int
}

func NewBuilder() QuizAnswerBuilder {
	return QuizAnswerBuilder{}
}

func (q *QuizAnswerBuilder) SetAnswerId(answerId string) *QuizAnswerBuilder {
	q.answerId = answerId
	return q
}

func (q *QuizAnswerBuilder) ClearAnswerId() *QuizAnswerBuilder {
	q.answerId = ""
	return q
}

func (q *QuizAnswerBuilder) SetQuestion(question string) *QuizAnswerBuilder {
	q.question = question
	return q
}

func (q *QuizAnswerBuilder) ClearQuestion() *QuizAnswerBuilder {
	q.question = ""
	return q
}

func (q *QuizAnswerBuilder) SetAnswer(answer string) *QuizAnswerBuilder {
	q.answer = answer
	return q
}

func (q *QuizAnswerBuilder) ClearAnswer() *QuizAnswerBuilder {
	q.answer = ""
	return q
}

func (q *QuizAnswerBuilder) SetPosition(position int) *QuizAnswerBuilder {
	q.position = position
	return q
}

func (q *QuizAnswerBuilder) ClearPosition() *QuizAnswerBuilder {
	q.position = 0
	return q
}

func (q *QuizAnswerBuilder) Build() QuizAnswer {
	return QuizAnswer{
		answerId: q.answerId,
		question: q.question,
		answer:   q.answer,
		position: q.position,
	}
}

type Director struct {
	builder QuizAnswerBuilder
}

func NewDirector(q QuizAnswerBuilder) *Director {
	return &Director{
		builder: q,
	}
}

func (d *Director) BuildQuizAnswer(answerId string, question string, answer string, position int) QuizAnswer {
	d.builder.SetAnswerId(answerId)
	d.builder.SetQuestion(question)
	d.builder.SetAnswer(answer)
	d.builder.SetPosition(position)
	return d.builder.Build()
}
