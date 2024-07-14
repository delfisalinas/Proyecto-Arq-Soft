import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import '../assets/styles/CourseFiles.css';

const CourseFiles = () => {
  const { courseId } = useParams();
  const navigate = useNavigate();
  const [files, setFiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchFiles = async () => {
      const token = localStorage.getItem('token'); // Obtener el token del almacenamiento local
      try {
        const response = await axios.get(`http://localhost:8080/files/course/${courseId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          }
        });
        console.log("Archivos recibidos:", response.data); // Depuración
        setFiles(response.data);
      } catch (err) {
        console.error("Error fetching files:", err); // Depuración
        setError('Error fetching files: ' + err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchFiles();
  }, [courseId]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="course-files-container">
      <button className="back-button" onClick={() => navigate(-1)}>Volver</button>
      <h1>Archivos del Curso</h1>
      {files.length === 0 ? (
        <p>No hay archivos disponibles.</p>
      ) : (
        <ul>
          {files.map(file => (
            <li key={file.id}>
              <a href={`data:application/octet-stream;base64,${file.content}`} download={file.name}>
                {file.name}
              </a>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default CourseFiles;
