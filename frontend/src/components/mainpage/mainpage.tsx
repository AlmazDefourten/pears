import classes from './mainpage.module.css';
import EventPost from './posts/post';
import Filters from './sidebar-filters/sidebar-filters';
import meetup from './meetup.png'

const Mainpage: React.FC = () =>{
    return(
        // <div className={classes['main']}>This is the Mainpage</div>
        <div className={classes['main']}>
            <EventPost/>
            {/* <div className={classes['dgffgd']}>
                <img src={meetup} alt="" className={classes['asdasd']}/>
            </div> */}
            {/* <Filters/> */}
        </div >
    )
}

export default Mainpage;
