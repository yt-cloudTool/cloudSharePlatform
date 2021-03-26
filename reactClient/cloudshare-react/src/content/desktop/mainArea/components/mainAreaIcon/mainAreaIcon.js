import React, { Component } from 'react'
import "./mainAreaIcon.css"

class MainAreaIcon extends Component {
    render () {
        return (
            <div className="mainAreaIcon_container">
            	{
					this.props.type != -1 ?
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
					:
						this.props.img ? <img className="mainAreaIcon_img" src={this.props.img}/> : ''
				}
				<p className="mainAreaIcon_label">{this.props.label}</p>
            </div>
        )
    }
    componentDidMount () {
    }
}

export default MainAreaIcon