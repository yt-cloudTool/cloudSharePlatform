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

export default {
    set_sideToolBar_show,
    set_loginStatus,
    set_loginPop_show
}