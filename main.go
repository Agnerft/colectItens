package main

import (
	"fmt"

	"github.com/agnerft/colectItens/savefile"
)

func main() {

	// nomeArquivo := filepath.Base("https://assets.mixkit.co/music/preview/mixkit-tech-house-vibes-130.mp3")
	// err := savefile.SaveFile("https://assets.mixkit.co/music/preview/mixkit-tech-house-vibes-130.mp3")
	// if err != nil {
	// 	fmt.Println("Deu erro")
	// }

	fileUrl, err := savefile.ExtractFilenameFromURL("https://audio-previews.elements.envatousercontent.com/files/184174962/preview.mp3?response-content-disposition=attachment%3B+filename%3D%22AEC3SH5-ambient-piano.mp3%22")
	if err != nil {
		fmt.Println("erro ao modificar o nome.")
	}

	fmt.Println(fileUrl)

	// url := "https://mixkit.co/free-stock-music/ambient/" // Substitua pela URL real
	// // page := "?page=1"
	// // Faça uma solicitação HTTP para obter o conteúdo da página
	// response, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()

	// // Use goquery para analisar o HTML
	// doc, err := goquery.NewDocumentFromReader(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Crie um arquivo para salvar os valores
	// file, err := os.Create("valores_audio1.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // Selecione todas as divs com a classe audio-player
	// doc.Find("[data-controller=audio-player]").Each(func(i int, audioPlayer *goquery.Selection) {
	// 	// Obtenha o valor do atributo data-audio-player-preview-url-value
	// 	previewURL, exists := audioPlayer.Attr("data-audio-player-preview-url-value")
	// 	if exists {
	// 		// Escreva o valor no arquivo
	// 		fmt.Println(previewURL)

	// 		fmt.Fprintf(file, "%s\n", fileUrl)
	// 		// err := savefile.SaveFile(previewURL)
	// 		// if err != nil {
	// 		// 	fmt.Println("Erro ao baixar")
	// 		// }

	// 		time.Sleep(6 * time.Second)
	// 	} else {
	// 		fmt.Fprintf(file, "Item %d: Atributo data-audio-player-preview-url-value não encontrado\n", i+1)
	// 	}
	// })

	// fmt.Println("Valores salvos em 'valores_audio.txt'")
}
