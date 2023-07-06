# Sobre

Este é um projeto de esqueleto (ou modelo de projeto) para Golang, estruturado com base na arquitetura Domain-Driven
Design (DDD). Ele contém uma organização de pastas com as camadas Application, Domain, Infraestrutura e Presenter,
além de binários para executar tanto uma aplicação web quanto uma CLI.

## Requisitos

- Golang (versão 1.18 ou superior)

## Estrutura do Projeto

```
.
├── app
│   ├── handlers
│   │   └── some_handler.go
│   ├── services
│   │   └── some_service.go
│   ├── providers  
│   │   └── some_provider.go
│   ├── usecases
│   │   └── some_usecase.go
│   └── app.go
├── cmd
│   ├── main.go
│   ├── root.go
│   └── serve.go
├── domain
│   ├── entities
│   │   └── some_entity.go
│   ├── repositories
│   │   └── some_repository.go
│   └── services
│       └── some_service.go
├── infra
│   ├── database
│   │   └── some_database.go
│   ├── repositories
│   │   └── some_repository.go
│   └── services
│       └── some_service.go
└── presenter
    ├── cli
    │   └── some
    │       └── index_presenter.go
    └── http
        └── some
            └── index_presenter.go
```

## Camadas

### app

A camada app contém a lógica de aplicação, responsável por orquestrar as interações entre as diferentes camadas.
Ela contém três subpastas: ```handlers```, ```services``` e ```usecases```.

- ```handlers```: contém os manipuladores de requisições HTTP, responsáveis por receber as requisições e executar os
casos de uso apropriados;
- ```services```: contém a lógica de negócios específica da aplicação, como validações e cálculos; 
- ```usecases```: contém os casos de uso da aplicação, responsáveis por orquestrar as interações entre os serviços 
e as entidades do domínio.

### cmd

A camada cmd contém os binários para executar a aplicação tanto como uma CLI quanto como uma aplicação web.

- ```cli```: contém o binário para executar a aplicação como uma CLI;
- ```http```: contém o binário para executar a aplicação como uma aplicação web.

### domain

A camada domain contém a lógica de domínio da aplicação, incluindo as entidades e serviços do domínio,
além dos repositórios que são responsáveis pela persistência dos dados.

- ```entities```: contém as entidades do domínio, que encapsulam as informações relevantes para a aplicação e as

### infra

A camada Infraestrutura é responsável por fornecer implementações concretas para as interfaces definidas na camada
Domain, como por exemplo, o acesso ao banco de dados e a configuração do servidor HTTP. A organização interna dessa
camada pode variar de acordo com a necessidade do projeto, no entanto, aqui é proposta a organização em subpastas.

- ```config```: esta subpasta contém as configurações gerais da aplicação.
- ```repository```: esta subpasta contém as implementações concretas das interfaces definidas na camada Domain.
- ```service```: esta subpasta contém as implementações concretas das interfaces definidas na camada Domain.

### presenter

A camada Presenter é responsável por fornecer interfaces de entrada e saída da aplicação, como a interface de linha
de comando (CLI) e a interface web (HTTP). A organização interna dessa camada pode variar de acordo com a necessidade
do projeto, no entanto, aqui é proposta a organização em subpastas.

- ```cli```: esta subpasta contém o código responsável por processar os comandos da interface de linha de comando (CLI).
- ```http```: esta subpasta contém o código responsável por processar as requisições HTTP.

# Como Executar

Execute no Terminal ou configure a IDE para executar o comando abaixo:

```
go run main.go
```
