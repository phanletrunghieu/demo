import { put, all, call, takeEvery } from 'redux-saga/effects'
import {SignupAction, SignupTypes} from '../redux/auth'
import { signup } from '../api/auth'

function* signupRequest({user}) {
    try {
        const result = yield call(signup, user)
        yield put(SignupAction.signupSuccess())
    } catch (error) {
        console.error(error);
        yield put(SignupAction.signupFailure(error))
    }
}

export default function* discoverUsersSaga() {
    yield all([
        takeEvery(SignupTypes.SIGNUP_REQUEST, signupRequest),
    ])
}