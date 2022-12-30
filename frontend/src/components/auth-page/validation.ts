const REQUIRED_FIELD = 'Обязательно для заполнения';

export const loginValidation = {
    required: REQUIRED_FIELD, // обязательный для дополнения
    validate: (value: string) => { // доп. условия
        if(value.match(/[а-яА-Я]/)) {
            return 'Логин не может содержать русские буквы'
        }

        return true;
    }
};

export const passwordValidation = {
    required: REQUIRED_FIELD,
    validate: (value: string) => {
        if(value.length < 6) {
            return 'Пароль должен быть длиннее 6-ти символов'
        }

        return true;
    }
};

export const nickValidation = {
    required: REQUIRED_FIELD,
    validate: (value: string) => {
        if(value.match(/а-яА-Яa-zA-Z0-9_-/)) {
            return 'Ник может содержать только _ - цифры, русские и латинские символы'
        }

        return true;
    }
};