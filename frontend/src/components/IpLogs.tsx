import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import AdminSideBar from './AdminSidebar';

const ipLogs = [
    { ip: '192.168.1.1', timestamp: '2023-06-15 14:30:22', country: 'United States', os: 'Windows' },
    { ip: '10.0.0.1', timestamp: '2023-06-15 14:28:15', country: 'Canada', os: 'Mac' },
    { ip: '172.16.0.1', timestamp: '2023-06-15 14:25:03', country: 'United Kingdom', os: 'Linux' },
    { ip: '192.168.0.5', timestamp: '2023-06-15 14:20:18', country: 'Germany', os: 'Windows' },
    { ip: '10.0.0.2', timestamp: '2023-06-15 14:15:30', country: 'France', os: 'Mac' },
];

const IPLogs: React.FC = () => {
    const [searchTerm, setSearchTerm] = useState('');

    const filteredLogs = ipLogs.filter(log =>
        log.ip.includes(searchTerm) ||
        log.country.toLowerCase().includes(searchTerm.toLowerCase()) ||
        log.os.toLowerCase().includes(searchTerm.toLowerCase())
    );

    return (
        <div className="flex h-screen bg-gray-100">
            {/* Sidebar */}
            <AdminSideBar />
            <div className="flex-1 p-10 overflow-y-auto">
                <h2 className="text-3xl font-semibold text-gray-800 mb-6">IP Logs</h2>

                <div className="mb-4">
                    <input
                        type="text"
                        placeholder="Search by IP, Country, or OS"
                        className="w-full p-2 border border-gray-300 rounded-md"
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                    />
                </div>

                <div className="bg-white rounded-lg shadow-md overflow-hidden">
                    <table className="min-w-full">
                        <thead className="bg-gray-50">
                            <tr>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">IP Address</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Timestamp</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Country</th>
                                <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">OS</th>
                            </tr>
                        </thead>
                        <tbody className="bg-white divide-y divide-gray-200">
                            {filteredLogs.map((log, index) => (
                                <tr key={index}>
                                    <td className="px-6 py-4 whitespace-nowrap">{log.ip}</td>
                                    <td className="px-6 py-4 whitespace-nowrap">{log.timestamp}</td>
                                    <td className="px-6 py-4 whitespace-nowrap">{log.country}</td>
                                    <td className="px-6 py-4 whitespace-nowrap">{log.os}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    );
};

export default IPLogs;