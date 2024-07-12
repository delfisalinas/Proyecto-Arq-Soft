import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, useParams } from 'react-router-dom';
import '../assets/styles/AddComment.css';

function AddComment() {
    const { courseId } = useParams();
    const [content, setContent] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleCommentSubmit = async (e) => {
        e.preventDefault();
        const userId = localStorage.getItem('userId');
        const token = localStorage.getItem('token');

        if (!content) {
            setError('Content is required.');
            return;
        }

        try {
            console.log('Sending request to /comments');
            const response = await axios.post('http://localhost:8080/comments', {
                user_id: parseInt(userId),
                course_id: parseInt(courseId),
                content: content
            }, {
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                }
            });
            console.log('Response:', response);

            alert('Comment added successfully.');
            navigate(`/courses/${courseId}`);
        } catch (err) {
            console.error('Error adding comment:', err);
            if (err.response) {
                setError(`Error adding comment: ${err.response.data.error}`);
            } else if (err.request) {
                setError('No response received from server');
            } else {
                setError(`Error in setting up request: ${err.message}`);
            }
        }
    };

    return (
        <div className="add-comment-container">
            <button className="back-button" onClick={() => navigate(`/courses/${courseId}`)}>Volver</button>
            <h1>Agregar Comentario</h1>
            {error && <p className="error-message">{error}</p>}
            <form onSubmit={handleCommentSubmit}>
                <textarea
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                    placeholder="Escribe tu comentario aquí..."
                    className="comment-textarea"
                />
                <button type="submit" className="submit-button">Enviar Comentario</button>
            </form>
        </div>
    );
}

export default AddComment;
