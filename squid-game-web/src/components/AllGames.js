import React from 'react'
import Card from './Card'
import socketIOClient from "socket.io-client";

const API = "http://35.225.182.66:4001/api/games";
const ENDPOINT = "http://35.225.182.66:4001";

class AllGames extends React.Component {
    
    state = {
        games: [

        ]
    }

    fetchGames = () => {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(gamesList => {
            this.setState({ games: gamesList });
        });
    }

    componentDidMount() {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(gamesList => {
            console.log(gamesList);
            this.setState({ games: gamesList });
        }).catch(err=>{
            console.log(err)
        });

        const socket = socketIOClient(ENDPOINT);

        /*socket.on("FromAPI", data => {
        //('hello geeks!');
        console.log(data);
        //console.log('hola');
        //setResponse(data);
        });*/
      
        socket.on("NewGamesNotify", data => {
            console.log('the db has changed!')
            this.fetchGames();
        //toast.info('New Tweets have been pushed!', 
        //{position: toast.POSITION.TOP_CENTER});
        });
    }

    render() {
        return (
            
                <div className="row">
                    {
                        this.state.games.map(game=> (
                            <div className="col-md-4 p-4" key={game._id}>
                                <Card Identificador={game.ID} 
                                Juego={game.juego} 
                                Ganador={game.max}/>
                            </div>
                        ))
                    }
                </div>
            
        )
    }
}

export default AllGames
