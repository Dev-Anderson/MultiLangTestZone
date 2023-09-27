# Autenticação JWT para API 
## Como rodar o projeto
1. Clonar o projeto 
2. Rodar o comando `go mod tidy`
3. Criar uma base de dados dentro do banco de dados postgres 
4. Configurar as variáveis de ambiente
5. Rodar o projeto com o comando `go run cmd/api/main.go`
6. Rota de criar usuário (localhost:8099/api/user), criar um usuário com o seguinte JSON `{"name":  "anderson","email":"anderson@gmail.com","password": "123"}`
7. Rota para login (localhost:8099/api/login), repassar qual usuário deseja fazer o login exemplo `{"email":  "anderson@gmail.com","password":  "123"}` logo depois deverá retornar um token 
8. Acessar a rota de Home (localhost:8099/api/home), dentro do header colocar key = Authorization, Value = Token que foi gerado, logo depois deverá vim a mensagem API rodando
