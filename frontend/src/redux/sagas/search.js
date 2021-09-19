import { put, call, takeLatest } from 'redux-saga/effects'
import { SEARCH_PLAYER_START, SEARCH_PLAYER_ERROR, SEARCH_PLAYER_COMPLETE } from  '../../consts/actionTypes'

import { apiCall } from '../api'
export function* searchPlayer({ payload }){
    try {
     const results = yield call(apiCall, null, null, 'GET')
     yield put({ type: SEARCH_PLAYER_COMPLETE, results })   
    } catch (error) {
        yield put({ type: SEARCH_PLAYER_ERROR, error });
    }
}

export default function* search(){
    yield takeLatest(SEARCH_PLAYER_START, searchPlayer);
}
