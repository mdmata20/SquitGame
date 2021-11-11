import React, { Component } from "react";
import { Link } from 'react-router-dom'

export default class Navbar extends Component {
  render() {
    return (
      <nav className="navbar navbar-expand-lg navbar-dark bg-dark p-3">
        <div className="container-fluid">
          <Link className="navbar-brand" to="/">
            Squid Game
            </Link>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav ms-auto">
              <li className="nav-item active">
                <Link to="/" className="nav-link">All Games</Link>
              </li>
              <li className="nav-item">
                <Link to="/last10games" className="nav-link">Last 10 Games</Link>
              </li>
              <li className="nav-item">
                <Link to="/best10" className="nav-link">Best 10 Players</Link>
              </li>
              <li className="nav-item">
                <Link to="/reports" className="nav-link">Reports</Link>
              </li>
            </ul>
          </div>
        </div>
      </nav>
    );
  }

};
