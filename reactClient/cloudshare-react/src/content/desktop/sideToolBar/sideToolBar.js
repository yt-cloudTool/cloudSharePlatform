import React, { Component } from 'react'
import { connect } from "react-redux"
import { set_sideToolBar_show } from '../../../store/actions.js'
import "./sideToolBar.css"

class SideToolBar extends Component {
    render () {
        return (
            <div className="sideToolBar_container">
            </div>
        )
    }
}

const mapStateToProps = (state) => {
	return {
		'sideToolBar_show': state.sideToolBar_show
	}
}

const mapDispatchToProps = (dispatch, ownProps) => {
    return {
        set_sideToolBar_show (data) {
            dispatch(set_sideToolBar_show(data))
        }
    }
}


export default connect(mapStateToProps, mapDispatchToProps)(SideToolBar)