# Camada de Infraestrutura

A camada de infraestrutura no Domain-Driven Design (DDD) é responsável por fornecer os mecanismos necessários para
persistência e recuperação de dados, além de prover recursos como comunicação com serviços externos, envio de e-mails,
armazenamento em cache, entre outros.

Dentro da camada de infraestrutura, é comum ter um subconjunto de objetos dedicados exclusivamente à manipulação dos
dados, conhecidos como "Data Access Objects" (DAO) ou "Repositories". Esses objetos são responsáveis por implementar
as operações de leitura e gravação dos dados no banco de dados ou em outro tipo de armazenamento.

O "Data Mapper" é outro padrão comum utilizado na camada de infraestrutura do DDD, responsável por mapear os objetos
de domínio para as estruturas de dados utilizadas pelo armazenamento. Ele é responsável por garantir que as informações
armazenadas estejam em sincronia com as regras de negócio definidas na camada de domínio.

Em resumo, a camada de infraestrutura no DDD é responsável por fornecer os recursos necessários para que a aplicação
possa interagir com o mundo externo e persistir seus dados de forma confiável, isolando as complexidades técnicas da
aplicação e permitindo que ela se concentre nas regras de negócio e comportamentos do domínio.

### Adaptadores

São objetos que implementam as interfaces definidas na camada de domínio, permitindo que os objetos dessa camada
se comuniquem com os mecanismos de infraestrutura. Os adaptadores podem ser implementados para diferentes tecnologias,
como bancos de dados, serviços de mensagens, serviços de autenticação, entre outros. Eles podem estar localizados em
pacotes como ```adapter```, ```gateway``` ou ```provider```.

### Implementações de Repositórios

São objetos que implementam as interfaces definidas na camada de domínio para fornecer mecanismos de persistência dos
objetos do domínio. Eles são responsáveis por traduzir as operações de persistência de dados para o formato exigido
pelo mecanismo de armazenamento, como um banco de dados relacional ou um banco de dados NoSQL. As implementações de
repositórios podem estar localizadas em pacotes como ```repository```.

### Mecanismos de Comunicação

São objetos responsáveis por fornecer mecanismos de comunicação entre a aplicação e outras aplicações ou serviços
externos. Eles podem implementar diferentes protocolos de comunicação, como HTTP, TCP, UDP, ou utilizar tecnologias
de mensageria, como RabbitMQ ou Kafka. Os mecanismos de comunicação podem estar localizados em pacotes como
```communication``` ou ```gateway```.

### Módulos de Segurança

São objetos responsáveis por fornecer recursos de segurança à aplicação, como autenticação, autorização, criptografia 
e controle de acesso. Esses módulos podem ser implementados em diferentes camadas da aplicação, mas geralmente estão
localizados na camada de infraestrutura. Eles podem estar localizados em pacotes como ```security``` ou ```auth```.