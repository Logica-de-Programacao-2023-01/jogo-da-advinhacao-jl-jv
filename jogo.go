package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var tentativa, tentativas int
	var listatentativas []int
	var resposta string
	rand.Seed(time.Now().UnixNano())
	for resposta != "não" {
		numeroaleatorio := rand.Intn(100) + 1
		fmt.Println("Bem vindo ao jogo da advinhação!")
		fmt.Print("Digite um número de 1 a 100: ")
		fmt.Scan(&tentativa)
		tentativas = 1
		for tentativa != numeroaleatorio {
			if tentativa < numeroaleatorio {
				fmt.Printf("O número é maior que %d\n", tentativa)
				fmt.Print("Digite um número de 1 a 100: ")
				fmt.Scan(&tentativa)
				tentativas++
			} else if tentativa > numeroaleatorio {
				fmt.Printf("O número é menor que %d\n", tentativa)
				fmt.Print("Digite um número de 1 a 100: ")
				fmt.Scan(&tentativa)

				tentativas++
			}
		}
		listatentativas = append(listatentativas, tentativas)
		fmt.Println("Parabéns você acertou!")
		fmt.Printf("Você utilizou %d tentativas\n", tentativas)
		fmt.Print("Deseja jogar denovo?(sim/não): ")
		fmt.Scan(&resposta)
	}
	for x := 0; x < len(listatentativas); x++ {
		fmt.Printf("\nNa jogada %d você utilizou %d tentativas", x+1, listatentativas[x])
	}
}
