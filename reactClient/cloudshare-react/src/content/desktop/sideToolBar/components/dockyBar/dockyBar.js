import React, { Component } from 'react'
import "./dockyBar.css"

class DockyBar extends Component {
    render () {
        return (
            <div className="dockyBar_container">
            </div>
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(DockyBar)