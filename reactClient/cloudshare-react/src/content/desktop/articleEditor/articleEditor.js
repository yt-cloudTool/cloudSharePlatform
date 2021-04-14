import React, { Component } from 'react'
import BraftEditor from 'braft-editor'
import 'braft-editor/dist/index.css'
import "./articleEditor.css"

class ArticleEditor extends Component {
    constructor (props) {
        super(props)

        this.state = {
            articleEditor_show: false,
            // 表单
            form_type: 			"normal", // 文章类型
            form_iconImg: 		"", // 自定义图标
            form_title: 		"",	// 标题
            form_content: 		null, // 内容 | 富文本内容
        }

        this.handle_form_change 			= this.handle_form_change.bind(this)
        this.handle_formBraftEditor_change 	= this.handle_formBraftEditor_change.bind(this)
        this.handleCancel      				= this.handleCancel.bind(this)
        this.handleSubmit	   				= this.handleSubmit.bind(this)

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            this.setState({
            	"articleEditor_show": _state_.articleEditor_show
            })
        })

    }

	// 表单中数据改变时
    handle_form_change (e) {
        this.setState({ [e.target.name]: e.target.value })
        // 文章类型改变时
        if (e.target.name === "form_type") {
        	// 判断哪些表单显示哪些表单清空内容
			switch (e.target.value) {
				case 'normal': 
				break
				case 'article':
				break
				case 'note':
				break
				case 'important':
				break
				case 'filebox':
				break
			}
		}
		// 自定义图标上传输入框改变时
		else if (e.target.name === "form_iconImg") {
			
		}
		// 文章标题
		else if (e.target.name === "form_label") {
			
		}
		// 文章内容
		else if (e.target.name === "form_content") {
			
		}
//		// 如果是富文本编辑器
//		else if (e.target.name === "braftEditor") {
			
//		}
    }
    // 表单中富文本方法
    handle_formBraftEditor_change (editorState) {
	   	this.setState({"form_content": editorState})
		console.log("e ==========>", this.state.form_content)

    }
    
    // 隐藏
    handle_articleEditor_hide () {
    	window.$store.dispatch(window.$actions.set_articleEditor_show(false))
    }
    
    // 显示上传文件框
    handle_fileBox_show () {
    	
    }
    
    // 取消按钮
    handleCancel () {
    	this.handle_articleEditor_hide()
    }
    
    // 确定按钮
    handleSubmit () {
    	window.$axios.post(
        	window.$api.articleInsert, 
			window.$qs.stringify({ 					// json转formData
				type: 	   this.state.form_type, 	// 类型 (判断图标类型)
				img: 	   this.state.form_iconImg, // 图标 (如果type==-1则用此字段显示图标)
				label: 	   this.state.form_label,   // 标签名称 (可重复)
				content:   this.state.form_content.toHTML ? this.state.form_content.toHTML() : this.state.form_content, // 内容
				fileboxid: '',             		 	// 字符串数组
			})
		).then((res) => {
        	if (res.data.status === 1) {
				this.handle_articleEditor_hide()
				// 状态还原
				this.setState({
		            form_type: 	  	"normal", 	// 文章类型
		            form_iconImg: 	"", 		// 自定义图标
		            form_title:   	"",			// 标题
		            form_content: 	null, 		// 内容 | 富文本内容
		        })
			}
			// 刷新主列表
			window.$store.dispatch(window.$actions.set_mainData_refresh(new Date().getTime()))
        }).catch((err) => {
        	
        })
    }

    render () {
    	let containerStyle = {
			display: this.state.articleEditor_show ? "flex" : "none"
		}
        return <div className="articleEditor_container" style={containerStyle}>
          	<div className="articleEditor_inner">
				<div className="articleEditor_titleArea loginPop_header">
					<span>New Article</span>
					<span className="loginPop_closeBtn_inner" onClick={this.handle_articleEditor_hide}></span>
				</div>
				
				<div className="articleEditor_contentArea">
					<div className="articleEditor_toolArea">
						<ul>
							
							<li className="articleEditor_articleTypeSelector">
								<select className="articleEditor_form" value={this.state.form_type} name="form_type" onChange={this.handle_form_change}>
									<option value="normal"> 	Normal	  	</option>
									<option value="article"> 	Article	  	</option>
									<option value="note"> 		Note	  	</option>
									<option value="important"> 	Important 	</option>
									<option value="filebox"> 	FileBox	  	</option>
								</select>
							</li>
							
							{/*自定义图标的上传*/}
							<li className="articleEditor_iconUpdate">
								<input type="file" value={this.state.iconImg} name="form_iconImg" onChange={this.handle_form_change}/>
							</li>
							
							{/*如果 type != FileBox 则显示此项*/}
							<li id="articleEditor_fileBoxShow_btn" className="articleEditor_fileBoxShow_btn" style={{display: this.state.form_type !== "filebox" ? "inline-block" : "none"}}>
								<button className="headerPop_actionBtn" onClick={this.handle_fileBox_show}>UploadFile</button>
							</li>
						</ul>
						
					</div>
					<div className="articleEditor_formArea">
						<div className="articleEditor_formContainer">
							<input id="articleEditor_normalTitle_input" name="form_label" onChange={this.handle_form_change} className="articleEditor_form articleEditor_normalTitle_input" placeholder="Label"/>
						</div>
						{/*如果 type == Normal 则显示此项*/}
						<div id="articleEditor_formArea_article" className="articleEditor_formContainer" style={{display: this.state.form_type === "normal" ? "flex" : "none"}}>
							<textarea id="articleEditor_normalContent_textarea" name="form_content" onChange={this.handle_form_change} className="articleEditor_form articleEditor_normalContent_textarea" placeholder="Content"></textarea>
						</div>
						
						{/*如果 type == Note 则显示此项*/}
						<div id="articleEditor_formArea_article" className="articleEditor_formContainer" style={{display: this.state.form_type === "note" ? "block" : "none"}}>
							<textarea id="articleEditor_noteContent_textarea" name="form_content" onChange={this.handle_form_change} className="articleEditor_form articleEditor_noteContent_textarea" placeholder="Content"></textarea>
						</div>
						
						{/*如果 type == Article 则显示此项*/}
						<div id="articleEditor_formArea_article" className="articleEditor_formContainer" style={{display: this.state.form_type === "article" ? "block" : "none"}}>
							{/*富文本编辑器*/}
							<BraftEditor
					        	value={this.state.form_content}
					        	onChange={this.handle_formBraftEditor_change}
					        	onSave={()=>{}}
					        />
						</div>
						
						{/*如果 type == Important 则显示此项*/}
						<div id="articleEditor_formArea_article" className="articleEditor_formContainer" style={{display: this.state.form_type === "important" ? "block" : "none"}}>
							{/*待定*/}
						</div>
						
						{/*如果 type == FileBox 则显示此项*/}
						<div id="articleEditor_formArea_article" className="articleEditor_formContainer" style={{display: this.state.form_type === "filebox" ? "block" : "none"}}>
							<button>New Folder</button> {/*每添加一个folder就添加一个fileBox组件*/}
							<ul>
								{/*fileBox list*/}
							</ul>
							
						</div>
						
					</div>
					<div className="articleEditor_formContainer articleEditor_actionBtnArea">
						<button className="headerPop_actionBtn" onClick={this.handleCancel}>Cancel</button>
						<button className="headerPop_actionBtn" onClick={this.handleSubmit}>Submit</button>
					</div>
				</div>
			</div>
        </div>
    }
}

export default require('react-redux').connect((state) => { return state }, null)(ArticleEditor)