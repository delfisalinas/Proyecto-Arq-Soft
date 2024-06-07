import React, { useState, useContext } from 'react';
import axios from 'axios';
import { UserContext } from '../context/UserContext';
import { useNavigate } from 'react-router-dom';

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const { setUser } = useContext(UserContext);
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/users/login', {
                username,
                password
            });
            const { id,user_type, token } = response.data;

            
            localStorage.setItem('userId', id);
            localStorage.setItem('usertype', user_type);
            localStorage.setItem('token', token);
            console.log(user_type);
            setUser(response.data);
            alert('Login successful' + user_type);
            navigate('/home');
        } catch (error) {
            setError('Failed to login: ' + error);
        }
    };

    const handleRegisterRedirect = () => {
        navigate('/register'); // Asegúrate de que esta es la ruta correcta para el registro
    };

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleLogin}>
                <input type="text" value={username} onChange={e => setUsername(e.target.value)} placeholder="Username" required />
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" required />
                <button type="submit">Login</button>
                {/* Botón adicional para registrarse */}
                <button type="button" onClick={handleRegisterRedirect}>Registrarse</button>
            </form>
            {error && <p>{error}</p>}
        </div>
    );
}

export default Login;
