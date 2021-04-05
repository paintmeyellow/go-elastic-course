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
            <el-input-number
                v-model="minValue"
                size="mini"
                :min="min"
                controls-position="right"
            />
            <el-input-number
                v-model="maxValue"
                size="mini"
                :max="max"
                controls-position="right"
            />
        </div>
    </el-card>
</template>

<script>
import _ from 'lodash'

export default {
    name: "RangeFilter",
    props: {
        name: {
            type: String,
            required: true
        },
        min: {
            type: Number,
            required: true
        },
        max: {
            type: Number,
            required: true
        },
    },
    data() {
        return {
            minValue: this.min,
            maxValue: this.max,
        }
    },
    computed: {
        values: {
            set([min, max]) {
                this.minValue = min
                this.maxValue = max
            },
            get() {
                return [this.minValue, this.maxValue]
            }
        }
    },
    watch: {
        values: _.debounce(function () {
            this.$emit('onChange', this.values)
        }, 1000)
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