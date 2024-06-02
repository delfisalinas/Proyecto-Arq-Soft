// src/context/AppProvider.js
import React, { useState } from 'react';
import AppContext from './AppContext';

const AppProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [myCourses, setMyCourses] = useState([]);

    const login = (userData) => {
        setUser(userData);
    };

    const logout = () => {
        setUser(null);
        setMyCourses([]);
    };

    const enrollInCourse = (course) => {
        setMyCourses([...myCourses, course]);
    };

    return (
        <AppContext.Provider value={{ user, login, logout, myCourses, enrollInCourse }}>
            {children}
        </AppContext.Provider>
    );
};

export default AppProvider;
