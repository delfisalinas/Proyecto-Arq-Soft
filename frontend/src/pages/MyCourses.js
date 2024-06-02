// src/pages/MyCourses.js
import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import AppContext from '../context/AppContext';
import './MyCourses.css';

const MyCourses = () => {
    const { myCourses } = useContext(AppContext);

    return (
        <div className="my-courses">
            <div className="header">
                <h1>Mis Cursos</h1>
            </div>
            <div className="container">
                <div className="course-list">
                    {myCourses.length > 0 ? (
                        myCourses.map(course => (
                            <div key={course.id} className="course-card">
                                <h2 className="course-title">{course.title}</h2>
                                <Link to={`/course/${course.id}`} className="details-link">Ver detalles</Link>
                            </div>
                        ))
                    ) : (
                        <p>No estás inscrito en ningún curso.</p>
                    )}
                </div>
            </div>
            <div className="footer">
                <p>© 2024 Gestión de Cursos</p>
            </div>
        </div>
    );
};

export default MyCourses;
