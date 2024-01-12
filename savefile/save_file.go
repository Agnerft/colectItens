package savefile

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func SaveFile(url string) error {
	//url := "https://assets.mixkit.co/music/preview/mixkit-tech-house-vibes-130.mp3" // Substitua pela URL real do arquivo que você deseja baixar

	// Extrai o nome do arquivo da URL
	nomeArquivo := filepath.Base(url)

	fmt.Println(url)
	// Faça uma solicitação HTTP para obter o arquivo
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação HTTP:", err)
		return err
	}
	defer response.Body.Close()

	// Crie um arquivo para salvar o conteúdo
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return err
	}
	defer arquivo.Close()

	// Copie o conteúdo da resposta HTTP para o arquivo
	_, err = io.Copy(arquivo, response.Body)
	if err != nil {
		fmt.Println("Erro ao copiar o conteúdo para o arquivo:", err)
		return err
	}

	fmt.Printf("Arquivo '%s' baixado com sucesso!\n", nomeArquivo)

	return nil
}
