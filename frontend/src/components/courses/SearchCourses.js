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
        <div className="search-container">
             <button className="back-button" onClick={() => navigate('/home')}>Volver</button>
            <form onSubmit={handleSearch} className="search-form">
                <input
                    type="text"
                    value={query}
                    onChange={e => setQuery(e.target.value)}
                    placeholder="Buscar cursos"
                    className="search-input"
                />
                <button type="submit" className="search-button">Buscar</button>
            </form>
            {error && <p className="error-message">{error}</p>}
            {(!results || results.length === 0) && !error && <p className="no-results">No hay cursos disponibles.</p>}
            <ul className="results-list">
                {results && results.map(course => (
                    <li key={course.id} className="result-item">
                        <div className="course-info">
                            <h3>{course.name}</h3>
                            <p>{course.description}</p>
                        </div>
                        <button className="details-button" onClick={() => navigate(`/courses/${course.id}`)}>Click para conocer más detalles</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default SearchCourses;
