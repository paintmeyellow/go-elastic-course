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
                    :key="item"
                    v-for="item in items"
                    @change="onChange"
                    :label="item.value">
                    {{ item.value }}
                    <span class="badge">{{item.count}}</span>
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
        items: Array
    },
    data() {
        return {
            checkList: [],
        }
    },
    methods: {
        onChange(){
            console.log(this.checkList)
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