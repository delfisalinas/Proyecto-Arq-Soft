import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../assets/styles/Home.css';


function Home() {
  const [cursos, setCursos] = useState([]);

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
      <h1>Cursos Disponibles</h1>
      <ul className="course-list">
        {cursos.map(curso => (
          <li key={curso.id} className="course-item">
            <span className="course-name">{curso.name}</span> - <span className="course-description">{curso.description}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Home;

