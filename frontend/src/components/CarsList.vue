<template>
    <el-row id="cars-list">
        <el-space wrap :size="10" style="justify-content: flex-start">
            <el-card class="box-card"
                     shadow="never"
                     :body-style="{ padding: '0px' }"
                     :key="product.id"
                     v-for="product in carsList"
            >
                <img src="../assets/placeholder.png"
                     class="image"
                     alt="{{product.make}}"
                >
                <div style="padding: 14px;">
                    <span>{{ product.make }} {{ product.model }}</span>
                    <div class="bottom">
                        <p><strong>${{ product.params.price }}</strong></p>
                    </div>
                </div>
            </el-card>
        </el-space>

        <el-empty v-if="isCarsListEmpty"
                  description="Nothing Found – Sorry, but nothing matched your search criteria."
        />
    </el-row>
</template>

<script>
import {mapState} from "vuex/dist/vuex.mjs";

export default {
    name: "CarsList",
    computed: {
        ...mapState({
            carsList: state => state.carsList,
        }),
        isCarsListEmpty() {
            return this.carsList === null || this.carsList.length === 0
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