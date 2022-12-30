import React, { useState } from 'react';
import './auth-page.css';
import { AuthForm } from './auth-form';
import { RegForm } from './reg-form';

export const AuthPage: React.FC = () => {
    const [rendr, setRender] = useState("auth");
    

    return (
        <div className='auth-page'>
            {(rendr === "auth") ? <AuthForm setRender={setRender}/> : <RegForm setRender={setRender}/>}
        </div>
    )
}