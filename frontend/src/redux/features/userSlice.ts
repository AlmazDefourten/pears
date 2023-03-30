import { createSlice } from '@reduxjs/toolkit';
import { IUserReducerState, IUserReducerAction } from '../Interfaces';


const initState: IUserReducerState = {
    isAuth: false
}
//kinda reducer itself
export const userSlice = createSlice({
    name: 'user',
    initialState: initState,
    reducers:{
        change_auth: (state: IUserReducerState, action: IUserReducerAction) =>{
            state.isAuth = action.payload
        }
    }
})
//actions (kinda pulled out from userSlice)
export const {change_auth} = userSlice.actions
