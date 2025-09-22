package ynov

import "fmt"

/*SI ne < 0 OU ne > 20
ALORS AFFICHER ("erreur:la note doit etre entre 0 et 20")
SINON SI ne <=7 ET ne >=1
ALORS AFFICHER ("echec")
SINON SI ne <= 13 ET ne >= 8
ALORS AFFICHER ("passable")
SINON SI ne == 14
ALORS AFFICHER ("bien")
SINON SI ne <= 20 ET ne >=15
ALORS AFFICHER ("excellent")
EXIT()*/

func main() {

	var ne int
	fmt.Scan(&ne)

	if ne < 0 || ne > 20 {
		fmt.Println("erreur:la note doit etre entre 0 et 20")
	} else if ne == 14 {
		fmt.Println("bien")
	} else if ne <= 13 && ne >= 8 {
		fmt.Println("passable")

	} else if ne <= 7 && ne >= 1 {
		fmt.Println("echec")

	} else if ne <= 20 && ne >= 15 {
		fmt.Println("excellent")
	}

}
