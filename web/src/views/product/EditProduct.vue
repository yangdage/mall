<template>
  <el-row>
    <el-col :span="18" :offset="3"><br>
      <el-form ref="product"
               :model="product"
               :rules="rules"
               label-width="100px">
        <el-descriptions direction="vertical" :column="4" border>
          <el-descriptions-item label="基本信息"><br>
            <el-form-item label="类目" prop="categoryId">
              <el-cascader v-model="product.categoryId"
                           :options="categoryOptions"
                           placeholder="请选择"
                           clearable/>
            </el-form-item>
            <el-form-item label="类型" prop="kind">
              <el-radio v-model.number="product.kind" :label="1">全新</el-radio>
              <el-radio v-model.number="product.kind" :label="2">二手</el-radio>
            </el-form-item>
            <el-form-item label="标题" prop="title">
              <el-input v-model="product.title"
                        maxlength="30"
                        style="width: 80%;"
                        show-word-limit></el-input>
            </el-form-item>
            <el-form-item label="品牌" label-width="100px" prop="brandId">
              <el-select v-model.number="product.brandId"
                         placeholder="请选择">
                <el-option
                    v-for="item in brandOptions"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="名称" prop="name">
              <el-input v-model="product.name"
                        maxlength="20"
                        style="width: 60%;"
                        show-word-limit></el-input>
            </el-form-item>
            <el-form-item label="售价" label-width="100px" prop="price">
              <el-input v-model.number="product.price" style="width: 30%;">
                <template #append>元</template>
              </el-input>
            </el-form-item>
            <el-form-item label="库存" label-width="100px" prop="amount">
              <el-input v-model.number="product.amount" style="width: 30%;">
                <template #append>件</template>
              </el-input>
            </el-form-item>
          </el-descriptions-item>
        </el-descriptions>
        <el-descriptions direction="vertical" :column="4" border>
          <el-descriptions-item label="图文描述"><br>
            <el-form-item label="商品主图" prop="imageUrl">
              <el-upload
                  action="http://localhost:8000/upload"
                  :headers="{'token': token}"
                  list-type="picture-card"
                  limit="1"
                  :on-success="handleAlbumSuccess"
                  name="file">
                <i class="el-icon-plus"></i>
                <el-dialog v-model="dialogVisible">
                  <el-image
                      style="width: 100%; height: 100%"
                      :src="product.imageUrl"></el-image>
                  <img width="100%" :src="product.imageUrl" alt="">
                </el-dialog>
              </el-upload>
            </el-form-item>
            <br>
            <el-form-item label="提示">
              <el-alert
                  title="只能上传1张商品图片，图片大小不能超过3MB，支持任意格式的图片"
                  type="info"
                  show-icon>
              </el-alert>
            </el-form-item>
          </el-descriptions-item>
        </el-descriptions>
        <el-descriptions direction="vertical" :column="4" border>
          <el-descriptions-item label="物流信息"><br>
            <el-form-item label="发货地址" prop="sendAddress">
              <el-input v-model="product.sendAddress" style="width: 60%;"></el-input>
            </el-form-item>
            <el-form-item label="快递类型" prop="parcelType">
              <el-input v-model="product.parcelType" style="width: 60%;"></el-input>
            </el-form-item>
          </el-descriptions-item>
        </el-descriptions>
        <el-descriptions direction="vertical" :column="4" border>
          <el-descriptions-item label="售后服务"><br>
            <el-form-item label="选择服务" prop="salesService">
              <el-checkbox-group v-model="product.salesService">
                <el-checkbox label="提供发票" name="type"></el-checkbox>
                <el-checkbox label="保修服务" name="type"></el-checkbox>
                <el-checkbox label="退换货承诺" name="type"></el-checkbox>
                <el-checkbox label="服务承诺：该类商品，必须支持【七天退货】服务" name="type"></el-checkbox>
              </el-checkbox-group>
            </el-form-item>
          </el-descriptions-item>
        </el-descriptions>
        <el-form-item label-width="200px"><br>
          <el-button type="primary" @click="submitForm('product')">保存草稿</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "EditProduct",
  data() {
    return {
      product: {
        categoryId: '',
        kind: '',
        title: '',
        brandId: '',
        name: '',
        price: '',
        amount: '',
        imageUrl: '',
        sendAddress: '',
        parcelType: '',
        salesService: [],
      },
      rules: {
        categoryId: [
          {
            required: true,
            message: '请选择一个类目',
            trigger: 'change',
          },
        ],
        kind: [
          {
            required: true,
            message: '请选择一种类型',
            trigger: 'change',
          },
        ],
        title: [
          {
            required: true,
            message: '请输入一个标题',
            trigger: 'blur'
          },
          {
            min: 1,
            max: 30,
            message: '标题长度不能超过30个字',
            trigger: 'blur'
          }
        ],
        brandId: [
          {
            required: true,
            message: '请选择一个品牌',
            trigger: 'change'
          }
        ],
        price: [
          {
            required: true,
            message: '价格不能为空'
          },
          {
            type: 'number',
            message: '价格必须为数字'
          }
        ],
        amount: [
          {
            required: true,
            message: '库存不能为空'
          },
          {
            type: 'number',
            message: '库存必须为数字'
          }
        ],
        imageUrl: [
          {required: true,
            message: '请选择一张图片',
            trigger: 'blur'}
        ],
        sendAddress: [
          {
            required: true,
            message: '请选择一张图片',
            trigger: 'blur'
          }
        ],
        parcelType: [
          {
            required: true,
            message: '请选择一张图片',
            trigger: 'blur'
          }
        ],
        salesService: [
          {
            type: 'array',
            required: true,
            message: '至少选择一个服务',
            trigger: 'change',
          },
        ],
      },
      categoryOptions: null,
      brandOptions: null,
      dialogVisible: false,
      token: '',
    }
  },
  mounted() {
    this.categoryOption();
    this.brandOption();
    this.getUpdateInfo();
    this.token = localStorage.getItem('token')
  },
  methods: {

    // 获取类目选项
    categoryOption() {
      this.$axios.get('/category/option').then((response) => {
        this.categoryOptions = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },

    // 获取品牌选项
    brandOption() {
      this.$axios.get('/brand/option').then((response) => {
        this.brandOptions = response.data.data;
      }).catch((error) => {
        console.log(error)
      })
    },

    // 图片响应结果处理
    handleAlbumSuccess(response) {
      if (response.code === 200) {
        this.product.imageUrl = response.data;
      }
      console.log(response.data);
    },

    // 重置表单
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },

    getUpdateInfo() {
      this.$axios.get('/product/info', {
        params: {
          id: this.$route.params.id
        }
      }).then(response => {
        let res = response.data.data;
        this.product.categoryId = res.categoryId;
        this.product.kind = res.kind;
        this.product.title = res.title;
        this.product.brandId = res.brandId;
        this.product.name = res.name;
        this.product.price = res.price;
        this.product.amount = res.amount;
        this.product.imageUrl = res.imageUrl;
        this.product.sendAddress = res.sendAddress;
        this.product.parcelType = res.parcelType;
        this.product.salesService = res.salesService.split(",");
      })
    },

    // 表单提交（保存或发布商品）
    submitForm(formName, status) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let salesServiceStr = '';
          for (let i = 0; i < this.product.salesService.length; i++) {
            salesServiceStr = salesServiceStr + this.product.salesService[i] + ',';
          }
          if (status === 0) {
            this.$store.commit('setPageTitle', '保存成功')
          } else {
            this.$store.commit('setPageTitle', '发布成功')
          }
          let pid = this.$route.params.id.toString()
          this.$axios.put('/product/update', {
            id: parseInt(pid),
            categoryId: this.product.categoryId,
            kind: this.product.kind,
            title: this.product.title,
            brandId: this.product.brandId,
            name: this.product.name,
            price: this.product.price,
            amount: this.product.amount,
            imageUrl: this.product.imageUrl,
            sendAddress: this.product.sendAddress,
            parcelType: this.product.parcelType,
            salesService: salesServiceStr,
            creatorId: parseInt(localStorage.getItem("uid")),
            status: 1
          }).then((response) => {
            if (response.data.code === 200) {
              this.$router.push({name: 'resultPage'})
            }
          }).catch((error) => {
            console.log(error)
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style scoped>

</style>