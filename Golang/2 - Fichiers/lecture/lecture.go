package lecture

import (
	"fmt"
	"os"
)

func LectureFichier(nomFichier string) {
	// Lecture du contenu du fichier
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Erreur : impossible de lire le fichier.")
	} else {
		fmt.Println(string(contenu))
	}
}
