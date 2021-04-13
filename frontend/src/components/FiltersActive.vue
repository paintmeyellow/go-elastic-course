<template>
    <el-row id="filters-active">
        <el-col>
            <el-space>
                <el-tag
                    @close="clear(filter)"
                    :key="filter"
                    v-for="filter in filters"
                    closable
                    size="medium">
                    {{ filter.name }}: {{ filter.value }}
                </el-tag>
                <el-button
                    @click="clearAll"
                    v-if="filters.length > 0"
                    size="mini"
                    type="default">
                    Clear all
                </el-button>
            </el-space>
        </el-col>
    </el-row>
</template>

<script>
export default {
    name: "FiltersActive",
    computed: {
        filters() {
            let filters = []

            let query = {...this.$route.query}
            let range = JSON.parse(query.range || "{}")
            let checkbox = JSON.parse(query.checkbox || "{}")

            Object.keys(range).forEach(name => {
                let values = range[name]
                filters.push({name: name, value: `${values.min}-${values.max}`})
            })

            Object.keys(checkbox).forEach(name => {
                let values = checkbox[name]
                values.forEach(value => {
                    filters.push({name: name, value: value})
                })
            })

            return filters
        }
    },
    methods: {
        clear({name, value}) {
            let query = {...this.$route.query}

            //clear range filters
            let range = JSON.parse(query.range || "{}")

            Object.keys(range).forEach(_name => {
                if (name === _name) {
                   delete range[_name]
                }
            })

            if (Object.keys(range).length > 0) {
                query.range = JSON.stringify(range)
            } else {
                delete query.range
            }

            //clear checkbox filters
            let checkbox = JSON.parse(query.checkbox || "{}")
            let newCheckbox = {...checkbox}

            Object.keys(checkbox).forEach(_name => {
                if (name === _name) {
                    let values = checkbox[_name]
                    let newValues = {...values}
                    values.forEach((_value, i) => {
                        if (value === _value) {
                            delete newValues[i]
                        }
                    })
                    newCheckbox[_name] = Object.values(newValues)

                    if (Object.keys(newCheckbox[_name]).length === 0) {
                        delete newCheckbox[_name]
                    }
                }
            })

            if (Object.keys(newCheckbox).length > 0) {
                query.checkbox = JSON.stringify(newCheckbox)
            } else {
                delete query.checkbox
            }

            this.$router.push({query: query})
                .then(() => this.$store.dispatch('search'))
        },
        clearAll() {
            let query = {...this.$route.query}
            delete query.checkbox
            delete query.range

            this.$router.push({query: query})
                .then(() => this.$store.dispatch('search'))
        }
    }
}
</script>

<style scoped>
#filters-active {
    margin-top: 15px;
}
</style>