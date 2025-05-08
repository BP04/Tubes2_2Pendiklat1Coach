import { BrowserRouter as Router, Routes, Route, Link, useLocation } from 'react-router-dom';
import Home from './components/Home';
import Main from './components/Main';
import AboutUs from './components/AboutUs';
import './index.css';

function App() {
  const location = useLocation();
  const isHomePage = location.pathname === '/';

  return (
    <div>
      <header className="header">
        <nav className="flex justify-center gap-4">
          <Link to="/" className="nav-link">
            Home
          </Link>
        
            <Link to="/about" className="nav-link">
              About Us
            </Link>
        
            <Link to="/main" className="nav-link">
              Recipe Finder
            </Link>

        </nav>
      </header>
      <div className="app">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/main" element={<Main />} />
          <Route path="/about" element={<AboutUs />} />
        </Routes>
      </div>
    </div>
  );
}

function AppWrapper() {
  return (
    <Router>
      <App />
    </Router>
  );
}

export default AppWrapper;