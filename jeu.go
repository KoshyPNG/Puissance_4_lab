package lab

import "fmt"

func Tour_de_jeu(p *Tableau) {
	var colonne int
	for !Plein() {
		fmt.Scan(&colonne)
		if colonne < 1 || colonne > 7 {
			fmt.Println("Colonne invalide, veuillez entrer un numero entre 1 et 7")
			continue
		}
		p.Colonne = colonne-1
		if p.Tour%2 != 0 {
			p.Ajouter(1)
		} else {
			p.Ajouter(2)
		}
		if p.Ligne != -1 {
			p.Tour++
		}
		
		p.Ligne = -1
		p.Afficher_plateau()
	}
}

func (p *Tableau) Ajouter(player int) {
	col := p.Colonne
	for i:=5; i>=0; i-- {

		if p.Tab[i][col-1] == 0 {
			p.Tab[i][col] = player
			p.Ligne = i
			fmt.Println("Pion ajouter colonne:",col+1)
			return

		} else if i == 0 {
			fmt.Println("Pion ne peut pas etre ajouter colonne:",col+1)
			return
		}
	}
}