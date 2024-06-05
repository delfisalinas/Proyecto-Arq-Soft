import React, { useState } from 'react';
import axios from 'axios';

function AddCourse() {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [category, setCategory] = useState('');
    const [duration, setDuration] = useState('');
    const [instructorId, setInstructorId] = useState('');
    const [error, setError] = useState('');

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
                instructor_id: instructorId  // Asegúrate de que este campo coincida con lo que espera tu backend
            });
            // Lógica post-creación
            setError('');
            alert('Curso agregado con éxito');
            // Limpia los campos o redirige según sea necesario
        } catch (error) {
            setError('Error al agregar curso: ' + (error.response?.data?.message || error.message));
        }
    };

    return (
        <div>
            <h1>Agregar nuevo curso</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <form onSubmit={handleSubmit}>
                <input type="text" value={name} onChange={e => setName(e.target.value)} placeholder="Nombre del curso" />
                <input type="text" value={description} onChange={e => setDescription(e.target.value)} placeholder="Descripción" />
                <input type="text" value={category} onChange={e => setCategory(e.target.value)} placeholder="Categoría" />
                <input type="text" value={duration} onChange={e => setDuration(e.target.value)} placeholder="Duración" />
                <input type="text" value={instructorId} onChange={e => setInstructorId(e.target.value)} placeholder="ID del instructor" />
                <button type="submit">Agregar curso</button>
            </form>
        </div>
    );
}

export default AddCourse;
