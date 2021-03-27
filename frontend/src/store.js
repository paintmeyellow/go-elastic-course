import {createStore} from 'vuex'

const store = createStore({
    state() {
        return {
            carsList: [],
        }
    },
    mutations: {
        updateCarsList(state, list) {
            state.carsList = list
        }
    }
})

export default store