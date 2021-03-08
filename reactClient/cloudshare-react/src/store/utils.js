import { connect } from 'react-redux'

export default {
    store_connect: function (STATE, MAPPER) {
        return connect(
            // 1
            (state) => {
                return (STATE ? STATE : state)
            },
            // 2
            MAPPER ? MAPPER : null
        )
    }
}