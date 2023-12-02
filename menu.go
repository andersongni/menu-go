package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibirIntroducao()

	for {
		exibirMenu()

		comando := lerComando()

		switch comando {
		case 1:
			fazerMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}

func exibirIntroducao() {
	versao := 1.1
	fmt.Println("Este programa está na versão", versao)

}

func exibirMenu() {
	fmt.Println()
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println("O valor da variável comando é:", comando)
	return comando
}

func fazerMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://httpbin.org/status/200", "https://httpbin.org/status/404", "https://www.alura.com.br"}
	// sites := []string{}
	// sites = append(sites, "https://httpbin.org/status/200")
	// sites = append(sites, "https://httpbin.org/status/404")
	// sites = append(sites, "https://www.alura.com.br")

	sites := lerArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println(testaSite(site), ":", site)
		}
		time.Sleep(delay * time.Second)
	}

}

func testaSite(site string) string {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao testar o site", site, ":", err)
	}

	// fmt.Println("Resposta com status:", resp.StatusCode)

	if resp.StatusCode == 200 {
		registraLog(site, true)
		return "Site no ar"
	} else {
		registraLog(site, false)
		return "Site indisponivel"
	}
}

func lerArquivo() []string {
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)

	var sites []string
	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	fmt.Println(sites)

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	arquivo.WriteString(time.Now().Format(time.DateTime) + " | " + site + " | " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

	fmt.Println(string(arquivo))

}
