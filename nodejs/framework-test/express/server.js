const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
    res.send('Ola mundo!')
})

app.listen(port, () => {
    console.log(`Exemplo de um servidor na porta ${port}`)
})