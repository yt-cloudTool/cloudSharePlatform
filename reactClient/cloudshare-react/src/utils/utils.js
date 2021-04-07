export default {
    
    // 返回特定格式时间日期
    GetDate: function (timestamp) {
        let dateObj = null
        if (timestamp) {
            dateObj = new Date(timestamp)
        } else {
            dateObj = new Date()
        }
        const year    = dateObj.getFullYear()
        const month   = (dateObj.getMonth() + 1) >= 10 ? (dateObj.getMonth() + 1) : ('0' + (dateObj.getMonth() + 1))
        let day       = dateObj.getDay()
        switch (day) {
            case 7: day = 'SUN'; break
            case 1: day = 'MON'; break
            case 2: day = 'TUE'; break
            case 3: day = 'WED'; break
            case 4: day = 'THU'; break
            case 5: day = 'FRI'; break
            case 6: day = 'SAT'; break
            default: break
        }
        const date    = dateObj.getDate() >= 10 ? dateObj.getDate() : ('0' + dateObj.getDate())
        const hour    = dateObj.getHours()
        const minute  = dateObj.getMinutes() >= 10 ? dateObj.getMinutes() : ('0' + dateObj.getMinutes())
        const seconds = dateObj.getSeconds() >= 10 ? dateObj.getSeconds() : ('0' + dateObj.getSeconds())
        const dateTimeStr = `${year}-${month}-${date} ${day} ${hour}:${minute}:${seconds}` 
        return dateTimeStr
    }
}