# Camada de Aplicação

Responsável por orquestrar as ações da aplicação, utilizando as regras de negócio definidas na camada de domínio e
os serviços fornecidos pela camada de infraestrutura. Essa camada é o ponto de entrada para as requisições externas,
sejam elas via interface do usuário, API, serviços ou outras aplicações.

Os objetos presentes na camada de aplicação atuam como intermediários entre as requisições externas e os objetos de
domínio e infraestrutura, permitindo a separação de responsabilidades e facilitando a manutenção do código. Essa camada
é responsável por implementar as operações definidas no contexto de aplicação e por garantir que as regras de negócio
sejam respeitadas.

### Serviços de Aplicação

São objetos que encapsulam uma operação específica do domínio, utilizando os objetos de domínio e infraestrutura
necessários para sua execução. Eles podem ser implementados para diferentes contextos de aplicação, como cadastro de
usuários, processamento de pedidos, geração de relatórios, entre outros.

### Casos de Uso

São objetos que encapsulam uma sequência de operações do domínio, necessárias para realizar uma tarefa específica.
Eles podem ser compostos por um ou mais serviços de aplicação, e podem ser implementados para diferentes fluxos de
trabalho da aplicação.

### Controladores

São objetos que recebem as requisições externas e direcionam as operações correspondentes para os serviços de
aplicação e casos de uso apropriados. Eles também são responsáveis por traduzir as informações da requisição para o
formato utilizado pelos objetos de domínio e infraestrutura.

### DTOs (Data Transfer Objects)

São objetos utilizados para transportar dados entre as diferentes camadas da aplicação, geralmente representando
informações obtidas a partir de uma requisição externa ou a serem enviadas como resposta. Eles podem ser utilizados
para mapear os dados da requisição para objetos de domínio, ou para representar informações em diferentes formatos,
como JSON, XML ou HTML.

