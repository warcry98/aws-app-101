import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './Pages/home'
import Login from './Pages/login'
import './App.css'
import { useEffect, useState } from 'react'
import axios from 'axios'

function App() {
  useEffect(() => {
    const token = localStorage.getItem("token")

    if (!token) {
      setLoggedIn(false)
      return
    }

    axios.get(
      process.env.REACT_APP_BACKEND_URL+"/validate",
      {
        headers: {
          Accept: "application/json",
          Authorization: token,
        }
      }
    )
    .then((res) => {
      if (res.status === 200) {
        setLoggedIn(true)
        setEmail(res.data["message"])
      }
    })
    .catch((err) => {
      window.alert("Cannot connect to Network")
      localStorage.removeItem("token")
      setLoggedIn(false)
  })
  }, [])

  const [loggedIn, setLoggedIn] = useState(false)
  const [email, setEmail] = useState('')

  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route
            path='/'
            element={<Home email={email} loggedIn={loggedIn} setLoggedIn={setLoggedIn} />}
          />
          <Route path='/login' element={<Login setLoggedIn={setLoggedIn} setEmail={setEmail} />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
