import { Avatar, Container, Card, CardContent, CardHeader, Typography, CardActions } from "@mui/material";
import IconButton from '@mui/material/IconButton';
import MoreVertIcon from "@mui/icons-material/MoreVert";
import ShareIcon from "@mui/icons-material/Share";
import ThumbUpIcon from "@mui/icons-material/ThumbUp";
import VisibilityIcon from '@mui/icons-material/Visibility';

import { Box } from "@mui/system";
import avatar from './imgs/ava.png';

import Slider from "./Slider";



const cardHeaderStyle={
    "& .MuiCardHeader-title": {
        fontWeight: 400,
        fontSize: '25.6364px',
        lineHeight: '32px',
        letterSpacing: '0.02em',
        color: 'white'
    },

    "& .MuiCardHeader-subheader": {
        fontWeight: 400,
        fontSize: '15.7273px',
        lineHeight: '24px',
        letterSpacing: '0.1em',
        color: 'white'
    },
}

const FriendPost: React.FC = () =>{
    return(
        <Card sx={{
            height: 'auto',
            backgroundColor: 'white',
            borderRadius: '15px',
            boxShadow: '0px 3.4591px 3.4591px rgba(0, 0, 0, 0.25)'
        }}>
            <CardHeader avatar={
                            <Avatar alt="ProjectPostAva" src={avatar}/>}
                        action={
                            <IconButton aria-label="settings" size="large" sx={{
                                color: 'white'
                            }}>
                              <MoreVertIcon />
                            </IconButton>
                        }
                        title="Особенности npm и хранение node_m..."
                        subheader="Lowellda"
                        sx={{
                            ...cardHeaderStyle,
                            backgroundColor: '#7EABB9'
                        }}
            />
            {/* <CardMedia component="img"
                       height="337"
                       image={media}
                       alt="ProjectPostImg" /> */}
            <Slider/>
            
            <CardContent>
                <Typography variant="body2" color="text.secondary">
                Для управления зависимостями в проекте, node.js, 
                как и многие другие платформы, предоставляет собственный 
                пакетный менеджер — npm. И несмотря на то, что он внешне похож, 
                например, на Ruby Gems, и вроде бы выполняет те же самые функции, 
                npm обладает некоторыми особенностями, которые стоит учитывать....
                </Typography>
            </CardContent>
            <Container sx={{
                width: '100%',
                display: 'flex',
                justifyContent: 'space-between',
                pb: 1
            }}>
                <CardActions disableSpacing>
                  <IconButton size="large" sx={{color: '#00c760'}} aria-label="add to favorites">
                    <ThumbUpIcon />
                  </IconButton>
                  <IconButton size="large" sx={{color: '#00c760'}} aria-label="share">
                  <ShareIcon />
                  </IconButton>
                  <IconButton size="large" sx={{color: '#00c760'}} aria-label="share">
                  <VisibilityIcon aria-label="view"/>
                  </IconButton>
                </CardActions>
                <Box sx={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center'
                }}>
                    <Typography sx={{backgroundColor: '#7EABB9',
                                     padding: 0.7,
                                     borderRadius: '15px'}}>Frontend, React</Typography>
                </Box>
            </Container>
        </Card>
    )
}

export default FriendPost;