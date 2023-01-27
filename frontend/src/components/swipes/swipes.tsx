import { useState } from 'react';
import Card from './card';
import css from "./swipes.module.css";

const Swipes =() =>{

    //we store users here
    //don`t have redux yet, so useState hook only
    const [characters, setCharacters] = useState([
        {
            name: 'Richard Hendricks',
        },
        {
            name: 'Erlich Bachman',
        },
        {
            name: 'Monica Hall',
        },
        {
            name: 'Jared Dunn',
        },
        {
            name: 'Dinesh Chugtai',
        },
    ]);

    return(
        <div className={css.cardContainer}>
            <h1>Cards</h1>
            {characters.map((char) =>
            // name - just some random data for card. This is just an example, so only name
            // name - какие-то данные для карточки. Т.к. пример это только, то только имя. Так-то стоит передавать весь объект
            // setCharacters - передаем функцию для изменения массива пользователей
            // setCharacters - pass a function that changes users state
            <Card key={char.name} name={char.name} setCharacters={setCharacters}/>
            )}
        </div>
    )
}

export default Swipes;