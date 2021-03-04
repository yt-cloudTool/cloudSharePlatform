import React, { Component } from 'react'
import { connect } from "react-redux"
import "./sideToolBar.css"

class SideToolBar extends Component {
    render () {
        return (
            <div className="sideToolBar_container">
            	sideToolBar	
            </div>
        )
    }
}

const mapStateToProps = (state) => {
	return {
		'sideToolBar_show': state.sideToolBar_show
	}
}

export default connect()(SideToolBar)