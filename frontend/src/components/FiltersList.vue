<template>
    <div>

        <range-filter name="Price" :min="0" :max="1000"/>

        <checkbox-filter name="Color"/>
        <checkbox-filter name="Year"/>

    </div>
</template>

<script>
import RangeFilter from "@/components/filter-types/RangeFilter";
import CheckboxFilter from "@/components/filter-types/CheckboxFilter";

export default {
    name: "FiltersList",
    components: {
        "range-filter": RangeFilter,
        "checkbox-filter": CheckboxFilter
    }
}
</script>

<style scoped>

</style>