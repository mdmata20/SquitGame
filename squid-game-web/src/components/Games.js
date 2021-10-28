import React from 'react'
import Card from './Card'
import './games.css'

const API = "https://20211002t224125-dot-vibrant-tree-324821.appspot.com/api/tweets";

class Games extends React.Component {
    state = {
        games: [

            {
                "_id" : 1,
                "request_number":30001,
                "game":17,
                "gamename": "Green light Red light",
                "winner": "7",
                "players": 456,
                "worker": "RabbitMQ"
            },
            {
                "_id" : 2,
                "request_number":30002,
                "game":17,
                "gamename": "Green light Red light",
                "winner": "7",
                "players": 456,
                "worker": "RabbitMQ"
            },
            {
                "_id" : 3,
                "request_number":30003,
                "game":17,
                "gamename": "Green light Red light",
                "winner": "7",
                "players": 456,
                "worker": "RabbitMQ"
            },
            {
                "_id" : 4,
                "request_number":30004,
                "game":17,
                "gamename": "Green light Red light",
                "winner": "7",
                "players": 456,
                "worker": "RabbitMQ"
            }
        ]
    }

    /*componentDidMount() {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(tweetsList => {
            this.setState({ tweets: tweetsList });
        });
    }

    fetchTweets = () => {
        fetch(`${API}`)
        .then((response) => response.json())
        .then(tweetsList => {
            this.setState({ tweets: tweetsList });
        });
    }*/
    
    render() {
        return (
            
            <div className="container d-flex flex-column justify-content-center align-items-center">
                <div className="row">
                    {
                        this.state.games.map(game=> (
                            <div className="col-md-4" key={game._id}>
                                <Card RequestNumber={game.request_number} 
                                GameID={game.game} 
                                GameName={game.gamename}
                                GameWinner={game.winner} 
                                PlayersAmount={game.players} 
                                Worker={game.worker}/>
                            </div>
                        ))
                    }
                </div>
            </div>
            
        )
    }
}

export default Games
