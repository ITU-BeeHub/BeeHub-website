import { Navigate, Route, Routes } from "react-router-dom";
import AdminLayout from "./AdminLayout";
import DashboardPage from "./DashboardPage";
import DownloadsPage from "./DownloadsPage";
import LoginPage from "./LoginPage";
import RequireAuth from "./RequireAuth";

/** Routes under /admin. Loaded lazily so visitors never download them. */
export default function AdminApp() {
    return (
        <Routes>
            <Route path="login" element={<LoginPage />} />
            <Route element={<RequireAuth><AdminLayout /></RequireAuth>}>
                <Route index element={<DashboardPage />} />
                <Route path="downloads" element={<DownloadsPage />} />
            </Route>

            {/* Paths of the previous admin panel. */}
            <Route path="dashboard" element={<Navigate to="/admin" replace />} />
            <Route path="download-stats" element={<Navigate to="/admin" replace />} />
            <Route path="ip-logs" element={<Navigate to="/admin/downloads" replace />} />
            <Route path="*" element={<Navigate to="/admin" replace />} />
        </Routes>
    );
}
