package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codeaspiras/goquiz/internal/entity"
	"github.com/codeaspiras/goquiz/internal/views/setquestion"
	"github.com/codeaspiras/goquiz/internal/views/showresult"
)

func main() {
	// load quiz
	questions, err := readQuiz()
	if err != nil {
		fmt.Printf("não foi possível ler o arquivo 'quiz.json': %s\n", err)
		return
	}

	// get answers
	answer := entity.Answer{
		Score:   0,
		Choices: make([]int, len(questions)),
	}
	for questionIndex := 0; questionIndex < len(questions); questionIndex++ {
		question := questions[questionIndex]
		choice, err := setquestion.Run(questionIndex, question)
		if err != nil {
			fmt.Printf("não foi possível obter a resposta da pergunta %d: %s\n", (questionIndex + 1), err)
			return
		}

		if choice == -1 {
			fmt.Printf("usuário abortou o quiz\n")
			return
		}

		if choice == question.RightAnswerIndex {
			answer.Score++
		}

		answer.Choices[questionIndex] = choice
	}

	// write answers
	err = writeAnswers(answer)
	if err != nil {
		fmt.Printf("não foi possível escrever o arquivo 'respostas.json': %s\n", err)
		return
	}

	// show results
	err = showresult.Run(questions, answer)
	if err != nil {
		fmt.Printf("não foi possível exibir o resultado: %s\n", err)
		return
	}
}

func readQuiz() (entity.Quiz, error) {
	bytes, err := os.ReadFile("./quiz.json")
	if err != nil {
		return nil, err
	}

	var quiz entity.Quiz
	err = json.Unmarshal(bytes, &quiz)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func writeAnswers(answers entity.Answer) error {
	bytes, err := json.Marshal(answers)
	if err != nil {
		return err
	}

	return os.WriteFile("respostas.json", bytes, os.ModePerm)
}
