import React from 'react'

function Card({Identificador, Juego, Ganador}) {
    return (
        <div className="card bg-dark">
            <div className="card-header text-light d-flex justify-content-between">
                <h5>Juego: {Juego}</h5>
            </div>
            <div className="card-body text-light">
                    <p>Ganador: {Ganador}</p>
            </div>
        </div>
    )
}

export default Card
