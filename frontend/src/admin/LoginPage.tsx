import React, { useState } from "react";
import { Navigate, useLocation } from "react-router-dom";
import BeakerIcon from "../components/icons/BeakerIcon";
import { login } from "./api";
import { saveSession, useSession } from "./session";

export default function LoginPage() {
    const session = useSession();
    const location = useLocation();
    const redirectTo = (location.state as { from?: string } | null)?.from ?? "/admin";

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [submitting, setSubmitting] = useState(false);

    if (session) return <Navigate to={redirectTo} replace />;

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        setSubmitting(true);
        setError("");
        try {
            // Saving the session re-renders this page, which then redirects.
            saveSession(await login(username, password));
        } catch (err) {
            setError(err instanceof Error ? err.message : "Sign in failed");
            setSubmitting(false);
        }
    };

    const inputClass = "mt-1 block w-full rounded-lg border border-slate-300 px-3 py-2 text-slate-900 shadow-sm focus:border-[#0372CE] focus:outline-none focus:ring-2 focus:ring-[#0372CE]/20";

    return (
        <div className="flex min-h-screen items-center justify-center bg-slate-50 px-4">
            <div className="w-full max-w-sm">
                <div className="mb-6 flex items-center justify-center gap-2">
                    <BeakerIcon className="h-9 w-9 text-[#FDC003]" />
                    <span className="text-2xl font-bold text-slate-900">BeeHub Admin</span>
                </div>

                <form onSubmit={handleSubmit} className="space-y-4 rounded-xl border border-slate-200 bg-white p-6 shadow-sm">
                    <label className="block text-sm font-medium text-slate-700">
                        Username
                        <input
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            autoComplete="username"
                            autoFocus
                            required
                            className={inputClass}
                        />
                    </label>
                    <label className="block text-sm font-medium text-slate-700">
                        Password
                        <input
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            autoComplete="current-password"
                            required
                            className={inputClass}
                        />
                    </label>

                    {error && <p role="alert" className="text-sm text-red-600">{error}</p>}

                    <button
                        type="submit"
                        disabled={submitting}
                        className="w-full rounded-lg bg-[#0372CE] py-2 font-medium text-white hover:opacity-90 disabled:opacity-60"
                    >
                        {submitting ? "Signing in…" : "Sign in"}
                    </button>
                </form>
            </div>
        </div>
    );
}
