import React, { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import socketIOClient from "socket.io-client";
import AllGames from './components/AllGames'
import Games from './components/Games'
import BestPlayers from './components/BestPlayers'
import Navbar from "./components/Navbar"
import PieChart from './components/PieChart'
import PieChartWorkers from './components/PieChartWorkers'
import UserStatistics from './components/UserStatistics'
import './App.css'

const ENDPOINT = "http://35.225.182.66:4001";

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
        <Route path="/reportTop3Games" component={PieChart} />
        <Route path="/reportWorkers" component={PieChartWorkers} />
        <Route path="/viewUser/:id" component={UserStatistics} />
      </div>
    </Router>
  )
}

export default App
