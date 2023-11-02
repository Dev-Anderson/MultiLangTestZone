# Estruturas de pasta para API Exemplo 

 arquivo "main.go" pode ficar na raiz do projeto. 

## Ideia do projeto para teste
API simples para demonstrar a estruturação das pasta que deve ser utilizada para as próximas APIs. 


### config 
- Carregamento das variaveis
- Configuração do Log

### controllers/hadler
O nome da pasta pode ser controllers ou handler tanto faz neste caso. <br>

### database
Conexão com o banco de dados e também outras informações pertinentes ao banco de dados, não deve colocar SQL dentro dessa pasta, para isso cria-se uma nova pasta com o nome "sql" na raiz do projeto. 

### assets
Pasta onde cica coisas como export do postman entre outros arquivos que precisam ser enviado. 

### models 
Onde faz as consultas, inserts, delete e update no banco de dados, para depois ser utilizado nos controllers/hadler

### routes 
Configurações das rotas e também o inicio do servidor

### helpers/utils
Pasta onde deve conter as funções gerais do projeto, exemplo conversão de inteiro para string, geração de algum tipo de código

### templates 
Pasta onde fica os arquivos HTML por exemplo, caso sua API retorna algum tipo de struct HTML. 