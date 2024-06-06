import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useParams } from 'react-router-dom';
import '../assets/styles/CourseDetails.css';


function CourseDetails() {
  const { courseId } = useParams();
  const navigate = useNavigate();
  const [course, setCourse] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchCourse();
  }, [courseId]);

  const fetchCourse = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/courses/${courseId}`);
      setCourse(response.data);
      setLoading(false);
    } catch (err) {
      setError('Error fetching course details');
      setLoading(false);
    }
  };

  const handleEnroll = async () => {
    try {
      const response = await axios.post(`http://localhost:8080/inscriptions`, {
        userId: 1, // Este ID debería ser el del usuario logueado
        courseId: courseId
      });
      alert('Inscripción exitosa!');
      navigate('/my-courses'); // Redirecciona a la lista de cursos del usuario
    } catch (err) {
      alert('Error en la inscripción');
    }
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  if (!course) return <div>Curso no encontrado</div>;

  return (
    <div className="course-details">
      <h1>{course.name}</h1>
      <p><strong>Description:</strong> {course.description}</p>
      <p><strong>Category:</strong> {course.category}</p>
      <p><strong>Duration:</strong> {course.duration}</p>
      <p><strong>Instructor ID:</strong> {course.instructor_id}</p>
      <button onClick={handleEnroll}>Inscribirse</button>
    </div>
  );
}

export default CourseDetails;

