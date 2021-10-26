<template>
  <el-row>
    <el-col :span="16" :offset="4"><br><br><br>
      <el-card shadow="never" class="card-box">
        <el-form ref="query" :model="orderSet"><br>
          <el-form-item label="订单付款超过" label-width="160px">
            <el-select v-model.number="orderSet.paymentOvertime" placeholder="请选择">
              <el-option label="12小时" value="12">12小时</el-option>
              <el-option label="24小时" value="24">24小时</el-option>
              <el-option label="48小时" value="48">48小时</el-option>
              <template #append>.com</template>
            </el-select>
            <span>（ 未付款，订单自动关闭 ）</span>
          </el-form-item>
          <el-form-item label="订单收货超过" label-width="160px">
            <el-select v-model.number="orderSet.receiveOvertime" placeholder="请选择">
              <el-option label="12小时" value="12">12小时</el-option>
              <el-option label="24小时" value="24">24小时</el-option>
              <el-option label="48小时" value="48">48小时</el-option>
            </el-select>
            <span>（ 未确认，订单自动确认收货 ）</span>
          </el-form-item>
          <el-form-item label="订单完成超过" label-width="160px">
            <el-select v-model.number="orderSet.finishOvertime" placeholder="请选择">
              <el-option label="7天" value="7">7天</el-option>
              <el-option label="14天" value="14">14天</el-option>
              <el-option label="21天" value="21">21天</el-option>
            </el-select>
            <span>（ 未评价，订单自动五星好评 ）</span>
          </el-form-item>
        </el-form>
        <div style="text-align: center;margin: 50px;">
          <el-button type="primary" @click="resetForm">重置</el-button>
          <el-button type="primary" @click="submitForm">保存</el-button>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "OrderSet",
  data() {
    return {
      orderSet: {
        id: '',
        paymentOvertime: '',
        receiveOvertime: '',
        finishOvertime: ''
      }
    }
  },
  mounted() {
    this.showInfo();
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    showInfo() {
      this.$axios.get('/order/set/info', {
        params: {
          id: parseInt(localStorage.getItem("uid"))
        }
      }).then((response) => {
        this.orderSet.id = response.data.data.id;
        this.orderSet.paymentOvertime = response.data.data.paymentOvertime;
        this.orderSet.receiveOvertime = response.data.data.receiveOvertime;
        this.orderSet.finishOvertime = response.data.data.finishOvertime;
      }).catch((error) => {
        console.log(error);
      })
    },
    submitForm() {
      this.$axios.post('/order/set/save', {
        id: this.orderSet.id,
        paymentOvertime: this.orderSet.paymentOvertime,
        receiveOvertime: this.orderSet.receiveOvertime,
        finishOvertime: this.orderSet.finishOvertime,
        adminId: parseInt(localStorage.getItem("uid"))
      }).then().catch()
    },
  }
}
</script>

<style scoped>
.card-box {
  background-color: #FAFAFA;
  margin: 18px;
  border: none;
  border-radius: 6px;
}
</style>