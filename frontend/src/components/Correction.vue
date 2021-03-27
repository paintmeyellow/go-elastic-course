<template>
    <el-row id="correction">
        <el-col>
            <span class="title">Did you mean:</span>
            <el-space size="small">
                <el-link type="primary" href="https://element.eleme.io" target="_blank">Ford</el-link>
            </el-space>
        </el-col>
    </el-row>
</template>

<script>
export default {
    name: "Correction"
}
</script>

<style scoped>
#correction {
    margin-top: 15px;
}

a {
    font-style: italic;
}

.title {
    margin-right: 10px;
    font-size: 11pt;
}
</style>