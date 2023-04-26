import { Avatar, Container, Card, CardContent, CardHeader, CardMedia, Typography, CardActions } from "@mui/material";
import IconButton from '@mui/material/IconButton';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import ShareIcon from '@mui/icons-material/Share';
import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import VisibilityIcon from '@mui/icons-material/Visibility';
import { Box } from "@mui/system";
import avatar from './imgs/avatar1.png';
import media from './imgs/image_4.png';

const cardHeaderStyle={
    "& .MuiCardHeader-title": {
        fontFamily: 'IBM Plex Sans',
        fontWeight: 400,
        fontSize: '27.6364px',
        lineHeight: '32px',
        letterSpacing: '0.02em',
        color: 'white'
    },

    "& .MuiCardHeader-subheader": {
        fontFamily: 'IBM Plex Sans',
        fontWeight: 400,
        fontSize: '20.7273px',
        lineHeight: '24px',
        letterSpacing: '0.1em',
        color: 'white'
    }
}

const ProjectPost: React.FC = () =>{
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
                        title="Project “Pears”"
                        sx={{
                            ...cardHeaderStyle,
                            backgroundColor: '#006C0B'
                        }}
            />
            <CardMedia component="img"
                       height="337"
                       image={media}
                       alt="ProjectPostImg" />
            <CardContent>
                <Typography variant="body2" 
                            sx={{fontFamily: 'IBM Plex Sans',
                                 fontStyle: 'normal',
                                 fontWeight: 300,
                                 fontSize: '20px',
                                 lineHeight: '26px',
                                 letterSpacing: '0.02em'}}
                            color="text.secondary">
                🍐Pears - это платформа для совместной проектной деятельности, 
                поиска команды, мероприятий и нетворкинга - все что нужно для плодотворной командной работы, 
                новых знакомств и получения мотивации
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
                    <Typography sx={{fontFamily: 'IBM Plex Sans',
                                     fontStyle: 'normal',
                                     fontWeight: 300,
                                     fontSize: '20.7273px',
                                     lineHeight: '27px',
                                     backgroundColor: '#7EABB9',
                                     padding: 0.7,
                                     borderRadius: '15px'}}>Startup, Go, React</Typography>
                </Box>
            </Container>
        </Card>
    )
}

export default ProjectPost;