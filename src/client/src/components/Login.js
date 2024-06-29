import React, { useState } from 'react';
import './../css/login.css';
import { useNavigate } from 'react-router-dom';
import loginimg from './../img/URLShortener.png';
import { useAuth } from '../context/AuthContext';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();
    const {login} = useAuth();

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log('Username:', username);
        console.log('Password:', password);
        if (username != "admin" && password != "admin") {
            alert("Invalid Credentials");
            return;
        }
        login();
        navigate("/home");
    };

    return (
        <div className="login-page">
            <div className="login-box">
                <form name="login-form" className='login-form' onSubmit={handleSubmit}>
                    <p className="form-title">URL SHORTENER</p>
                    <p>Login to the Dashboard</p>
                    <div className="form-item">
                        <input
                            type="text"
                            placeholder="Username"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required
                        />
                    </div>
                    <div className="form-item">
                        <input
                            type="password"
                            placeholder="Password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required
                        />
                    </div>
                    <div className="form-item">
                        <button type="submit" className="login-form-button">
                            LOGIN
                        </button>
                    </div>
                </form>
                <div className="illustration-wrapper">
                    <img src={loginimg} alt="Login" style={{height:"100%"}} />
                </div>
            </div>
        </div>
    );
};

export default Login;
