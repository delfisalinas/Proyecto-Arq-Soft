import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../assets/styles/AddCourse.css';

function AddCourse() {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [category, setCategory] = useState('');
    const [duration, setDuration] = useState('');
    const [instructorId, setInstructorId] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();  // Prevenir el comportamiento por defecto del formulario

        if (!name || !description || !category || !duration || !instructorId) {
            setError('Todos los campos son obligatorios');
            return;
        }

        try {
            const response = await axios.post('http://localhost:8080/courses', {
                name,
                description,
                category,
                duration,
                instructor_id: parseInt(instructorId)
            });
            // Lógica post-creación
            setError('');
            alert('Curso agregado con éxito');
            navigate('/manage-courses'); // Redirigir a la página de gestión de cursos
        } catch (error) {
            setError('Error al agregar curso: ' + (error.response?.data?.message || error.message));
        }
    };

    return (
        <div className="add-course-container">
             <button className="back-button" onClick={() => navigate('/home')}>Volver</button>
            <h1>Agregar nuevo curso</h1>
            {error && <p className="error-message">{error}</p>}
            <form onSubmit={handleSubmit} className="add-course-form">
                <input type="text" value={name} onChange={e => setName(e.target.value)} placeholder="Nombre del curso" />
                <input type="text" value={description} onChange={e => setDescription(e.target.value)} placeholder="Descripción" />
                <input type="text" value={category} onChange={e => setCategory(e.target.value)} placeholder="Categoría" />
                <input type="text" value={duration} onChange={e => setDuration(e.target.value)} placeholder="Duración" />
                <input 
                    type="number" 
                    value={instructorId} 
                    onChange={e => setInstructorId(e.target.value)} 
                    placeholder="ID del instructor" 
                    min="0"  // Asegura que solo se permitan valores positivos
                />
                <button type="submit" className="submit-button">Agregar curso</button>
            </form>
        </div>
    );
}

export default AddCourse;
