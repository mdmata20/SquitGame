package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func endpoint(game string, players int, rungames int, concurrence int, timeout int) {
	rand.Seed(time.Now().UnixNano())

	var conc []string
	cont_corridas := rungames

	cont_tiempo := timeout * 60

	for {
		time.Sleep(1 * time.Second)

		maxJugadores := rand.Intn(players)
		jugadores := strconv.Itoa(maxJugadores)
		endpoint := "/game/" + game + "/gamename/Random/players/" + jugadores
		conc = append(conc, endpoint)
		fmt.Println(cont_tiempo)
		cont_tiempo = cont_tiempo - 1
		cont_corridas = cont_corridas - concurrence

		if cont_tiempo == 0 || cont_corridas == 0 {
			break
		}

	}
	fmt.Println(conc)
	fmt.Println("")
	fmt.Println("")
	var wg sync.WaitGroup

	endpointMax := len(conc)
	partes := endpointMax / concurrence

	for i := 0; i < concurrence; i++ {
		inicio := i * partes
		fin := inicio + partes

		parte := conc[inicio:fin]

		wg.Add(1)

		go func(endpoints []string) {
			defer wg.Done()
			//var accesos string

			for _, n := range endpoints {
				fmt.Println("==========================")
				fmt.Println(endpoints)
				fmt.Println("****" + n + "***")
				fmt.Println("aqui tenemos que enviar")
				fmt.Println("==========================")
			}

			//	atomic.AddInt64(&int64(cont_tiempo), int64(concurrence))
		}(parte)
	}

	wg.Wait()

}

func main() {
	rand.Seed(time.Now().UnixNano())

	gamename1 := os.Args[3]
	players := os.Args[5]
	rungames := os.Args[7]
	concurrence := os.Args[9]
	time := os.Args[11]

	reg, err := regexp.Compile("( \\| [a-zA-Z]+[0-9]?)")
	reg2, err2 := regexp.Compile(" \\| ")
	reg3, err3 := regexp.Compile("m")
	processed := reg.ReplaceAllString(gamename1, "")
	doublep := reg2.ReplaceAllString(processed, ",")
	timeout := reg3.ReplaceAllString(time, "")
	game := strings.Split(doublep, ",")

	jugadores, err1 := strconv.Atoi(players)
	corridas, err1 := strconv.Atoi(rungames)
	tiempo, err1 := strconv.Atoi(timeout)
	concurrencia, err1 := strconv.Atoi(concurrence)

	if err != nil {
		log.Fatal(err)
	} else if err2 != nil {
		log.Fatal(err2)
	} else if err3 != nil {
		log.Fatal(err3)
	} else if err1 != nil {
		log.Fatal(err1)
	}

	randomGame := rand.Intn(len(game))
	pick := game[randomGame] //! PARAMETRO 1

	//fmt.Println(pick, jugadores, corridas, concurrence, tiempo)
	endpoint(pick, jugadores, corridas, concurrencia, tiempo)

}
