// Navbar.js
import React from 'react';
import { Link } from 'react-router-dom';
import logo from "../assets/logo.png";

const Navbar = () => {
  return (
    <nav className="bg-green-500 p-4 fixed top-0 w-full z-10">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="flex items-center space-x-2">
          <img
            src={logo}  
            alt="Logo"
            className="w-12 h-12 rounded-full"
          />
          <span className="text-white font-bold text-lg">Hask.inc</span>
        </Link>

        {/* Navigation Links */}
        <div className="space-x-4">
          <Link to="/" className="text-white">Home</Link>
          <Link to="/contact" className="text-white">Contact</Link>
          <Link to="/analytics" className="text-white">Analytics</Link>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
