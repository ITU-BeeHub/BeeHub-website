import { useState } from "react"

export default function Login() {

    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const handleLogin = (e: any) => {
        e.preventDefault()
        console.log("Logging in with email:", email, "and password:", password)
    }
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
                    <div className="flex items-center justify-center"> {/* Buton ortalandı */}
                        <button
                            type="submit"
                            className="bg-[#FDC003] hover:bg-[#e8b400] text-[#0372CE] font-bold py-2 px-8 rounded-lg focus:outline-none focus:shadow-outline" /* Köşeler yuvarlatıldı */
                        >
                            Sign In
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}
