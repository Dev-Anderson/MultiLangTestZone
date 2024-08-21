# Produtividade com Makefile 

## 1. O que e Makefile? 

makefile e um arquivo (por padrao chamado de "Makefile") que contem um conjunto de diretivas (linkers) usadas pela ferrament de automacao de compilacao GNU Make para executar um "alvo/meta/script". 

### 1.1. Instalacao 

Primeiro vamos verificar se o make ja esta instaldo em sua distribuicao:

```
make -v
```

O seu resultado, caso o GNU Make ja esteja instalado em sua distro, sera similar ao mostrado abaixo: 

```
GNU Make 3.81
Copyright (C) 2006  Free Software Foundation, Inc.
This is free software; see the source for copying conditions.
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

This program built for i386-apple-darwin11.3.0
```

Se por acaso, o comando make nao for reconhecido, voce pode instalar com o seguinte comando: 

```
sudo apt install make
```

### 1.2. Como usar? 

Crie um diretorio com nome de sua escolha, e dentro dele crie um arquivo de nome "Makefile", com o seguinte conteudo: 

```
hello:
	@echo "Hello World"
```

Na primeira linha, temos a diretiva (liker), e na segunda o comando a ser executando quando essa diretiva for acionada, para isso basta usar o comando make + a diretiva desejada, como no exemplo a seguir: 

```
make hello
```

Print sera : 
```
Hello World
```

Comandao @
o @ no inicio da linha indica que voce deseja ocultar o comando durante a execucao, ou seja, caso voce nao inicie o comando com @ no script do docmento, ele sera exibido antes do resultado durante a execucao, como no exemplo abaixo:

Makefile:

```
hello:
    echo "Hello world"
```

Resultado:

```
make hello 
echo "Hello World"
Hello World
```

Parametros
E possivel tambem usar parametros, como por exemplo:

```
hello:
    @echo "Hello World $(name)"
```

Resultado: 

```
make hello name=Antonio
Hello World Antonio
```

Include:

```
make hello name=Anderson
Hello World Anderson
```

Include:

As vezes o arquivo Makefile, pode ser torna bem grande, uma das formas de organizar isso, e separar os comandos de acordo com o contexto que melhor se encaixe e usar o coso comando include. 
Por exemplo:

```
include= .targets-docker 
include= .targets-sqs
```

Help 
Essa e uma sugestao bem legal, e possivel criar um comando help no makefile, para nos ajudar a entender para que serve cada comando que criamos. 
Existem varias formas de fazer isso, buscando um pouco na comunidade, encontrei alguns exemplos, e customizei da seguinte forma:

1. Crie um arquivo .help no mesmo diretorio do Makefile em questao, com o conteudo abaixo:
2. Inclua esse arquivo iniciando seu makefile com o seguinte comando: 

```
include .help
```

Agora basta fazer seus comandos, com comentario acima com ##, por exemplo:

```
## Escreve Ola mundo!
hello:
    echo "Hello World"
```