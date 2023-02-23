package allumette

import (
	"fmt"
	"strconv"
)

type gameAllumette struct {
	nbAllum     int
	nbAllumAct  int
	actualRound int
}

func Allumette() {
	g1 := gameAllumette{}
	g1.nbAllum = 10
	g1.nbAllumAct = 10
	g1.actualRound = 10

	game := true
	var inputInt int
	var input string
	displayAllum := make([]string, g1.nbAllum)

	var inputAllumInt int
	var inputAllum string

	choiceAllum := true

	for choiceAllum {
		fmt.Println("Entrez une valeur personnalisé du nombre d'allumette (4 minimum - 99 maximum)")
		fmt.Println("-Entrer 'e' pour prendre la valeur par défault (10)")
		fmt.Print("-> ")
		_, err := fmt.Scanln(&inputAllum)
		if err != nil {
			return
		}
		if inputAllum == "e" {
			choiceAllum = false
			continue
		} else {
			if checkNum(inputAllum) {
				inputAllumInt, _ = strconv.Atoi(inputAllum)
				if inputAllumInt >= 4 && inputAllumInt <= 99 {
					g1.nbAllum = inputAllumInt
					choiceAllum = false
					break
				} else {
					fmt.Println("Valeur non valide.")
				}
			} else {
				fmt.Println("Valeur non valide.")
			}
		}
	}

	for g1.nbAllumAct != 0 && game {
		if g1.actualRound%2 == 0 {
			fmt.Println("Joueur 1 :")
		} else {
			fmt.Println("Joueur 2 :")
		}

		for i := 0; i < g1.nbAllumAct; i++ {
			displayAllum[i] = "|"
			fmt.Print(displayAllum[i], " ")
		}
		fmt.Println("\nPlus que", g1.nbAllumAct, "allumettes.\n")

		fmt.Println("Entrez un entier compris entre 1 & 3")
		fmt.Print("-> ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		if len(input) > 2 {
			input = "a"
		} else if !checkNum(input) {
			fmt.Println("Valeur non valide.")
		} else {
			inputInt, _ = strconv.Atoi(input)

			if inputInt < 1 || inputInt > 3 {
				fmt.Println("Valeur non valide.")
			} else {
				if inputInt == 1 {
					g1.nbAllumAct--
				} else if inputInt == 2 {
					g1.nbAllumAct -= 2
				} else {
					g1.nbAllumAct -= 3
				}
				fmt.Println()
				if g1.nbAllumAct == 0 {
					if g1.actualRound%2 != 0 {
						fmt.Println("Joueur 1 as gagné")
					} else {
						fmt.Println("Joueur 2 as gagné")
					}
				}
				g1.actualRound--
			}
		}
	}
}

func checkNum(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
