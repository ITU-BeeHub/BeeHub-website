import React, { useEffect, useState } from 'react';
import {
    BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer
} from 'recharts';
import AdminSideBar from '../components/AdminSidebar';

type OsDataItem = {
    name: string;
    downloads: number;
};

type RecentIP = {
    ip: string;
    timestamp: string;
};

const AdminPanel: React.FC = () => {
    const [osData, setOsData] = useState<OsDataItem[]>([]);
    const [recentIPs, setRecentIPs] = useState<RecentIP[]>([]);
    const [totalDownloads, setTotalDownloads] = useState<number>(0);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const token = sessionStorage.getItem("token");
                const headers = {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                };

                // OS verilerini çekmek için
                const osResponse = await fetch('http://localhost:8080/admin/download-stats', { headers });
                const osDataJson: { [key: string]: number } = await osResponse.json();
                const osChartData: OsDataItem[] = Object.keys(osDataJson).map(os => ({
                    name: os,
                    downloads: osDataJson[os],
                }));
                setOsData(osChartData);

                // Toplam indirme sayısını hesaplayın
                const total = Object.values(osDataJson).reduce((sum: number, count: number) => sum + count, 0);
                setTotalDownloads(total);

                // Recent IP logs verilerini çekmek için
                const ipResponse = await fetch('http://localhost:8080/admin/ip-logs', { headers });
                const ipDataJson: any[] = await ipResponse.json();
                // IP loglarını uygun formata dönüştürün
                const ipLogs: RecentIP[] = ipDataJson.map((log: any) => ({
                    ip: log.ip_address,
                    timestamp: new Date(log.created_at).toLocaleString(),
                }));
                setRecentIPs(ipLogs);

            } catch (error) {
                console.error('Error fetching admin data:', error);
            }
        };

        fetchData();
    }, []);

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
                    <p className="text-4xl font-bold text-blue-600">{totalDownloads}</p>
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
