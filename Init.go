package main

import "fmt"
import "strconv"


type Processo struct {
	nome string
	burst int
	tcheg int
	prioridade int
	quantum int
	tespera int 
	tretorno int 
	tfim int 
	tresp int 
}

func main(){
	var algoritmo int
	var modo string
	processos := []*Processo{}

	fmt.Printf("Escolha seu algoritmo:\n1.FCFS\n2.SJF\n3.SRTF\n4.Round Robin\n5.Multinível\n\ta.primeiro nível - RR\n\tb.segundo nível - FCFS\n")
	fmt.Scanln(&algoritmo)

	fmt.Printf("Inserção de processos:\n 1.use M para manual\n 2.Use L para ler arquivos\n")
	fmt.Scanln(&modo)

	//Trabalhar o modo de leitura de dados
	if modo == "M" {
		var qtd int
		var burst int

		fmt.Printf("Digite a quantidade de processos: \n")
		fmt.Scanln(&qtd)
		for i:= 1; i<=qtd; i++ {
			fmt.Printf("Digite o burst do Processo %d:\n", i)
			fmt.Scanln(&burst)
			p := new(Processo)
			p.nome = "P" + strconv.Itoa(i)
			p.burst = burst

			processos = append(processos, p)
		}

	} else if modo == "L" {
		fmt.Printf("Modo leitura")
	} else {
		fmt.Printf("Modo de entrada de processos inválida\n")
	}

	if algoritmo == 1{
		fmt.Printf("algoritmo FCFS")
		p := processos[0]
		fmt.Println(p.nome)
		fmt.Println(p.burst)
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