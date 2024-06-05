import React, { useState, useEffect } from 'react';
import axios from 'axios';

function ManageCourses() {
    const [cursos, setCursos] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

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
    const handleAddCourse = async (courseData) => {
        try {
            await axios.post('http://localhost:8080/courses', courseData);
            fetchCursos();  // Refetch the courses after adding
        } catch (error) {
            setError('Error adding course');
        }
    };
    
    const handleEditCourse = async (id, courseData) => {
        try {
            await axios.put(`http://localhost:8080/courses/${id}`, courseData);
            fetchCursos();  // Refetch the courses after editing
        } catch (error) {
            setError('Error editing course');
        }
    };

    return (
        <div>
            <h1>Manage Courses</h1>
            {loading ? <p>Loading...</p> : error ? <p>{error}</p> : (
                <ul>
                    {cursos.map(curso => (
                        <li key={curso.id}>
                            {curso.name}
                            <button onClick={() => handleDelete(curso.id)}>Borrar</button>
                            <button onClick={() => handleEditCourse(curso.id)}>Editar</button>
                        
                        </li>
                    ))}
                    <button onClick={() => handleAddCourse()}>Agregar</button>
                </ul>
            )}
        </div>
    );

  
    
}


export default ManageCourses;
