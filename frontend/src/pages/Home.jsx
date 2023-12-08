// src/components/LandingPage.js
import React from 'react';
import { Link } from 'react-router-dom';
import logo from '../assets/logo.png';

const Home = () => {
  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-100">
      <img src={logo} alt="Logo" className="w-24 h-24 rounded-full mb-4" />
      <h1 className="text-4xl font-bold text-gray-800 mb-2">Welcome to Hask.inc</h1>
      <p className="text-lg text-gray-600 mb-8">Your Creative Advertising Partner</p>
      <Link to="/contact" className="bg-green-500 text-white py-2 px-4 rounded-full hover:bg-green-600 transition duration-300">
        Contact Us
      </Link>
    </div>
  );
};

export default Home;
