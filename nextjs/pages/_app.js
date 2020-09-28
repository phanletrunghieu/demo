import React from 'react'
import App from "next/app";
import { Provider as ReduxProvider } from "react-redux";
import { PersistGate } from 'redux-persist/integration/react'
import withRedux from "next-redux-wrapper";
import { store, persistor } from '../redux/store'

class MyApp extends App {
    static async getInitialProps({Component, ctx}) {
        const pageProps = Component.getInitialProps ? await Component.getInitialProps(ctx) : {};
        return {pageProps};
    }
    render() {
        const {Component, pageProps, store} = this.props;
        return (
            <ReduxProvider store={store}>
                <PersistGate loading={null} persistor={persistor}>
                    <Component {...pageProps} />
                </PersistGate>
            </ReduxProvider>
        );
    }
}

const makeStore = (initialState, options) => {
    return store;
};

export default withRedux(makeStore, { debug: true })(MyApp);