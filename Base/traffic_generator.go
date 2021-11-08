package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type juegoStruct struct {
	Id    int    `json:"Id"`
	Juego string `json:"Juego"`
	Max   int    `json:"Max"`
}

type decide struct {
	id   int
	name string
}

func endpoint(game int, gamename string, players int, rungames int, concurrence int, timeout int) {
	rand.Seed(time.Now().UnixNano())
	var mem []juegoStruct
	cont_corridas := rungames

	cont_tiempo := timeout * 60

	for {
		time.Sleep(1 * time.Second)

		maxJugadores := rand.Intn(players)

		cuerpo := juegoStruct{
			Id:    game,
			Juego: gamename,
			Max:   maxJugadores,
		}

		mem = append(mem, cuerpo)

		cont_tiempo = cont_tiempo - 1
		cont_corridas = cont_corridas - concurrence

		if cont_tiempo == 0 || cont_corridas == 0 {
			break
		}

	}

	endpointMax := len(mem)

	partes := endpointMax / concurrence
	//partes2 := int(partes)
	var wg sync.WaitGroup
	for i := 0; i < concurrence; i++ {
		inicio := i * partes
		fmt.Println("INICIO ", inicio)
		fin := inicio + partes

		fmt.Println("FIN ", fin)
		parte := mem[inicio:fin]
		fmt.Println("PARTE ", parte)
		wg.Add(1)

		go func(jsonJ []juegoStruct) {
			defer wg.Done()

			for _, n := range jsonJ {
				time.Sleep(500 * time.Millisecond)
				fmt.Println("===HACER POST AQUI============ \n \n ")
				json_p, err := json.Marshal(n)
				if err != nil {
					fmt.Println("Error en MArshal")
				}
				resp, err := http.Post("http://localhost:8080/grpc", "application/json", bytes.NewBuffer(json_p))
				if err != nil {
					fmt.Println("Error de conexión")
				}

				defer resp.Body.Close()

				if resp.StatusCode == http.StatusCreated {
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						panic(err)
					}

					jsonStr := string(body)
					fmt.Println("Response: ", jsonStr)
				} else {
					fmt.Println("Fail: ", resp.Status)
				}
			}

			time.Sleep(1 * time.Second)

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
	reg4, err := regexp.Compile("[0-9] \\| ")

	processed := reg.ReplaceAllString(gamename1, "")
	doublep := reg2.ReplaceAllString(processed, ",")
	game := strings.Split(doublep, ",")

	ident := reg4.ReplaceAllString(gamename1, "")
	ident2 := reg2.ReplaceAllString(ident, ",")
	ident3 := strings.Split(ident2, ",")

	timeout := reg3.ReplaceAllString(time, "")

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

	//game array num
	// ident 3 array juego

	randomGame := rand.Intn(len(game))
	pick := game[randomGame] //! PARAMETRO 1
	pick2, err := strconv.Atoi(pick)
	pickG := ident3[randomGame]

	//fmt.Println(pick2, pickG, jugadores, corridas, concurrencia, tiempo)
	endpoint(pick2, pickG, jugadores, corridas, concurrencia, tiempo)

}
