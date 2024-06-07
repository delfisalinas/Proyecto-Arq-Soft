import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import '../assets/styles/MyCourses.css';

function MyCourses() {
    const [courses, setCourses] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchMyCourses = async () => {
            try {
                const response = await axios.get('http://localhost:8080/my-courses');
                setCourses(response.data);
                setLoading(false);
            } catch (err) {
                setError('Error fetching my courses');
                setLoading(false);
            }
        };

        fetchMyCourses();
    }, []);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div className="my-courses-container">
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
