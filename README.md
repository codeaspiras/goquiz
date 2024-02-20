# goquiz

Solução do desafio "Quiz player" do [CodeAspiras](https://codeaspiras.dev/discord). Para desenhar as views, utilizei o projeto [Bubbletea](https://github.com/charmbracelet/bubbletea).

## Desafio:

Faça um programa que lê um arquivo quiz.json, faz todas as perguntas programadas no arquivo e salva as respostas no arquivo respostas.json.

## Regras

1. As perguntas do quiz devem ser todas de única escolha, com no mínimo 2 opções e máximo 4;
2. As perguntas devem mostrar apenas o índice da resposta (0 = primeira opção, 1 = segunda opção);
3. Os arquivos devem seguir o formato abaixo:

`quiz.json`

```json
[
  {
    "question": "Essa seria a primeira pergunta...?",
    "options": ["Sim...?", "Não....?"],
    "right_answer": 0
  },
  {
    "question": "Essa seria a terceira pergunta...?",
    "options": ["Sim...?", "Não....?"],
    "right_answer": 1
  }
]
```

Onde question tem a pergunta a ser exibida, options tem a lista de opções e right_answer indica qual seria o índice da opção correta.

===

`respostas.json`

```json
{
  "score": 1,
  "choices": [0, 0]
}
```

Onde score representa o número de respostas corretas e choices tem todas as opções marcadas pelo usuário. Cada número representa o índice de uma resposta. O primeiro zero seria a resposta a pergunta do índice 0. O segundo zero seria a resposta a pergunta do índice 1.

## Como rodar a solução

1. Você precisar ter o [Golang v1.20+](https://go.dev/) instalado na sua máquina. Caso não o tenha, acesse o link informado anteriormente, baixe e instale!
2. Abra um terminal no diretório raiz desse projeto;
3. Rode `go run .` e siga as instruções do programa!

Você também pode compilar o projeto em um executável rodando `go build .`. Um arquivo executável será criado na raiz do projeto para que possa fazer o que quiser.