import React, { Component } from 'react'
import Icon from "./components/mainAreaIcon/mainAreaIcon.js"
import "./mainArea.css"
import FileBox from "../../pubComponent/normal/fileBox/fileBox.js"

class MainArea extends Component {
	constructor (props) {
        super(props)

        this.state = {
            dataList: [],
            page: 1,
            size: 100,
            mainData_refresh: 0, // 主列表刷新时间戳标识
        }

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            // 如果判断有更新主列表数据状态
            if (_state_.mainData_refresh > this.state.mainData_refresh) {
	            this.setState({ "mainData_refresh": _state_.mainData_refresh })
				this.getMainDataList() // 更新数据      	
            }
        })
    }
    
    // 获取主数据
    getMainDataList () {
    	window.$axios.get(window.$api.articleList, {
			params: {
				page: this.state.page,
				size: this.state.size
			}
		}).then((res) => {
			console.log('res =>', res)
        	if (res.data.status === 1) {
				this.setState({ dataList: res.data.data })
			}
        }).catch((err) => {
            
        })
    }
    
    // 跳转页数
    handlePageChange (pageNum) {
    	this.setState({ page: pageNum})
    }
    
    // 每页数量
    handleSizeChange (sizeNum) {
    	this.setState({ size: sizeNum })
    }
    
    render () {
        return (
            <div className="mainArea_container">
            	<div className="mainArea_inner">
					<FileBox/>
		            {
		            	this.state.dataList.map((ite, ind) => {
							return (
								<Icon key={ind} label={ite.label} type={ite.type} img={ite.img}/>
							)
						})
		            }
				</div>
            </div>
        )
    }
    componentDidMount () {
   		this.getMainDataList()
    }
}

export default MainArea