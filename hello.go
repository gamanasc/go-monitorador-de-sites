package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	exibeIntroducao()

	// For sem parâmetros roda indefinidamente
	for {
		exibeMenu()

		comando := leComando()

		// Em Go, o switch não precisa de break
		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Lucas"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/500",
		"https://httpbin.org/status/200"}

	for i := 0; i < 5; i++ {
		// For range é o equivalente ao foreach
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(5 * time.Second)
	}

	fmt.Println()

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}
