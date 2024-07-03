import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/courses/Home';
import Login from './components/auth/Login';
import Register from './components/auth/Register';
import CourseDetails from './components/courses/CourseDetails';
import MyCourses from './components/courses/MyCourses';
import Files from './components/courses/Files';
import SearchCourses from './components/courses/SearchCourses';
import ManageCourses from './components/courses/ManageCourses';
import AddCourse from './components/courses/AddCourse';
import EditCourse from './components/courses/EditCourse';
import { Navigate } from 'react-router-dom';
import ProtectedRoute from './components/ProtectedRoute';



function App() {
  return (
<Router>
    <Routes>
        <Route path="/" element={<Navigate replace to="/login" />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route element={<ProtectedRoute />}>
          <Route path="/home" element={<Home />} />
          <Route path="/courses/:courseId" element={<CourseDetails />} />
          <Route path="/my-courses" element={<MyCourses />} />
          <Route path="/search" element={<SearchCourses />} />
          <Route path="/manage-courses" element={<ManageCourses />} />
         <Route path="/add-course" element={<AddCourse />} />
         <Route path="/edit-course/:courseId" element={<EditCourse />} />
         <Route path="/upload/:courseId" element={<Files />} />
        </Route>
               
    </Routes>
</Router>
  );
}

export default App;

