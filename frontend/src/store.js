import {createStore} from 'vuex'

const store = createStore({
    state() {
        return {
            cars: [],
            searched: false,
        }
    },
    mutations: {
        updateCars(state, cars) {
            state.cars = cars
            state.searched = true
        }
    }
})

export default store