import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../assets/styles/ManageCourses.css';

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
            setCursos(response.data || []);
        } catch (error) {
            setError('Error fetching courses');
            setCursos([]);
        } finally {
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
        <div className="manage-courses-container">
            <button className="back-button" onClick={() => navigate('/home')}>Volver</button>
            <h1>Manage Courses</h1>
            {loading ? <p>Loading...</p> : error ? <p>{error}</p> : null}
            <button onClick={() => navigate('/add-course')} className="add-course-button">Add New Course</button>
            {cursos.length === 0 ? (
                <p>No hay cursos disponibles.</p>
            ) : (
                <ul className="course-list">
                    {cursos.map(curso => (
                        <li key={curso.id} className="course-item">
                            <div className="course-info">
                                <h3>{curso.name}</h3>
                            </div>
                            <div className="course-actions">
                                <button onClick={() => handleDelete(curso.id)} className="delete-button">Delete</button>
                                <button onClick={() => navigate(`/edit-course/${curso.id}`)} className="edit-button">Edit</button>
                                <button onClick={() => navigate(`/courses/${curso.id}`)} className="details-button">Click para conocer más detalles</button>
                            </div>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}

export default ManageCourses;
