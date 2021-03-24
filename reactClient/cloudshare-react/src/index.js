import React            from 'react'
import ReactDOM         from 'react-dom'
import { Provider }     from "react-redux"
import App              from './App'
import store            from './store/store.js'
import actions          from './store/actions.js'
import reducers         from './store/reducers.js'
import reportWebVitals  from './reportWebVitals'
import axios            from 'axios'
import api				from './common/api.js'
import qs				from "qs"
import './index.css'

// ===============================================
// 请求拦截
axios.interceptors.request.use((config) => {
  const token = window.localStorage.getItem('token')
  if (token) {
      config.headers.Authorization = token
  }
  return config
}, (err) => {
  return Promise.reject(err)
})
// 响应拦截
axios.interceptors.response.use((res) => {
  return res
}, (err) => {
  if (err.response) {
    switch (err.response.status) {
      case 400:
        break
      case 401:
        break
      default: break
    }
  }
  return Promise.reject(err.response.data)
})

// ================================================
window.$store     = store
window.$actions   = actions
window.$reducers  = reducers
window.$axios     = axios
window.$api		  = api
window.$qs		  = qs	// json转为formData

// ================================================
ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
