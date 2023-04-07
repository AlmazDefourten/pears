import { Container, Checkbox, Button, FormGroup, FormControlLabel } from '@mui/material';

const formControlLabelStyle = {
    "& .MuiFormControlLabel-label": {
      fontSize: '20px',
      fontWeight: 300,
      lineHeight: '23px',
      letterSpacing: '0.02em'
    }
}

const Filters: React.FC = () =>{
    return(
        <Container sx={{ position: 'sticky',
                   top: 120,
                   width: '13vw !important',
                   height: '26vh',
                   backgroundColor: 'white',
                   borderRadius: '25px',
                   filter: 'drop-shadow(0px 3.80416px 3.80416px rgba(0, 0, 0, 0.25))',
                   paddingTop: '20px',
                   padding: '20px 0px 0px 28px',
                   fontFamily: "'Ubuntu', sans-serif"}}>
            <FormGroup sx={{...formControlLabelStyle}}>
                <FormControlLabel control={<Checkbox  defaultChecked/>} label="Мероприятия"/>
                <FormControlLabel control={<Checkbox  defaultChecked/>} label="Проекты"/>
                <FormControlLabel control={<Checkbox  defaultChecked/>} label="Друзья"/>
                <FormControlLabel control={<Checkbox  defaultChecked/>} label="Сообщества"/>
                <Button sx={{
                    marginTop: 1.5,
                    backgroundColor: '#006C0B',
                    color: 'white',
                    boxShadow: '0px 1.90208px 0.951039px rgba(0, 0, 0, 0.12)',
                    borderRadius: '2.85312px',
                    fontWeight: 500,
                    fontSize: '20px',
                    lineHeight: '21px',
                    textTransform: 'none',
                    '&:hover': {
                        color: 'black'
                    }
                }}>Применить</Button>
            </FormGroup>        
        </Container>
    )
}

export default Filters;
