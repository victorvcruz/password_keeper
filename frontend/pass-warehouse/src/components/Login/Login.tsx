import React, { useState } from 'react';
import './Login.css';
import { useLogin } from './LoginContext';
import { useNavigate } from "react-router-dom";

import { Google, Github } from "react-bootstrap-icons";


function Login() {
    const navigate = useNavigate();

    const { user, login } = useLogin();
    const initialState = {
        email: '',
        password: '',
    };

    const [inputs, setInputs] = useState(initialState);

    const handleLogin = () => {
        const loginData = {
            email: inputs.email,
            password: inputs.password,
        };
        login(loginData);
        navigate('/');
    };

    const handleOnchange = (text: React.ChangeEvent<HTMLInputElement>, input: string) => {
        setInputs(prevState => ({ ...prevState, [input]: text.target.value }));
        console.log(user);
    };


    return (
        <div>
            <section className="vh-100">
                <div className="container-fluid h-custom center-section">
                    <div className="row d-flex justify-content-center align-items-center h-100">
                        <div className="col-md-8 col-lg-6 col-xl-4 offset-xl-1">

                            <form>
                                <div className="d-flex flex-row align-items-center justify-content-center align-items-center">
                                    <p className="lead fw-normal mb-0 me-3">Sign in with</p>
                                </div>

                                <div className="form-outline mb-4">
                                    <input type="email" id="form3Example3" className="form-control form-control-lg"
                                        placeholder="Enter a valid email address" onChange={text => handleOnchange(text, 'email')} />
                                    <label className="form-label" htmlFor="form3Example3">Email address</label>
                                </div>

                                <div className="form-outline mb-3">
                                    <input type="password" id="form3Example4" className="form-control form-control-lg"
                                        placeholder="Enter password" onChange={text => handleOnchange(text, 'password')} />
                                    <label className="form-label" htmlFor="form3Example4">Password</label>
                                </div>

                                <div className="d-flex justify-content-between align-items-center">
                                    {/* Checkbox */}
                                    <div className="form-check mb-0">
                                        <input className="form-check-input me-2" type="checkbox" value="" id="form2Example3" />
                                        <label className="form-check-label" htmlFor="form2Example3">
                                            Remember me
                                        </label>
                                    </div>
                                    <a href="#!" className="text-body">Forgot password?</a>
                                </div>

                                <div className="text-center text-lg-start mt-4 pt-2">
                                    <button type="button" className="btn btn-primary btn-lg"
                                        style={{ paddingLeft: '2.5rem', paddingRight: '2.5rem' }} onClick={handleLogin}>Login</button>
                                    <p className="small fw-bold mt-2 pt-1 mb-0">Don't have an account? <a href="#!"
                                        className="link-danger">Register</a></p>
                                </div>


                                <div className="divider d-flex align-items-center my-4">
                                    <p className="text-center fw-bold mx-3 mb-0">Or</p>
                                </div>

                                <div className="d-flex flex-row align-items-center justify-content-center align-items-center">
                                    <button type="button" className="btn btn-primary btn-floating mx-3">
                                        <Google />
                                    </button>

                                    <button type="button" className="btn btn-primary btn-floating mx-3">
                                        <Github />
                                    </button>
                                </div>

                            </form>
                        </div>
                    </div>
                </div>
            </section>
        </div>
    );
}

export default Login;
