# Como usar migrations 

lib utilizada para criar as migrations

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


- Comando para fazer a instalacao do projeto

```
brew install golang-migrate
```

Com o CLI instalado na sua maquina e estando dentro da pasta do projeto, execute o seguinte comando

```
migrate create -ext sql -dir migrations -seq users
```

Se voce olhar nas pasta migrations vera que foram cirados os arquivos ```000001_users.down``` e ```000001_users.up```

No arquivo ```000001_users.up``` vamos adicionar o SQL para criar nossa tabela users

No arquivo ```000001_users.down``` vamos adicionar o SQL para dropar a tabela

Com os dois arquivos criados, e tendo o CLI instalado na maquina, ja e possivel aplicar a migration utilizando o comando 

Comando para criar o migrate dentro da pasta "migrations"

```
migrate create -ext sql -dir migratoins -seq users
```

```
migrate -path ./migrations -database "postgresql://user:password@localhost/dbname?sslmode=disable" -verbose up
```



Embora para realizar o DOWN de forma automatizada seja um pouco complexo, ja que teriamos que ter alguma especie de controle de versao do app, o UP e simples 
e como o migrate controla as versoes e so aplica o que ha de novo, podemos aplicar o up sempre que a app subir. 
Para fazer isso, vamos fazer uma pequena modificacao na nossa funcao que faz a conexao com o banco de dados. 
