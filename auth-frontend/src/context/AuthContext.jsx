
import { createContext, useEffect, useState } from "react";
import { authProfile } from "../api/auth";
import { useLocation } from "react-router-dom";

export const AuthContext = createContext();

function AuthProvider({ children }) {

    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [loading, setLoading] = useState(true);

    const loc=useLocation()

    useEffect(() => {

        

        if(loc.pathname==="/login"  || loc.pathname==="/register"){
            setLoading(false)
            return
        }


        checkSession();
    }, [loc.pathname]);

    async function checkSession() {

        try {

        const res = await authProfile();

        if (res.status === 401) {
            setIsAuthenticated(false);
            return;
        }

        if (res.ok) {
            setIsAuthenticated(true);
        }

    } catch (error) {

        console.error("Session check failed:", error);
        setIsAuthenticated(false);

    } finally {

        setLoading(false);

    }
}

    function login() {
        setIsAuthenticated(true);
    }

    function logout() {
        setIsAuthenticated(false);
    }

    return (
        <AuthContext.Provider
            value={{
                isAuthenticated,
                loading,
                login,
                logout,
                checkSession,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
}

export default AuthProvider;