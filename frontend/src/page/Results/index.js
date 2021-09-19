import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Container } from "@material-ui/core";
import queryString from 'query-string';


import { searchPlayer } from '../../redux/actions/search'
export default ({ location }) => {
    const dispatch = useDispatch();

    const players = useSelector(state => {
        
    });
    
    useEffect(() => {
        const { playerName } = queryString.parse(location.search);
        dispatch(searchPlayer({playerName}))
    });
    
    
    return(
        <Container>
            Results
        </Container>
    )
}