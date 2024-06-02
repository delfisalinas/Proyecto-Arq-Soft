// src/services/authService.js
const API_URL = 'http://tu-backend-url/api'; // Reemplaza con la URL de tu backend

export const login = async (username, password) => {
    const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
    });

    if (!response.ok) {
        throw new Error('Error en la autenticación');
    }

    const data = await response.json();
    localStorage.setItem('token', data.token); // Almacena el token en localStorage
    return data;
};

export const logout = () => {
    localStorage.removeItem('token'); // Elimina el token del almacenamiento
};

export const getCurrentUser = () => {
    return localStorage.getItem('token');
};
