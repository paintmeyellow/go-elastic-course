<template>
    <el-row>
        <el-col id="cars-list" v-if="!isCarsEmpty">

            <div class="card" :key="car.id" v-for="car in cars">
                <img src="../assets/placeholder.png" class="image" alt="{{car.make}}">
                <div class="detail title">
                    {{ car.make }} {{ car.model }} {{ car.params.color }} {{ car.params.year }}
                </div>
                <div class="detail price">
                    ${{ formatPrice(car.params.price) }}
                </div>
            </div>
        </el-col>

        <el-col id="cars-empty" v-if="isCarsEmpty">
            <el-empty description="Nothing Found – Sorry, but nothing matched your search criteria."/>
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
        },
    },
    methods: {
        formatPrice(price) {
            return new Intl.NumberFormat('fr-FR').format(price)
        }
    }
}
</script>

<style scoped>
#cars-list {
    margin-top: 15px;
    display: flex;
    justify-content: flex-start;
    align-items: stretch;
    flex-wrap: wrap;
}

#cars-empty {
    margin-top: 15px;
    display: flex;
    justify-content: center;
}

.card img {
    max-width: 280px;
}

.card {
    padding: 0;
    margin: 0 10px 10px 0;
    width: 280px;
    border-radius: 4px;
    border: 1px solid #EBEEF5;
    background-color: #FFF;
    overflow: hidden;
    color: #303133;
    -webkit-transition: .3s;
    transition: .3s;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.card .detail {
    padding: 14px;
    display: flex;
    flex-direction: column;
}

.card .detail.price {
    font-size: 22px;
    font-weight: bold;
    color: #3e3e3e;
}
</style>