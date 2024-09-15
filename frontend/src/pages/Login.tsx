import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e: any) => {
        e.preventDefault();
        try {
            const response = await fetch("http://localhost:8080/admin/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username: email, password }),
            });
            const data = await response.json();

            if (response.ok) {
                // JWT token'ı sessionStorage'a kaydediyoruz
                sessionStorage.setItem("token", data.token);

                // Admin dashboard'a yönlendir
                navigate("/admin/dashboard");
            } else {
                setError(data.error || "Login failed. Please check your credentials.");
            }
        } catch (err) {
            console.error("Login error:", err);
            setError("An error occurred. Please try again.");
        }
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-white relative overflow-hidden">
            <div className="absolute inset-0 bg-repeat bg-[url('/bee.svg')] opacity-10" />
            <div className="absolute inset-0 bg-repeat bg-[url('/bee.svg')] opacity-5" />
            <div className="max-w-md w-full z-10 px-6 py-8 bg-white rounded-lg shadow-md">
                <h1 className="mb-4 text-4xl font-bold text-[#212121]">Admin Login</h1>
                <form onSubmit={handleLogin}>
                    <div className="mb-4">
                        <label htmlFor="email" className="block text-[#212121] font-bold mb-2">
                            Email
                        </label>
                        <input
                            id="email"
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-[#212121] leading-tight focus:outline-none focus:shadow-outline"
                            placeholder="Enter your email"
                        />
                    </div>
                    <div className="mb-6">
                        <label htmlFor="password" className="block text-[#212121] font-bold mb-2">
                            Password
                        </label>
                        <input
                            id="password"
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="shadow appearance-none border rounded w-full py-2 px-3 text-[#212121] leading-tight focus:outline-none focus:shadow-outline"
                            placeholder="Enter your password"
                        />
                    </div>
                    {error && <p className="text-red-500">{error}</p>}
                    <div className="flex items-center justify-center">
                        <button
                            type="submit"
                            className="bg-[#FDC003] hover:bg-[#e8b400] text-[#0372CE] font-bold py-2 px-8 rounded-lg focus:outline-none focus:shadow-outline"
                        >
                            Sign In
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}
