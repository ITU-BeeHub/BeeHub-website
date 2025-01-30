import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Home from './pages/Home'
import Layout from "./components/Layout";
import Login from "./pages/Login";
import AdminPanel from "./pages/AdminPanel";
import DownloadState from "./components/DownloadState";
import IpLogs from "./components/IpLogs";
import PrivateRoute from "./components/PrivateRoute";
import Documentation from "./pages/Documentation";
import BeePickerDoc from "./pages/BeePickerDoc";
import BeeArchiveDoc from "./pages/BeeArchiveDoc";
import Contact from "./pages/Contact";
import PageNotFound from "./pages/PageNotFound";
import WhatsNew from "./pages/WhatsNew";

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
        <Route path="/documentation" element={<Documentation />} />
        <Route path="/beepicker" element={<BeePickerDoc />} />
        <Route path="/beearchive" element={<BeeArchiveDoc />} />
        <Route path="/contact" element={<Contact />} />
        <Route path="/whats-new" element={<WhatsNew />} />
        <Route path="*" element={<PageNotFound />} /> {/* Catch-all for undefined routes */}
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
