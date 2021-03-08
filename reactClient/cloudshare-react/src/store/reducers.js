import { combineReducers } from "redux"
import defaultState from "./state.js"

// 侧边栏显示
const sideToolBar_show = (state = defaultState.sideToolBar_show, action) => {
	switch (action.type)  {
		case 'SET_SIDETOOLBAR_SHOW': return action.data
		default: return state
	}
}

// 登录状态
const loginStatus = (state = defaultState.loginStatus, action) => {
	switch (action.type)  {
		case 'SET_LOGINSTATUS': return action.data
		default: return state
	}
}

// 登录弹窗显示
const loginPop_show = (state = defaultState.loginPop_show, action) => {
	switch (action.type)  {
		case 'SET_LOGINPOP_SHOW': return action.data
		default: return state
	}
}

export default combineReducers({
	sideToolBar_show,
	loginStatus,
	loginPop_show
})