# Backend

Para iniciar o backend é necesssário ter o Go 1.18

```bash
go get
go run main.go
```

## Resultados

Os resultados dos testes são armazenados em um banco de dados SQL.

A coluna responses armazena um json contendo as respostas que o usuário fez na
seguinte forma:

```json
[
    {
      "question_id": "asldjkhas1238412",
      "response": true
    },
    {
      "question_id": "96f789vcbx6bx6c8",
      "response": false
    },
    {
      "question_id": "a0s8796asfasdf6f",
      "response": true
    }
]
```

Para calcular a eficácia de uma questão basta fazer a razão entre a quantidade
de vezes que a o status da response foi igual ao status de confirmação pela
quantidade total de vezes que a questão foi apresentada, quanto mais próximo
de 1, mais precisa é a questão.
