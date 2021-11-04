import React from 'react'

function BestPlayer({IdJugador, Conteo}) {
    return (
        <div className="card bg-dark">
            <div className="card-body text-light">
                <div>
                    <table>
                        <tbody>
                            <tr>
                                <td>
                                    <b className="card-title">Jugador:</b>
                                </td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td>
                                    <p className="card-text text-secondary">{IdJugador}</p>
                                </td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                            </tr>
                            <tr>
                                <td>
                                    <b className="card-title">Juegos Ganados:</b>
                                </td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td>
                                    <p className="card-text text-secondary">{Conteo}</p>
                                </td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                                <td></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    )
}

export default BestPlayer
