<template>
    <el-row id="filters-active">
        <el-col>
            <el-space>
                <el-tag closable size="medium">Color: Red</el-tag>
                <el-tag closable size="medium">Color: Green</el-tag>
                <el-tag closable size="medium">Year: 2008</el-tag>
                <el-button size="mini" type="default">Clear all</el-button>
            </el-space>
        </el-col>
    </el-row>
</template>

<script>
export default {
    name: "FiltersActive"
}
</script>

<style scoped>
#filters-active {
    margin-top: 15px;
}
</style>