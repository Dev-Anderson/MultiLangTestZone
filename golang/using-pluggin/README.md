# Sobre o recurso de plugin em Golang 

A partir do Go 1.8, os desenvolvedores podem escrever plug-ins e carregá-loa dinamicamente em tempo de execuçaõ. Esse recurso adicona um novo nível de flexibilidade aos programas Go, permitindo designs modulares e extensíveis. Com plug-ins, você pode estender a funcionalidade do seu aplicativo sem a necessidade de recompilar ou reimplantar toda a base de código. 

## Principais benefícios 

- Carregamento dinâmico: os plug-ins podem ser carregados e descarregados dinamicamente em tempo de execução, permitindo a adição ou remoção de funcionalidades sem interromper o programa principal. 
- Flexibilidade: Ao contrário das importações de bibliotecas tradicionais, que são resolvida no momento da construçaõ, os plug-ins oferecem uma abordagem mais flexível. Eles podem ser trocados ou atualizados de forma independente, dando a você a liberdade de iterar e experimentar diferentes implementações. 
- Modularidade: O recurso de plugin promove um padrão de design modular encapsulado funcionalidades específicas em componentes separados. Isso melhora a organização, a reutilização e a capacidade de manuteção do código. 
- Encapsulamento: Os plug-ins são executados em seus próprios namespaces isolados, evitando conflitos entre diferentes plug-ins ou o programa principal. Isso permite extensibilidade segura e confiável. 

### build do plugin 

```
go build -buildmode=plugin -o ./simple-plugin/plugin.so ./plugin.go
```

## Como utilizar
1. Criar a seguinte função

```
func loadPlugin() func() {
    // carregando o plugin gerado
	plugin, err := plugin.Open("plugins/simple-plugin.so")
	if err != nil {
		log.Fatal(err)
	}

    // carregando a função para ser utilizado
	simplePluginFunc, err := plugin.Lookup("SimplePluginFunc")
	if err != nil {
		log.Fatal(err)
	}

    // se a função foi carregada com sucesso  
	f, ok := simplePluginFunc.(func())
	if !ok {
		log.Fatal("Tipo nao suportado")
	}
	return f
}
```

2. Passo 02 só utilizar a função exemplo

```
func main() {
	simplePlugin := loadPlugin()
	simplePlugin()
}
```