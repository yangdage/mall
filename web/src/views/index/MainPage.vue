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
    <div style="padding: 30px;">
      <el-row :gutter="30">
        <el-col :span="8">
          <el-card class="box-card" shadow="never">
              <el-descriptions title="今日订单" :column="1" border>
                <el-descriptions-item v-for="item in todayOrder"
                                      :key="item.title"
                                      :label="item.title">
                  <span class="span-data">{{item.data}}</span>
                </el-descriptions-item>
              </el-descriptions>
          </el-card>
        </el-col>
        <el-col :span="16">
          <el-card class="box-card" shadow="never">
            <el-descriptions title="每周订单总览" :column="1">
              <el-descriptions-item>
                <div id="main" style="width: 100%;height:300px;"></div>
              </el-descriptions-item>
            </el-descriptions>
          </el-card>
        </el-col>
      </el-row>
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
      todayOrder: [
        {
          title: '待付款订单',
          data: ''
        },
        {
          title: '待发货订单',
          data: ''
        },
        {
          title: '已发货订单',
          data: ''
        },
        {
          title: '待收货订单',
          data: ''
        },
        {
          title: '待评价订单',
          data: ''
        },
        {
          title: '已完成订单',
          data: ''
        },
      ],
    }
  },
  mounted() {
    this.getInfo()
    this.getTodayOrderInfo();

    const echarts = require('echarts/lib/echarts');
    require('echarts/lib/component/title');
    require('echarts/lib/component/toolbox');
    require('echarts/lib/component/tooltip');
    require('echarts/lib/component/grid');
    require('echarts/lib/component/legend');
    require('echarts/lib/chart/line');

    let chartDom = document.getElementById('main');
    let myChart = echarts.init(chartDom);
    let option;

    this.$axios.get("/week/order/info", {
      params: {
        id: localStorage.getItem("uid")
      }
    }).then(response => {
      if (response.data.code === 200) {
        let res = response.data.data
        option = {
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
              type: 'value',
              name: '订单量',
              min: 0,
              max: function (value) {
                return value.max * 2;
              },
              position: 'left',
            },
            {
              type: 'value',
              name: '销售额',
              min: 0,
              max: function (value) {
                return value.max * 2;
              },
              position: 'right',
            }
          ],
          series: [
            {
              name: '订单量',
              type: 'line',
              stack: 'Total',
              yAxisIndex: 0,
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
              yAxisIndex: 1,
              areaStyle: {},
              emphasis: {
                focus: 'series'
              },
              data: res.amount
            }
          ]
        };
        option && myChart.setOption(option);
      }
    })
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
    },
    getTodayOrderInfo() {
      this.$axios.get("/today/order/info", {
        params: {
          id: localStorage.getItem("uid")
        }
      }).then(response => {
        if (response.data.code === 200) {
          let res = response.data.data.data;
          for (let i = 0; i < 6; i++) {
            this.todayOrder[i].data = res[i]
          }
        }
      })
    },
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
.span-data{
  color: #409EFF;
}

.box-card {
  background-color: #FAFAFA;
  height: 420px;
  border: none;
}
</style>