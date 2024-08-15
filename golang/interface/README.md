Interface sao colecoes de assinaturas de metodos. 
Aqui esta uma interface basica para formas geometricas. 
Para nosso exemplo, implementaremos essa interface nos tipos "recte circle"

Para implementar uma interface em Go, precisamos apenas implementar todos os metodos na interface. Aqui implementamos
"geometrym rects". 
A implementacao para "circles"
Se uma variavel tiverum tipo de interface, entao podemos chamar metodos que estao na interface nomeadas. Aqui esta 
uma "measure" funcao generica aproveitando isso para trabalhar em qualquer "geometry". 

Os tipos "struct" "circle" "rect struct" implementam a "geometry" "interface" para que possamos usar instancias dessas
"structs" como argumentos para "mesure". 

## Simplificando
As interfaces em Go são uma maneira de definir um conjunto de métodos que um tipo deve implementar. elas permitem que você escreva código mais flexível e reutilizável, já que você pode trabalhar com diferentes tipos de dados desde que eles implemetem a interface requerida. 

Explicando o código "interface01.go"

- interface "Animal": Declara um método "Falar" que retorna uma string; 
- Tipos "Cachorro" e "Gato": Ambos implementam o método "Falar", portanto, ambos satisfazem a interface "Animal"; 
- Uso da interface; No "main", podemos atribuir instâncias de "Cachorro" e "Gato" à variável "animal" do tipo "Animal" e chamar o método "Falar" de forma polimórfica. 

### Exemplo bobo do que é uma interface 

Imagine que você tem uma caixinha mágica de brinquedos. Essa caixinha pode se transformar em qualquer brinquedo, mas só se esse brinquedo puder fazer uma coisa especial: emitir som.

Então, você diz para a caixinha:

"Eu quero brincar com algo que faz som de animal!"
Agora, a caixinha pode virar qualquer brinquedo que saiba emitir som de animal, como:

Um cachorro de brinquedo que faz "Au Au!"
Um gato de brinquedo que faz "Miau!"
A caixinha não se importa se é um cachorro ou um gato, desde que o brinquedo saiba fazer um som.

Em Go, isso é uma interface! 
* Caixinha mágina: É como uma interface. Ela não precisa saber exatamente que tipo de brinquedo ela é, só precisa saber que ele pode fazer som. 
* Brinquedos (cachorro e gato): São tipos diferentes que sabem fazer o som, então eles "funcionam" na caixinha mágica. 

Em programação, as interfaces permitem que a gente use diferentes coisas de uma maneira parecida, contando que elas saibam fazer o que a gente precisa (como emitir som). Assim, podemos brincar com muitos brinquedos diferentes, sem precisar de caixinhas diferentes para cada um. 


-----------------

O que é uma interface
Uma interface são duas coisas um conjunto de métodos, mas também é um tipo. Vamos focar no aspecto do conjunto de métodos das interfaces primeiro. 
Normalmente, somos apresentados a com algum exemplo artificial. vamos com o exemplo artificical de escrever algum aplicativo onde você está definindo tipos de dados Animal, porque essa é uma situação totalmente realista que acontece o tempo todo. O "Animal" tipo será uma interface, e definiremos um "Animal" como sendo qualuqer coisa que possa falar. este é um conceito central no sistema de tipos do Go; em vez de projetar nossas abastrações em termos de que tipo de dados nossos tipos podem conter, projetamos nossas abstrações em termos de quais ações nossos tipos podem executar. 
Começamos definindo nossa "Animal" interface: 

type Interface Animal {
    Falar() string
}

bem simples: definimos "Animal" como sendo qualquer tipo que tenha um método chamado "Speak". O "Speak" método não aceita argumentos e retorna uma "string". Qualquer tipo que define esse métodos é dito satisfazer a interface "Animal". Não há "implements" palavra-chave em Go. se um tipo satisfaz ou não uma interface é determinado automaticamente. Vamos criar alguns tipos que satisfazem essa interface 

tipo Cão struct {
}

func (d Cachorro) Fala() string {
    retornar "Uau!"
}

tipo Cat struct {
}

func (c Gato) Fala() string {
    retornar "Miau!"
}

tipo Llama struct {
}

func (l Lhama) Fala() string {
    retornar "?????"
}

tipo JavaProgrammer struct {
}

func (j JavaProgrammer) Fala() string {
    retornar "Padrões de design!"
}

Agora temos 4 tpios diferentes de animais: um cachorro, um gato, uma Ilhama e um programador Java. Em nossa "main()" função, podemos criar uma fatia de Animals, e colocar um de cada tipo nessa fatia, e ver o que cada animal diz. Vamos fazer isso agora: 

função principal() {
    animais := []Animal{Cão{}, Gato{}, Lhama{}, JavaProgrammer{}}
    para _, animal := intervalo animais {
        fmt.Println(animal.Falar())
    }
}

Ótimo, agora você sabe como usar interfaces, e não preciso mais falar sobre elas, certo? Bem, não não realmente. Vamos dar uma olhada em algumas coisas que não são muito óbvias para o gopher iniciante. 

## O interface{} tipo

O interface{} tipo, a interface vazia, é a fonte de muita confusão. O interface{} tip é a interface que não tem métodos. Como não há "implements" palavra-chave, todos os tipos implementam pelo menos zero métodos, e satisfazer uma interface é feito automaticamente, todos os tipos satisfazem a interface vazia. Isso significa que se você escrever uam função que recebe uma "interface{}" vazia. Isso significa que se você escrever uma função pode fornecer a essa função qualquer valor. Então, essa funçaõ:

func DoSomething(v interface{}) {}

Aceitará qualquer parâmetro. 