// src/services/courseService.js
const API_URL = 'http://tu-backend-url/api'; // Reemplaza con la URL de tu backend

export const getCourses = async () => {
    const response = await fetch(`${API_URL}/courses`, {
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    });

    if (!response.ok) {
        throw new Error('Error obteniendo cursos');
    }

    return await response.json();
};

export const enrollInCourse = async (courseId) => {
    const response = await fetch(`${API_URL}/courses/${courseId}/enroll`, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`,
            'Content-Type': 'application/json'
        }
    });

    if (!response.ok) {
        throw new Error('Error inscribiéndose en el curso');
    }

    return await response.json();
};
