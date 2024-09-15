import React, { useEffect, useState } from 'react';
import AdminSideBar from './AdminSidebar';
import {
    BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer, LineChart, Line
} from 'recharts';

type OsDataItem = {
    name: string;
    downloads: number;
};

type MonthlyDataItem = {
    name: string;
    downloads: number;
};

const DownloadStats: React.FC = () => {
    const [monthlyData, setMonthlyData] = useState<MonthlyDataItem[]>([]);
    const [osData, setOsData] = useState<OsDataItem[]>([]);

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

                // Aylık verileri çekmek için
                const monthlyResponse = await fetch('http://localhost:8080/admin/monthly-download-stats', { headers });
                const monthlyDataJson: { month: string; count: number }[] = await monthlyResponse.json();
                console.log('Monthly data:', monthlyDataJson);
                const monthlyChartData: MonthlyDataItem[] = monthlyDataJson.map((item) => ({
                    name: new Date(item.month).toLocaleString('default', { month: 'short', year: 'numeric' }),
                    downloads: item.count,
                }));
                setMonthlyData(monthlyChartData);

            } catch (error) {
                console.error('Error fetching download stats:', error);
            }
        };

        fetchData();
    }, []);

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
            </div>
        </div>
    );
};

export default DownloadStats;
