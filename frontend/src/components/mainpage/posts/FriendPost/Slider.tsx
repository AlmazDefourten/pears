import { Navigation, Pagination, Scrollbar, A11y } from 'swiper';
import {Swiper, SwiperSlide} from 'swiper/react';
import { CardMedia } from '@mui/material';

import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';
import 'swiper/css/scrollbar';

import slide_1 from "./imgs/carousels/item_1.png";
import slide_2 from "./imgs/carousels/item_2.png";
import slide_3 from "./imgs/carousels/item_3.png";
import slide_4 from "./imgs/carousels/item_4.png";

// const slides = {
//     {
//         image: item_1,
//         title: "React photo"
//     }
// }

const slides = [slide_1, slide_2, slide_3, slide_4];

const Slider: React.FC = () =>{
    return (
        <Swiper slidesPerView={1}
                modules={[Navigation, Pagination, Scrollbar, A11y]}
                navigation
                pagination={{ clickable: true,
                              type: "fraction" }}
                scrollbar={{ draggable: true }}
                onSlideChange={() => console.log('slide change')}
                onSwiper={(swiper) => console.log(swiper)}
        >
            {slides.map((slide) => (
                <SwiperSlide key={slide}>
                    <CardMedia component="img"
                       height="337"
                       image={slide}
                       alt="FriendImage" />
                </SwiperSlide>
            ))}
        </Swiper>
    )
}

export default Slider;