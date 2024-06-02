// src/App.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/Login';
import CourseDetail from './pages/CourseDetail';
import MyCourses from './pages/MyCourses';
import './styles/global.css';

function App() {
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/course/:id" element={<CourseDetail />} />
                    <Route path="/my-courses" element={<MyCourses />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
