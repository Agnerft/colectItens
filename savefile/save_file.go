package savefile

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
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

func ExtractFilenameFromURL(fileURL string) (string, error) {
	// Parse a URL
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", err
	}

	// Extrai o nome do arquivo da última barra no caminho
	filePath := filepath.Base(parsedURL.Path)
	fmt.Println(filePath)
	// Divide o caminho usando o símbolo de percentagem (%)
	parts := strings.Split(filePath, "%")
	fmt.Println(parts[0], parts[1], parts[2])
	// Verifica se há pelo menos quatro partes (antes e depois dos percentuais)
	if len(parts) < 4 {
		return "", fmt.Errorf("A URL não contém quatro partes separadas por percentuais")
	}

	// Junta as partes 2, 3 e 4 para obter o nome do arquivo desejado
	fileName := parts[2] + "%" + parts[3]

	// Decodifica os percentuais para obter o nome do arquivo correto
	fileName, err = url.QueryUnescape(fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
