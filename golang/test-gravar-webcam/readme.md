# Passo a passo 

Instalacoes

- Instalado o pkg-config
 
"brew install pkg-config"

- Instalando o opencv 

"brew install opencv"

- Verificando a instalacao
"pkg-config --modversion opencv4"

# Explicacao

- Abrindo a camera: "gocv.OpenVideoCapture(0)" abre a camera padrao 
- Criando a janela "gocv.NewWindow("Hello")" cria uma janela para exibir a captura de video 
- Criando um objeto MAT: "gocv.NewMat()" e usado para armazenar as imagens do video 
- Criando o writer de video: "gocv.VideoWriterFile" cria um writer para salvar o video em um arquivo
- Loop principal le o frame da camera, exibe na janela e escreve no arquivo de video 
- Esperando a teclada "window.Waitkey(1) == 27" espera a tecla ESC para sair do loop

Com isso, voce tera um programa simples que captura video da webcam e o salva em um arquvio AVI. 