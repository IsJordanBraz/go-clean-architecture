# go-clean-architecture
Challenge 03 - Pós Go Expert 2024 - FullCycle

Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

# GraphQL - https://gqlgen.com/
- go run github.com/99designs/gqlgen generate
- go run server.go

# Proto files
- protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto


- show databases;

# docker
- docker-compose up -d
- docker-compose exec mysql bash
- mysql -uroot -p orders

# Personal Notes
- Use case - Intenção
- Frameworks, banco de dados, apis, não devem impactar as regras de negócio.
- Detalhes não devem impactar nas regras.
- # Use Cases - SRP(Single Responsibility Principle)
- Alterar vs Inserir, ambos verificam se uma entidade existe mas são casos de usos diferentes
- Resistir a vontade de reaproveitar codigo pois são casos de uso diferentes.
- Duplicação real vs acidental.
- Use case conta uma história
- # Limites Arquiteturais.
- Tudo que não impacta nas regras de negócio deve estar em um limite arquitetural diferente
- Ex: o Banco de dados conhece as regras de negócio mas as regras não conhecem o banco.
- # Presenters 
- transforma um output em um formato de entrega: json, xml, protobuf, graphql
- input = new CategoryInputDTO("NAME")
- output = CreateCategoryUseCase(input)
- json = CategoryPresenter(output).toJson()
- xml = CategoryPresenter(output).toXML()
- # Entities
- Entites Clean arch <> Entities DDD
