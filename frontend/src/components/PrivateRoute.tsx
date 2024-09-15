import React from 'react';
import { Navigate } from 'react-router-dom';

interface PrivateRouteProps {
    children: React.ReactElement;
}

const PrivateRoute: React.FC<PrivateRouteProps> = ({ children }) => {
    const token = sessionStorage.getItem('token');

    if (!token) {
        // Kullanıcı oturum açmamışsa login sayfasına yönlendir
        return <Navigate to="/admin/login" replace />;
    }

    return children;
};

export default PrivateRoute;
