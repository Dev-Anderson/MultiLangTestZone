# Express JS 

Expresss.js e um dos frameworks de back-en mais famoso do nodejs. E uma estrutura de aplicativo da web de codigo aberto, disponivel gratuitamente e contruida na plataforma nodejs. Por ser uma estrutura minima, desenvollvedores web novatos e experitentes tendem a optar pelo express.js. E usado principalmente para criar aplicativos web e Api RestFul. 

## Principais recursos: O que faz com que ele se destaque?
1. Roteamente eficiente 
Express.js oferece um método limpo e simples para gerenciar várias solicitações HTTP e atruibuí-las a tarefas específicas. Vejamos um exemplo. 

```
// app.js 
const express = require('express'); 
const app=express(); 
porta const = 3000; 

// Rota para a página inicial 
app.get('/', (req, res) => { 
  res.send('Bem-vindo à página inicial!'); 
}); 

// Rota 2 
app.get('/user/:id', (req, res) => { 
  const userId = req.params.id; 
  res.send(Página de perfil do usuário - ID: ${userId} ); 
} );
```

2. Suporte de middleware 
Express.js permite suporte de middleware para lidar com solicitações hTTP. Vejamos um exemplo simples de criaçaõ de um middleware para registrar detalhes de solicitações hTTP. 


3. Fácil integração com banco de dados
Express.js é independente de banco de dados. Não impõe uma escolha específica de banco de dados. Os desenvolvedores podem escolher seu banco de dados preferido. A facilidade de integração de bancos de dados com express.js se deve à ua natureza modular e flexível e ao rico ecosistema de pacotes npm que fornecem conectivaide de banco de dados. 

4. Fácil de aprender 
Express.js é conhecido por sua simplicidade e design minimalista, tornando-o muito fácil de aprender para os desenvolvedores, especialmente se eles já estiverem familiarizaods com javascript e nodejs. 
Além disso, você pode começar facilmente a usar express.js com ferramentas com Bit. Se vocÊ nunca usou o bit antes, é um sistema de construção de próxima geração para software combinável. 

Você pode ver dois componentes: um autoreizaodr e também um aplicativo API. Esses dois componentes foram implementados como cpomnentes Bit indepdnentes e são mantidos e versionados em seu espaço isolado.
Ao fazer isso, você pode projetoar seu aplicativo de maneira combinável e rapidamente. 
