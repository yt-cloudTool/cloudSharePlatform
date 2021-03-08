import React, { Component } from 'react'
import "./loginPop.css"

class LoginPop extends Component {
    constructor (props) {
        super(props)

        this.state = {
            loginStatus:    	'NOLOGIN',
            loginPop_show:  	false,
            username:			'',
            password:			''
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
            	<div className="loginPop_inner">
				
					<div className="loginPop_header">
						<span>Login</span>
						<span className="loginPop_closeBtn" onClick={this.hidePop}></span>
					</div>
					
					<div className="headerPop_body">
						<div className="headerPop_usernameInput_container">
							<input className="headerPop_username_input" value={this.state.username} placeholder="Username"/>
						</div>
						<div className="headerPop_passwordInput_container">
							<input className="headerPop_password_input" value={this.state.password} placeholder="Password"/>
						</div>
						<div className="headerPop_btnArea">
							<span class="headerPop_actionBtn">v</span>
						</div>
					</div>
					
				</div>
            </div>
            
        } else {
            return null
        }
    }
}

export default require('react-redux').connect((state) => { return state }, null)(LoginPop)