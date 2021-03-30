<template>

    <range-filter :key="filter"
                  v-for="filter in filters.range"
                  :name="filter.name"
                  :min="filter.min"
                  :max="filter.max"/>

    <checkbox-filter :key="filter"
                     v-for="filter in filters.checkbox"
                     :name="filter.name"
                     :items="filter.items"/>

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
    }
}
</script>

<style scoped>

</style>