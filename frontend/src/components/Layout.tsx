import Footer from "./Footer";
import Header from "./Header";

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <div className=" h-screen overflow-hidden flex-col">

            <Header />
            <main className="flex-1 overflow-y-auto   bg-[#F5FDFD] justify-center">
                {children}
            </main>
            <Footer />
        </div>
    );
}