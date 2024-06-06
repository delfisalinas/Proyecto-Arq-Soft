import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function ManageCourses() {
    const [cursos, setCursos] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        fetchCursos();
    }, []);

    const fetchCursos = async () => {
        try {
            const response = await axios.get('http://localhost:8080/courses');
            setCursos(response.data);
            setLoading(false);
        } catch (error) {
            setError('Error fetching courses');
            setLoading(false);
        }
    };

    const handleDelete = async (id) => {
        try {
            await axios.delete(`http://localhost:8080/courses/${id}`);
            fetchCursos();  // Refetch the courses after deletion
        } catch (error) {
            setError('Error deleting course');
        }
    };

    return (
        <div>
            <h1>Manage Courses</h1>
            {loading ? <p>Loading...</p> : error ? <p>{error}</p> : null}
            <button onClick={() => navigate('/add-course')} style={{ margin: '10px', padding: '5px' }}>Add New Course</button>
            <ul>
                {cursos.map(curso => (
                    <li key={curso.id}>
                        {curso.name}
                        <button onClick={() => handleDelete(curso.id)} style={{ margin: '5px', padding: '5px' }}>Delete</button>
                        <button onClick={() => navigate(`/edit-course/${curso.id}`)} style={{ margin: '5px', padding: '5px' }}>Edit</button>
                        <button onClick={() => navigate(`/courses/${curso.id}`)} style={{ margin: '5px', padding: '5px' }}>Click para conocer más detalles</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default ManageCourses;
