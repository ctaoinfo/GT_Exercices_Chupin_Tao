package main

import (
	"fichier/ecriture"
	"fichier/lecture"
	"fmt"
)

func main() {
	var choix int
	var nomFichier string
	var texte string

	for {
		fmt.Println("Pour les tests utiliser le fichier \"fichier.txt\"")
		fmt.Println()
		fmt.Println("1. Récupérer tout le texte contenu dans un fichier")
		fmt.Println("2. Ajouter du texte dans un fichier")
		fmt.Println("3. Supprimer tout le contenu d'un fichier")
		fmt.Println("4. Remplacer le contenu d'un fichier par du texte donné par l'utilisateur")
		fmt.Println("5. Quitter")
		fmt.Print("Choix : ")

		fmt.Scanln(&choix)

		switch choix {
		case 1:
			fmt.Print("Nom du fichier : ")
			_, err := fmt.Scanln(&nomFichier)
			if err != nil {
				return
			}
			lecture.LectureFichier(nomFichier)

		case 2:
			fmt.Print("Nom du fichier : ")
			_, err := fmt.Scanln(&nomFichier)
			if err != nil {
				return
			}
			fmt.Print("Texte à ajouter : ")
			_, err = fmt.Scanln(&texte)
			if err != nil {
				return
			}
			ecriture.EcritureSuiv(nomFichier, texte)

		case 3:
			fmt.Print("Nom du fichier : ")
			_, err := fmt.Scanln(&nomFichier)
			if err != nil {
				return
			}
			ecriture.SuppressionTotale(nomFichier)

		case 4:
			fmt.Print("Nom du fichier : ")
			_, err := fmt.Scanln(&nomFichier)
			if err != nil {
				return
			}
			fmt.Print("Nouveau contenu : ")
			_, err = fmt.Scanln(&texte)
			if err != nil {
				return
			}
			ecriture.EcritureTotale(nomFichier, texte)

		case 5:
			fmt.Println("A la prochaine !")
			return

		default:
			fmt.Println("Choix invalide.")
		}
		fmt.Println()
	}
}
