/*
-----------------
用于显示各种详情 比如
    -- 帐号详情
    -- 服务器数据
    -- 临时文件上传
    -- 小工具快捷方式


 */
import React, { Component } from 'react'
import "./sideToolBar.css"

class SideToolBar extends Component {
    constructor (props) {
        super(props)

        this.state = {
            
        }

        window.$store.subscribe(() => {

        })

    }
    
    render () {
        return (
            <div className="sideToolBar_container">
            </div>
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(SideToolBar)
