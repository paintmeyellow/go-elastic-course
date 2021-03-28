import {createStore} from 'vuex'

const store = createStore({
    state() {
        return {
            cars: [],
            filters: [],
            searched: false,
        }
    },
    mutations: {
        updateCars(state, cars) {
            state.cars = cars
            state.searched = true
        },
        updateFilters(state, filters) {
            state.filters = filters
        }
    }
})

export default store