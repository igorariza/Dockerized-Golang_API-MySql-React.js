import React, { useState } from 'react';
import {Button, Card, Container, Grid, TextField, Typography} from '@material-ui/core';
import { PlayerIcon } from '../../icons';

import styles from './style';
const Home = ({ history }) => {

  const [searchText, setSearchText] = useState('');
  const classes = styles(

  );
  const handleSearchTextChange = event => {
    setSearchText(event.target.value)
  };

  const handleCleanTextClick = event => {
    setSearchText('');
  };

  const handleSearchTextClick  = event => {
    history.push(`/results?playerName=${searchText}`) 
  };


  return (
      <Container className={classes.container}>
        <Card className={classes.cardContainer}>
          <Grid container className={classes.titleGridContainer}>
            <Grid>
              <Typography className={classes.title}>Bienvenido!</Typography>
            </Grid>
            <Grid>
            <PlayerIcon className={classes.playericon} />
            </Grid>           
          </Grid>
          <TextField
            value={searchText}
            placeholder="Buscar..."
            className={classes.textFielSearch}
            onChange={handleSearchTextChange} />
          <Grid className={classes.buttonsContainer}>
            <Button variant="contained" onClick={handleCleanTextClick}>Limpiar</Button>
            <Button variant="contained" className={classes.searchButton} color="primary" size="large" onClick={handleSearchTextClick}>Buscar</Button>  
          </Grid>  
        </Card>
        
          
      </Container>
    )
  };

export default Home;
