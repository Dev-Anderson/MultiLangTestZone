package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	//abre o dispositivo de captura (0 e geralmente a webcam padrao)
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("Erro operando video capture device: ", err)
		return
	}
	defer webcam.Close()

	// cria uma janela para exibir a captura de videco
	window := gocv.NewWindow("Ola")
	defer window.Close()

	// cria um objeto mat para armazenar a imagem do video
	img := gocv.NewMat()
	defer img.Close()

	// define o codec e cria o writer para slavar o video
	codec := "avc1"
	outFile := "output.mp4"
	writer, err := gocv.VideoWriterFile(outFile, codec, 20, 640, 480, true)
	if err != nil {
		fmt.Println("Erro ao operar o video na gravacao", err)
		return
	}
	defer writer.Close()

	fmt.Println("Pressione ESC para sair")

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Camera fechada")
			return
		}
		if img.Empty() {
			continue
		}
		//mostra a imagem na janela
		window.IMShow(img)

		//escreve a imagem no arquivo de video
		writer.Write(img)

		// espera po 1 ms para a tecla de saida
		if window.WaitKey(1) == 27 {
			break
		}
	}

}
