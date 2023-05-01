import { Box, Link, Typography, IconButton } from '@mui/material';
import AccountCircleOutlinedIcon from '@mui/icons-material/AccountCircleOutlined';
// import NewspaperIcon from '@mui/icons-material/Newspaper';
import NewspaperIcon from '@mui/icons-material/Newspaper';
import MessageIcon from '@mui/icons-material/Message';

const FontStyles = {
    fontFamily: 'Ubuntu',
    fontWeight: 400,
    ml: 1
}

const Navmenu: React.FC = () =>{
    return(
        <>
            <Box sx={{ width: '14vw', height: '10vh', position: 'sticky', top: 120}}>
                    <Link href='#' underline='none'>
                        <IconButton sx={{color: 'black', padding: 1,
                                        '&:hover': {
                                            color: '#006C0B',
                                            borderRadius: 5,
                                        }}}>
                            <AccountCircleOutlinedIcon sx={{fontSize: 35}}/>
                            <Typography variant='h6' sx={{...FontStyles}}>Моя страница</Typography>
                        </IconButton>
                    </Link>
                    <Link href='#' underline='none'>
                        <IconButton sx={{color: 'black', padding: 1,
                                        '&:hover': {
                                            color: '#006C0B',
                                            borderRadius: 5,
                                        }}}>
                            <NewspaperIcon sx={{fontSize: 35}}/>
                            <Typography variant='h6' sx={{...FontStyles}}>Новости</Typography>
                        </IconButton>
                    </Link>
                    <Link href='#' underline='none'>
                        <IconButton sx={{color: 'black', padding: 1,
                                        '&:hover': {
                                            color: '#006C0B',
                                            borderRadius: 5,
                                        }}}>
                            <MessageIcon sx={{fontSize: 35}}/>
                            <Typography variant='h6' sx={{...FontStyles}}>Сообщения</Typography>
                        </IconButton>
                    </Link>
            </Box>
        </>
        )
}

export default Navmenu;