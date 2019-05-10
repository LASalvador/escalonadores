package main

import "fmt"

func main(){
	var algoritmo int
	var modo string
	fmt.Printf("Escolha seu algoritmo:\n1.FCFS\n2.SJF\n3.SRTF\n4.Round Robin\n5.Multinível\n\ta.primeiro nível - RR\n\tb.segundo nível - FCFS\n")
	fmt.Scanln(&algoritmo)

	fmt.Printf("Inserção de processos:\n Use L para ler arquivos\n use M para manual\n")
	fmt.Scanln(&modo)

	//Trabalhar o modo de leitura de dados
	if modo == "M" {
		var qtd int
		fmt.Printf("Digite a quantidade de processos")
		fmt.Scanln(&qtd)
	} else if modo == "L" {
		fmt.Printf("Modo leitura")
	} else {
		fmt.Printf("Modo de entrada de processos inválida\n")
	}

	if algoritmo == 1{
		fmt.Printf("algoritmo FCFS")
	} else if algoritmo == 2{
		fmt.Printf("algoritmo SJF")
	} else if algoritmo == 3{
		fmt.Printf("algoritmo SRTF")
	} else if algoritmo == 4{
		fmt.Printf("algoritmo RR")
	} else if algoritmo == 5{
		fmt.Printf("algoritmo Multinível")
	} else {
		fmt.Printf("Escolha inválida!!\n")
	}

}