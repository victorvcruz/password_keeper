import React from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

import Login from './components/Login/Login'; // Importe o componente de Login
import Home from './components/Home/Home'; // Importe o componente de Login

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/home" element={<Home />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;