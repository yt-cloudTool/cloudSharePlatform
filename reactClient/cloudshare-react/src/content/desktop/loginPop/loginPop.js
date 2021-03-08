import React, { Component } from 'react'
import "./loginPop.css"

class LoginPop extends Component {
    constructor (props) {
        super(props)

        this.state = {
            loginStatus:    'NOLOGIN',
            loginPop_show:  false
        }

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            
            this.state.loginPop_show = _state_.loginPop_show
            this.state.loginStatus   = _state_.loginStatus
        })

    }

    showPop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(true))
    }
    hidePop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(false))
    }

    render () {
        if (this.state.loginPop_show === true) {
            return <div className="loginPop_container">
                loginPop {String(this.state.loginStatus)}
            </div>
            
        } else {
            return null
        }
    }
}

export default require('react-redux').connect((state) => { return state }, null)(LoginPop)