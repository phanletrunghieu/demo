import { applyMiddleware, createStore } from "redux";
import { createLogger } from "redux-logger";
import createSagaMiddleware from 'redux-saga'
import { persistStore, persistReducer } from 'redux-persist'
import autoMergeLevel2 from 'redux-persist/lib/stateReconciler/autoMergeLevel2'
import storage from 'redux-persist/lib/storage'
import rootReducer from "./rootReducer";
import rootSaga from "../saga";

// logger middleware
let logger = createLogger({
    timestamps: true,
    duration: true
});


// saga middleware
const sagaMiddleware = createSagaMiddleware()

// persist middleware
const persistConfig = {
    key: 'root',
    storage,
    stateReconciler: autoMergeLevel2,
}
const persistedReducer = persistReducer(persistConfig, rootReducer)

const store = createStore(
    persistedReducer,
    applyMiddleware(logger, sagaMiddleware)
);

sagaMiddleware.run(rootSaga)

let persistor = persistStore(store)

export {
    store,
    persistor,
};