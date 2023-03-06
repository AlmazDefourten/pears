import { Routes, Route, Navigate } from "react-router-dom";
import Swipes from '../swipes/swipes';
import Mainpage from "../mainpage/mainpage";
import Profile from "../profile/profile";
import { AuthPage } from "../auth-page";

interface IFlag{
    flag: boolean
}

const Router: React.FC<IFlag> = ({flag}) =>{

    return(
        <div>
            {flag ? 
                <Routes>
                    <Route path="/" element={ <Mainpage /> }/>
                    <Route path="/swipes" element={ <Swipes /> }/>
                    <Route path="/profile" element={ <Profile /> }/>
                    <Route path="*" element={ <Navigate to='/' replace /> }/>
                </Routes> :
                <Routes>
                    <Route path="/AuthPage" element={ <AuthPage /> }/>
                    <Route path="*" element={ <Navigate to='/AuthPage' replace /> }/>
                </Routes>
            }
        </div>
    )
}

export default Router;
