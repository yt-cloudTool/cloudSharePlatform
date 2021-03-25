/*
-----------------
用于显示各种详情 比如
    -- 帐号详情
    -- 服务器数据
    -- 临时文件上传
    -- 小工具快捷方式


 */
import React, { Component } from 'react'
import Icon from "./components/sideToolBarIcon/sideToolBarIcon.js"
import "./sideToolBar.css"
import IMG_newarticle from "../../../assets/newarticle.svg"
import IMG_sharenote from "../../../assets/sharenote.svg"
import IMG_articlehis from "../../../assets/articlehis.svg"
import IMG_todobook from "../../../assets/todobook.svg"

class SideToolBar extends Component {
    constructor (props) {
        super(props)

        this.state = {
            sideToolBar_toggle: false
        }

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            this.state.sideToolBar_toggle = _state_.sideToolBar_toggle
        })

    }
    
    render () {
    		let containerStyle = {
			left: this.state.sideToolBar_toggle ? 0 : '-280px'
		}
        return (
            <div className="sideToolBar_container" style={containerStyle}>
            		{ /* 功能区域 */ }
				<div className="sideToolBar_functionArea">
	            	<Icon label="New Article" img={IMG_newarticle}/>
				<Icon label="Article His" img={IMG_articlehis}/>
	            	<Icon label="Share Note" img={IMG_sharenote}/>
	            	<Icon label="Todo Book" img={IMG_todobook}/>
				</div>
				{ /* 通知区域 */ }
				<div className="sideToolBar_notifyArea">
				
				</div>
            </div>
        )
    }
}

export default require('react-redux').connect((state) => { return state }, null)(SideToolBar)
