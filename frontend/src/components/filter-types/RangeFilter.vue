<template>
    <el-card class="box-card" shadow="false">
        <template #header>
            <div class="card-header">
                <span>{{ name }}</span>
            </div>
        </template>
        <el-slider
            v-model="values"
            @change="onChange"
            :min="min"
            :max="max"
            range
        >
        </el-slider>
        <div class="ranges">
            <el-input-number
                v-model="values[0]"
                @change="onChange"
                size="mini"
                :min="min"
                controls-position="right"
            />
            <el-input-number
                v-model="values[1]"
                @change="onChange"
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
            values: [this.min, this.max]
        }
    },
    created() {
        let re = /(\w+)\[(-?\d+)-(-?\d+)]/
        let res = re.exec(this.$route.query.range || "")
        if (res !== null) {
            let [min, max] = res.splice(2)

            min = parseInt(min)
            max = parseInt(max)

            this.values = [min, max]
        }
    },
    methods: {
        onChange: _.debounce(function () {
            let query = {...this.$route.query}

            if (this.values[0] === this.min && this.values[1] === this.max) {
                delete query.range
            } else {
                query.range = `${this.name}[${this.values[0]}-${this.values[1]}]`
            }

            this.$router.push({
                ...this.$router.currentRoute,
                query: query,
            })
                .then(() => this.$store.dispatch('search'))
        }, 500)
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