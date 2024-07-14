import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate, Link } from 'react-router-dom';
import { useParams } from 'react-router-dom';
import '../assets/styles/CourseDetails.css';

function CourseDetails() {
    const { courseId } = useParams();
    const navigate = useNavigate();
    const [course, setCourse] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const CourseId = parseInt(courseId, 10);
        if (isNaN(CourseId)) {
            setError('No se convierte a entero el id');
            setLoading(false);
            return;
        }
        fetchCourse(CourseId);
    }, [courseId]);

    const fetchCourse = async (CourseId) => {
        try {
            const response = await axios.get(`http://localhost:8080/courses/${CourseId}`);
            setCourse(response.data);
            setLoading(false);
        } catch (err) {
            setError('Error fetching course details ' + err);
            setLoading(false);
        }
    };

    const handleEnroll = async () => {
        const user_id = localStorage.getItem('userId');
        try {
            await axios.post(`http://localhost:8080/inscriptions`, {
                user_id: parseInt(user_id),
                course_id: parseInt(courseId)
            });
            alert('Inscripción exitosa!');
            navigate('/home');
        } catch (err) {
            alert('Error en la inscripción ' + err);
        }
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;
    if (!course) return <div>Curso no encontrado</div>;

    return (
        <div className="course-details-container">
            <button className="back-button" onClick={() => navigate('/home')}>Volver</button>
            <div className="course-details">
                <h1>{course.name}</h1>
                <p><strong>Description:</strong> {course.description}</p>
                <p><strong>Category:</strong> {course.category}</p>
                <p><strong>Duration:</strong> {course.duration}</p>
                <p><strong>Instructor ID:</strong> {course.instructor_id}</p>
                <button onClick={handleEnroll}>Inscribirse</button>

                <div className="course-files">
                    <h2>Archivos del Curso</h2>
                    <Link to={`/course-files/${courseId}`} className="files-button">Ver Archivos</Link>
                </div>
                <div className="course-comments">
                    <h2>Comentarios del Curso</h2>
                    <Link to={`/courses/${courseId}/comments`} className="comments-button">Ver Comentarios</Link>
                </div>
            </div>
        </div>
    );
}

export default CourseDetails;
