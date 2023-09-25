import React from 'react'
import { useLocation } from "react-router-dom";


const Home = () => {

    const { state } = useLocation();

    return (
        <div>
            <h1>Home</h1>
            <p>Email from the previous page: {state.inputs.email}</p>
            <p>Password from the previous page: {state.inputs.password}</p>
        </div>
    );
}

export default Home;

