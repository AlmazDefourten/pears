import styles from './EventPost.module.css'
import eventimg from './imgs/eventimg.png'
import option from './imgs/option.png'
import meetup from './imgs/meetup.png'
import place from './imgs/place.png'
import date from './imgs/date.png'
import clock from './imgs/clock.png'
import like from './imgs/like.png'
import share from './imgs/share.png'
import glaz from './imgs/glaz.png'

const EventPost: React.FC = () =>{
    return(
        <div className={styles['event-content']}>
            <header className={styles['header-content']}>
                <div className={styles['event-author-and-name']}>
                    <img src={eventimg} alt="" className={styles['event-author-img']}/>
                    <span>RSHB Identity Management Meetup</span>
                </div>
                <div className={styles['event-time-and-options']}>
                    <span>Вчера в 19:29</span>
                    <img src={option} alt="" className={styles['event-options-img']} />
                </div>
            </header>
            <img src={meetup} alt="" className={styles['event-img']}/>
            {false ? <div className={styles['information']}>
                <div className={styles['icon-plus-text']}>
                    <img src={place} alt="" className={styles['information-icon']}/>
                    <div>
                        <p>Город</p>
                        <p>Казань</p>
                    </div>
                </div>
                <div className={styles['icon-plus-text']}>
                    <img src={date} alt="" className={styles['information-icon']}/>
                    <div>
                        <p>Дата</p>
                        <p>16 фев 2023</p>
                    </div>
                </div>
                <div className={styles['icon-plus-text']}>
                    <img src={clock} alt="" className={styles['information-icon']}/>
                    <div>
                        <p>Время</p>
                        <p>19:00-22:00</p>
                    </div>
                </div>
            </div> : ''}
            <div className={styles['text-description']}>
                <p>
                    Поговорим о работе с Keycloak и построим единую систему аутентификации 16 февраля на митапе RSHB Identity Management Meetup. Начало в онлайне в 19:00. Участие бесплатное. Нужно только зарегистрироваться.
                </p>
            </div>
            <footer className={styles['social-button-and-tags']}>
                <div className={styles['social-buttons']}>
                    <div className={styles['icon-plus-text']}>
                        <img src={like} alt="" className={styles['social-button-icon']}/>
                        <p>123</p>
                    </div>
                    <div className={styles['icon-plus-text']}>
                        <img src={share} alt="" className={styles['social-button-icon']}/>
                        <p>342</p>
                    </div>
                    <div className={styles['icon-plus-text']}>
                        <img src={glaz} alt="" className={styles['social-button-icon']}/>
                        <p>1</p>
                    </div>
                </div>
                <div className={styles['tags-container']}>
                    <div className={styles['tags']}>
                        <p>Design, Figma, WebFlow</p>
                    </div>
                </div>
            </footer>
        </div>
    )
}

export default EventPost;
