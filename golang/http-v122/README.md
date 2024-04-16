# Info sobre a versao 1.22

Golang 1.22 foi lancado em 6 de fevereiro de 2024. Muitas atualizacoes foram aplicadas a estava versao, e uma das coisas legais foi o roteamenteo aprimorado. Com esse recuros, podemos criar parametros de rota dinamico sem a necessidade de bibliotecas de terceiros. 

## Sobre rota dinamica
Como voce pode ver, parece quase igual a outas bibliotecas existente. A diferenca e que o metodo http e definido antes do caminho da URL. Para recuperar o valor do parametro de URL dinamico, podemos uar "r.PathValue" tenha cuidade com o metodo HTTP porque ele diferencia maiusculas de minusculas, portanto, o nome do metodo hTTP deve ser maiusculo. 