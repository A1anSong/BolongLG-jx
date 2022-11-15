<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <!--      <div class="gva-btn-list">-->
      <!--        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>-->
      <!--        <el-popover v-model:visible="deleteVisible" placement="top" width="160">-->
      <!--          <p>确定要删除吗？</p>-->
      <!--          <div style="text-align: right; margin-top: 8px;">-->
      <!--            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>-->
      <!--            <el-button size="small" type="primary" @click="onDelete">确定</el-button>-->
      <!--          </div>-->
      <!--          <template #reference>-->
      <!--            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>-->
      <!--          </template>-->
      <!--        </el-popover>-->
      <!--      </div>-->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        border
        size="small"
        table-layout="auto"
        @selection-change="handleSelectionChange"
      >
        <!--        <el-table-column type="selection" width="55" />-->
        <!--        <el-table-column align="left" label="日期" width="180">-->
        <!--          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>-->
        <!--        </el-table-column>-->
        <!--        <el-table-column align="left" label="订单编号" prop="orderNo" />-->
        <el-table-column align="center" label="订单编号" prop="apply.applyNo" />
        <el-table-column align="center" label="标段名称" prop="apply.projectName" />
        <!--        <el-table-column align="left" label="申请时间" prop="apply.CreatedAt" />-->
        <el-table-column align="center" label="申请时间">
          <template #default="scope">{{ date(scope.row.apply.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="员工信息">
          <el-table-column align="center" label="工号" prop="employeeNo" />
          <el-table-column align="center" label="业务员" prop="employeeNo" />
        </el-table-column>
        <el-table-column align="center" label="申请企业" prop="apply.insureName" />
        <!--        <el-table-column align="left" label="投保方社会信用代码" prop="apply.insureCreditCode" />-->
        <!--        <el-table-column align="left" label="投保方法人" prop="apply.insureLegalName" />-->
        <!--        <el-table-column align="left" label="投保方社会信用代码" prop="apply.insureLegalIdCard" />-->
        <!--        <el-table-column align="left" label="投保方地址" prop="apply.insureAddress" />-->
        <!--        <el-table-column align="left" label="经办人" prop="apply.applicantName" />-->
        <!--        <el-table-column align="left" label="经办人身份证" prop="apply.applicantIdCard" />-->
        <!--        <el-table-column align="left" label="经办人电话" prop="apply.applicantTel" />-->
        <el-table-column align="center" label="标段编号" prop="apply.projectNo" />
        <!--        <el-table-column align="left" label="保函产品费率" prop="apply.productRate" />-->
        <!--        <el-table-column align="left" label="保函费用" prop="apply.elogAmount" />-->
        <!--        <el-table-column align="left" label="项目标识" prop="apply.projectGuid" />-->
        <el-table-column align="center" label="受益方名称" prop="apply.insuredName" />
        <!--        <el-table-column align="left" label="受益方社会信用代码" prop="apply.insuredCreditCode" />-->
        <!--        <el-table-column align="left" label="受益方地址" prop="apply.insuredAddress" />-->
        <!--        <el-table-column align="left" label="担保金额" prop="apply.tenderDeposit" />-->
        <el-table-column align="center" label="担保金额" min-width="120">
          <template #default="scope">{{ amount(scope.row.apply.tenderDeposit) }}</template>
        </el-table-column>
        <el-table-column align="center" label="所属市" prop="project" />
        <!--        <el-table-column align="left" label="保证金开始缴纳时间" prop="apply.depositStartDate" />-->
        <!--        <el-table-column align="left" label="保证金截止缴纳时间" prop="apply.depositEndDate" />-->
        <el-table-column align="center" label="开标时间" prop="apply.openBeginDate" />
        <el-table-column align="center" label="审核状态" min-width="120">
          <template #default="scope">{{ auditStatus(scope.row.apply.auditStatus) }}</template>
        </el-table-column>
        <el-table-column align="center" label="审核时间" prop="apply.openBeginDate" />
        <el-table-column align="center" label="付款信息" min-width="120">
          <template #default="scope">{{ scope.row.pay != null ? "已付款" : "未付款" }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="付款金额" prop="pay.payAmount" />-->
        <el-table-column align="center" label="付款金额" min-width="120">
          <template #default="scope">{{ scope.row.pay === null ? '' : amount(scope.row.pay.payAmount) }}</template>
        </el-table-column>
        <el-table-column align="center" label="付款时间" prop="pay.payTime" />
        <el-table-column align="center" label="开函时间" min-width="120">
          <template #default="scope">{{ scope.row.letter === null ? '' : date(scope.row.letter.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="开函状态" min-width="120">
          <template #default="scope">{{ scope.row.letter === null ? '未开函' : '已开函' }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="保函格式编号" prop="apply.elogTemplateNo" />-->
        <el-table-column align="center" label="保函格式名称" prop="apply.elogTemplateName" />
        <el-table-column align="center" label="来源" min-width="120">
          <template #default="scope">{{ productType(scope.row.apply.productType) }}</template>
        </el-table-column>
        <!--        <el-table-column align="left" label="附件信息" prop="apply.attachInfo" />-->
        <el-table-column align="center" label="操作" min-width="240" fixed="right">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              size="small"
              class="table-button"
              @click="updateOrderFunc(scope.row)"
            >变更
            </el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单编号:" prop="orderNo">
          <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TestOrder'
}
</script>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/testOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { date } from '@/utils/jxlg/date'
import { auditStatus } from '@/utils/jxlg/auditStatus'
import { productType } from '@/utils/jxlg/productType'
import { amount } from '@/utils/jxlg/amount'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  orderNo: '',
})

// 验证规则
const rule = reactive({})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteOrderFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
  multipleSelection.value.map(item => {
    ids.push(item.ID)
  })
  const res = await deleteOrderByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateOrderFunc = async(row) => {
  const res = await findOrder({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reorder
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteOrderFunc = async(row) => {
  const res = await deleteOrder({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    orderNo: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createOrder(formData.value)
        break
      case 'update':
        res = await updateOrder(formData.value)
        break
      default:
        res = await createOrder(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>
</style>
