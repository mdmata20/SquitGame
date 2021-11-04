import React from 'react'
import BestPlayer from './BestPlayer'
import './games.css'

const API = "http://localhost:4001/api/games/top10players";

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
