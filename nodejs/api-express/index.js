const express = require('express'); 
const app = express(); 

app.get('/api/hello', (req, res) => {
    res.json({ message: 'Hello, world! expresss'}); 
}); 

const port = 3000; 
app.listen(port, () => {
    console.log(`Server is running on port ${port}`); 
}); 