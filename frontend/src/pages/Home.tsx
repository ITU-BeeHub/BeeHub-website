import React, { useState } from 'react';
import { BiLogoWindows } from "react-icons/bi";
import { FaApple } from "react-icons/fa";
import { Link } from 'react-router-dom';

const Button: React.FC<React.ButtonHTMLAttributes<HTMLButtonElement>> = ({ children, className, ...props }) => (
    <button
        className={`inline-flex items-center justify-center transition-all duration-300 ease-in-out ${className}`}
        {...props}
    >
        {children}
    </button>
);

export default function Home() {
    const [selectedOS, setSelectedOS] = useState("windows");

    const handleDownload = async () => {
        const downloadUrl = `https://beehubapp.com/api/download?os=${selectedOS}`;
        try {
            window.location.href = downloadUrl;
        } catch (error) {
            console.error("Download error:", error);
            alert("An error occurred while downloading. Please try again.");
        }
    };

    return (
        <div className="flex flex-col items-center justify-center h-[calc(100vh-108px)] bg-gradient-to-b from-white to-gray-50 relative overflow-hidden">
            <div className="max-w-4xl mx-auto px-4 py-16 sm:px-6 lg:px-8 text-center">
                <div className="mb-12">
                    <h1 className="text-5xl font-extrabold text-gray-900 mb-6">
                        Welcome to <span className="text-[#0372CE]">Bee</span><span className="text-[#FDC003]">Hub</span>
                    </h1>
                    <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
                        Your ultimate university companion for managing your academic journey
                    </p>
                </div>

                <div className="space-y-8">
                    <div className="inline-flex items-center justify-center p-1 rounded-full bg-gray-100">
                        <button
                            className={`flex items-center space-x-2 px-6 py-3 rounded-full transition-all duration-300 ${selectedOS === "windows"
                                ? "bg-white shadow-md text-[#0372CE]"
                                : "text-gray-500 hover:bg-white/50"
                                }`}
                            onClick={() => setSelectedOS("windows")}
                        >
                            <BiLogoWindows className="h-6 w-6 " />
                            <span>Windows</span>
                        </button>
                        <button
                            className={`flex items-center space-x-2 px-6 py-3 rounded-full transition-all duration-300 ${selectedOS === "mac"
                                ? "bg-white shadow-md text-[#0372CE]"
                                : "text-gray-500 hover:bg-white/50"
                                }`}
                            onClick={() => setSelectedOS("mac")}
                        >
                            <FaApple className="h-6 w-6" />
                            <span>macOS</span>
                        </button>
                    </div>

                    <div className="space-y-4">
                        <Button
                            onClick={handleDownload}
                            className="bg-[#0372CE] hover:opacity-90 text-white px-8 py-4 rounded-2xl text-lg font-medium shadow-md hover:shadow-lg w-full sm:w-auto"
                        >
                            Download for {selectedOS === "windows" ? "Windows" : "macOS"}
                        </Button>

                        <div className="flex justify-center">
                            <Link to="/whats-new">
                                <Button className="text-gray-600 hover:text-[#FDC003] px-6 py-3 rounded-lg font-medium transition-colors duration-300">
                                    See what's new in the latest version →
                                </Button>
                            </Link>
                        </div>
                    </div>
                </div>
            </div>

            {/* Subtle decorative elements */}
            <div className="absolute top-1/2 left-0 w-64 h-64 bg-[#0372CE]/5 rounded-full filter blur-3xl"></div>
            <div className="absolute bottom-0 right-0 w-96 h-96 bg-[#FDC003]/5 rounded-full filter blur-3xl"></div>
        </div>
    );
}
