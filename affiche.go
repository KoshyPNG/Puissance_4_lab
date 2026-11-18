package lab

import "fmt"

func (p *Tableau) Afficher_plateau() {
	fmt.Println("TOUR:  ", p.Tour)
	for _, ligne := range p.Tab {
		fmt.Print("|")
		for _, val := range ligne {
			switch val {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("X")
			case 2:
				fmt.Print("O")
			}
			fmt.Print("|")
		}
		fmt.Println(" ")
	}
	for i := 1; i < 8; i++ {
		fmt.Print(" ")
		fmt.Print(i)
	}
	fmt.Println(" ")
	if p.Tour%2 == 1 {
		fmt.Print("Joueur 1 de jouer")
	} else {
		fmt.Print("Joueur 2 de jouer")
	}
	 fmt.Println(" ecrire le numero de la colonne ou vous voulez jouer")
}