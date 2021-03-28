<template>
    <el-card class="box-card" shadow="false">
        <template #header>
            <div class="card-header">
                <span>{{ name }}</span>
            </div>
        </template>
        <el-slider
            v-model="values"
            :min="min"
            :max="max"
            range
        >
        </el-slider>
        <div class="ranges">
            <el-input-number @change="changeMinValue"
                             :model-value="values[0]"
                             size="mini"
                             :min="min"
                             controls-position="right"
            />
            <el-input-number @change="changeMaxValue"
                             :model-value="values[1]"
                             size="mini"
                             :max="max"
                             controls-position="right"
            />
        </div>
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
            values: [this.min, this.max]
        }
    },
    methods: {
        changeValues(vv) {
            this.values = vv
        },
        changeMinValue(v) {
            this.values = [v, this.values[1]]
        },
        changeMaxValue(v) {
            this.values = [this.values[0], v]
        }
    }
}
</script>

<style scoped>
.box-card {
    margin-bottom: 15px;
}

.ranges {
    display: flex;
    justify-content: space-evenly;
    flex-wrap: wrap;
}

.el-input-number {
    margin-top: 10px;
    max-width: 110px;
}
</style>