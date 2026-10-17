import { Outlet } from "react-router-dom";
import Footer from "./Footer";
import Header from "./Header";

export default function Layout() {
    return (
        <div className="min-h-screen flex flex-col">
            <Header />
            <main className="flex-grow overflow-auto">
                <Outlet />
            </main>
            <Footer />
        </div>
    );
}