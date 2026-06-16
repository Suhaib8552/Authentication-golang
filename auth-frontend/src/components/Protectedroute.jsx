// import React, { useContext } from 'react'
// import { Navigate } from 'react-router-dom'
// import { AuthContext } from '../context/AuthContext'

// function Protectedroute({children}) {

    

//     const token=document.cookie.includes("token")

//     if(!token){

//         return <Navigate to="/login"/>

//     }

//   return children
// }

// export default Protectedroute


import { useContext } from "react";
import { Navigate } from "react-router-dom";
import { AuthContext } from "../context/AuthContext";

function ProtectedRoute({ children }) {

    const { isAuthenticated, loading } =
        useContext(AuthContext);

    if (loading) {
        return <h3>Loading...</h3>;
    }

    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }

    return children;
}

export default ProtectedRoute;