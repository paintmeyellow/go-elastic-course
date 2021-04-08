import {createStore} from 'vuex'
import {ElLoading} from "element-plus";
import axios from "axios";
import router from "@/router.js";

const store = createStore({
    state() {
        return {
            cars: [],
            filters: {
                "range": [],
                "checkbox": []
            },
            searched: false,
        }
    },
    getters: {
        checkboxByName: (state) => (name) => {
            return state.filters.checkbox.find(item => item.name === name)
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
    },
    actions: {
        search({commit}) {
            let queryParams = router.currentRoute.value.query

            let activeFilters = {}

            if (queryParams.checkbox) {
                activeFilters.checkbox = JSON.parse(queryParams.checkbox || "{}")
            }

            if (queryParams.range) {
                activeFilters.range = JSON.parse(queryParams.range || "{}")
            }

            let loading = ElLoading.service({fullscreen: true})
            axios.post("http://localhost:8081/search/cars", JSON.stringify({
                query: queryParams.query,
                active_filters: activeFilters
            }))
                .then(response => {
                    setTimeout(() => {
                        commit("updateCars", response.data["cars"])
                        commit("updateFilters", response.data["filters"])
                        loading.close();
                    }, 200)
                })
        }
    }
})

export default store