package lab

type Tableau struct {
	Tour int
	Tab  [6][7]int
	Colonne int
	Ligne int
}

func Creer_plateau() Tableau {
	var plateau Tableau
	plateau.Tour = 1
	for i, ligne := range plateau.Tab {
		for v, _ := range ligne {
			plateau.Tab[i][v] = 0
		}
	}
	return plateau
}
