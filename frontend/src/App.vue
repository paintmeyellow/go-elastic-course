<template>
  <el-container>
    <el-aside>
    </el-aside>
    <el-main>
      <el-row :gutter="40">

        <el-col :span="5" id="filters">

          <el-card class="box-card" shadow="false">
            <template #header>
              <div class="card-header">
                <span>Price</span>
              </div>
            </template>
            <el-slider
                v-model="value"
                range
                :max="100"
                :marks="marks"
            >
            </el-slider>

          </el-card>

          <el-card class="box-card" shadow="false">
            <template #header>
              <div class="card-header">
                <span>Color</span>
              </div>
            </template>
            <el-checkbox-group v-model="checkList">
              <el-space direction="vertical" alignment="left">
                <el-checkbox label="Red">Red<span class="badge">131</span></el-checkbox>
                <el-checkbox label="Green">Green<span class="badge">52</span></el-checkbox>
                <el-checkbox label="Blue">Blue<span class="badge">8</span></el-checkbox>
              </el-space>
            </el-checkbox-group>
          </el-card>

          <el-card class="box-card" shadow="false">
            <template #header>
              <div class="card-header">
                <span>Year</span>
              </div>
            </template>
            <el-checkbox-group v-model="checkList">
              <el-space direction="vertical" alignment="left">
                <el-checkbox label="2008"></el-checkbox>
                <el-checkbox label="1995"></el-checkbox>
                <el-checkbox label="2021"></el-checkbox>
              </el-space>
            </el-checkbox-group>
          </el-card>

        </el-col>

        <el-col :span="13">
          <el-autocomplete
              class="inline-input el-input-group"
              v-model="state2"
              :fetch-suggestions="querySearch"
              placeholder="Please Input"
              :trigger-on-focus="false"
              @select="handleSelect"
          >
            <template #append>
              <el-button icon="el-icon-search" type="primary">Search</el-button>
            </template>
          </el-autocomplete>

          <el-row id="correction">
            <el-col>
              <span class="title">Did you mean:</span>
              <el-space size="small">
                <el-link type="primary" href="https://element.eleme.io" target="_blank">Ford</el-link>
              </el-space>
            </el-col>
          </el-row>

          <el-row id="tags">
            <el-col>
              <el-space>
                <el-tag closable size="medium">Color: Red</el-tag>
                <el-tag closable size="medium">Color: Green</el-tag>
                <el-tag closable size="medium">Year: 2008</el-tag>
                <el-button size="mini" type="default">Clear all</el-button>
              </el-space>
            </el-col>
          </el-row>

          <el-row id="products">
            <el-space wrap :size="10" style="justify-content: space-between">
              <el-card class="box-card"
                       shadow="never"
                       :body-style="{ padding: '0px' }"
                       :key="product.id"
                       v-for="product in products"
              >
                <img src="./assets/placeholder.png"
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
          </el-row>

          <el-empty description="Nothing Found – Sorry, but nothing matched your search criteria."></el-empty>
        </el-col>

      </el-row>

    </el-main>
  </el-container>
</template>

<script>

import {onMounted, ref} from "@vue/runtime-core";

export default {
  name: 'App',
  components: {},
  data() {
    return {
      value: [30,69],
      checkList: ['Option A'],
      marks: {
        0: '0',
        25: '25',
        50: '50',
        75: '75',
        100: "100"
      },
      products: [
        {
          "id": 994,
          "make": "Toyota",
          "model": "4Runner",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Blue",
            "year": 1992,
            "price": 85030
          }
        },
        {
          "id": 995,
          "make": "GMC",
          "model": "3500",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Aquamarine",
            "year": 1997,
            "price": 46618
          }
        },
        {
          "id": 996,
          "make": "Mercury",
          "model": "Tracer",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Teal",
            "year": 1997,
            "price": 95290
          }
        },
        {
          "id": 997,
          "make": "Suzuki",
          "model": "SJ 410",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Teal",
            "year": 1984,
            "price": 64626
          }
        },
        {
          "id": 998,
          "make": "Mitsubishi",
          "model": "Precis",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Fuscia",
            "year": 1994,
            "price": 97613
          }
        },
        {
          "id": 999,
          "make": "Toyota",
          "model": "Corolla",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Khaki",
            "year": 2007,
            "price": 24375
          }
        },
        {
          "id": 1000,
          "make": "Buick",
          "model": "Riviera",
          "image": "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png",
          "params": {
            "color": "Khaki",
            "year": 1998,
            "price": 38161
          }
        }
      ]
    }
  },
  setup() {
    const restaurants = ref([]);
    const querySearch = (queryString, cb) => {
      let results = queryString
          ? restaurants.value.filter(createFilter(queryString))
          : restaurants.value;
      // call callback function to return suggestions
      cb(results);
    };
    const createFilter = (queryString) => {
      return (restaurant) => {
        return (
            restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) ===
            0
        );
      };
    };
    const loadAll = () => {
      return [
        {"value": "vue", "link": "https://github.com/vuejs/vue"},
        {"value": "element", "link": "https://github.com/ElemeFE/element"},
        {"value": "cooking", "link": "https://github.com/ElemeFE/cooking"},
        {"value": "mint-ui", "link": "https://github.com/ElemeFE/mint-ui"},
        {"value": "vuex", "link": "https://github.com/vuejs/vuex"},
        {"value": "vue-router", "link": "https://github.com/vuejs/vue-router"},
        {"value": "babel", "link": "https://github.com/babel/babel"}
      ];
    };
    const handleSelect = (item) => {
      console.log(item);
    };
    onMounted(() => {
      restaurants.value = loadAll();
    });
    return {
      restaurants,
      state1: ref(''),
      state2: ref(''),
      querySearch,
      createFilter,
      loadAll,
      handleSelect,
    };
  },
}
</script>

<style>

* {
  font-family: "Helvetica", sans-serif;
}

.el-row {
  margin-top: 50px;
}

.el-card {
  margin-bottom: 15px;
}

#correction .el-link {
  font-style: italic;
}

#correction, #tags, #products {
  margin-top: 15px;
}

#correction .title {
  margin-right: 10px;
  font-size: 11pt;
}

#products .box-card {
  width: 250px;
}

#products .box-card img {
  max-width: 250px;
}

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

#filters .el-card__header {
  padding: 15px 20px;
}

</style>

