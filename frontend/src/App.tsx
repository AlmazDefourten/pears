import './App.css';
import Router  from './components/router/router';
import Navbar from './components/navbar/navbar';

function App() {

  const IsAccessTokenValid = () => {
    //let flag = localStorage.getItem('') 
    let flag: boolean = true;
    return flag;
  }

  const flag = IsAccessTokenValid();

  return (
    <div>
      {flag ? <Navbar /> : ''}
      <Router flag={flag}/>
    </div>
  );
}

export default App;
