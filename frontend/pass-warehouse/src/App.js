import React from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

import Login from './Login'; // Importe o componente de Login
import NullPage from './NullPage'; // Importe o componente de Login

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/null-page" element={<NullPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;