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
                // Asume que el backend puede identificar al usuario por el token y devuelve solo sus cursos
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
        <div className="my-courses">
            <h1>Mis Cursos</h1>
            {courses.length > 0 ? (
                <ul>
                   {courses.map(course => (
                        <li key={course.id}>
                            {course.name} - {course.description}
                            <Link to={`/courses/${course.id}`}><Link to={`/courses/${course.id}`}>{course.name}</Link> - {course.description}
                                <button style={{ marginLeft: '10px' }}>Click para conocer más detalles</button>
                            </Link>
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
