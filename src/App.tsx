// src/App.tsx
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import Layout from "./components/Layout";
import Beepicker from "./pages/BeePicker";
import BeeSync from "./pages/BeeSync";
import BeeArchive from "./pages/BeeArchive";
import BeeCalendar from "./pages/BeeCalendar";
import BeeChat from "./pages/BeeChat";
import PageNotFound from "./pages/PageNotFound";
import Contact from "./pages/Contact";
import PrivacyPolicy from "./pages/PrivacyPolicy";
import TermsOfService from "./pages/TermsOfService";

function App() {
  return (
      <Router>
        <Layout>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/beepicker" element={<Beepicker />} />
            <Route path="/beesync" element={<BeeSync />} />
            <Route path="/beecalendar" element={<BeeCalendar />} />
            <Route path="/beearchive" element={<BeeArchive />} />
            <Route path="/beeChat" element={<BeeChat />} />

            {/* Version route which returns a json*/}
            {/* It can be found at https://url/version.tsx */}

            <Route path="/contact" element={<Contact />} />
            <Route path="/privacy" element={<PrivacyPolicy />} />
            <Route path="/terms" element={<TermsOfService />} />

            <Route path="/*" element={<PageNotFound />} />
          </Routes>
        </Layout>
      </Router>
  );
}

export default App;
