package main

import (
	"fmt"
 	"strconv"
 )


type Processo struct {
	nome string
	burst int
	trest int
	tcheg int
	quantum int
	prioridade int
	tespera int 
	tretorno int 
	tMod int 
	tresp int 
}

func main(){
	var algoritmo int
	var modo string
	burst_total := 0
	processos := []Processo{}

	for ;; {
	fmt.Println("Escolha seu algoritmo:\n1.FCFS\n2.SJF\n3.SRTF\n4.Round Robin\n5.Multinível\n\ta.primeiro nível - RR\n\tb.segundo nível\n6 - SAIR - FCFS")
	fmt.Scanln(&algoritmo)

	if algoritmo == 6{ 
		break
	}	
	fmt.Println("Inserção de processos:\n 1.use M para manual\n 2.Use L para ler arquivos\n  3. Use S para Sair")
	fmt.Scanln(&modo)

	if modo == "S"{
		break
	}

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
			burst_total+=burst
			processos = append(processos, Processo{nome,burst,burst,tcheg,0,0,0,0,0,0})
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
		processos = sjf(processos)
		_ = processos
	} else if algoritmo == 3{
		processos = srtf(processos, burst_total)
		_ = processos
	} else if algoritmo == 4{
		var quantum int
		fmt.Println("Digite o quantum:",)
		fmt.Scanln(&quantum)
		processos = rr(processos, burst_total, quantum)
		_ = processos
	} else if algoritmo == 5{
		var quantum int
		fmt.Println("Digite o quantum:",)
		fmt.Scanln(&quantum)
		processos = mn(processos, burst_total, quantum)
		_ = processos
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
}

func fcfs(processos []Processo) []Processo {
	var time int
	time = 0
	// Simulando processamentos
	for i := 0; i < len(processos); i++  {
		
		processos[i].tespera = time - processos[i].tcheg
		time += processos[i].burst
		processos[i].tretorno = time - processos[i].tcheg
	}

	return processos
}

func sjf(processos []Processo) []Processo {
	var time int
	time = 0
	// Ordenando o vetor de acordo com tcheg e burst
	increasingTime := func(p1, p2 *Processo) bool {
		return p1.tcheg < p2.tcheg
	}
	increasingBurst := func(p1, p2 *Processo) bool {
		return p1.burst < p2.burst
	}
	OrderedBy(increasingTime, increasingBurst).Sort(processos)
	// Simulando Processamento
	for i := 0; i < len(processos); i++  {	
		processos[i].tespera = time - processos[i].tcheg
		time += processos[i].burst
		processos[i].tretorno = time - processos[i].tcheg
	}

	return processos
}

func srtf(processos []Processo, burst_total int) []Processo {
	completos := 0
	time := 0
	menorTrest := 60000
    menorPosicao := 0 
   	bandeira := false
   	tfim :=0
   	n := len(processos)
   	// Encontrando o tempo de espera
   	for ;completos != n; {
   		//  Encontrando o processo com menor tempo restante no tempo atual 
       for j:=0; j < n ; j++ {
       	if (processos[j].tcheg <= time) && (processos[j].trest < menorTrest) && processos[j].trest > 0 {
       		menorTrest = processos[j].trest
       		menorPosicao = j
       		bandeira = true
       	}
       }

       if bandeira == false {
       		time++
       		continue
       }
      // Executando o processo uma unidade de tempo
       processos[menorPosicao].trest--

      // Atualizando menor burst
       menorTrest = processos[menorPosicao].trest
       if(menorTrest == 0){
       		menorTrest = 60000
       }

       // Caso o processo seja concluído
       if(processos[menorPosicao].trest == 0 ) {
       	  // Incrementando completos
       	  completos++
       	  bandeira = false
       	  // encontrando tempo atual
       	  tfim = time + 1

       	  // Calculando tempo de espera
       	  processos[menorPosicao].tespera = tfim - processos[menorPosicao].burst - processos[menorPosicao].tcheg

       	  if(processos[menorPosicao].tespera < 0){
       	  	processos[menorPosicao].tespera = 0
       	  }

       }
       // Incrementando tempo
       time++
   	}

   	// Calculando t de retorno
   	for i:=0; i<n; i ++ {
   		processos[i].tretorno = processos[i].burst + processos[i].tespera
   	}

   	return processos
}

func rr(processos []Processo, burst_total int, quantum int) []Processo {
	time := 0
	pos := 0
	tamanho := len(processos)

	for ;time < burst_total; {
		// p := processos[pos]
		// fmt.Printf("%d %d %d\n", pos, processos[pos].trest, time)
		if processos[pos].trest == 0 {
			pos = (pos + 1) % tamanho
			continue
		}

		processos[pos].tespera =  processos[pos].tespera + (time - processos[pos].tMod)

		if processos[pos].trest < quantum {
			time += processos[pos].trest
			processos[pos].trest -= processos[pos].trest
		} else 	{
			time +=quantum
			processos[pos].trest -= quantum
		}

		if processos[pos].trest == 0	{
			processos[pos].tretorno =  time - processos[pos].tcheg
		}

		processos[pos].tMod = time
		pos = (pos + 1) % tamanho
	}
	return processos
}


func mn(processos []Processo, burst_total int, quantum int) []Processo {
	time := 0
	tamanho := len(processos)
	// Round Robin
	for pos := 0; pos < tamanho; pos++  {
		if processos[pos].trest == 0 {
			continue
		}

		processos[pos].tespera =  processos[pos].tespera + (time - processos[pos].tMod)

		if processos[pos].trest < quantum {
			time += processos[pos].trest
			processos[pos].trest -= processos[pos].trest
		} else 	{
			time +=quantum
			processos[pos].trest -= quantum
		}

		if processos[pos].trest == 0	{
			processos[pos].tretorno =  time - processos[pos].tcheg
		}

		processos[pos].tMod = time	
	}

	time = 0
	// First Come first Served
	for i := 0; i < tamanho; i++ {
		if processos[i].trest == 0 {
			continue
		}

		processos[i].tespera += time - processos[i].tcheg
		time += processos[i].burst
		processos[i].tretorno += time - processos[i].tcheg
	}

	return processos

}