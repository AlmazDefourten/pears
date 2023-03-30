export interface IUserReducerState{
    isAuth: boolean
}

export interface IUserReducerAction{
    type: string,
    payload: boolean
}