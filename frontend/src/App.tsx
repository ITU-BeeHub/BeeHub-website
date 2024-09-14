
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Home from './pages/Home'
import Layout from "./components/Layout";
import Login from "./pages/Login";
import AdminPanel from "./pages/AdminPanel";
import DownloadState from "./components/DownloadState";
import IpLogs from "./components/IpLogs";


function AppContent() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admin/login" element={<Login />} />
        <Route path="/admin/dashboard" element={<AdminPanel />} />
        <Route path="/admin/download-stats" element={<DownloadState />} />
        <Route path="/admin/ip-logs" element={<IpLogs />} />
      </Routes>
    </Layout>
  );
}



function App() {
  return (
    <Router>
      <AppContent />
    </Router>
  );
}

export default App
