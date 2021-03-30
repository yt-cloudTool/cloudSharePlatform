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
            <div className="sideToolBarIcon_container" onClick={this.props.onClick}>
            	{
					this.props.img ? <img className="sideToolBarIcon_img" src={this.props.img}/> : ''
				}
				<p className="sideToolBarIcon_label">{this.props.label}</p>
            </div>
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(SideToolBarIcon)