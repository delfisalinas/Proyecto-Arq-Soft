import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import '../assets/styles/CourseComments.css';

const CourseComments = () => {
  const { courseId } = useParams();
  const navigate = useNavigate();
  const [comments, setComments] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchComments = async () => {
      const token = localStorage.getItem('token'); // Obtener el token del almacenamiento local
      try {
        const response = await axios.get(`http://localhost:8080/comments/courses/${courseId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        });
        console.log("Comentarios recibidos:", response.data); // Depuración
        setComments(response.data);
      } catch (err) {
        console.error("Error fetching comments:", err); // Depuración
        setError('Error fetching comments: ' + err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchComments();
  }, [courseId]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="course-comments-container">
      <button className="back-button" onClick={() => navigate(-1)}>Volver</button>
      <h1>Comentarios del Curso</h1>
      {comments && comments.length === 0 ? (
        <p>No hay comentarios disponibles.</p>
      ) : (
        <ul>
          {comments && comments.map(comment => (
            <li key={comment.id}>
              <p>{comment.content}</p>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default CourseComments;
