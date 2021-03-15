import React, { Component } from 'react'
import "./sideToolBarIcon.css"

class SideToolBarIcon extends Component {
    constructor (props) {
        super(props)

        this.state = {
            
        }

        window.$store.subscribe(() => {

        })

    }

    render () {
        return (
            <div className="sideToolBarIcon_container">
            </div>
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(SideToolBarIcon)