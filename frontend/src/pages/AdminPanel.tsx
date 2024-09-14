import React from 'react';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { Link } from 'react-router-dom'; // Link bileşenini import edin
import AdminSideBar from '../components/AdminSidebar';

// Mock data for the chart
const osData = [
    { name: 'Windows', downloads: 1234 },
    { name: 'Mac', downloads: 987 },
];

// Mock data for recent IP logs
const recentIPs = [
    { ip: '192.168.1.1', timestamp: '2023-06-15 14:30:22' },
    { ip: '10.0.0.1', timestamp: '2023-06-15 14:28:15' },
    { ip: '172.16.0.1', timestamp: '2023-06-15 14:25:03' },
];

const AdminPanel: React.FC = () => {
    return (
        <div className="flex h-screen bg-gray-100">
            {/* Sidebar */}
            <AdminSideBar />

            {/* Main content */}
            <div className="flex-1 p-10 overflow-y-auto">
                <h2 className="text-3xl font-semibold text-gray-800 mb-6">Dashboard</h2>

                {/* Total Downloads Card */}
                <div className="bg-white rounded-lg shadow-md p-6 mb-6">
                    <h3 className="text-xl font-semibold text-gray-700 mb-2">Total Downloads</h3>
                    <p className="text-4xl font-bold text-blue-600">2,221</p>
                </div>

                {/* OS Breakdown Chart */}
                <div className="bg-white rounded-lg shadow-md p-6 mb-6">
                    <h3 className="text-xl font-semibold text-gray-700 mb-4">Downloads by OS</h3>
                    <div className="h-64">
                        <ResponsiveContainer width="100%" height="100%">
                            <BarChart data={osData}>
                                <CartesianGrid strokeDasharray="3 3" />
                                <XAxis dataKey="name" />
                                <YAxis />
                                <Tooltip />
                                <Bar dataKey="downloads" fill="#3B82F6" />
                            </BarChart>
                        </ResponsiveContainer>
                    </div>
                </div>

                {/* Recent IP Logs */}
                <div className="bg-white rounded-lg shadow-md p-6">
                    <h3 className="text-xl font-semibold text-gray-700 mb-4">Recent IP Logs</h3>
                    <table className="min-w-full">
                        <thead>
                            <tr>
                                <th className="text-left text-gray-600">IP Address</th>
                                <th className="text-left text-gray-600">Timestamp</th>
                            </tr>
                        </thead>
                        <tbody>
                            {recentIPs.map((log, index) => (
                                <tr key={index} className="border-t">
                                    <td className="py-2">{log.ip}</td>
                                    <td className="py-2">{log.timestamp}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    );
};

export default AdminPanel;
