import React, { Component } from 'react'
import "./mainAreaIcon.css"

class MainAreaIcon extends Component {
    render () {
        return (
            <div className="mainAreaIcon_container">
            	{
					{/* type==-1时按img字段显示图片 否则按标准来显示 */}
					this.props.type === -1
					?
						this.props.img ? <img className="mainAreaIcon_img" src={this.props.img}/> : ''
					:
						<img src={
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
							
						}/>
				}
				<p className="mainAreaIcon_label">{this.props.label}</p>
            </div>
        )
    }
    componentDidMount () {
    }
}

export default MainAreaIcon