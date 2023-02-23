package ecriture

import (
	"fmt"
	"os"
)

// Fonction de rajout de texte a la fin du fichier
func EcritureSuiv(nomFichier string, texte string) {
	f, err := os.OpenFile(nomFichier, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erreur : impossible d'ouvrir le fichier.")
	} else {
		_, err := f.WriteString(texte)
		if err != nil {
			return
		}
		err = f.Close()
		if err != nil {
			return
		}
	}
}

// Fonction pour supprimer le contenu du fichier
func SuppressionTotale(nomFichier string) {
	err := os.WriteFile(nomFichier, []byte{}, 0644)
	if err != nil {
		fmt.Println("Erreur : impossible de supprimer le contenu du fichier.")
	} else {
		fmt.Println("Contenu du fichier supprimé.")
	}
}

// Fonction de remplacement du contenu d'un fichier
func EcritureTotale(nomFichier string, texte string) {
	err := os.WriteFile(nomFichier, []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur : impossible d'écrire dans le fichier.")
	} else {
		fmt.Println("Nouveau contenu écrit.")
	}
}
