import { get } from 'lodash';
import { SEARCH_PLAYER_START, SEARCH_PLAYER_ERROR, SEARCH_PLAYER_COMPLETE } from  '../../consts/actionTypes';


const initialState = {};

export default function(state = initialState, action){
    switch (action.type) {
        case SEARCH_PLAYER_START:
            return {...state, isloading: true};
            break;
        case SEARCH_PLAYER_ERROR:
            return {...state, isloading: false, players: null};
            break;
        case SEARCH_PLAYER_COMPLETE:
            return {...state, isloading: false, playerResults: action.results.data};
            break;
        default:
            return {...state};
    }
}

