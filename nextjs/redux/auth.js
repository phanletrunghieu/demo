import { createActions, createReducer } from 'reduxsauce'

const INITIAL_STATE ={
    isLoading: false,
    success: false,
    error: null
}

// action
const { Types, Creators } = createActions({
    signupRequest: ['user'],
    signupSuccess: null,
    signupFailure: ['error'],
})

export const SignupAction = Creators
export const SignupTypes = Types

// reducer
export const request = (state = INITIAL_STATE, action) => {
    return {...state, isLoading: true}
}

export const success = (state = INITIAL_STATE, action) => {
    return {
        ...state,
        isLoading: false,
        success: true,
    }
}

export const failure = (state = INITIAL_STATE, action) => {
    return {
        ...state,
        isLoading: false,
        error: action.error,
    }
}

export default createReducer(INITIAL_STATE, {
    [Types.SIGNUP_REQUEST]: request,
    [Types.SIGNUP_SUCCESS]: success,
    [Types.SIGNUP_FAILURE]: failure,
})