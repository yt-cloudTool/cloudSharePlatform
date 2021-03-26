import React, { Component } from 'react'
import "./topStatusBar.css"
import utils  from "../../../utils/utils.js"

class TopStatusBa extends Component {
    constructor (props) {
        super(props)

        this.state = {
            loginStatus:    		"NOLOGIN",
            loginPop_show:  		false,
            dateTime:       		utils.GetDate(),
            sideToolBar_toggle: 	false,
            serverInfoText:			""
        }

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            this.setState({
            	"loginStatus": 		  _state_.loginStatus,
	            "loginPop_show": 	  _state_.loginPop_show,
	            "sideToolBar_toggle": _state_.sideToolBar_toggle
            })
            
        })
        
        // 循环时间
        setInterval(() => {
            this.setState({ "dateTime": utils.GetDate() })
        }, 1000);
    }

    showLoginPop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(true))
    }
    
    toggleMenuDrawer () {
    	window.$store.dispatch(window.$actions.set_sideToolBar_toggle(!window.$store.getState().sideToolBar_toggle))
    }
    
    getServerInfo () {
		window.$axios.get(window.$api.serverinfo).then((res) => {
        	if (res.data.status === 1) {
				this.setState({
					serverInfoText: res.data.data.numCPU + "cpu"
				})
			}
        }).catch((err) => {
            
        })
    }

    render () {
    		let menuBtnStyle = {
				
		}
        return (
            <div className="topStatusBar_container">

                <div className="topStatusBar_inner">

                    <div className="topStatusBar_inner_left">
                    	{/* 菜单按钮 */}
                     	<span className="topStatusBar_item topStatusBar_menu" onClick={this.toggleMenuDrawer} style={menuBtnStyle}>
							Menu
						</span>
                        {/* 登录状态指示&登录选项功能 */}
                        <span className="topStatusBar_item topStatusBar_loginStatus">
                            
                            <span className="topStatusBar_loginActionArea">
                                { this.state.loginStatus === "NOLOGIN" ? <span onClick={this.showLoginPop}>Login</span> : <span>Account</span> }
                            </span>
                        </span>
                    </div>

                    <div className="topStatusBar_inner_right">
                        {/* 搜索 */}
                        {/* <span className="topStatusBar_item topStatusBar_search">
                            search
                        </span> */}
                        
                        {/* 服务器信息 */}
                        <span className="topStatusBar_item topStatusBar_serverInfo">
                            {this.state.serverInfoText}
                        </span>

                        {/* main区域显示效果调整&关闭所有窗口按钮&手机pc版本页面切换 */}
                        <span className="topStatusBar_item topStatusBar_action">
                            action
                        </span>

                        {/* 数据数量统计 */}
                        <span className="topStatusBar_item topStatusBar_dataSum">
                            dataSum
                        </span>

                        {/* 时间日期 */}
                        <span className="topStatusBar_item topStatusBar_timeDate">
                            {this.state.dateTime}
                        </span>
                    </div>
                </div>


            </div>
        )
    }
    
    componentDidMount () {
    	this.getServerInfo()
    }
}

export default require('react-redux').connect((state) => { return state }, null)(TopStatusBa)