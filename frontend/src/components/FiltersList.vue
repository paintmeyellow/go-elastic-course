<template>

    <template :key="filter" v-for="(filter, name) in filters">
        <range-filter v-if="isRange(filter)" :name="name" :min="filter.min" :max="filter.max"/>
        <checkbox-filter v-if="isCheckbox(filter)" :name="name" :variants="filter.variants"/>
    </template>

</template>

<script>
import RangeFilter from "@/components/filter-types/RangeFilter";
import CheckboxFilter from "@/components/filter-types/CheckboxFilter";
import {mapState} from "vuex/dist/vuex.mjs";

export default {
    name: "FiltersList",
    components: {
        "range-filter": RangeFilter,
        "checkbox-filter": CheckboxFilter
    },
    computed: {
        ...mapState({
            filters: state => state.filters,
        }),
    },
    methods: {
        isRange(filter) {
            return filter.type === "range"
        },
        isCheckbox(filter) {
            return filter.type === "checkbox"
        }
    }
}
</script>

<style scoped>

</style>