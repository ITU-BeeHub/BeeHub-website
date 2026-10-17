import React, { useEffect } from "react";
import { Navigate, useLocation } from "react-router-dom";
import { clearSession, useSession } from "./session";

/** Renders its children only while an admin session is active. */
export default function RequireAuth({ children }: { children: React.ReactElement }) {
    const session = useSession();
    const location = useLocation();

    useEffect(() => {
        if (!session) return;
        const timer = setTimeout(clearSession, new Date(session.expiresAt).getTime() - Date.now());
        return () => clearTimeout(timer);
    }, [session]);

    if (!session) {
        return <Navigate to="/admin/login" replace state={{ from: location.pathname }} />;
    }
    return children;
}
