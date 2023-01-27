import { motion, PanInfo, useAnimationControls } from 'framer-motion';
import css from "./swipes.module.css";
interface IProps{
    name: string, 
    setCharacters: React.Dispatch<React.SetStateAction<{name: string;}[]>>,
}

const Card = ({name, setCharacters}: IProps) =>{

    //animations hook
    //хук анимаций
    const controls = useAnimationControls();

    return(
        <motion.div
        className={css.swipe}
        // animate - values to animate to
        // animate - значения, к которым привести
        animate={controls}
        //drag only works in the x-axis
        //drag идет только по оси x
        drag='x'
        //nothing works without event, so do not remove
        //без event ничего не работает, не убирать
        //offset contains x and y values for the distance moved from the first pan event
        //offset - изменения в пикселях от позиции до начала drag
        onDragEnd={async (event: MouseEvent | TouchEvent | PointerEvent, info: PanInfo) => {
            const x:number = info.offset.x;
            switch(true){
                case(x > 200):
                    //controls.start() return a promise
                    //controls.start() возвращает promise, см документацию
                    await controls.start({x: window.screen.width});
                    setCharacters(prevState => prevState.filter(char => char.name !== name));
                    console.log("Swiped right");
                    break;
                case(x < -200):
                    await controls.start({x: -window.screen.width});
                    setCharacters(prevState => prevState.filter(char => char.name !== name));
                    console.log("Swiped left");
                    break;
            }
            // if (info.offset.x > 200) {
            // await controls.start({x: window.screen.width})
            // setCharacters(prevState => prevState.filter(char => char.name !== name))
            // console.log("Swiped right")
            // }
            // if (info.offset.x < -200) {
            // await controls.start({x: -window.screen.width})
            // setCharacters(prevState => prevState.filter(char => char.name !== name))
            // console.log("Swiped left")
            // }
        }}
        //does what it says
        //возвращает на место при отпускании
        dragSnapToOrigin
        //does what it says
        //ограничения в пикселях на drag от оригинального положения(самого первого, типа init), не позволит выйти за эти значения
        dragConstraints={{ left: -300, right: 300 }}
        //does not allows an object to slide after dropping\releasing it
        //не позволяет карточке скользить после отпускания
        dragMomentum={true}>
            <div className={css.card}>{name}</div>
        </motion.div>
    )
} 

export default Card;