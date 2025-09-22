package ynov

import "fmt"


/*ne == 15
ALORS AFFICHER ("bien")
SINON SI ne < 15 ET ne > 8 
ALORS AFFICHER ("passable")
SINON SI ne <8 ET ne >1 
ALORS AFFICHER ("echec")
SINON SI ne < 20 ET ne >15 
ALORS AFFICHR ("excellent")
EXIT()*/


if ne == 15 {
	fmt.Println ("bien")
}else{
	if ne <15 && ne 8{
		fmt.Println("passable")
	}
}else{
	if ne < 8 && ne > 1{
		fmt.Println("echec")
	}
}else{
	if ne < 20 && ne > 15{
		fmt.Println("excellent")
	}
}
