<template>
  <div>
      <el-card shadow="never" class="card-box">
        <el-form ref="query" :model="query" label-width="80px">
          <el-row :gutter="12">
            <el-col :span="8">
              <el-form-item label="订单ID" prop="orderId">
                <el-input v-model="query.orderId" size="small"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="订单状态" prop="orderStatus">
                <el-select v-model="query.orderStatus"
                           placeholder="请选择" size="small">
                  <el-option
                      v-for="item in options"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value">
                  </el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item style="margin-left: 10px;">
                <el-button @click="resetForm('query')" size="small">重置</el-button>
                <el-button type="primary" @click="queryOrder" size="small">查询</el-button>
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
      <el-card shadow="never" style="border-top: none;">
      <el-table
          :highlight-current-row="true"
          :data="tableData"
          border
          style="width: 100%;">
        <el-table-column
            type="selection"
            width="55">
        </el-table-column>
        <el-table-column
            prop="id"
            label="订单编号"
            width="180">
        </el-table-column>
        <el-table-column
            prop="created"
            label="提交时间"
            width="200">
          <template #default="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.created }}</span>
          </template>
        </el-table-column>
        <el-table-column
            prop="userName"
            label="用户账号"
            width="100">
        </el-table-column>
        <el-table-column
            prop="totalPrice"
            label="总计"
            width="100">
          <template #default="scope">
            <span style="margin-left: 5px;font-weight: 500;">¥ {{ scope.row.totalPrice }}</span>
          </template>
        </el-table-column>
        <el-table-column
            prop="status"
            label="订单状态"
            width="100">
          <template #default="scope">
            <el-tag v-if="scope.row.status === '待付款'" size="mini" type="danger">待付款</el-tag>
            <el-tag v-if="scope.row.status === '待发货'" size="mini" type="success">待发货</el-tag>
            <el-tag v-if="scope.row.status === '已发货'" size="mini" type="primary">已发货</el-tag>
            <el-tag v-if="scope.row.status === '待收货'" size="mini" type="primary">待收货</el-tag>
            <el-tag v-if="scope.row.status === '待评价'" size="mini" type="primary">待评价</el-tag>
            <el-tag v-if="scope.row.status === '已完成'" size="mini" type="success">已完成</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="300">
          <template #default="scope">
            <el-button
                size="mini"
                @click="checkOrder(scope.$index, scope.row)">查看
            </el-button>
            <el-button
                size="mini"
                v-if="scope.row.status === '待发货'"
                type="primary"
                @click="clickDelivery(scope.$index, scope.row)">发货
            </el-button>
            <el-button
                size="mini"
                v-if="scope.row.status === '待付款' ||
                scope.row.status === '已发货' ||
                scope.row.status === '待收货' ||
                scope.row.status === '待评价'"
                type="primary"
                @click="remindUser(scope.$index, scope.row)">提醒
            </el-button>
            <el-button
                size="mini"
                v-if="scope.row.status === '已完成'"
                type="danger"
                @click="deleteOrder(scope.$index, scope.row)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table><br><br>
      <el-pagination
          background
          @current-change="handleCurrentChange"
          @prev-click="handleCurrentChangePrev"
          @next-click="handleCurrentChangeNext"
          :currentPage="currentPage"
          :page-size="size"
          layout="total, prev, pager, next"
          :total="total">
      </el-pagination>
      </el-card>
    <!-- 订单发货弹出框 -->
    <el-dialog title="订单发货" width="30%" v-model="dialogFormVisible" center>
      <el-form :model="order">
        <el-form-item label="物流名称" label-width="80px">
          <el-input v-model="order.courierName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="物流单号" label-width="80px">
          <el-input v-model="order.shipmentNumber" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogFormVisible = false">取 消</el-button>
          <el-button type="primary" @click="orderDelivery">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import { ElNotification } from 'element-plus'
export default {
  name: "OrderList",
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
          value: '待付款',
          label: '待付款',
        },
        {
          value: '待发货',
          label: '待发货',
        },
        {
          value: '已发货',
          label: '已发货',
        },
        {
          value: '待收货',
          label: '待收货',
        },
        {
          value: '待评价',
          label: '待评价',
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
    handleCurrentChangePrev(val) {
      this.currentPage = val;
      console.log(`上一页: ${val}`);
    },
    handleCurrentChange(val) {
      this.currentPage = val;
      this.queryOrder();
      console.log(`当前页: ${val}`);
    },
    handleCurrentChangeNext(val) {
      this.currentPage = val;
      console.log(`下一页: ${val}`);
    },
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
    remindUser() {
      ElNotification.success({
        title: '提醒成功',
        message: '已向用户发送当前订单状态',
      })
    },
    deleteOrder(index, row) {
      console.log(index)
      this.$axios.delete('/order/delete', {
        params: {
          id: row.id,
        }
      }).then((response) => {
         if (response.data.code === 200){
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
        if (response.data.code === 200){
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
  border: none;
  border-radius: 6px;
}
</style>