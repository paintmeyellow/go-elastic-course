<template>
    <el-row id="cars-list">
        <el-col>
            <el-space wrap :size="10" style="justify-content: flex-start">
                <el-card class="box-card"
                         shadow="never"
                         :body-style="{ padding: '0px' }"
                         :key="car.id"
                         v-for="car in cars"
                >
                    <img src="../assets/placeholder.png" class="image" alt="{{car.make}}">
                    <div style="padding: 14px;">
                        <span>{{ car.make }} {{ car.model }}</span>
                        <p><strong>${{ car.params.price }}</strong></p>
                    </div>
                </el-card>
            </el-space>

            <el-empty v-if="isCarsEmpty"
                      description="Nothing Found – Sorry, but nothing matched your search criteria."/>
        </el-col>
    </el-row>
</template>

<script>
import {mapState} from "vuex/dist/vuex.mjs";

export default {
    name: "CarsList",
    computed: {
        ...mapState({
            cars: state => state.cars,
            searched: state => state.searched,
        }),
        isCarsEmpty() {
            return this.searched && (this.cars === null || this.cars.length === 0)
        }
    }
}
</script>

<style scoped>
#cars-list {
    display: flex;
    justify-content: center;
    margin-top: 15px;
}

.box-card img {
    max-width: 280px;
}

.box-card {
    width: 280px;
}
</style>