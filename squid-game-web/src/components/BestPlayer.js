import React from 'react'
import { Link } from 'react-router-dom'

function BestPlayer({IdJugador, Conteo}) {
    return (
        <div className="card bg-dark">
            <div className="card-header text-light d-flex justify-content-between" >
                <b className="card-title">Jugador: {IdJugador}</b>
                <Link to={"/viewUser/" + IdJugador} className="btn btn-secondary">View</Link>
            </div>
            <div className="card-body text-light">
                <b className="card-title">Juegos Ganados: {Conteo}</b>
            </div>
        </div>
    )
}

export default BestPlayer
