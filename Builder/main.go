package main

import (
	"Go/Builder/models"
	json2 "encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Builder Pattern \n")
	quizBuilder := models.NewBuilder()
	quiz := quizBuilder.SetAnswerId("Id1").SetQuestion("First question").SetAnswer("answer 1").SetPosition(1).Build()
	fmt.Printf("Built quiz %v", quiz)
	fmt.Printf("Built quiz answer id %v", quiz.AnswerId())

	quizBuilderV2 := models.NewBuilder()
	quizDirector := models.NewDirector(quizBuilderV2)
	quizV2 := quizDirector.BuildQuizAnswer("Id2", "Second question", "answer 2", 2)

	fmt.Printf("\nBuilt quiz v2 with director %v", quizV2)
	fmt.Printf("\nBuilt quiz v2 Answer -> %s", quizV2.Answer())

	quizV2 = quizV2.ToBuilder().SetAnswer("Edited Answer").Build()
	fmt.Printf("\nBuilt quiz v2 Answer -> %s", quizV2.Answer())

	inputJson := make(map[string]interface{})
	inputJson["name"] = "first"
	inputJson["grade"] = "fifth"
	fmt.Printf("\nInput json %v", inputJson)
	json, _ := json2.Marshal(inputJson)
	fmt.Printf("\nJson %v", json)
	fmt.Printf("\nJson String%v", string(json))

	inputJson2 := "{\"name\":\"test\"}"
	fmt.Printf("\nInput json2 %v", inputJson2)
	json2_1, _ := json2.Marshal(inputJson2)
	fmt.Printf("\nJson2 %v\n", json2_1)
}
