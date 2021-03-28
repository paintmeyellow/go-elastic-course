<template>
    <el-form @submit.prevent="search">
        <el-autocomplete
            class="el-input-group"
            v-model="query"
            clearable
            :fetch-suggestions="querySearch"
            placeholder="Please Input"
            :trigger-on-focus="false"
            @select="handleSelect"
        >
            <template #append>
                <el-button
                    native-type="submit"
                    icon="el-icon-search"
                    type="primary">
                    Search
                </el-button>
            </template>
        </el-autocomplete>
    </el-form>
</template>

<script>
import {onMounted, ref} from "@vue/runtime-core";
import axios from "axios";
import store from "@/store";
import {ElLoading} from "element-plus";

export default {
    name: "SearchInput",
    data() {
        return {
            query: "",
        }
    },
    mounted() {
        this.search()
    },
    methods: {
        search() {
            let loading = ElLoading.service({fullscreen: true})
            axios.post("http://localhost:8081/search/cars", JSON.stringify({query: this.query}))
                .then(response => {
                    setTimeout(() => {
                        console.log(response.data)
                        store.commit("updateCars", response.data["cars"])
                        store.commit("updateFilters", response.data["filters"])
                        this.$nextTick(() => loading.close());
                    }, 200)
                })
        }
    },
    setup() {
        const restaurants = ref([]);
        const querySearch = (queryString, cb) => {
            let results = queryString
                ? restaurants.value.filter(createFilter(queryString))
                : restaurants.value;
            // call callback function to return suggestions
            cb(results);
        };
        const createFilter = (queryString) => {
            return (restaurant) => {
                return (
                    restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) ===
                    0
                );
            };
        };
        const loadAll = () => {
            return [
                {"value": "vue", "link": "https://github.com/vuejs/vue"},
                {"value": "element", "link": "https://github.com/ElemeFE/element"},
                {"value": "cooking", "link": "https://github.com/ElemeFE/cooking"},
                {"value": "mint-ui", "link": "https://github.com/ElemeFE/mint-ui"},
                {"value": "vuex", "link": "https://github.com/vuejs/vuex"},
                {"value": "vue-router", "link": "https://github.com/vuejs/vue-router"},
                {"value": "babel", "link": "https://github.com/babel/babel"}
            ];
        };
        const handleSelect = (item) => {
            console.log(item);
        };
        onMounted(() => {
            restaurants.value = loadAll();
        });
        return {
            restaurants,
            state1: ref(''),
            state2: ref(''),
            querySearch,
            createFilter,
            loadAll,
            handleSelect,
        };
    }
    ,
}
</script>

<style scoped>

</style>