import React from 'react'
import { useNavigate } from 'react-router-dom'

const Home = (props) => {
    const { loggedIn, email } = props
    const navigate = useNavigate()

    const onButtonClick = () => {
        if (loggedIn) {
            localStorage.removeItem("token")
            props.setLoggedIn(false)
        } else {
            navigate("/login")
        }
    }

    return (
        <div className='mainContainer'>
            <div className='titleContainer'>
                <div>Welcome!</div>
            </div>
            <div>Home Page</div>
            <div className='buttonContainer'>
                <input
                    className='inputButton'
                    type='button'
                    onClick={onButtonClick}
                    value={loggedIn ? 'Log out': 'Log in'}
                />
                {loggedIn ? <div>Your username is {email}</div> : <div />}
            </div>
        </div>
    )
}

export default Home
