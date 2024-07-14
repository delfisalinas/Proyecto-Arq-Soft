import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, useParams } from 'react-router-dom';
import '../assets/styles/CommentForm.css';

function CommentForm() {
    const { courseId } = useParams();
    const navigate = useNavigate();
    const [comment, setComment] = useState('');
    const [error, setError] = useState(null);

    const handleInputChange = (e) => {
        setComment(e.target.value);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const userId = parseInt(localStorage.getItem('userId'), 10); // Convertir a número
        const token = localStorage.getItem('token');

        if (!userId) {
            setError('User ID not found');
            return;
        }

        try {
            const response = await axios.post(
                `http://localhost:8080/comments/`,
                {
                    user_id: userId, // Asegurarse de enviar como número
                    course_id: parseInt(courseId, 10), // Convertir a número
                    content: comment,
                },
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                }
            );
            alert('Comentario enviado exitosamente!');
            navigate(`/courses/${courseId}`);
        } catch (err) {
            console.error('Error al enviar el comentario:', err);
            setError(`Error al enviar el comentario: ${err.response ? err.response.data.error : err.message}`);
        }
    };

    return (
        <div className="comment-form-container">
            <button className="back-button" onClick={() => navigate(-1)}>Volver</button>
            <h1>Realizar Comentario</h1>
            {error && <p className="error-message">{error}</p>}
            <form onSubmit={handleSubmit}>
                <textarea
                    value={comment}
                    onChange={handleInputChange}
                    placeholder="Escribe tu comentario aquí"
                    required
                    className="comment-textarea"
                />
                <button type="submit" className="submit-button">Enviar Comentario</button>
            </form>
        </div>
    );
}

export default CommentForm;
