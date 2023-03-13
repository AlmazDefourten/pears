import { NavLink } from "react-router-dom";
import classes from './navbar.module.css';
import userpic from './userpic.png';
import logo from './logo.png';

const Navbar: React.FC = () =>{
    return(
        <header>
            <div className={classes['navbar-content']}>
                <div className={classes['logo-container']}>
                    <NavLink to="/" className={navData => navData.isActive ? `${classes['active-link']} ${classes['logo']}` : `${classes['inactive-link']} ${classes['logo']}`}>
                        <img src={logo} alt="" className={classes['image-logo']}/>
                        <div>
                            PEARS
                        </div>
                    </NavLink>
                </div>
                
                {/* <NavLink to="/swipes" className={navData => navData.isActive ? classes['active-link'] : classes['inactive-link']}>Swipes</NavLink> */}
                <div className={classes['image-profile-container']}>
                    <NavLink to="/profile" className={navData => navData.isActive ? `${classes['active-link']} ${classes['profile-navigation']}` : `${classes['inactive-link']} ${classes['profile-navigation']}`}>
                        <img src={userpic} alt="Profile" className={classes['image-profile']}/>
                    </NavLink>
                </div>
                
            </div>
        </header>    
    )
}

export default Navbar;