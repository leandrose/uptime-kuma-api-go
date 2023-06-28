# Camada de Domínio

Responsável por modelar as entidades de negócios e suas regras de negócios associadas.
A camada de domínio é o coração da aplicação, e é onde a lógica de negócios crítica é implementada.

Os objetos dentro da camada de Domain são definidos pelos especialistas do domínio e são representados
por entidades, objetos de valor e agregados.

### Entidades

São objetos que têm identidade única e que são reconhecidos por seu comportamento. Eles possuem atributos e operações
e são utilizados para modelar os objetos do domínio que têm ciclos de vida e históricos. Um exemplo de entidade pode
ser um Cliente, que é identificado por seu número de identificação e possui um histórico de pedidos e transações.

### Objetos de Valor

São objetos que não possuem identidade própria e são definidos apenas pelos seus atributos. Eles são usados para
representar valores conceituais do domínio, como uma data, um endereço ou um preço. Esses objetos geralmente são
imutáveis, ou seja, não podem ser alterados após a sua criação.

### Agregados

São objetos que consistem em um grupo de entidades e objetos de valor relacionados. Eles são usados para definir
uma unidade de transação no domínio da aplicação e garantir a consistência dos dados dentro do agregado. Um exemplo
de agregado pode ser uma Ordem de Compra, que inclui um Cliente, uma lista de produtos e um endereço de entrega.

### Interfaces (Repository, Service, UseCase)

Objetos de Interface para implementação de funcionalidades. Encontra-se as interface para o Repository, Service,
UseCase, entre outros.
