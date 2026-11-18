package main

import (
	"lab"
)



func main() {
	plateau := lab.Creer_plateau()
	plateau.Afficher_plateau()
	lab.Tour_de_jeu(&plateau)
	
}

