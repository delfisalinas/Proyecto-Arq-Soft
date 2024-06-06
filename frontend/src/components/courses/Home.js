import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../assets/styles/Home.css';

function Home() {
  const [cursos, setCursos] = useState([]);
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
      <div className="header">
        <h1>Mis Cursos</h1>
        <button onClick={() => navigate('/search')}>Buscar un curso</button>
        <button onClick={() => navigate('/manage-courses')}>Gestión de Cursos</button>
      </div>
      <h2>Cursos Disponibles</h2>
      <ul className="course-list">
        {cursos.map(curso => (
          <li key={curso.id} className="course-item" onClick={() => navigate(`/courses/${curso.id}`)}>
            {curso.name} - Click aquí para más detalles
            <button onClick={() => navigate(`/courses/${curso.id}`)} style={{ marginLeft: '10px' }}>Click para conocer más detalles</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Home;
