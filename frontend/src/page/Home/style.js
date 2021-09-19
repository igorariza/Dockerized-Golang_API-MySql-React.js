import { makeStyles } from "@material-ui/core";


const centerdStyleObj = {
    display: 'flex',
    alignItems: 'center',
}

export default makeStyles({
    container:{
        height:'100vh',
        flexDirection: 'column',
        ...centerdStyleObj        
    },
    cardContainer:{
        flexDirection:'column',
        borderRadius: '20px',
        maxWidth: '95%',
        maxHeight: '95%',
        padding: '2rem',
        ...centerdStyleObj
    },
    title: {
        fontSize: '3rem'
    },
    titleGridContainer:{
        ...centerdStyleObj
    },
    textFieldSearch:{
        width:'100%'
    },
    buttonsContainer:{
        marginTop: '.5rem'
    },
    playericon:{
        fontSize: '4rem'
    }


});