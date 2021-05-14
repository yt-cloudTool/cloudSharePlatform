/*
    上传组件的核心功能
    父组件参数 {
    	mode:模式<string> READONLY(只读) NOCREATE(仅无创建功能) NORMAL(全功能)
		tmpSelector: 显示tmp选择器<string> SHOW(显示) HIDE(隐藏后传的是临时状态)
	}
*/
import React, { Component } from 'react'
import "./fileBox.css"

class FileBox extends Component {
	constructor (props) {
		super(props)
		this.state = {
			mode: this.props['mode'] || "NORMAL", // 模式 READONLY(只读) NOCREATE(仅无创建功能) NORMAL(全功能)
			tmpSelector: this.props['tmpSelector'] || "SHOW", // 显示tmp选择器 SHOW(显示) HIDE(隐藏后传的是临时状态)
            box_id: this.props['boxId'] || '',

            selectedFiles: [],
            box_name: this.props['name'],
            is_tmp: 0, // 0正式 1临时
            is_pub: 0, // 0私有 1公开         
            fileInputer: null,
            cur_uploadedFileList: [], // 当前上传操作中返回的文件列表
        }
        
       	this.handle_inputChange 	= this.handle_inputChange.bind(this) 
		this.handle_selectChange 	= this.handle_selectChange.bind(this)
       	this.fileInputRef 			= this.fileInputRef.bind(this)
        this.handle_clickFileInput 	= this.handle_clickFileInput.bind(this)
        this.handle_fileInputChange = this.handle_fileInputChange.bind(this)
        this.handle_create			= this.handle_create.bind(this)
        this.handle_save			= this.handle_save.bind(this)
	}

	componentDidMount () {
    }

	static getDerivedStateFromProps (props, state) {
		return {
			mode: 			props['mode'] 			|| "NORMAL",
			tmpSelector: 	props['tmpSelector'] 	|| "SHOW",
            box_id: 		props['boxId'] 			|| '',
		}
	}
	
	handle_inputChange (e) {
		if (e.target.name == "boxname") {
			this.setState({ box_name: e.target.value })
		}
	}
	
	handle_selectChange (e) {
		if (e.target.name == "isTmp") {
			this.setState({is_tmp: e.target.value})
		}
		else if (e.target.name == "isPub") {
			this.setState({is_pub: e.target.value})
		}
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
			"headers": { "token": window.localStorage.getItem('token') }
		}).then((res) => {
        	if (res.data.status === 1) {
				// 返回文件存储信息 
				let data = res.data.data
				this.setState({
					"cur_uploadedFileList": data,
					"selectedFiles": [ ...this.state.selectedFiles, ...data ]
				})
			}
        }).catch((err) => {
        	
        })
	}
	
	// 返回fileBox id
	handle_create () {
		return new Promise((resolve, reject) => {
			window.$axios({
				"url": window.$api.createFileBox,
				"method": "POST",
				"data": window.$qs.stringify({ // json转formData
					box_name: 	this.state.box_name,
					is_tmp: 	this.state.tmpSelector === "SHOW" ? this.state.is_tmp : "0",
					is_pub: 	this.state.is_pub,
					files: 		this.state.selectedFiles,
				}),
				"headers": { "token": window.localStorage.getItem('token') }
			}).then((res) => {
				if (res.data.status === 1) {
					this.setState({ box_id: res.data.data.InsertedID })
					resolve(res)
				} else {
					reject(res)
				}
			}).catch((err) => {
				reject(err)
			})
		})
	}
	
	// 保存更改
	handle_save () {
		return new Promise((resolve, reject) => {
			window.$axios({
				"url": window.$api.fileBoxInsert,
				"method": "POST",
				"data": window.$qs.stringify({ // json转formData
					box_id: 	this.state.box_id,
					files: 		this.state.selectedFiles,
				}),
				"headers": { "token": window.localStorage.getItem('token') }
			}).then((res) => {
				if (res.data.status === 1) {
					resolve(res)
				} else {
					reject(res)
				}
			}).catch((err) => {
				reject(err)
			})
		})
	}
	
	// 获取已传数据
	getList (box_id) {
		window.$axios({
        	"url": window.$api.fileBoxInsert,
			"method": "POST",
	        "data": window.$qs.stringify({ // json转formData
				box_id: 	this.state.box_id,
		        files: 		this.state.selectedFiles,
			}),
			"headers": { "token": window.localStorage.getItem('token') }
		}).then((res) => {
        	if (res.data.status === 1) {
			}
        }).catch((err) => {
        	
        })
	}
	
    render () {
        return (
            <div className="fileBox_container">
            	<div>
					{
		            	(this.state.mode == 'NORMAL') ? (
							<div className="fileBox_controlArea">
								<div className="fileBox_controlArea_boxname">
									<input className="fileBox_input_boxName normalInput" name="boxname" value={this.state.box_name} onChange={this.handle_inputChange} placeholder="Box Name"/>
									
									{
										this.state.tmpSelector === "SHOW" ? (<select className="fileBox_typeSelect normalSelect" name="isTmp" value={this.state.is_tmp} onChange={this.handle_selectChange}>
											<option value="0">Tmp</option>
											<option value="1">Regu</option>
										</select>) : ''
									}
									<select className="fileBox_typeSelect normalSelect" name="isPub" value={this.state.is_pub} onChange={this.handle_selectChange}>
										<option value="0">Priv</option>
										<option value="1">Pub</option>
									</select>
									{
										!this.state.box_id ? (<button className="normalButton" onClick={this.handle_create}>Create</button>) : ''
									}
								</div>
							</div>
						) : ''
					}
					{
						(this.state.mode === 'NORMAL' || this.state.mode === 'NOCREATE') ? (
							<div className="flex">
								<input ref={this.fileInputRef} style={{"display": "none"}} type="file" multiple="true" onChange={this.handle_fileInputChange}/>
								<button className="normalButton w100" onClick={this.handle_clickFileInput}>Select Files</button>
							</div>
						) : ''
					}
					<div className="fileBox_fileListContainer">
						<ul className="fileBox_fileList">
							{
								this.state.selectedFiles.map((ite, ind) => {
									return (
										<li className="fileBox_fileList_li">{ite.filename}</li>
									)
								})				
							}
						</ul>
					</div>
				</div>
				{
					(this.state.mode === 'NORMAL' || this.state.mode === 'NOCREATE') && this.state.box_id ? (
						<div className="fileBox_footer">
							<button className="fileBox_footer_saveBtn normalButton w100" onClick={this.handle_save}>Save</button>
						</div>
					) : ''
				}
            </div>
        )
    }
    
    
}

FileBox.defaultProps = {
	mode: "READONLY"
}

export default FileBox