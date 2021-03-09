import React, { Component } from 'react'
import "./dockyBar.css"

class DockyBar extends Component {
    constructor (props) {
        super(props)

        this.state = {
            
        }

        window.$store.subscribe(() => {

        })

    }

    render () {
        return (
            <div className="dockyBar_container">
            </div>
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(DockyBar)