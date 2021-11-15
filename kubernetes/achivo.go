func JuegoMongo1(w http.ResponseWriter, r http.Request) {
    w.Header().Set("content-type", "application/json")

    var body map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
        fmt.Println(err)
    }
    body["way"] = "RabbitMQ"
    data, err := json.Marshal(body)

	//observar la info
    nuevo := string(data)
    fmt.Println(nuevo)

    var juego JuegoMongo
     = json.NewDecoder(r.Body).Decode(&juego)

    clintOptions := options.Client().ApplyURI("mongodb://34.125.189.71:27017")
    client, err := mongo.Connect(context.TODO(), clintOptions)
    if err != nil {
        fmt.Println("Mongo.connect() ERROR: ", err)
        os.Exit(1)
    }
    col := client.Database("squid-game").Collection("games")

    fmt.Println("ClientOptions Type: ", reflect.TypeOf(clintOptions), "\n")

    ctx,  := context.WithTimeout(context.Background(), 15time.Second)

    fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

    result, insertErr := col.InsertOne(ctx, bson.D{
		{Key: "ID", Value: body["ID"]},
		{Key: "juego", Value: body["juego"]},
		{Key: "max", Value: body["max"]},
		{Key: "players", Value: body["players"]},
		{Key: "worker", Value: body["worker"]}},
	})
    if insertErr != nil {
        fmt.Println("InsertONE Error:", insertErr)
        os.Exit(1)
    } else {
        fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
        fmt.Println("InsertOne() api result type: ", result)

        newID := result.InsertedID
        fmt.Println("InsertedOne(), newID", newID)
        fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

    }

    json.NewEncoder(w).Encode(result)

    //fmt.Fprintf(w, "Se Funcion")
}


func (s *server) RegGame(ctx context.Context, in gamepb.GameRequest) (gamepb.GameResponse, error) {

    var winner int
    result := ""

    id := in.GetGame().GetId()
    juego := in.GetGame().GetJuego()
    max := in.GetGame().GetMax()

    desicion := int(id)
    maximo := int(max)

    if desicion == 1 {
        winner = juego1(maximo)
    } else if desicion == 2 {
        winner = juego2(maximo)
    } else if desicion == 3 {
        winner = juego3(maximo)
    } else {
        winner = 0
    }

    //    identificador := strconv.FormatInt(id, 10)
    //ganador := strconv.Itoa(winner)

    byt := []byte(`{"ID":0}`)

    var dat map[string]interface{}
    err := json.Unmarshal(byt, &dat); 

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(dat)

    dat["ID"] = desicion
    dat["juego"] = juego
    dat["max"] = winner
    dat["players"] = maximo
    dat["worker"] = "kafka"

    data, err := json.Marshal(dat)

    peticion, _ := json.Marshal(gameStruct{
        Id:      desicion,
        juego:   juego,
        max: winner,
        players: maximo,
        worker:  "kafka",

    })

    petition := bytes.NewBuffer(data)

    response, err := http.Post("http://34.122.191.135:2062/", "application/json", petition)
    if err != nil {
        log.Fatalln("Error sending info", err)
    }
    //response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }
    // SE NECESITA OTRO ENDPOINT
    result += string(body)
    fmt.Println(result + "------------")
    //result := " ID: " + identificador + " Juego: " + juego + " Ganador!: " + ganador

    res := &gamepb.GameResponse{
        Result: result,
    }

    return res, nil
}