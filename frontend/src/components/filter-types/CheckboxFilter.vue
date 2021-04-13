<template>
    <el-card class="box-card" shadow="false" body-style="overflow-y: auto; max-height: 300px;">
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
                    <span class="badge">{{ item.count }}</span>
                </el-checkbox>
            </el-space>
        </el-checkbox-group>
    </el-card>
</template>

<script>
import _ from "lodash";

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
    created() {
      let query = {...this.$route.query}
      let checkbox = JSON.parse(query.checkbox || "{}")
      this.checkList = checkbox[this.name] || []
    },
    methods: {
        onChange: _.debounce(function () {
            let query = {...this.$route.query}
            let checkbox = JSON.parse(query.checkbox || "{}")

            if (this.checkList.length === 0) {
                delete checkbox[this.name]
                if (Object.keys(checkbox).length > 0) {
                    query.checkbox = JSON.stringify(checkbox)
                } else {
                    delete query.checkbox
                }
            } else {
                checkbox[this.name] = this.checkList
                query.checkbox = JSON.stringify(checkbox)
            }

            this.$router.push({query: query})
                .then(() => this.$store.dispatch('search'))

        }, 500)
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