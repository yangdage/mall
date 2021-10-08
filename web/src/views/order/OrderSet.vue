<template>
  <el-row>
    <el-col :span="12" :offset="6"><br><br><br>
      <el-card shadow="never" class="card-box">
        <el-form ref="query" :model="query" label-width="80px">
          <el-form-item label="订单超过" prop="orderId">
            <el-input v-model="input2">
              <template #append>小时，未付款，订单自动关闭</template>
            </el-input>
          </el-form-item>
          <el-form-item label="发货超过" prop="orderId">
            <el-input v-model="input2">
              <template #append>小时，未收货，订单自动完成</template>
            </el-input>
          </el-form-item>
          <el-form-item label="评价超过" prop="orderId">
            <el-input v-model="input2">
              <template #append>小时，未评价，订单自动评价</template>
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button @click="resetForm('query')" size="small">重置</el-button>
            <el-button type="primary" @click="queryOrder" size="small">提交</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "OrderSet",
  data() {
    return {
      query: {
        orderId: '',
        orderStatus: '',
      },
      order: {
        id: '',
        courierName: '',
        shipmentNumber: '',
      },
      dialogFormVisible: false,
      tableData: null,
      currentPage: 1,
      size: 10,
      total: '',
      options: [
        {
          value: '未支付',
          label: '未支付',
        },
        {
          value: '已支付',
          label: '已支付',
        },
        {
          value: '已发货',
          label: '已发货',
        },
        {
          value: '已完成',
          label: '已完成',
        }
      ],
    }
  },
  mounted() {
    this.queryOrder();
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields();
      this.queryOrder();
    },
    queryOrder() {
      this.$axios.get('/order/list', {
        params: {
          id: this.query.orderId,
          status: this.query.orderStatus,
          pageNum: this.currentPage,
          pageSize: this.size,
          adminId: localStorage.getItem("uid")
        }
      }).then((response) => {
        this.total = response.data.data.total;
        this.tableData = response.data.data.list;
      }).catch((error) => {
        console.log(error);
      })
    },
    checkOrder(index, row) {
      console.log(index)
      this.$router.push({
        name: 'orderDetail',
        params: {
          id: row.id,
          totalPrice: row.totalPrice
        }
      });
    },
    deleteOrder(index, row) {
      console.log(index)
      this.$axios.get('/order/delete', {
        params: {
          orderId: row.id,
        }
      }).then((response) => {
        if (response.data.code === 200) {
          this.$message({
            type: 'success',
            message: '删除成功!'
          });
          this.queryOrder();
        }
      }).catch((error) => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
        console.log(error);
      })

    },
    clickDelivery(index, row) {
      this.dialogFormVisible = true;
      this.order.id = row.id;
    },
    orderDelivery() {
      this.$axios.put('/order/update', {
        courierName: this.order.courierName,
        shipmentNumber: this.order.shipmentNumber,
        status: '已发货',
        id: this.order.id,
        adminId: parseInt(localStorage.getItem("uid"))
      }).then((response) => {
        this.dialogFormVisible = false;
        if (response.data.code === 200) {
          this.queryOrder();
        }
      }).catch((error) => {
        console.log(error);
      })
    }
  }
}
</script>

<style scoped>
.card-box {
  background-color: #F2F4F7;
  margin: 18px;
  border-radius: 6px;
}
</style>