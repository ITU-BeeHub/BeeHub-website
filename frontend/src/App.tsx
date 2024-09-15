import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Home from './pages/Home'
import Layout from "./components/Layout";
import Login from "./pages/Login";
import AdminPanel from "./pages/AdminPanel";
import DownloadState from "./components/DownloadState";
import IpLogs from "./components/IpLogs";
import PrivateRoute from "./components/PrivateRoute";

function AppContent() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/admin/login" element={<Login />} />

        {/* Admin dashboard rotalarını koruyoruz */}
        <Route path="/admin/dashboard" element={<PrivateRoute><AdminPanel /></PrivateRoute>} />
        <Route path="/admin/download-stats" element={<PrivateRoute><DownloadState /></PrivateRoute>} />
        <Route path="/admin/ip-logs" element={<PrivateRoute><IpLogs /></PrivateRoute>} />
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

export default App;
