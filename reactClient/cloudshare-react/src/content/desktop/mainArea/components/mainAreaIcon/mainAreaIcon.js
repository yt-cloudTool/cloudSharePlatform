import React, { Component } from 'react'
import "./mainAreaIcon.css"

class MainAreaIcon extends Component {
    render () {
//    	type=="article"时内容为富文本
        return (
            <div className="mainAreaIcon_container">
            	{
					this.props.type === "-1"
					?
						this.props.img ? (<img className="mainAreaIcon_img" src={this.props.img}/>) : ''
					:
						(<img className="mainAreaIcon_img" src={
							this.props.type === 'note'
							?
								'note'
							:
								this.props.type === 'important'
								?
									'important'
								:
									this.props.type === 'filebox'
									?
										'filebox'
									:
										this.props.type === 'article'
										?
											'article'
										:
											'normal'
							
						}/>)
				}
				<p className="mainAreaIcon_label">{this.props.label}</p>
            </div>
        )
    }
    componentDidMount () {
    }
}

export default MainAreaIcon