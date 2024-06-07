import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import '../assets/styles/SearchCourses.css';

function SearchCourses() {
    const [query, setQuery] = useState('');
    const [results, setResults] = useState([]);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSearch = async (event) => {
        event.preventDefault();
        try {
            const response = await axios.get(`http://localhost:8080/search/courses?q=${query}`);
            setResults(response.data);
            setError('');
        } catch (error) {
            console.error('Error during data fetching', error);
            setError('Error al cargar los datos. Por favor, intenta nuevamente.');
            setResults([]);
        }
    };

    return (
        <div>
            <form onSubmit={handleSearch}>
                <input
                    type="text"
                    value={query}
                    onChange={e => setQuery(e.target.value)}
                    placeholder="Buscar cursos"
                />
                <button type="submit">Buscar</button>
            </form>
            {error && <p>{error}</p>}
            {(!results || results.length === 0) && !error && <p>No hay cursos disponibles.</p>}
            <ul>
                {results && results.map(course => (
                    <li key={course.id}>
                        {course.name} - {course.description}
                        <button onClick={() => navigate(`/courses/${course.id}`)} style={{ marginLeft: '10px' }}>Click para conocer más detalles</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default SearchCourses;
