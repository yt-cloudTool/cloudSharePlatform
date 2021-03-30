import React, { Component } from 'react'
import "./articleEditor.css"

class ArticleEditor extends Component {
    constructor (props) {
        super(props)

        this.state = {
            articleEditor_show: false,
        }

        this.showPop           = this.showPop.bind(this)
        this.hidePop           = this.hidePop.bind(this)
        this.handleInputChange = this.handleInputChange.bind(this)
        this.confirm           = this.confirm.bind(this)

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            this.setState({
            	"articleEditor_show": _state_.articleEditor_show
            })
        })

    }

    showPop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(true))
    }
    hidePop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(false))
        this.setState({ username: "", password: "" })
    }
    handleInputChange (e) {
        this.setState({ [e.target.name]: e.target.value })
    }
    confirm () {
        window.$axios.post(
        	window.$api.login, 
	        window.$qs.stringify({ // json转formData
				"loginname": this.state.username,
				"password":  this.state.password
			})
		).then((res) => {
        	if (res.data.status === 1) {
				// 存储 token 到 localStorage
				localStorage.setItem("token", 	  res.data.data.token)
				localStorage.setItem("loginname", res.data.data.loginname)
				localStorage.setItem("nickname",  res.data.data.nickname)
			}
        }).catch((err) => {
        	
        })
    }

    render () {
    	let containerStyle = {
			display: this.state.articleEditor_show ? "flex" : "none"
		}
        return <div className="articleEditor_container" style={containerStyle}>
          	<div className="articleEditor_inner">
			111
			</div>
        </div>
    }
}

export default require('react-redux').connect((state) => { return state }, null)(ArticleEditor)