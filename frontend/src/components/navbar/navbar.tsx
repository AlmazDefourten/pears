import { NavLink } from "react-router-dom";
import classes from './navbar.module.css';
import userpic from './userpic.png';
import logo from './logo.png';

const Navbar: React.FC = () =>{
    return(
        <header>
            <div className={classes.content}>
                <NavLink to="/" className={navData => navData.isActive ? `${classes['active-link']} ${classes['logo']}` : `${classes['inactive-link']} ${classes['logo']}`}>
                    <img src={logo} alt="" className={classes['image-logo']}/>
                    <div>
                        PEARS
                    </div>
                    </NavLink>
                {/* <NavLink to="/swipes" className={navData => navData.isActive ? classes['active-link'] : classes['inactive-link']}>Swipes</NavLink> */}
                <NavLink to="/profile" className={navData => navData.isActive ? classes['active-link'] : classes['inactive-link']}>
                    <img src={userpic} alt="Profile" className={classes['image-profile']}/>
                </NavLink>
            </div>
        </header>    
    )
}

export default Navbar;