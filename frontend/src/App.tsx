
import { HashRouter as Router, Route, Routes } from "react-router-dom";

import Home from './pages/Home'
import Layout from "./components/Layout";


function AppContent() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
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
