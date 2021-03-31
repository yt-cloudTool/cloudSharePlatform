import React, { Component } from 'react'
import "./articleEditor.css"

class ArticleEditor extends Component {
    constructor (props) {
        super(props)

        this.state = {
            articleEditor_show: false,
        }

        this.showPop           = this.showPop.bind(this)
        this.hidePop           = this.hidePop.bind(this)
        this.handleInputChange = this.handleInputChange.bind(this)
        this.confirm           = this.confirm.bind(this)

        window.$store.subscribe(() => {
            const _state_ = window.$store.getState()
            this.setState({
            	"articleEditor_show": _state_.articleEditor_show
            })
        })

    }

    showPop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(true))
    }
    hidePop () {
        window.$store.dispatch(window.$actions.set_loginPop_show(false))
        this.setState({ username: "", password: "" })
    }
    handleInputChange (e) {
        this.setState({ [e.target.name]: e.target.value })
    }
    confirm () {
        window.$axios.post(
        	window.$api.login, 
	        window.$qs.stringify({ // json转formData
				"loginname": this.state.username,
				"password":  this.state.password
			})
		).then((res) => {
        	if (res.data.status === 1) {
				// 存储 token 到 localStorage
				localStorage.setItem("token", 	  res.data.data.token)
				localStorage.setItem("loginname", res.data.data.loginname)
				localStorage.setItem("nickname",  res.data.data.nickname)
			}
        }).catch((err) => {
        	
        })
    }
    
    // 隐藏
    handle_articleEditor_hide () {
    		window.$store.dispatch(window.$actions.set_articleEditor_show(false))
    }
    
    // 显示上传文件框
    handle_fileBox_show () {
    	
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
								<select>
									<option value="Normal"> 		Normal	  	</option>
									<option value="Article"> 	Article	  	</option>
									<option value="Note"> 		Note	  		</option>
									<option value="Important"> 	Important 	</option>
									<option value="FileBox"> 	FileBox	  	</option>
								</select>
							</li>
							
							{/*如果 type != FileBox 则显示此项*/}
							<li id="articleEditor_fileBoxShow_btn" className="articleEditor_fileBoxShow_btn">
								<button onClick={this.handle_fileBox_show}>uploadFile</button>
							</li>
						</ul>
						
					</div>
					<div className="articleEditor_formArea">
					
						{/*如果 type == Normal 则显示此项*/}
						<div id="articleEditor_formArea_article">
							<input id="articleEditor_normalTitle_input" className="articleEditor_normalTitle_input"/>
							<textarea id="articleEditor_normalContent_textarea" className="articleEditor_normalContent_textarea"></textarea>
						</div>
						
						{/*如果 type == Note 则显示此项*/}
						<div id="articleEditor_formArea_article">
							<textarea id="articleEditor_noteContent_textarea" className="articleEditor_noteContent_textarea"></textarea>
						</div>
						
						{/*如果 type == Article 则显示此项*/}
						<div id="articleEditor_formArea_article">
							{/*富文本编辑器*/}
						</div>
						
						{/*如果 type == Important 则显示此项*/}
						<div id="articleEditor_formArea_article">
							{/*待定*/}
						</div>
						
						{/*如果 type == FileBox 则显示此项*/}
						<div id="articleEditor_formArea_article">
							<button>New Folder</button> {/*每添加一个folder就添加一个fileBox组件*/}
							<ul>
								{/*fileBox list*/}
							</ul>
							
						</div>
						
					</div>
					<div className="articleEditor_actionBtnArea">
					
					</div>
				</div>
			</div>
        </div>
    }
}

export default require('react-redux').connect((state) => { return state }, null)(ArticleEditor)