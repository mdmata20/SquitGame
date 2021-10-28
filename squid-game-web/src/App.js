import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import Games from './components/Games'
import Navbar from "./components/Navbar";

function App() {
  return (
    <Router>
      <Navbar/>
      <div className="App">
        <Route path="/" exact component={Games} />
        <Route path="/reports" component={Games} />
      </div>
    </Router>
  )
}

export default App
