# Estrutura de um API com arquitetura limpa 

Pastas
- Handlers: Contém os manipuladores HTTP (controladores) que lidam com as requisições HTTP. 
- models: Contém as definições de estrutura de dados (modelos) usadas pela aplicação
- repositores: Contém os repositórios que lidam com as operações de banco de dados
- services: Contém a lógica de negócio principal da aplicação
- config: Contém o código para configurar a aplicação, como configurações de banco de dados etc. 
- db: Contém o código relacionado ao banco de dados,como conexão com o banco de dados

Pasta principais 
pkg -> Pasta onde fica todas as funções que serão compartilhadas
internal -> Pasta com coisas internas 


Ainda falta implementar os teste 

## Arquivo de configuração

api-test-host=
api-test-port=
api-test-user=
api-test-password=
api-test-dbname=
api-test-porthttp=:8080

## Teste realizados nas rotas
