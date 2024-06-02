// src/pages/Home.js
import React, { useEffect, useState } from 'react';
import { getCourses } from '../services/courseService';
import './Home.css';

const Home = () => {
    const [courses, setCourses] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');

    useEffect(() => {
        const fetchCourses = async () => {
            try {
                const data = await getCourses();
                setCourses(data);
            } catch (error) {
                console.error('Error obteniendo cursos', error);
            }
        };

        fetchCourses();
    }, []);

    const handleSearch = (event) => {
        setSearchTerm(event.target.value);
    };

    const filteredCourses = courses.filter(course =>
        course.title.toLowerCase().includes(searchTerm.toLowerCase())
    );

    return (
        <div className="home">
            <div className="header">
                <h1>Listado de Cursos</h1>
                <input
                    type="text"
                    placeholder="Buscar cursos"
                    value={searchTerm}
                    onChange={handleSearch}
                    className="search-input"
                />
            </div>
            <div className="container">
                <div className="course-list">
                    {filteredCourses.map(course => (
                        <div key={course.id} className="course-card">
                            <h2 className="course-title">{course.title}</h2>
                            <button>Inscribirse</button>
                        </div>
                    ))}
                </div>
            </div>
            <div className="footer">
                <p>© 2024 Gestión de Cursos</p>
            </div>
        </div>
    );
};

export default Home;

