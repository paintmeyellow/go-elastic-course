<template>
    <el-card class="box-card" shadow="false">
        <template #header>
            <div class="card-header">
                <span>Color</span>
            </div>
        </template>
        <el-checkbox-group v-model="checkList">
            <el-space direction="vertical" alignment="left">
                <el-checkbox label="Red">Red<span class="badge">131</span></el-checkbox>
                <el-checkbox label="Green">Green<span class="badge">52</span></el-checkbox>
                <el-checkbox label="Blue">Blue<span class="badge">8</span></el-checkbox>
            </el-space>
        </el-checkbox-group>
    </el-card>
</template>

<script>
export default {
    name: "CheckboxFilter",
    data() {
        return {
            checkList: ['Option A'],
        }
    }
}
</script>

<style scoped>
.badge {
    display: inline-block;
    padding: .25em .6em .25em .6em;
    font-size: 75%;
    font-weight: 700;
    line-height: 1;
    text-align: center;
    border-radius: 10rem;
    color: #fff;
    background-color: #cdd3d4;
    margin-left: 3px;
}

.box-card {
    margin-bottom: 15px;
}
</style>