import React from 'react';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { useForm, Controller, SubmitHandler, useFormState } from 'react-hook-form';
import '../auth-page.css';
import { loginValidation, passwordValidation, nickValidation } from '../validation';
import $api from '../../http';

interface ISignUpForm {
    nick: string;
    login: string;
    password: string;
}

interface IComp{
    setRender: React.Dispatch<React.SetStateAction<string>>
}

export const RegForm = ({setRender}:  IComp) => {
    const { handleSubmit, control} = useForm<ISignUpForm>(); // useForm is a custom hook for managing forms with ease. It takes one object as optional argument
    const onSubmit: SubmitHandler<ISignUpForm> = data => $api.post('/user/registration/', {
        headers: {
                'Content-Type': 'application/json'
            },
        body: JSON.stringify({
                login: data.login,
                password: data.password,
                nick: data.nick
            }
        )
    })
    .then(
        (result) => {
            console.log(result);
        });
    // const onSubmit: SubmitHandler<ISignUpForm> = data => fetch("http://localhost:8080/api/v1/user/registration",
    //     {
    //         method: 'POST',
    //         headers: {
    //             'Content-Type': 'application/json'
    //         },
    //         body: JSON.stringify({
    //                 login: data.login,
    //                 password: data.password,
    //                 nick: data.nick
    //             }
    //         )
    //     })
    //     // .then(res => res.json())
    //     .then(
    //         (result) => {
    //             console.log(result);
    //         });// Validation will trigger on the submit event and inputs will attach onChange event listeners to re-validate them.
    const { errors } = useFormState({ // This custom hook allows you to subscribe to each form state, and isolate the re-render at the custom hook level
        control
    })

    return (
        // Typography - Tag for text from MUI
        // Controller - Avoids problems with external controlled components. In our case with MUI
        // TextField - Text Fields let users enter and edit text

        <div className='auth-form'>
            <Typography variant="h4" gutterBottom> 
                Регистрация
            </Typography>
            <Typography variant="subtitle1" gutterBottom className="auth-form__subtitle">
                Прикоснись к будущему
            </Typography>
            <form className="auth-form__form" onSubmit={handleSubmit(onSubmit)}>
                <Controller
                    control={ control } // This object contains methods for registering components into React Hook Form
                    name="nick"
                    rules={nickValidation} // The rules for validation are in the validation.ts file
                    render={({ field }) => ( // render - A function that returns a React element and provides the ability to attach events and value into the component
                        <TextField
                            label="Ник"
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
                    name="login"
                    rules={loginValidation} // The rules for validation are in the validation.ts file
                    render={({ field }) => (
                        <TextField
                            label="Почта"
                            size="small"
                            margin="normal"
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
                    Зарегистрироваться
                </Button>
            </form>
            <div className="auth-form__footer">
                <Typography variant="subtitle1" component="span">
                    Есть аккаунт?{" "}
                </Typography>
                <Typography variant="subtitle1" component="span" className='link' onClick={() => setRender("auth")}>
                    Войдите
                </Typography>
            </div>
        </div>
    )
}