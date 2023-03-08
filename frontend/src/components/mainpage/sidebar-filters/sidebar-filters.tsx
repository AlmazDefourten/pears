import classes from './sidebar-filters.module.css';

const Filters: React.FC = () =>{
    return(
        <div className={classes['sidebar']}>
            {/* <div>This is Filters</div> */}
            {/* <label>
                <input type="checkbox" name="happy" value="yes"/>Happy
            </label> */}
            <label className={classes['main']}>CodeX
                <input type="checkbox"/>
                <span className={classes['geekmark']}></span>
            </label>
            <button>Применить</button>
            
        </div>
    )
}

export default Filters;
