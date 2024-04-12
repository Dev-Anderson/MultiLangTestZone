const { Pool } = require('pg'); 


//configs da conexao com o banco de dados
const pool = new Pool({
  user: 'postgres', 
  host: 'localhost', 
  database: 'teste', 
  password: '1234', 
  port: 5432
})

//teste conexao 
pool.query('select now()', (err, res) => {
  if (err) {
    console.error('Erro ao fazer a conexao com o banco de dados: ', err)
  } else {
    console.log('Conexao com o banco de dados realizada com sucesso!')
  }
}) 

//Exemplo de consulta
pool.query('select * from anderson', (err, res) => {
  if (err) {
    console.error('Erro ao executar a consulta: ', err); 
  } else {
    console.log('Resultado a consulta: ', res.rows)
  }
})

//fechando a conexao
pool.end(); 