const fastify = require('fastify')(); 

fastify.get('/api/hello', async (request, reply) => {
    return { message: 'Hello, World! com o Fastfy'}; 
}); 

const port = 3000; 
fastify.listen(port, (err) => {
    if (err) throw err; 
    console.log(`Server is running on port ${port}`); 
}); 