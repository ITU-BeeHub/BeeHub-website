import { NavLink, Outlet } from "react-router-dom";
import BeakerIcon from "../components/icons/BeakerIcon";
import { clearSession } from "./session";

const NAVIGATION = [
    { to: "/admin", label: "Overview", end: true },
    { to: "/admin/downloads", label: "Download log", end: false },
];

export default function AdminLayout() {
    return (
        <div className="min-h-screen bg-slate-50 md:flex">
            <aside className="border-b border-slate-200 bg-white md:sticky md:top-0 md:flex md:h-screen md:w-60 md:flex-col md:border-b-0 md:border-r">
                <div className="flex items-center justify-between px-5 py-4 md:block">
                    <div className="flex items-center gap-2">
                        <BeakerIcon className="h-7 w-7 text-[#FDC003]" />
                        <span className="text-lg font-bold text-slate-900">BeeHub</span>
                        <span className="rounded bg-slate-100 px-1.5 py-0.5 text-xs font-medium text-slate-500">Admin</span>
                    </div>
                    <button onClick={clearSession} className="text-sm text-slate-500 hover:text-slate-900 md:hidden">
                        Sign out
                    </button>
                </div>

                <nav className="flex gap-1 px-3 pb-3 md:flex-1 md:flex-col md:pb-0">
                    {NAVIGATION.map((item) => (
                        <NavLink
                            key={item.to}
                            to={item.to}
                            end={item.end}
                            className={({ isActive }) => `rounded-lg px-3 py-2 text-sm font-medium transition-colors ${isActive
                                ? "bg-[#0372CE]/10 text-[#0372CE]"
                                : "text-slate-600 hover:bg-slate-100 hover:text-slate-900"}`}
                        >
                            {item.label}
                        </NavLink>
                    ))}
                </nav>

                <div className="hidden border-t border-slate-100 px-3 py-3 md:block">
                    <a href="/" className="block rounded-lg px-3 py-2 text-sm text-slate-500 hover:bg-slate-100 hover:text-slate-900">
                        View website
                    </a>
                    <button onClick={clearSession} className="w-full rounded-lg px-3 py-2 text-left text-sm text-slate-500 hover:bg-slate-100 hover:text-slate-900">
                        Sign out
                    </button>
                </div>
            </aside>

            <main className="min-w-0 flex-1 px-4 py-6 md:px-8 md:py-8">
                <div className="mx-auto max-w-6xl">
                    <Outlet />
                </div>
            </main>
        </div>
    );
}
