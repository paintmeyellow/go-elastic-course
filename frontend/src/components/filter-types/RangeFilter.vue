<template>
    <el-card class="box-card" shadow="false">
        <template #header>
            <div class="card-header">
                <span>{{ name }}</span>
            </div>
        </template>
        <el-slider
            v-model="value"
            range
            :min="min"
            :max="max"
            :marks="marks"
        >
        </el-slider>
    </el-card>
</template>

<script>
export default {
    name: "RangeFilter",
    props: {
        name: String,
        min: Number,
        max: Number,
    },
    data() {
        return {
            value: [0, 569],
            marks: {
                0: '0',
                1000: "1000"
            },
        }
    }

}
</script>

<style scoped>
.box-card {
    margin-bottom: 15px;
}
</style>