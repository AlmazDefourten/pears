import React from 'react';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { useForm, Controller, SubmitHandler, useFormState } from 'react-hook-form';
import './auth-form.css';
import { loginValidation, passwordValidation } from './validation';

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
            <form className="auth-form__form" onSubmit={handleSubmit(onSubmit)}>
                <Controller
                    control={ control }
                    name="login"
                    rules={loginValidation}
                    render={({ field }) => (
                        <TextField
                            label="Логин"
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
                    type="submit"
                    variant='contained'
                    fullWidth={ true }
                    disableElevation={ true }
                    sx={{
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