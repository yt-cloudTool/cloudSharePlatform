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

        this.handleInputChange = this.handleInputChange.bind(this)

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
    handleInputChange (e) {
        this.setState({ [e.target.name]: e.target.value })
    }
    confirm () {
        
    }

    render () {
        if (this.state.loginPop_show === true) {
            return <div className="loginPop_container">
            	<div className="loginPop_inner">
				
					<div className="loginPop_header">
						<span>Login</span>
						<div className="loginPop_closeBtn" onClick={this.hidePop}>
							<div className="loginPop_closeBtn_inner"></div>
						</div>
					</div>
					
					<div className="headerPop_body">
						<div className="headerPop_usernameInput_container">
							<input className="headerPop_username_input" name="username" value={this.state.username} onChange={this.handleInputChange} placeholder="Username"/>
						</div>
						<div className="headerPop_passwordInput_container">
							<input className="headerPop_password_input" name="password" value={this.state.password} onChange={this.handleInputChange} placeholder="Password"/>
						</div>
						<div className="headerPop_btnArea">
							<span class="headerPop_actionBtn" onClick={this.confirm}>Confirm</span>
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