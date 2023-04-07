import './App.css';
import Router  from './components/router/router';
import Navbar from './components/navbar/navbar';
import { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { change_auth } from './redux/features/userSlice';
import { IUserReducerState } from './redux/Interfaces';

interface IState{
  user: IUserReducerState
}

function App() {

  const dispatch = useDispatch();

  const isAuth = useSelector((state: IState) => state.user.isAuth)

  useEffect(() =>{
    dispatch(change_auth(true))
  }, [dispatch])

  return (
    <div>
      {isAuth ? <Navbar /> : ''}
      <Router isAuth={isAuth}/>
    </div>
  );
}

export default App;
