import React from 'react';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { useForm, Controller, SubmitHandler, useFormState } from 'react-hook-form';
import '../auth-page.css';
import { loginValidation, passwordValidation } from '../validation';

interface ISignInForm {
    login: string;
    password: string;
}

interface IComp{
    setRender: React.Dispatch<React.SetStateAction<string>>
}

export const AuthForm = ({setRender}:  IComp) => {
    const { handleSubmit, control} = useForm<ISignInForm>(); // useForm is a custom hook for managing forms with ease. It takes one object as optional argument
    const onSubmit: SubmitHandler<ISignInForm> = data => fetch("http://localhost:8080/api/v1/user/authorization",
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                    login: data.login,
                    password: data.password,
                }
            )
        })
        .then(res => res.json())
        .then(
            (result) => {
                console.log(result);
            },
            (error) => {
                console.log(error);
            }); // Validation will trigger on the submit event and inputs will attach onChange event listeners to re-validate them.
    const { errors } = useFormState({ // This custom hook allows you to subscribe to each form state, and isolate the re-render at the custom hook level
        control
    })

    return (
        // Typography - Tag for text from MUI
        // Controller - Avoids problems with external controlled components. In our case with MUI
        // TextField - Text Fields let users enter and edit text

        <div className='auth-form'>
            <Typography variant="h4" gutterBottom>
                Вход
            </Typography>
            <Typography variant="subtitle1" gutterBottom className="auth-form__subtitle"> 
                Прикоснись к будущему
            </Typography>
            {/* handleSubmit - This function will receive the form data if form validation is successful */}
            <form className="auth-form__form" onSubmit={handleSubmit(onSubmit)}>
                <Controller
                    control={ control } // This object contains methods for registering components into React Hook Form
                    name="login" 
                    rules={loginValidation} // The rules for validation are in the validation.ts file
                    render={({ field }) => ( // render - A function that returns a React element and provides the ability to attach events and value into the component
                        <TextField
                            label="Логин"
                            size="small"
                            className="auth-form__input"
                            fullWidth={ true }
                            onChange={(e) => field.onChange(e)}
                            value={ field.value }
                            error={ !!errors.login?.message }
                            helperText={ errors?.login?.message }
                        />
                    )}
                />
                <Controller
                    control={ control }
                    name="password"
                    rules={passwordValidation} // The rules for validation are in the validation.ts file
                    render={({ field }) => (
                        <TextField
                            label="Пароль"
                            type="password"
                            size="small"
                            margin="normal"
                            className="auth-form__input"
                            fullWidth={ true }
                            onChange={(e) => field.onChange(e)}
                            value={ field.value }
                            error={ !!errors?.password?.message }
                            helperText={ errors?.password?.message }
                        />
                    )}
                />
                
                <Button
                    type="submit"
                    variant='contained'
                    fullWidth={ true }
                    disableElevation={ true }
                    sx={{ // The system prop that allows defining system overrides as well as additional CSS styles
                        marginTop: 2
                    }} 
                >
                    Войти
                </Button>
            </form>
            <div className="auth-form__footer">
                <Typography variant="subtitle1" component="span">
                    Нету аккаунта?{" "}
                </Typography>
                <Typography variant="subtitle1" component="span" className='link' onClick={() => setRender("reg")}>
                    Зарегистрируйтесь
                </Typography>
            </div>
        </div>
    )
}