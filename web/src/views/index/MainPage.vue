<template>
  <div>
  <div class="card-container">
  <el-row :span="24" :gutter="30">
    <el-col :span="6">
      <el-card shadow="never">
        <div style="display: flex;">
          <img src="../../assets/goods.png" class="card-img" alt="logo"/><br>
          <div style="display: block;">
          <span class="card-span-1">商品数</span>
          <span class="card-span-2">{{statistics.goodsCount}}</span>
            </div>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card shadow="never">
        <div style="display: flex;">
          <img src="../../assets/order.png" class="card-img" alt="logo"/><br>
          <div style="display: block;">
          <span class="card-span-1">订单量</span>
            <span class="card-span-2">{{statistics.orderCount}}</span>
          </div>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card shadow="never">
        <div style="display: flex;">
          <img src="../../assets/amount.png" class="card-img" alt="logo"/><br>
          <div style="display: block;">
          <span class="card-span-1">交易金额（元）</span>
            <span class="card-span-2">{{statistics.amount}}</span>
          </div>
        </div>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card shadow="never">
        <div style="display: flex;">
          <img src="../../assets/customer.png" class="card-img" alt="logo"/><br>
          <div style="display: block;">
          <span class="card-span-1">访客数</span>
            <span class="card-span-2">{{statistics.visitorCount}}</span>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
  </div>
  <div class="card-container1">
    <el-row :span="24">
      <el-col :span="24">
        <el-card shadow="never" style="border: none;">
          <div id="main" :style="{with: '100%', height: '360px'}"></div>
        </el-card>
      </el-col>
    </el-row>
  </div>
  </div>
</template>

<script>
import * as echarts from 'echarts';
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
    // 获取商品数、订单量、交易金额、访客数
    this.$axios.get("/statistics/info", {
      params: {
        id: localStorage.getItem("uid")
      }
    }).then(response => {
      if (response.data.code === 200){
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

    this.$axios.get("/week/info", {
      params: {
        id: localStorage.getItem("uid")
      }
    }).then(response => {
      if (response.data.code === 200){
        let res = response.data.data;
        let option = {
          tooltip: {
            trigger: 'axis',
                axisPointer: {
              type: 'cross',
                  label: {
                backgroundColor: '#6a7985'
              }
            }
          },
          legend: {
            data: ['订单量', '销售额']
          },
          toolbox: {
            feature: {
              saveAsImage: {}
            }
          },
          grid: {
            left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
          },
          xAxis: [
            {
              type: 'category',
              boundaryGap: false,
              data: ['周一', '周二', '周三', '周四', '周五', '周六', '周日']
            }
          ],
              yAxis: [
            {
              type: 'value'
            }
          ],
              series: [
            {
              name: '订单量',
              type: 'line',
              stack: 'Total',
              areaStyle: {},
              emphasis: {
                focus: 'series'
              },
              data: res.orders
            },
            {
              name: '销售额',
              type: 'line',
              stack: 'Total',
              areaStyle: {},
              emphasis: {
                focus: 'series'
              },
              data: res.amount
            }
          ]
        }

        let chartDom = document.getElementById('main');
        let myChart = echarts.init(chartDom);
        option && myChart.setOption(option);
      }
    })
  }
}
</script>

<style scoped>
.card-container{
  border: none;
  background-color: #f0f3f6;
  padding: 30px;
}
.card-container1{
  border: none;
  background-color: white;
  padding: 20px;
}
.card-img{
  width: 60px;
  height: 60px;
}
.card-span-1{
  padding-top: 3px;
  padding-left: 10px;
}
.card-span-2{
  display: block;
  font-size: 24px;
  font-weight: 500;
  padding-left: 10px;
  padding-top: 8px;
}
</style>