<template>
    <el-card class="box-card" shadow="false">
        <template #header>
            <div class="card-header">
                <span>{{ name }}</span>
            </div>
        </template>
        <el-checkbox-group v-model="checkList">
            <el-space direction="vertical" alignment="left">
                <el-checkbox
                    :key="variant"
                    v-for="variant in variants"
                    :label="variant.value">
                    {{ variant.value }}
                    <span class="badge">{{variant.count}}</span>
                </el-checkbox>
            </el-space>
        </el-checkbox-group>
    </el-card>
</template>

<script>
export default {
    name: "CheckboxFilter",
    props: {
        name: String,
        variants: Array
    },
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