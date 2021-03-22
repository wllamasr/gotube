import {UPLOAD_FILE} from './actionTypes'

const initialState = {}

export default function (state = initialState, action: any) {
    switch (action.type) {
        case UPLOAD_FILE:
            return {...state}
        default:
            return state
    }
}