import React, { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import socketIOClient from "socket.io-client";
import AllGames from './components/AllGames'
import Games from './components/Games'
import BestPlayers from './components/BestPlayers'
import Navbar from "./components/Navbar"
import Reports from './components/Reports'
import './App.css'

const ENDPOINT = "http://localhost:4001";

function App() {
  const [response, setResponse] = useState("");

  useEffect(() => {
    const socket = socketIOClient(ENDPOINT);
    
  }, []);

  return (
    <Router>
      <Navbar/>
      <div className="App">
        <Route path="/" exact component={AllGames} />
        <Route path="/last10games" component={Games} />
        <Route path="/best10" component={BestPlayers} />
        <Route path="/reports" component={Reports} />
      </div>
    </Router>
  )
}

export default App
