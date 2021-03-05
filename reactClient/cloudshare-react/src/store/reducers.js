import { combineReducers } from "redux"
import defaultState from "./state.js"

const sideToolBar_show = (state = defaultState.sideToolBar_show, action) => {
	switch (action.type)  {
		case 'SET_SIDETOOLBAR_SHOW':
			return action.data
		default:
			return state
	}
}

export default combineReducers({
	sideToolBar_show
})