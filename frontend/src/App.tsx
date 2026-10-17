import { lazy, Suspense } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Home from './pages/Home'
import Layout from "./components/Layout";
import Documentation from "./pages/Documentation";
import BeePickerDoc from "./pages/BeePickerDoc";
import BeeArchiveDoc from "./pages/BeeArchiveDoc";
import Contact from "./pages/Contact";
import PageNotFound from "./pages/PageNotFound";
import WhatsNew from "./pages/WhatsNew";

const AdminApp = lazy(() => import("./admin/AdminApp"));

function App() {
  return (
    <Router>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
          <Route path="/documentation" element={<Documentation />} />
          <Route path="/beepicker" element={<BeePickerDoc />} />
          <Route path="/beearchive" element={<BeeArchiveDoc />} />
          <Route path="/contact" element={<Contact />} />
          <Route path="/whats-new" element={<WhatsNew />} />
          <Route path="*" element={<PageNotFound />} />
        </Route>
        <Route path="/admin/*" element={<Suspense fallback={null}><AdminApp /></Suspense>} />
      </Routes>
    </Router>
  );
}

export default App;
