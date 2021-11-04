import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import Games from './components/Games'
import BestPlayers from './components/BestPlayers'
import Navbar from "./components/Navbar";
import './App.css'

function App() {
  return (
    <Router>
      <Navbar/>
      <div className="App">
        <Route path="/" exact component={Games} />
        <Route path="/best10" component={BestPlayers} />
        <Route path="/reports" component={Games} />
      </div>
    </Router>
  )
}

export default App
