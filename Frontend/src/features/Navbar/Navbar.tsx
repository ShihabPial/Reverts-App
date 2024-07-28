import React, { useState } from 'react';
import './Navbar.scss';
import Hamburger from 'hamburger-react'

const Navbar: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <nav className="navbar">
        <Hamburger toggled={isOpen} toggle={setIsOpen} />
        {isOpen && (
        <ul className="menu">
          <li><a href="#home">Home</a></li>
          <li><a href="#about">About</a></li>
          <li><a href="#contact">Contact</a></li>
        </ul>
      )}
    </nav>
  );
};

export default Navbar;
