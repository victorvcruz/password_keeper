import React from 'react';
import { useLogin } from '../../components/Login/LoginContext';
import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();
  const { user, logout } = useLogin();

  const login = () => {
    navigate('/login');
  };

  return (
    <div>
      {user ? (
        <div>
          <h1>Bem-vindo, {user.email}!</h1>
          <button onClick={logout}>Sair</button>
        </div>
      ) : (
        <div>
          <h1>Bem-vindo!</h1>
          <p>Faça login para continuar.</p>
          <button onClick={login}>Login</button>
        </div>
      )}
    </div>
  );
}

export default Home;