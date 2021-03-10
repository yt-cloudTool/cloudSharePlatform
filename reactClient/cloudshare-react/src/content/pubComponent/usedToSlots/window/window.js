import React, { Component } from 'react'
import "./window.css"

class Window extends Component {
    constructor (props) {
        super(props)

        this.state = {
            win_X: 0,
            win_Y: 0,
            win_W: 200,
            win_H: 200,
        }

        this.containerRef = null

        this.handle_container_mouseDown = this.handle_container_mouseDown.bind(this)
        this.handle_container_mouseMove = this.handle_container_mouseMove.bind(this)
        this.handle_container_mouseUp   = this.handle_container_mouseUp.bind(this)
    }

    handle_container_mouseDown (e) {
        let mouse_X = e.clientX
        let mouse_Y = e.clientY
        let dom_W   = e.target.offsetWidth
        let dom_H   = e.target.offsetHeight
        console.log(`mouse_X:${mouse_X} mouse_Y:${mouse_Y} dom_W:${dom_W} dom_H:${dom_H}`, this.containerRef.style)
        this.setState({ win_W: 400, win_H: 400 })
    }
    handle_container_mouseMove (e) {

    }
    handle_container_mouseUp (e) {

    }

    componentDidMount () {
        // 窗口绑定事件
        this.containerRef.onmousedown = this.handle_container_mouseDown
        this.containerRef.onmousemove = this.handle_container_mouseMove
        this.containerRef.onmouseup   = this.handle_container_mouseUp
    }

    render () {
        return (
            <div ref={(ref) => {this.containerRef = ref}} className="window_container" style={{
                width:  this.state.win_W + 'px',
                height: this.state.win_H + 'px'
            }}>
                <div className="window_header">
                    <span className="window_title">
                        
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
            </div>
        )
    }
}

export default Window