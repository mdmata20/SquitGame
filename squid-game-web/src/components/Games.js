import React from 'react'
import Card from './Card'
import './games.css'

const API = "http://localhost:4001/api/games/top10";

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
    }

    fetchTweets = () => {
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
                                <Card Identificador={game.identificador} 
                                Juego={game.juego} 
                                Ganador={game.ganador}/>
                            </div>
                        ))
                    }
                </div>
            </div>
            
        )
    }
}

export default Games
