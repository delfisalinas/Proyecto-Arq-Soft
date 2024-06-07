import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import '../assets/styles/EditCourse.css';

function EditCourse() {
    const { courseId } = useParams();
    const navigate = useNavigate();
    const [courseData, setCourseData] = useState({
        name: '',
        description: '',
        category: '',
        duration: '',
        instructor_id: ''
    });
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchCourseData = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/courses/${courseId}`);
                setCourseData({
                    name: response.data.name,
                    description: response.data.description,
                    category: response.data.category,
                    duration: response.data.duration,
                    instructor_id: response.data.instructor_id.toString()  // Convertir a string si es necesario
                });
            } catch (err) {
                setError('Error fetching course details: ' + err.message);
            }
        };
        fetchCourseData();
    }, [courseId]);

    const handleChange = (e) => {
        setCourseData({ ...courseData, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.put(`http://localhost:8080/courses/${courseId}`, courseData);
            alert('Curso actualizado con éxito');
            navigate('/manage-courses'); // Redirigir a la gestión de cursos
        } catch (error) {
            setError(`Error updating course: ${error.response?.data?.message || error.message}`);
        }
    };

    return (
        <div className="edit-course-container">
            <h1>Editar Curso</h1>
            {error && <p className="error-message">{error}</p>}
            <form onSubmit={handleSubmit} className="edit-course-form">
                <label>Nombre del curso</label>
                <input name="name" value={courseData.name} onChange={handleChange} />

                <label>Descripción</label>
                <input name="description" value={courseData.description} onChange={handleChange} />

                <label>Categoría</label>
                <input name="category" value={courseData.category} onChange={handleChange} />

                <label>Duración</label>
                <input name="duration" value={courseData.duration} onChange={handleChange} />

                <label>ID del Instructor</label>
                <input name="instructor_id" value={courseData.instructor_id} onChange={handleChange} />

                <button type="submit" className="submit-button">Actualizar Curso</button>
            </form>
        </div>
    );
}

export default EditCourse;
