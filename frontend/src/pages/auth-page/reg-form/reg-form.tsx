import React from 'react';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import { useForm, Controller, SubmitHandler, useFormState } from 'react-hook-form';
import './reg-form.css';
import { loginValidation, passwordValidation, nickValidation } from './validation';

interface ISignUpForm {
    nick: string;
    login: string;
    password: string;
}

interface IComp{
    setRender: React.Dispatch<React.SetStateAction<string>>
}

export const RegForm = ({setRender}:  IComp) => {
    const { handleSubmit, control} = useForm<ISignUpForm>();
    const onSubmit: SubmitHandler<ISignUpForm> = data => console.log(data);
    const { errors } = useFormState({ 
        control
    })

    return (
        <div className='auth-form'>
            <Typography variant="h4" gutterBottom>
                Регистрация
            </Typography>
            <Typography variant="subtitle1" gutterBottom className="auth-form__subtitle">
                Прикоснись к будущему
            </Typography>
            <form className="auth-form__form" onSubmit={handleSubmit(onSubmit)}>
                <Controller
                    control={ control }
                    name="nick"
                    rules={nickValidation}
                    render={({ field }) => (
                        <TextField
                            label="Ник"
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
                    name="login"
                    rules={loginValidation}
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
                    Зарегистрироваться
                </Button>
            </form>
            <div className="auth-form__footer">
                <Typography variant="subtitle1" component="span">
                    Есть аккаунт?{" "}
                </Typography>
                <Typography variant="subtitle1" component="span" className='link'onClick={() => setRender("auth")}>
                    Войдите
                </Typography>
            </div>
        </div>
    )
}