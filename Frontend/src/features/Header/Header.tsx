import React from 'react';
import './Header.scss';
import Navbar from '../Navbar/Navbar';

const Header: React.FC = () => {
  return (
    <header className="header">
      <div className="logo">
        {/* Add your logo here */}
        <img src="path/to/logo.png" alt="Logo" />
      </div>
      <div className="app-name">
        <h1>Reverts App</h1>
      </div>
      <Navbar />
    </header>
  );
};

export default Header;
