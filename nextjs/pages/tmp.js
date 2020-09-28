import React, {Component} from 'react'
import { connect } from 'react-redux';
import Head from 'next/head'
import Nav from '../components/nav'
import {SignupAction} from "../redux/auth"

class Home extends Component {
    componentDidMount(){
        this.props.signup({
            phone: 123,
            pass: 123
        })
    }

    render() {
        return (
            <div>
                <Head>
                    <title>Home</title>
                    <link rel='icon' href='/favicon.ico' />
                </Head>
                <Nav />
                <div>{this.props.isLoading ? "" : "Singup OK"}</div>
            </div>
        )
    }
}

const mapStateToProps = (state) => ({
    isLoading: state.auth.isLoading,
    signupSuccess: state.auth.success,
})

const mapDispatchToProps = dispatch => ({
    signup: (user)=>dispatch(SignupAction.signupRequest(user)),
})

export default connect(mapStateToProps, mapDispatchToProps)(Home);