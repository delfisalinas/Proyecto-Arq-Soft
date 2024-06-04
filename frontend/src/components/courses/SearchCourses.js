import React, { useState } from 'react';
import axios from 'axios';
import '../assets/styles/SearchCourses.css';



function SearchCourses() {
    const [query, setQuery] = useState('');
    const [results, setResults] = useState([]);

    const handleSearch = async (event) => {
        event.preventDefault();
        try {
            const response = await axios.get(`http://localhost:8080/search/courses?q=${query}`);
            setResults(response.data);
        } catch (error) {
            console.error('Error during data fetching', error);
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
            <ul>
                {results.map(course => (
                    <li key={course.id}>
                        {course.name} - {course.description}
                    </li>
                ))}
            </ul>
        </div>
    );
}

export default SearchCourses;
