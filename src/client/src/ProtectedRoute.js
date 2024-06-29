import React from 'react';
import { Route, Navigate } from 'react-router-dom';
import { useAuth } from './context/AuthContext';

const ProtectedRoute = ({ element: Element, ...rest }) => {
    const { isAuthenticated } = useAuth();

    return (
        <Route
            {...rest}
            element={isAuthenticated ? <Element /> : <Navigate to="/login" />}
        />
    );
};

export default ProtectedRoute;
