import React, { Component } from 'react'
import Desktop from "./desktop/desktop.js"
import "./content.css"

class Content extends Component {
    render () {
        return (
            <div className="content_container">
            	<Desktop/>	
            </div>
        )
    }
}

export default Content