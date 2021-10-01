<template>
  <div>
    <div class="card-container">
      <el-row :span="24" :gutter="30">
        <el-col :span="8">
          <el-card shadow="never">
            <div style="display: flex;">
              <img src="../../assets/goods.png" class="card-img" alt="logo"/><br>
              <div style="display: block;">
                <span class="card-span-1">商品数</span>
                <span class="card-span-2">{{ statistics.goodsCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="never">
            <div style="display: flex;">
              <img src="../../assets/order.png" class="card-img" alt="logo"/><br>
              <div style="display: block;">
                <span class="card-span-1">订单量</span>
                <span class="card-span-2">{{ statistics.orderCount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="never">
            <div style="display: flex;">
              <img src="../../assets/amount.png" class="card-img" alt="logo"/><br>
              <div style="display: block;">
                <span class="card-span-1">交易金额（元）</span>
                <span class="card-span-2">{{ statistics.amount }}</span>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div>
      <el-empty class="empty" description="暂无数据"></el-empty>
    </div>
  </div>
</template>

<script>
export default {
  name: "MainPage",
  data() {
    return {
      statistics: {
        goodsCount: '',
        orderCount: '',
        amount: '',
        visitorCount: ''
      },
    }
  },
  mounted() {
    this.getInfo()
  },
  methods: {
    getInfo() {
      // 获取商品数、订单量、交易金额、访客数
      this.$axios.get("/statistics/info", {
        params: {
          id: localStorage.getItem("uid")
        }
      }).then(response => {
        if (response.data.code === 200) {
          let res = response.data.data;
          this.statistics.goodsCount = res.goodsCount;
          this.statistics.orderCount = res.orderCount;
          this.statistics.amount = res.amount;
          this.statistics.visitorCount = res.visitorCount;
          if (res.goodsCount === 0 || res.orderCount === 0 || res.amount === 0 || res.visitorCount === 0) {
            this.statistics.goodsCount = '__';
            this.statistics.orderCount = '--';
            this.statistics.amount = '--';
            this.statistics.visitorCount = '--';
          }
        }
      })
    }
  }
}
</script>

<style scoped>
.card-container {
  border: none;
  background-color: #f0f3f6;
  padding: 30px;
}

.card-img {
  width: 60px;
  height: 60px;
}

.card-span-1 {
  padding-top: 3px;
  padding-left: 10px;
}

.card-span-2 {
  display: block;
  font-size: 24px;
  font-weight: 500;
  padding-left: 10px;
  padding-top: 8px;
}
.empty {
  margin-top: 100px;
}
</style>