import * as React from 'react';
import EventPost from './posts/EventPost/EventPost';
import ProjectPost from './posts/ProjectPost/ProjectPost';
import Filters from './sidebar-filters/sidebar-filters';
import Navmenu from './navmenu/navmenu';
import { Container } from '@mui/system';
import { Stack, Box} from '@mui/material';
import FriendPost from './posts/FriendPost/FriendPost';

const Mainpage: React.FC = () =>{
    return(
        <Box sx={{background: "#F4F4F4"}}>
            <Container maxWidth='xl' sx={{
                display: "flex",
                justifyContent: "center",
                padding: "35px 0vw 7vh 0vw",
                pt: 15,
                gridGap: "10px"
            }}>
                <Navmenu />
                <Stack spacing={5} sx={{width: '40vw'}}>
                    <EventPost/>
                    <ProjectPost/>
                    <FriendPost/>
                </Stack>
                <Filters/>
            </Container>
        </Box>
    )
}

export default Mainpage;
