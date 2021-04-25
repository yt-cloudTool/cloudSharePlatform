/*
    上传组件的核心功能
*/
import React, { Component } from 'react'
import "./fileBox.css"

class FileBox extends Component {
	constructor (props) {
		super(props)
		this.state = {
            selectedFile: [],
            fileInputer: null
        }
        
       this.fileInputRef = this.fileInputRef.bind(this)
        this.handle_clickFileInput = this.handle_clickFileInput.bind(this)
        this.handle_fileInputChange = this.handle_fileInputChange.bind(this)
	}
	fileInputRef (el) {
		this.setState({ fileInputer: el })
	}
	handle_clickFileInput (e) {
		this.state.fileInputer.click()
	}
	handle_fileInputChange (e) {
		console.log("e =======>", e)
		let files = e.target.files
		if (files.length < 1) {
			return false
		}
		
		// =====================================================================
		// 上传操作
		// 文件
		let formData = new FormData()
		for (let i=0; i<files.length; i++) {
			formData.append("file", files[i])
		}
		// 参数
		formData.append("count", files.length)
		formData.append("is_tmp", 0)
		formData.append("is_pub", 1)
		// 进行上传
		window.$axios({
			"url": window.$api.fileUpload, 
        	"method": "POST",
	        "data": formData,
			"headers": {
				"token": window.localStorage.getItem('token')
			}
		}).then((res) => {
        	if (res.data.status === 1) {
				// 返回文件存储信息 
				let data = res.data.data
			}
        }).catch((err) => {
        	
        })
	}
    render () {
        return (
            <div className="fileBox_container">
            	<div className="fileBox_fileSelector">
					<input ref={this.fileInputRef} style={{"display": "none"}} type="file" multiple="true" onChange={this.handle_fileInputChange}/>
					<button onClick={this.handle_clickFileInput}>Select Files</button>
				</div>
				<div className="fileBox_fileList">
				
				</div>
            </div>
        )
    }
}

export default FileBox