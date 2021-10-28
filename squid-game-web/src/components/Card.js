import React from 'react'

function Card({RequestNumber, GameID, GameName, GameWinner, PlayersAmount, Worker}) {
    return (
        <div className="card bg-dark">
            <div className="card-body text-light">
                <div>
                    <table>
                        <tbody>
                            <tr>
                                <td>
                                    <h4 className="card-title">Request #:</h4>
                                </td>
                                <td>
                                    <p className="card-text text-secondary">{RequestNumber}</p>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <h4 className="card-title">Game ID:</h4>
                                </td>
                                <td>
                                    <p className="card-text text-secondary">{GameID}</p>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <h4 className="card-title">Game Name:</h4>
                                </td>
                                <td>
                                    <p className="card-text text-secondary">{GameName}</p>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <h4 className="card-title">Game Winner:</h4>
                                </td>
                                <td>
                                    <p className="card-text text-secondary">{GameWinner}</p>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <h4 className="card-title">Players #:</h4>
                                </td>
                                <td>
                                    <p className="card-text text-secondary">{PlayersAmount}</p>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <h4 className="card-title">Worker:</h4>
                                </td>
                                <td>
                                    <p className="card-text text-secondary">{Worker}</p>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    )
}

export default Card
