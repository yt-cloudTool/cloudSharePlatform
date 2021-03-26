import React, { Component } from 'react'
import Icon from "./components/mainAreaIcon/mainAreaIcon.js"
import "./mainArea.css"

class MainArea extends Component {
	constructor (props) {
        super(props)

        this.state = {
            mainDataList: []
        }

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
        })
    }
    
    // 获取主数据
    getMainDataList () {
    	window.$axios.get(window.$api.serverinfo).then((res) => {
        	if (res.data.status === 1) {
				this.setState({
					mainDataList: res.data.data
				})
			}
        }).catch((err) => {
            
        })
    }
    
    render () {
        return (
            <div className="mainArea_container">
	            {
	            	this.state.mainDataList.map((ite, ind) => {
						return (
							<Icon label={ite.label} type={ite.type} img={ite.img}/>
						)
					})
	            }
            	<Icon/>
            </div>
        )
    }
    componentDidMount () {
    	this.getMainDataList()
    }
}

export default MainArea