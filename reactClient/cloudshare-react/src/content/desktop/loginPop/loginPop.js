import React, { Component } from 'react'
import { connect }          from 'react-redux'
import { set_sideToolBar_show } from '../../../store/actions.js'
import "./loginPop.css"

class LoginPop extends Component {
    constructor (props) {
        super(props)

        this.state = {
            barshow: false
        }

        window.$store.subscribe(() => {
            const state = window.$store.getState()
            this.state.barshow = state.sideToolBar_show
        })

        this.testChangeStore()
    }
    render () {
        if (this.state.barshow) {
            return (
                <div className="loginPop_container">
                    loginPop	{String(this.state.barshow)}
                </div>
            )
        } else {
            return ""
        }
    }

    testChangeStore () {
        // setInterval(() =>{
        //     window.$store.dispatch(set_sideToolBar_show(!window.$store.getState().sideToolBar_show))
        // },100)
    }
}

const mapStateToProps = (state) => {
    return {
        "sideToolBar_show": state.sideToolBar_show
    }
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        set_sideToolBar_show (data) {
            dispatch(set_sideToolBar_show(data))
        }
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(LoginPop)