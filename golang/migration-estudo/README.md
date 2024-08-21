# Migrations Go

## O que são migratoins?

As migratoin são muiteo utilziada para manter um controle de versão da estrutura do seu banco de dados, é uma solução muito útil para manter seu banco de dados organizado. 
Imagine que você tenha uma tabela de usuário e precise inserir um novo campo nessa tabela, sem migrations você precisaria rodar um SQL na mão, como por exemplo: 

ALTER TABLE "usere" ADD COLUMN "phone" VARCHAR(255) NULL;

Agora, toda vez que precisar recriar a tabela "users", você vai precisar lembrar de criar este campo, a menos que alterar a criação original da tabeal "users", porém isso começa a fica inviável a medida que sua tabela e sua aplicação cresem, por isso utilzar as migrations são uma ótima opção. 

## As migrations

O funcionamento das migratinso são relativamente simples, geralmente temos uma arquivo de "up" e outr de "down", alguns ORMS como o PrimaORM criam apenas 1 arquivo, no arquivo de "up" criamos nosso SQL que vai criar ou alterar nosso banco de dados, no arquivo de "down" criamoso o SQL que desfaz a alteração. 

## Qual a vantagem 

Agora com esses arquivos, mantemos um histórico de alterações do banco de dados, cada alteração tem seu arquiv de "up" e "down", aogra se precisarmos criar as tabelas, rodamos todos os arquivos "up" e tudo é criado, se precisar reverter, basta rodar o "down". 

## Migrations em Go 
O Go não oferece nativamente suporte ao uso de migrations, mas poderíamos utilizar o ORM que tenha essa funcionalidade, como o GORM que é o mais utilizado pela comunicdade, mas podemos utilziar as migrations sem o uso de um ORM, para isso vamos utilizar o pacote "golang-migrate". 

## Pacote Golang Migrate 
O pacote "golang-migrate" é o mais recomendado para isso, já temos tudo que precisamos para gerenciar nossas migrations e oferece suprote a praticamente todos os banco de dados, para o nosso exemplo vamos utilziar o PostgresSQL. 

## Nosso projeto de exemplo

Já criei previamente um projeto simples, mas vou explicar rapidamenteo, pois o foco é utilizar as migrations. 

Teremos essa estrutura, bem simples, o código para exemplo deve ficar interido em "main.go"

## Utilizando o golang-migrate

Precisamo instalar o CLI do pacite, veja como instalar aqui, rode o comando 

```
migrate --version
```

Todo certo para continuar, o próximo passo é criar nossa primeira migrations com o comando: 

```
migrate create -ext=sql -dir=internal/database/migrations -sql init
```

- ext: determina a extensao, vaoms usar o sql. 
- dir: Aqui fica o diretorio onde vai ser criado as nossas migrations 
- sq: determina a sequencia do nome do arquivo da migrations, vamos usar numerico, pode ser usado timestamp. 

Com isso, voce vai perceber que foi criado uma pasta chamada migratoins dentro pasta database. 
Foi criado o arquivo "up"e o arquivo "down", na sequencia, como e a primeira fica 000001, se rodar novamente o comando "migrate create", vai criar a migration 000002. Agora vamos criar nosso SQL e rodar as migrations:
No arquivo de "up", vamos criar o create table. 

```
 CREATE TABLE users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) NOT NULL
  );
```

Dentro do arquivo de "down" vamos remover a tabela. 

```
DROP TABLE IF EXISTS users; 
```

Com isso SQL pronto, podemos rodar o "up"da nossa migrations, nao se esqueca de garantir que seu banco de dado esteja rodando, para isos deixei no projeto um arquivo docker compose para rodar uma imagem do PostgreSQL. 

```
  migrate -path=internal/database/migrations -database "postgresql://golang_migrate:golang_migrate@localhost:5432/golang_migrate?sslmode=disable" -verbose up
```

- path: informa onde estao as nossas migrations 
- database: URL da conexao com o banco de dados
- -verbose: apenas para exibir todas as execucoes 

Se acessarmos algum cliente com o PGADMIN ou Beekeeper, ou acessando seu container via bash e verificando via CLI, poderemos ver que a tabela foi criada com sucesso: 

Agora podemos rodar o "down",e exatamenteo o mesmo comando, porem alterando de "up"para "down"

```
  migrate -path=internal/database/migrations -database "postgresql://golang_migrate:golang_migrate@localhost:5432/golang_migrate?sslmode=disable" -verbose down
```

Com isso a tabela e removida

## Adicionado mais campos

Vamos ver agora como ficaria se fosse necessario adicionar mais um campo na tabela "users", sem migrations, teriamos que alterar diretamente a tabela original, mas com migrations nao precisamos, vamos criar outra migration: 

```
migrate create -ext=sql -dir=interna/database/migration -seq init
```

Vair criar a migration "up"e "down" com o sequencial 2, vamos adicioanr o campo "phone":

Dentro do arquivo UP: 

```
ALTER TABLE "users" ADD COLUMN "phone" VARCHAR(255) NULL;
```

Dentro do arquivo DOWN:

```
ALTER TABLE "users" DROP COLUMN "phone"; 
```

Ao rodar novamente: 

  migrate -path=internal/database/migrations -database "postgresql://golang_migrate:golang_migrate@localhost:5432/golang_migrate?sslmode=disable" -verbose up

Nosso campo "phone e adicionado a table user, mas se eu quiser fazer o "down" apenas na migration que adiciona o campo "phone"? e possivel, basta usar o mesmo comando, passadno o valor 1, que significa que deseja desfazer a ultima migration

```
  migrate -path=internal/database/migrations -database "postgresql://golang_migrate:golang_migrate@localhost:5432/golang_migrate?sslmode=disable" -verbose down 1
```

O campo "phone" e removido. 

## Facilitando o uso da CLI 

Como voce pode perceber, os comandos do golang-migrate podem ser um pouco cansativos de usar, podemos facilitar usando um arquivo makefile. 

```
  include .env

  create_migration:
    migrate create -ext=sql -dir=internal/database/migrations -seq init

  migrate_up:
    migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

  migrate_down:
    migrate -path=internal/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

  .PHONY: create_migration migrate_up migrate_down
```

Criamos atalhos que rodam o comando de "migrate", precisamos incluir nossas envs usando o "include .env", depois criamos os comandos. 

- create_migration: Cria nossos arquivos de migration
- migrate_up: Executa nossas migrations "up"
- migrate_dow: Execute nossas migrates "down"
- PHONY: garante que vain executar um comando, o makefile pode tentar pegar um arquivo, caso existe arquivo com o nome "migrate_up" por exemplo. 

Com isso, basta usar o comando:

```
make create_migrate
```