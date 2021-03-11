import React, { Component } from 'react'
import "./window.css"

class Window extends Component {
    constructor (props) {
        super(props)

        this.state = {
            win_X: 60,
            win_Y: 80,
            win_W: 200,
            win_H: 200,
            // 存储开始拖动前窗口端点与鼠标正交距离
            origin_Xlen: 0,
            origin_Ylen: 0,
            // 原窗口offset
            origin_offsetLeft: 0,
            origin_offsetRight: 0,
            origin_offsetTop: 0,
            origin_offsetBottom: 0
        }

        this.windowContainerRef = null 	// 窗口容器ref
        this.windowCornerTLRef  = null 	// 左上角ref
        this.windowCornerTRRef  = null		// 右上角ref
        this.windowCornerBLRef  = null 	// 左下角ref
        this.windowCornerBRRef  = null 	// 右下角ref

		// container event
        this.handle_container_mouseDown = this.handle_container_mouseDown.bind(this)
        this.handle_container_mouseMove = this.handle_container_mouseMove.bind(this)
        this.handle_container_mouseUp   = this.handle_container_mouseUp.bind(this)
        // corner event
        this.handle_corner_mouseDown = this.handle_corner_mouseDown.bind(this)
        this.handle_corner_mouseMove = this.handle_corner_mouseMove.bind(this)
        this.handle_corner_mouseUp   = this.handle_corner_mouseUp.bind(this)
    }

	// =========================================================================
	//								实现窗口移动
	// =========================================================================
    handle_container_mouseDown (e) {
		// 如果是window_header则使用移动事件
		if (e.target.className && (e.target.className == "window_header" || (e.target.parentNode && e.target.parentNode.className == "window_header"))) {
			this.setState({
				origin_Xlen: e.clientX - this.windowContainerRef.offsetLeft,
				origin_Ylen: e.clientY - this.windowContainerRef.offsetTop,
			})
			window.onmousemove = this.handle_container_mouseMove
			window.onmouseup = this.handle_container_mouseUp
		}
    }
    handle_container_mouseMove (e) {
		this.setState({
			win_X: e.clientX - this.state.origin_Xlen,
			win_Y: e.clientY - this.state.origin_Ylen,
		})
        
    }
    handle_container_mouseUp (e) {
    	window.onmousemove = null
		window.onmouseup = null
    }
    
    // =========================================================================
	//								实现窗口调整
	// =========================================================================
	handle_corner_mouseDown (e) {
		if (e.target.parentNode && e.target.parentNode.getAttribute('name')) {
			switch (e.target.parentNode.getAttribute('name')) {
				case 'tl':
					this.setState({
						origin_offsetLeft: 0,
			           	origin_offsetRight: 0,
			           	origin_offsetTop: 0,
			           	origin_offsetBottom: 0
					})
				break
				case 'tr':
				break
				case 'bl':
				break
				case 'br':
				break
				default:
					return false
			}
			
			this.setState({
				origin_Xlen: e.clientX - this.windowContainerRef.offsetLeft,
				origin_Ylen: e.clientY - this.windowContainerRef.offsetTop,
			})
			
			window.onmousemove = this.handle_corner_mouseMove
			window.onmouseup = this.handle_corner_mouseUp
		}
	}
	handle_corner_mouseMove (e) {
		switch (e.target.parentNode.getAttribute('name')) {
			case 'tl':
				this.setState({
					win_W: this.state.win_W + this.state.origin_offsetRight,
					win_H: this.state.win_H + this.state.origin_offsetBottom
				})
			break
			case 'tr':
			break
			case 'bl':
			break
			case 'br':
			break
			default:
				return false
		}
		this.setState({
			win_X: e.clientX - this.state.origin_Xlen,
			win_Y: e.clientY - this.state.origin_Ylen,
		})
	}
	handle_corner_mouseUp (e) {
		window.onmousemove = null
		window.onmouseup = null
	}

    componentDidMount () {
        // 窗口绑定事件
        this.windowContainerRef.onmousedown = this.handle_container_mouseDown
        this.windowCornerTLRef.onmousedown = this.handle_corner_mouseDown
        this.windowCornerTRRef.onmousedown = this.handle_corner_mouseDown
        this.windowCornerBLRef.onmousedown = this.handle_corner_mouseDown
        this.windowCornerBRRef.onmousedown = this.handle_corner_mouseDown
    }

    render () {
        return (
            <div ref={(ref) => {this.windowContainerRef = ref}} className="window_container" style={{
	            	top: this.state.win_Y + 'px',
					left: this.state.win_X + 'px',
					width:  this.state.win_W + 'px',
					height: this.state.win_H + 'px'
	            }}>
                <div className="window_header">
                    <span className="window_title">
                        Window Title
                    </span>
                    <span className="window_actionBtn">
                        <span className="window_actionBtn_outer">
                            <span className="window_actionBtn_min"></span>
                        </span>
                        <span className="window_actionBtn_outer">
                            <span className="window_actionBtn_maxAndResume"></span>
                        </span>
                        <span className="window_actionBtn_outer">
                            <span className="window_actionBtn_close"></span>
                        </span>
                    </span>
                </div>
                <div className="window_body">
                	Window	
                </div>
                
                <div className="window_corner window_corner_tl" name="tl" ref={(ref) => {this.windowCornerTLRef = ref}}>
	                <div className="window_corner_row"></div>
	                <div className="window_corner_col"></div>
				</div>
				<div className="window_corner window_corner_tr" name="tr" ref={(ref) => {this.windowCornerTRRef = ref}}>
	                <div className="window_corner_row"></div>
	                <div className="window_corner_col"></div>
				</div>
				<div className="window_corner window_corner_bl" name="bl" ref={(ref) => {this.windowCornerBLRef = ref}}>
	                <div className="window_corner_row"></div>
	                <div className="window_corner_col"></div>
				</div>
				<div className="window_corner window_corner_br" name="br" ref={(ref) => {this.windowCornerBRRef = ref}}>
	                <div className="window_corner_row"></div>
	                <div className="window_corner_col"></div>
				</div>
                
            </div>
        )
    }
}

export default Window