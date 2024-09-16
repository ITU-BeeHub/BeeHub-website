import React, { useState } from 'react';
import { BiLogoWindows } from "react-icons/bi";
import { FaApple } from "react-icons/fa";

// Custom Button component
const Button: React.FC<React.ButtonHTMLAttributes<HTMLButtonElement>> = ({ children, className, ...props }) => (
    <button
        className={`bg-[hsl(var(--primary))] hover:bg-[hsl(var(--primary)/90)] text-white text-lg font-bold py-3 px-8 rounded-full shadow-lg transition-all duration-300 ease-in-out ${className}`}
        {...props}
    >
        {children}
    </button>
);

export default function Home() {
    const [selectedOS, setSelectedOS] = useState("windows");

    // İndirme butonuna basıldığında backend'e GET isteği gönderme fonksiyonu
    const handleDownload = async () => {
        const downloadUrl = `https://beehubapp.com/api/download?os=${selectedOS}`; // Backend URL'sini güncelledik

        try {
            // İndirme işlemini başlatmak için yeni bir pencere veya sekme açabiliriz
            window.location.href = downloadUrl;
        } catch (error) {
            console.error("Download error:", error);
            alert("An error occurred while downloading. Please try again.");
        }
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-white relative overflow-hidden mt-[-80px]">
            <div>
                <h1 className="mb-4 text-4xl font-bold text-[#212121] z-10">Get BeeHub!</h1>
            </div>
            <div className="flex flex-col items-center z-10">
                <div className="flex items-center justify-center space-x-0 rounded-full overflow-hidden shadow-md  mt-8 mb-6">
                    <button
                        className={`p-4 w-14 h-11 flex items-center justify-center transition-colors ${selectedOS === "windows" ? "bg-[#FDC003] text-[#0372CE]" : "bg-[#E0E0E0] text-[#737373]"
                            }`}
                        onClick={() => setSelectedOS("windows")}
                    >
                        <BiLogoWindows className="h-6 w-6" />
                    </button>
                    <button
                        className={`p-4 w-14 h-11 flex items-center justify-center transition-colors ${selectedOS === "mac" ? "bg-[#FDC003] text-[#0372CE]" : "bg-[#E0E0E0] text-[#737373]"
                            }`}
                        onClick={() => setSelectedOS("mac")}
                    >
                        <FaApple className="h-6 w-6" />
                    </button>
                </div>
                <Button className="px-8 py-4 text-xl font-bold rounded-lg" onClick={handleDownload}>
                    Download {selectedOS === "windows" ? "Windows" : "Mac"}
                </Button>
            </div>
        </div>
    );
}
