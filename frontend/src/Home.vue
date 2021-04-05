<template>
    <el-container id="container">
        <el-main>
            <el-row :gutter="15">

                <el-col :span="4" :offset="4">
                    <filters-list/>
                </el-col>

                <el-col :span="12">
                    <search-input/>
                    <correction/>
                    <filters-active/>
                    <cars-list/>
                </el-col>

            </el-row>
        </el-main>
    </el-container>
</template>

<script>
import SearchInput from "@/components/SearchInput";
import FiltersList from "@/components/FiltersList";
import Correction from "@/components/Correction";
import FiltersActive from "@/components/FiltersActive";
import CarsList from "@/components/CarsList";

export default {
    name: 'Home',
    components: {
        "search-input": SearchInput,
        "filters-list": FiltersList,
        "correction": Correction,
        "filters-active": FiltersActive,
        "cars-list": CarsList
    }

}
</script>

<style>
* {
    font-family: "Helvetica", sans-serif;
}

#container {
    margin-top: 50px;
}
</style>

