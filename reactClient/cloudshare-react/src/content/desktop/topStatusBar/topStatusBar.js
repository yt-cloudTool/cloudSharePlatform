import React, { Component } from 'react'
import "./topStatusBar.css"

class TopStatusBa extends Component {
    render () {
        return (
            <div className="topStatusBar_container">

                {/* 登录状态指示&登录选项功能 */}
                <span className="topStatusBar_item topStatusBar_loginStatus">
                    status
                </span>

                {/* 搜索 */}
                <span className="topStatusBar_item topStatusBar_search">
                    search
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
                    timeDate
                </span>

            </div>
        )
    }
}

export default TopStatusBa