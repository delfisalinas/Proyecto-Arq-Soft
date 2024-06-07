import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../context/UserContext';
import '../assets/styles/Home.css';

function Home() {
    const [cursos, setCursos] = useState([]);
    const { user } = useContext(UserContext);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchCursos = async () => {
            try {
                const response = await axios.get('http://localhost:8080/courses');
                setCursos(response.data);
            } catch (error) {
                console.error('Error fetching data: ', error);
            }
        };

        fetchCursos();
    }, []);

    return (
        <div className="home-container">
            <div className="welcome-message">Bienvenido al Portal de Cursos, {user && user.name ? user.name : 'Usuario'}</div>
            <p className="description">Explora y administra tus cursos con facilidad. Aquí puedes encontrar información detallada sobre todos los cursos disponibles y gestionar tus cursos activos.</p>
            <div className="header">
                <button className="button" onClick={() => navigate('/search')}>Buscar un curso</button>
                <button className="button" onClick={() => navigate('/my-courses')}>Mis Cursos</button>
                <button className="button" onClick={() => navigate('/manage-courses')}>Gestión de Cursos</button>
            </div>
            <h2>Cursos Disponibles</h2>
            <ul className="course-list">
                {cursos.map(curso => (
                    <li key={curso.id} className="course-item">
                        {curso.name}
                        <button className="button" onClick={() => navigate(`/courses/${curso.id}`)}>Click para conocer más detalles</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default Home;
