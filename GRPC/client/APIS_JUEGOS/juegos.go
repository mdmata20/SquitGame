
package prueba


import (
	"fmt"
	"math/rand"
	"time"
)




func juego1(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int
	
	fmt.Println("--------- juego 1 ----------")

	fmt.Println("jugadores jugando")

	for i := 1; i <= jugadores; i++ {
		ganador = append(ganador,i)
	}

	fmt.Println("El ganador es: ")

	ganador1 := rand.Intn(len(ganador))
	Campeon := ganador[ganador1]
	

	fmt.Println(Campeon)
	return Campeon
}


func juego2(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int
	fmt.Println("--------- juego 2 ----------")

	fmt.Println("jugadores jugando")

	for i := 1; i <= jugadores; i++ {
		
		if(i%2 == 1){
			ganador = append(ganador,i)
			
		}

	}

	fmt.Println("El ganador es: ")
	ganador2 := rand.Intn(len(ganador))
	campeon2 := ganador[ganador2]

	fmt.Println(campeon2)

	return campeon2

}


func juego3(jugadores int) int {
	rand.Seed(time.Now().UnixNano())

	var ganador []int
	fmt.Println("--------- juego 2 ----------")

	fmt.Println("jugadores jugando")

	for i := 1; i <= jugadores; i++ {
		
		if(i%2 == 0){
			ganador = append(ganador,i)
			
		}

	}

	fmt.Println("El ganador es: ")
	ganador2 := rand.Intn(len(ganador))
	campeon2 := ganador[ganador2]

	fmt.Println(campeon2)

	return campeon2

}




func prueba() {
	juego3(20)
	//juego2(10)
	//juego1(20)
}

