import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';
import '../assets/styles/MyCourses.css';

function MyCourses() {
    const [courses, setCourses] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchMyCourses = async () => {
            try {
                const userId = localStorage.getItem('userId');
                if (!userId) {
                    throw new Error('User ID not found');
                }

                const response = await axios.get(`http://localhost:8080/user/${userId}/courses`);
                setCourses(response.data);
            } catch (err) {
                console.error('Error fetching my courses:', err);
                setError('Error fetching my courses');
            } finally {
                setLoading(false);
            }
        };

        fetchMyCourses();
    }, []);

    const handleUploadClick = (courseId) => {
        navigate(`/upload/${courseId}`);
    };

    const handleCommentClick = (courseId) => {
        navigate(`/courses/${courseId}/comment`);
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div className="my-courses-container">
            <button className="back-button" onClick={() => navigate('/home')}>Volver</button>
            <h1>Mis Cursos</h1>
            {courses.length > 0 ? (
                <ul className="course-list">
                    {courses.map(course => (
                        <li key={course.id} className="course-item">
                            <div className="course-info">
                                <h3>{course.name}</h3>
                                <p>{course.description}</p>
                            </div>
                            <Link to={`/courses/${course.id}`} className="details-button">Click para conocer más detalles</Link>
                            <button onClick={() => handleUploadClick(course.id)} className="upload-button">Subir Archivo</button>
                            <button onClick={() => handleCommentClick(course.id)} className="comment-button">Realizar Comentario</button>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No estás inscrito en ningún curso aún.</p>
            )}
        </div>
    );
}

export default MyCourses;
