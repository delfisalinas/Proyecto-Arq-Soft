import React, { useState, useContext } from 'react';
import axios from 'axios';
import { UserContext } from '../context/UserContext';
import { useNavigate } from 'react-router-dom';


function Register() {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [user_type, setUserType] = useState('alumno');
    const [error, setError] = useState('');
    const { setUser } = useContext(UserContext);
    const navigate = useNavigate();

    const handleRegister = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/users/register', {
                username,
                email,
                password,
                user_type
            });
            const {id} = response.data;
            localStorage.setItem('userId', id);
           localStorage.setItem('usertype', user_type);
           console.log(user_type);
            setUser(response.data);
            alert('Registration successful');
            navigate('/home');
        } catch (error) {
            setError('Failed to register: ' + error);
        }
    };

    return (
        <div>
            <h2>Register</h2>
            <form onSubmit={handleRegister}>
                <input type="text" value={username} onChange={e => setUsername(e.target.value)} placeholder="Username" required />
                <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="Email" required />
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" required />
                <select value={user_type} onChange={e => setUserType(e.target.value)} required>
                    <option value="alumno">Alumno</option>
                    <option value="administrador">Administrador</option>
                </select>

                <button type="submit">Register</button>
            </form>
            {error && <p>{error}</p>}
        </div>
    );
}

export default Register;
