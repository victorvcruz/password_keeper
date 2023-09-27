import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Login from './components/Login/Login'; // Importe o componente de Login
import Home from './pages/Home/Home';
import { LoginProvider } from './components/Login/LoginContext';

function App() {
  return (
    <LoginProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </BrowserRouter>
    </LoginProvider>
  );
}

export default App;
