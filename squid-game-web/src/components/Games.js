import React from 'react'
import Card from './Card'
import socketIOClient from "socket.io-client";
import './games.css'

const API = "http://localhost:4001/api/games/top10";
const ENDPOINT = "http://localhost:4001";

class Games extends React.Component {
    state = {
        games: [

        ]
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
      
        socket.on("NewGamesNotify", data => {
            console.log('the db has changed!')
            this.fetchLast10Games();
        //toast.info('New Tweets have been pushed!', 
        //{position: toast.POSITION.TOP_CENTER});
        });
    }

    fetchLast10Games = () => {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(gamesList => {
            this.setState({ games: gamesList });
        });
    }
    
    render() {
        return (
            
            <div className="container d-flex flex-column justify-content-center align-items-center" style={{padding: 20}}>
                <div className="row">
                    {
                        this.state.games.map(game=> (
                            <div className="col-md-4" key={game._id}>
                                <Card Identificador={game.ID} 
                                Juego={game.juego} 
                                Ganador={game.max}/>
                            </div>
                        ))
                    }
                </div>
            </div>
            
        )
    }
}

export default Games
