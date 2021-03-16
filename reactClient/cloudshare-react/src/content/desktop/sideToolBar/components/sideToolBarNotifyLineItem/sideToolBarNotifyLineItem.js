import React, { Component } from 'react'
import "./sideToolBarNotifyLineItem.css"

class SideToolBarNotifyLineItem extends Component {
    constructor (props) {
        super(props)

        this.state = {
            
        }

        window.$store.subscribe(() => {

        })

    }

    render () {
        return (
           
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(SideToolBarNotifyLineItem)