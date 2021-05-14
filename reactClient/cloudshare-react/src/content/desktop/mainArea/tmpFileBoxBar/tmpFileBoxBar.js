import React, { Component } from 'react'
import { GetYMDHMS } from "../../../../utils/utils"
import "./tmpFileBoxBar.css"

class TmpFileBox extends Component {
	constructor (props) {
        super(props)

        this.state = {
            dataList: [],
            page: 1,
            size: 100,
            customBoxName: this.createTmpBoxName(), // box名字
        }

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
        })
    }

    createTmpBoxName () {
        return GetYMDHMS() + "_BOX"
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
            <div className="tmpFileBox_container">
            	<h3 className="tmpFileBox_container_title">TmpFiles</h3>
            	<div className="tmpFileBox_inner">
                    <ul className="tmpFileBox_inner_list">

                    </ul>
				</div>
            </div>
        )
    }
    componentDidMount () {
   		this.getMainDataList()
    }
}

export default TmpFileBox