// src/pages/CourseDetail.js
import React, { useEffect, useState, useContext } from 'react';
import { useParams } from 'react-router-dom';
import AppContext from '../context/AppContext';
import { enrollInCourse } from '../services/courseService';
import './CourseDetail.css';

const CourseDetail = () => {
    const { id } = useParams();
    const { enrollInCourse: enroll, myCourses } = useContext(AppContext);
    const [course, setCourse] = useState(null);
    const [isEnrolled, setIsEnrolled] = useState(false);

    useEffect(() => {
        const fetchedCourse = {
            id,
            title: 'Curso de React',
            description: 'Aprende React desde cero y construye aplicaciones web modernas.',
            instructor: 'Juan Pérez',
            duration: '10 horas',
            requirements: 'Conocimientos básicos de JavaScript'
        };
        setCourse(fetchedCourse);
        setIsEnrolled(myCourses.some(c => c.id === id));
    }, [id, myCourses]);

    const handleEnroll = async () => {
        try {
            await enrollInCourse(course.id);
            enroll(course);
            setIsEnrolled(true);
            alert('Te has inscrito correctamente en el curso');
        } catch (error) {
            alert('Error inscribiéndose en el curso');
        }
    };

    if (!course) return <div>Cargando...</div>;

    return (
        <div className="course-detail">
            <div className="header">
                <h1>{course.title}</h1>
            </div>
            <div className="container">
                <p><strong>Descripción:</strong> {course.description}</p>
                <p><strong>Instructor:</strong> {course.instructor}</p>
                <p><strong>Duración:</strong> {course.duration}</p>
                <p><strong>Requisitos:</strong> {course.requirements}</p>
                {!isEnrolled && (
                    <button className="enroll-button" onClick={handleEnroll}>Inscribirse</button>
                )}
                {isEnrolled && <p>Ya estás inscrito en este curso.</p>}
            </div>
            <div className="footer">
                <p>© 2024 Gestión de Cursos</p>
            </div>
        </div>
    );
};

export default CourseDetail;
