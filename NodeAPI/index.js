//const { PubSub } = require('@google-cloud/pubsub');
const axios = require('axios');
//const mysqlConnection = require('./src/conexion');
const socket = require('socket.io')
const http = require("http");
const cors = require('cors')

var games = require('./src/routes/games');

// Acá escribimos la suscripción que creamos en Google Pub/Sub
//const SUB_NAME = 'projects/vibrant-tree-324821/subscriptions/twitterSubscription';

// Cantidad de segundos que estara prendido nuestro listener
// Solo para efectos practicos, realmente esto debería estar escuchando en todo momento!
const TIMEOUT = process.env.TIMEOUT || 18000;

// Crear un nuevo cliente de pubsub
//const client = new PubSub();

// En este array guardaremos nuestros datos
const messages = [];

var express = require('express'),
  app = express(),
  port = process.env.PORT || 4001;


//app.use(cors())

//app.use(express.urlencoded({ extended: true }));
//app.use(express.json());

app.use('/', games);

//app.listen(port);

//var server = app.listen(port);

const server = http.createServer(app);

const io = socket(server);

let interval

/*io.on("connection", (socket) => {
    console.log("New client connected");
    if (interval) {
        clearInterval(interval);
    }
    interval = setInterval(() => getApiAndEmit(socket,"hola"), 1000);
    socket.on("disconnect", () => {
        console.log("Client disconnected");
        clearInterval(interval);
    });
});
  
const getApiAndEmit = (socket, jsonMsg) => {
    const response = new Date();
    // Emitting a new message. Will be consumed by the client
    //console.log(response)
    //console.log(jsonMsg)
    if(messages.length > 0){
        socket.emit("NewGamesNotify", messages.pop());
    }
    socket.emit("FromAPI", response);
    
};

// Esta funcion se utilizara para leer un mensaje
// Se activara cuando se dispare el evento "message" del subscriber
const messageReader = async message => {

    console.log('¡Mensaje recibido!');
    console.log(`${message.id} - ${message.data}`);
    console.table(message.attributes);
    var jsonArray;

    messages.push({ msg: String(message.data), id: message.id, ...message.attributes });

    // Con esto marcamos el acknowledgement de que recibimos el mensaje
    // Si no marcamos esto, los mensajes se nos seguirán enviando aunque ya los hayamos leído!
    message.ack();

    try {
        console.log(`Agregando mensaje al servidor...`);
        const jsonMessage = JSON.parse(message.data) || {};
        const request_body = { guardados: jsonMessage.Guardados || jsonMessage.guardados || "unknown", 
                               api: jsonMessage.Api || jsonMessage.api || "unknown",
                               tiempoDeCarga: jsonMessage.tiempoDeCarga || jsonMessage.TiempoDeCarga || "unknown",
                               bd: jsonMessage.bd || jsonMessage.Bd || "unknown" };
        
        const topLimit = jsonMessage.guardados || jsonMessage.Guardados
    }
    catch (e) {
        console.log(`Error al realizar POST ${e.message}`);
    }
};

// Empezamos nuestro manejador de notificaciones
const notificationListener = () => {

    // Creamos un subscriptor
    // Pasamos el nombre de nuestro subscriptor (que encontramos en Google Cloud)
    const sub = client.subscription(SUB_NAME);

    // Conectar el evento "message" al lector de mensajes
    sub.on('message', messageReader);

    console.log("Esperando por nuevos mensajes...");

    setTimeout(() => {
        sub.removeListener('message', messageReader);

        if (messages.length > 0) {
            console.log(`${messages.length} mensajes recibidos: `);
            console.log("---------");
            console.table(messages);
        }
        else {
            console.log("No hubo ningún mensaje durante este tiempo.")
        }

    }, TIMEOUT * 1000);
    
};*/

//console.log(`Iniciando Subscriber, deteniendolo en ${TIMEOUT} segundos...`);

// Empezar a escuchar los mensajes
//notificationListener();

//app.listen(port);
server.listen(port, () => console.log(`Listening on port ${port}`));
//console.log('Node JS API started on: ' + port);







