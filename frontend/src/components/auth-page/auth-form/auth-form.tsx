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
    const { handleSubmit, control} = useForm<ISignInForm>();
    const onSubmit: SubmitHandler<ISignInForm> = data => console.log(data);
    const { errors } = useFormState({ 
        control
    })

    return (
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
                    rules={loginValidation}
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
                    rules={passwordValidation}
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
                    type="submit" // вызов события отправки формы
                    variant='contained'
                    fullWidth={ true }
                    disableElevation={ true } // отключение теней
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