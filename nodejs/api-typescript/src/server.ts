import express, { Request, Response } from 'express'; 

const app = express(); 
const port = 3000; 

app.use(express.json()); 

// Dados em memoria (para fins de exemplo)
interface Item {
    id: number; 
    name: string; 
}

let items: Item[] = [
    { id: 1, name: 'Item 1' }, 
    { id: 2, name: 'Item 2' }, 
    { id: 3, name: 'Item 3' }, 
]; 

// metodo GET para listar todos os itens 
app.get('/items', (req: Request, res: Response) => {
    res.json(items); 
}); 

// metodo GET para obter um item especifico pelo ID 
app.get('/items/:id', (req: Request, res: Response) => {
    const id = parseInt(req.params.id, 10); 
    const item = items.find(item => item.id === id); 
    if (item) {
        res.json(item); 
    } else {
        res.status(404).send('Item nao encontrado'); 
    }
}); 

// Metodo POST para criar um novo item 
app.post('/items', (req: Request, res: Response) => {
    const newItem: Item = {
        id: items.length + 1, 
        name: req.body.name
    }; 
    items.push(newItem); 
    res.status(201).json(newItem)
}); 

// Metodos PUT para atualizar um item existente pelo ID 
app.put('/items/:id', (req: Request, res: Response) => {
    const id = parseInt(req.params.id, 10); 
    const itemIndex = items.findIndex(item => item.id === id); 
    if (itemIndex !== -1) {
        items[itemIndex].name = req.body.name; 
        res.json(items[itemIndex]); 
    } else {
        res.status(404).send('Item nao encontrado'); 
    }
}); 

// Metodo DELETE para remover um item pelo ID
app.delete('/items/:id', (req: Request, res: Response) => {
    const id = parseInt(req.params.id, 10); 
    const itemIndex = items.findIndex(item => item.id === id); 
    if (itemIndex !== -1) {
        items.splice(itemIndex, 1); 
        res.status(204).send(); 
    } else {
        res.status(404).send('Item nao encontrado'); 
    }
}); 

app.listen(port, () => {
    console.log(`Servidor rodando em http://localhost:${port}`)
})