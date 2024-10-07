import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import axios from 'axios';

const Login = (props) => {
    const [email, setEmail] = useState('')
    const [emailError, setEmailError] = useState('')
    const [password, setPassword] = useState('')
    const [passwordError, setPasswordError] = useState('')

    const navigate = useNavigate()

    const onButtonClick = () => {
        setEmailError('')
        setPasswordError('')

        if ('' === email) {
            setEmailError('Please enter your email/username')
            return
        }

        if ('' === password) {
            setPasswordError('Please enter a password')
            return
        }

        if (password.length < 7) {
            setPasswordError('The password must be 8 character or longer')
            return
        }

        const logIn = () => {
            axios.post(
                process.env.REACT_APP_BACKEND_URL+"/login",
                {
                    "Email": email,
                    "Password": password,
                },
            )
            .then((res) => {
                if (res.status === 200) {
                    localStorage.setItem("token", res.data["token"])
                    props.setLoggedIn(true)
                    props.setEmail(email)
                    navigate('/')
                } else {
                    window.alert("Wrong Username/Email or Password")
                }
            })
            .catch((err) => {
                window.alert("Cannot connect to Network")
            })
        }

        logIn()
    }

    return (
        <div className="mainContainer">
            <div className="titleContainer">
                <div>Login</div>
            </div>
            <br />
            <div className="inputContainer">
                <input
                    value={email}
                    placeholder="Enter your email/username"
                    onChange={(e) => setEmail(e.target.value)}
                    className="inputBox"
                />
                <label className="errorLabel">{emailError}</label>
            </div>
            <br />
            <div className="inputContainer">
                <input
                    value={password}
                    placeholder="Enter your password here"
                    onChange={(e) => setPassword(e.target.value)}
                    className="inputBox"
                />
                <label className="errorLabel">{passwordError}</label>
            </div>
            <br />
            <div className="inputCntainer">
                <input className="inputButton" type="button" onClick={onButtonClick} value={"Log in"}/>
            </div>
        </div>
    )
}

export default Login
