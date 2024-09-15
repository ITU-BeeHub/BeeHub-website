import React, { useState, useEffect } from 'react';
import AdminSideBar from './AdminSidebar';

type IPLog = {
    ip: string;
    timestamp: string;
    country: string;
    os: string;
};

type DownloadLog = {
    IPAddress: string;
    CreatedAt: string;
    Country?: string;
    OS: string;
};

const IPLogs: React.FC = () => {
    const [searchTerm, setSearchTerm] = useState('');
    const [ipLogs, setIpLogs] = useState<IPLog[]>([]);

    useEffect(() => {
        const fetchIPLogs = async () => {
            try {
                const token = sessionStorage.getItem("token");
                const headers = {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                };

                const response = await fetch('http://localhost:8080/admin/ip-logs', { headers });
                const data: DownloadLog[] = await response.json();

                console.log("API'den gelen veriler:", data);

                // Gelen verinin bir dizi olup olmadığını kontrol edin
                if (Array.isArray(data)) {
                    const logs: IPLog[] = data.map((log) => ({
                        ip: log.IPAddress || 'Unknown',
                        timestamp: log.CreatedAt ? new Date(log.CreatedAt).toLocaleString() : 'Unknown',
                        country: log.Country || 'Unknown',
                        os: log.OS || 'Unknown',
                    }));
                    setIpLogs(logs);
                } else {
                    console.error('Beklenmeyen veri formatı:', data);
                }
            } catch (error) {
                console.error('Error fetching IP logs:', error);
            }
        };

        fetchIPLogs();
    }, []);

    const filteredLogs = ipLogs.filter((log) =>
        (log.ip || '').includes(searchTerm) ||
        (log.country || '').toLowerCase().includes(searchTerm.toLowerCase()) ||
        (log.os || '').toLowerCase().includes(searchTerm.toLowerCase())
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
                            {filteredLogs.length > 0 ? (
                                filteredLogs.map((log, index) => (
                                    <tr key={index}>
                                        <td className="px-6 py-4 whitespace-nowrap">{log.ip}</td>
                                        <td className="px-6 py-4 whitespace-nowrap">{log.timestamp}</td>
                                        <td className="px-6 py-4 whitespace-nowrap">{log.country}</td>
                                        <td className="px-6 py-4 whitespace-nowrap">{log.os}</td>
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan={4} className="px-6 py-4 text-center">
                                        No logs found.
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    );
};

export default IPLogs;
