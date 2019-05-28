package main

import "fmt"
import "strconv"


type Processo struct {
	nome string
	burst int
	tcheg int
	quantum int
	prioridade int
	tespera int 
	tretorno int 
	tfim int 
	tresp int 
}

func main(){
	var algoritmo int
	var modo string
	processos := []Processo{}

	fmt.Println("Escolha seu algoritmo:\n1.FCFS\n2.SJF\n3.SRTF\n4.Round Robin\n5.Multinível\n\ta.primeiro nível - RR\n\tb.segundo nível - FCFS")
	fmt.Scanln(&algoritmo)

	fmt.Println("Inserção de processos:\n 1.use M para manual\n 2.Use L para ler arquivos")
	fmt.Scanln(&modo)

	//Trabalhar o modo de leitura de dados
	if modo == "M" {
		var qtd int
		var burst int
		var nome string
		var tcheg int

		fmt.Println("Digite a quantidade de processos:")
		fmt.Scanln(&qtd)
		for i:= 1; i<=qtd; i++ {
			fmt.Println("Digite o burst do Processo:", i)
			fmt.Scanln(&burst)
			fmt.Println("Digite a tcheg do Processo:", i)
			fmt.Scanln(&tcheg)

			nome = "P" + strconv.Itoa(i)

			processos = append(processos, Processo{nome,burst,tcheg,0,0,0,0,0,0})
		}

	} else if modo == "L" {
		fmt.Println("Modo leitura")
	} else {
		fmt.Println("Modo de entrada de processos inválida")
	}

	if algoritmo == 1{
		processos = fcfs(processos)
		_ = processos
	} else if algoritmo == 2{
		fmt.Println("algoritmo SJF")
	} else if algoritmo == 3{
		fmt.Println("algoritmo SRTF")
	} else if algoritmo == 4{
		fmt.Println("algoritmo RR")
	} else if algoritmo == 5{
		fmt.Println("algoritmo Multinível")
	} else {
		fmt.Println("Escolha inválida!!")
	}
	fmt.Println("-------------Informações do processamento----------------")

	var esperatotal int
	var esperamedia int
	var retornototal int
	var retornomedio int
	esperatotal = 0
	retornototal = 0
	fmt.Println("|Processo\t|Burst\t|T.Espera\t|T.Retorno\t|")
	for i := 0; i < len(processos); i++ {
		fmt.Printf("|%s\t\t|%d\t|%d\t\t|%d\t\t|\n", processos[i].nome, processos[i].burst,processos[i].tespera, processos[i].tretorno)
		esperatotal += processos[i].tespera
		retornototal +=processos[i].tretorno
	}

	esperamedia = esperatotal / len(processos)
	retornomedio = retornototal / len(processos)
	fmt.Printf("T.Espera Medio:%d\n", esperamedia)
	fmt.Printf("T.Retorno Medio:%d\n", retornomedio)
}

func fcfs(processos []Processo) []Processo {
	var time int
	time = 0
	for i := 0; i < len(processos); i++  {
		
		processos[i].tespera = time - processos[i].tcheg
		processos[i].tresp = time - processos[i].tcheg
		time += processos[i].burst
		processos[i].tretorno = time - processos[i].tcheg
	}

	return processos
}