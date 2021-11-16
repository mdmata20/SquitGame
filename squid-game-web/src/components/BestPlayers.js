import React from 'react'
import BestPlayer from './BestPlayer'
import socketIOClient from "socket.io-client";
import './games.css'

const API = "http://35.225.182.66:4001/api/games/top10players";
const ENDPOINT = "http://35.225.182.66:4001";

class BestPlayers extends React.Component {
    state = {
        players: [

        ]
    }

    componentDidMount() {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(playersList => {
            console.log(playersList);
            this.setState({ players: playersList });
        }).catch(err=>{
            console.log(err)
        });

        const socket = socketIOClient(ENDPOINT);
      
        socket.on("NewGamesNotify", data => {
            console.log('the db has changed!')
            this.fetchBest10Players();
        //toast.info('New Tweets have been pushed!', 
        //{position: toast.POSITION.TOP_CENTER});
        });
    }

    fetchBest10Players = () => {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(playersList => {
            console.log(playersList);
            this.setState({ players: playersList });
        }).catch(err=>{
            console.log(err)
        });
    }

    render(){
        return (
            <div className="container d-flex flex-column justify-content-center align-items-center" style={{padding: 20}}>
                <div className="row">
                    {
                        this.state.players.map(player=> (
                            <div className="col-md-4 p-2" key={player._id} >
                                <BestPlayer IdJugador={player._id} 
                                Conteo={player.count} />
                            </div>
                        ))
                    }
                </div>
            </div>
        )
    }
}

export default BestPlayers
