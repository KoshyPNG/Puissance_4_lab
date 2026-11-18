package lab

func (p *Tableau) Gagnez() bool {
	if p.Gagne_ligne() || p.Gagne_colonne() || p.Gagne_diag() {
		return true
	}
	return false
}

func (p *Tableau) Gagne_ligne() bool {
	var max int
	var min int
	var player int
	if p.Tour%2 == 1 {
		player = 1
	} else {
		player = 2
	}

	if p.Ligne >= 5{
		min = 6 - p.Ligne
		max = 7
	} else {
		min = 0
		max = p.Ligne + 3
	}
	combo := 0
	quatre := min+3
	for min < max {
		i := min
		for i < quatre {
			if p.Tab[p.Ligne][min] == player {
				combo++
			}
			if combo == 4 {
				return true
			}
		min++
		combo = 0
		quatre = min+3
		}
	}
	return false
}
