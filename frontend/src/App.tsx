import './App.css';
import Router  from './components/router/router';
import Navbar from './components/navbar/navbar';
import { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { change_auth } from './redux/features/userSlice';
import { IUserReducerState } from './redux/Interfaces';
import axios from 'axios';
import {API_URL} from './components/http/index'

interface IState{
  user: IUserReducerState
}

function App() {

  const dispatch = useDispatch();

  const isAuth = useSelector((state: IState) => state.user.isAuth)

  useEffect(() =>{
    if (localStorage.getItem('token') && localStorage.getItem('token') !== "undefined"){
      axios({
        method:'get',
        url: `${API_URL}/user/refresh`,
        withCredentials: true
     })
     .then((result)=>{
        console.log('asdasd')
        const {data} = result;
        localStorage.setItem('token', data.access_token)
        dispatch(change_auth(true))
     })
    }
  }, [dispatch])

  return (
    <div>
      {isAuth ? <Navbar /> : ''}
      <Router isAuth={isAuth}/>
    </div>
  );
}

export default App;
