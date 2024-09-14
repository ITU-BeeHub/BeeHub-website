import React from 'react';
import { Link } from 'react-router-dom';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, LineChart, Line } from 'recharts';
import AdminSideBar from './AdminSidebar';

const monthlyData = [
    { name: 'Jan', downloads: 1200 },
    { name: 'Feb', downloads: 1900 },
    { name: 'Mar', downloads: 2400 },
    { name: 'Apr', downloads: 1800 },
    { name: 'May', downloads: 2800 },
    { name: 'Jun', downloads: 2600 },
];

const osData = [
    { name: 'Windows', downloads: 5400 },
    { name: 'Mac', downloads: 4200 },
    { name: 'Linux', downloads: 1100 },
];

const DownloadStats: React.FC = () => {
    return (
        <div className="flex h-screen bg-gray-100">
            {/* Sidebar */}
            <AdminSideBar />
            <div className="flex-1 p-10 overflow-y-auto">
                <h2 className="text-3xl font-semibold text-gray-800 mb-6">Download Statistics</h2>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div className="bg-white rounded-lg shadow-md p-6">
                        <h3 className="text-xl font-semibold text-gray-700 mb-4">Monthly Downloads</h3>
                        <div className="h-64">
                            <ResponsiveContainer width="100%" height="100%">
                                <LineChart data={monthlyData}>
                                    <CartesianGrid strokeDasharray="3 3" />
                                    <XAxis dataKey="name" />
                                    <YAxis />
                                    <Tooltip />
                                    <Line type="monotone" dataKey="downloads" stroke="#3B82F6" />
                                </LineChart>
                            </ResponsiveContainer>
                        </div>
                    </div>

                    <div className="bg-white rounded-lg shadow-md p-6">
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
                </div>

                <div className="mt-6 bg-white rounded-lg shadow-md p-6">
                    <h3 className="text-xl font-semibold text-gray-700 mb-4">Download Summary</h3>
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <div className="bg-blue-100 p-4 rounded-lg">
                            <p className="text-sm text-blue-600 font-semibold">Total Downloads</p>
                            <p className="text-3xl font-bold text-blue-800">10,700</p>
                        </div>
                        <div className="bg-green-100 p-4 rounded-lg">
                            <p className="text-sm text-green-600 font-semibold">This Month</p>
                            <p className="text-3xl font-bold text-green-800">2,600</p>
                        </div>
                        <div className="bg-purple-100 p-4 rounded-lg">
                            <p className="text-sm text-purple-600 font-semibold">Average per Day</p>
                            <p className="text-3xl font-bold text-purple-800">356</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default DownloadStats;