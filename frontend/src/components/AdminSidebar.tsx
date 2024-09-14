import { Link } from "react-router-dom";

export default function AdminSideBar() {
    return (
        <div className="w-64 bg-white shadow-md">
            <div className="p-4">
                <h1 className="text-2xl font-bold text-gray-800">BeeHub Admin</h1>
            </div>
            <nav className="mt-4">
                <Link to="/admin/dashboard" className="block py-2 px-4 text-gray-700 hover:bg-gray-200">Dashboard</Link>
                <Link to="/admin/download-stats" className="block py-2 px-4 text-gray-700 hover:bg-gray-200">Download Stats</Link>
                <Link to="/admin/ip-logs" className="block py-2 px-4 text-gray-700 hover:bg-gray-200">IP Logs</Link>
            </nav>
        </div>
    );

}