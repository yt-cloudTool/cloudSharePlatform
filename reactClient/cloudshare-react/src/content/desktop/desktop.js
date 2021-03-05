import React, { Component } from 'react'
import TopStatusBar 	from "./topStatusBar/topStatusBar.js"
import SideToolBar 	 	from "./sideToolBar/sideToolBar.js"
import LoginPop			from "./loginPop/loginPop.js"
import MainArea			from "./mainArea/mainArea.js"
import UploadPop        from "./uploadPop/uploadPop.js"
import "./desktop.css"

class Desktop extends Component {
    render () {
        return (
            <div className="desktop_container">
            	<TopStatusBar/>
				<SideToolBar/>
				<LoginPop/>
				<MainArea/>
                <UploadPop/>
            </div>
        )
    }
}

export default Desktop