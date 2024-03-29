const set_sideToolBar_show = (data)=>{
    return (dispatch, getState) => {
        dispatch({ "type": "SET_SIDETOOLBAR_SHOW", "data": data })
    }
}

const set_loginStatus = (data)=>{
    return (dispatch, getState) => {
        dispatch({ "type": "SET_LOGINSTATUS", "data": data })
    }
}

const set_loginPop_show = (data)=>{
    return (dispatch, getState) => {
        dispatch({ "type": "SET_LOGINPOP_SHOW", "data": data })
    }
}

const set_windowList = (data)=>{
    return (dispatch, getState) => {
        dispatch({ "type": "SET_WINDOWLIST", "data": data })
    }
}

const set_sideToolBar_toggle = (data)=>{
    return (dispatch, getState) => {
        dispatch({ "type": "SET_SIDETOOLBAR_TOGGLE", "data": data })
    }
}

const set_articleEditor_show = (data)=> {
	return (dispatch, getState) => {
        dispatch({ "type": "SET_ARTICLEEDITOR_SHOW", "data": data })
    }
}

const set_mainData_refresh = (data /*timestamp*/)=> {
	return (dispatch, getState) => {
        dispatch({ "type": "SET_MAINDATA_REFRESH", "data": data })
    }
}




export default {
    set_sideToolBar_show,
    set_loginStatus,
    set_loginPop_show,
    set_windowList,
    set_sideToolBar_toggle,
    set_articleEditor_show,
    set_mainData_refresh
}