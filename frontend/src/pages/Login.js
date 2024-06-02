// src/pages/Login.js
import React, { useState, useContext } from 'react';
import AppContext from '../context/AppContext';
import { login as loginService } from '../services/authService';
import './Login.css';

const Login = () => {
    const { login } = useContext(AppContext);
    const [userType, setUserType] = useState('student');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = async (event) => {
        event.preventDefault();
        try {
            const userData = await loginService(username, password);
            login({ username, userType, token: userData.token });
            alert(`Logueado como ${userType}`);
        } catch (error) {
            alert('Error en la autenticación');
        }
    };

    return (
        <div className="login">
            <div className="login-form">
                <h1>Iniciar Sesión</h1>
                <div>
                    <label>
                        <input
                            type="radio"
                            value="student"
                            checked={userType === 'student'}
                            onChange={() => setUserType('student')}
                        />
                        Alumno
                    </label>
                    <label>
                        <input
                            type="radio"
                            value="admin"
                            checked={userType === 'admin'}
                            onChange={() => setUserType('admin')}
                        />
                        Administrador
                    </label>
                </div>
                <form onSubmit={handleLogin}>
                    <input
                        type="text"
                        placeholder="Usuario"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <input
                        type="password"
                        placeholder="Contraseña"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <button type="submit">Ingresar</button>
                </form>
            </div>
        </div>
    );
};

export default Login;

