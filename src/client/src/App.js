import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Home from './components/Home';
import ProtectedRoute from './ProtectedRoute';  // Ensure this import is correct
import { AuthProvider } from './context/AuthContext';

const App = () => {
    return (
        <AuthProvider>
            <Router>
                <Routes>
                    {/* Route for login page */}
                    <Route path="/login" element={<Login />} />

                    {/* Protected route for home page */}
                    {/* <ProtectedRoute path="/home" element={<Home />} /> */}
                    <Route path="/home" element={<ProtectedRoute element={<Home />} />} />

                    {/* Default route to redirect to home */}
                    <Route path="/" element={<Navigate to="/home" />} />
                </Routes>
            </Router>
        </AuthProvider>
    );
};

export default App;
