/*
-----------------
用于显示各种详情 比如
    -- 帐号详情
    -- 服务器数据
    -- 临时文件上传
    -- 小工具快捷方式


 */
import React, { Component } from 'react'
import { connect } from "react-redux"
// import {  } from '../../../store/actions.js'
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
            dispatch(window.$actions.set_sideToolBar_show(data))
        }
    }
}


export default connect(mapStateToProps, mapDispatchToProps)(SideToolBar)