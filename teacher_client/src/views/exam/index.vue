<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.publisher_teacher_id" class="filter-item" placeholder="发布者工号" style="width: 200px;"
                @keyup.enter.native="handleFilter"/>
      <el-button v-waves class="filter-item" icon="el-icon-search" type="primary" @click="handleFilter">
        搜索
      </el-button>
      <el-button v-waves class="filter-item" icon="el-icon-edit" type="primary"
                 @click="handleCreate">
        添加
      </el-button>
      <el-tooltip class="item" content="勾选表格左侧以多选，删除所有选中项" effect="dark" placement="top-start">
        <el-button v-waves class="filter-item" icon="el-icon-delete" type="danger"
                   @click="handleMultiDelete">
          删除
        </el-button>
      </el-tooltip>
    </div>

    <el-table
      :key="tableKey"
      ref="examTable"
      v-loading="listLoading"
      :data="list"
      border
      fit
      show-header
      stripe
      style="width: 100%;"
    >

      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column align="center" label="ID" prop="id" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="详情" type="expand">
        <template slot-scope="{row}">
          <el-form>
            <el-form-item v-for="q of
            [['单选题',row.mcq_score,row.mcq_num],['多选题',row.maq_score,row.maq_num],['填空题',row.bfq_score,row.bfq_num],
            ['判断题',row.tfq_score,row.tfq_num],['代码阅读题',row.crq_score,row.crq_num],['编程题',row.cq_score,row.cq_num]]"
                          :label="q[0]"><span>
              {{ '每题 ' + q[1] + ' 分，共 ' + q[2] + ' 题，合计 ' + q[1] * q[2] + ' 分' }}</span></el-form-item>
          </el-form>
          <router-link v-if="checkStatus(row)==='已结束'" :to="'/example/edit/'+row.id">
            <el-button icon="el-icon-info" size="small" type="success">
              查看考生作答情况
            </el-button>
          </router-link>
          <span v-else-if="checkStatus(row)==='进行中'" class="link-type" @click="handleGetStudentList(row.id)">
            查看考生名单
          </span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="发布者工号">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.publisher_teacher_id }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="开始时刻">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.start_time | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="结束时刻">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.start_time | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="状态" width="100">
        <template slot-scope="{row}">
          <el-tag :type="new Map([['已结束', 'info'], ['进行中', 'success'], ['未开始', 'primary']]).get(checkStatus(row))">
            {{ checkStatus(row) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="230">
        <template slot-scope="{row}">
          <el-button :disabled="checkStatus(row)==='已结束'" size="mini" type="primary" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="handleDelete(row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :limit.sync="listQuery.page_size" :page.sync="listQuery.page_index" :total="total"
                align="center" @pagination="getList"/>

    <el-dialog :close-on-click-modal="false" :title="textMap[dialogStatus]"
               :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="dataForm" :model="temp" :rules="rules" label-position="left" label-width="125px"
               style="margin-left:50px;">
        <el-form-item label="开始时刻" prop="start_time">
          <el-date-picker v-model="temp.start_time" align="center" placeholder="选择日期与时刻" type="datetime">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="结束时刻" prop="end_time">
          <el-date-picker v-model="temp.end_time" align="center" placeholder="选择日期与时刻" type="datetime">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="考试限时（分钟）" prop="time_allowed">
          <el-input-number v-model="temp.time_allowed" :max="180" :min="10" :step="10"></el-input-number>
        </el-form-item>
        <el-form-item label="单选题分数" prop="mcq_score">
          <el-input-number v-model="temp.mcq_score" :max="6" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="单选题数目" prop="mcq_num">
          <el-input-number v-model="temp.mcq_num" :max="50" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="多选题分数" prop="maq_score">
          <el-input-number v-model="temp.maq_score" :max="6" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="多选题数目" prop="maq_num">
          <el-input-number v-model="temp.maq_num" :max="30" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="填空题分数" prop="bfq_score">
          <el-tooltip content="填空题固定为每小题 3 分" effect="dark" placement="top-start">
            <el-input-number v-model="temp.bfq_score" :max="3" :min="3" disabled></el-input-number>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="填空题数目" prop="bfq_num">
          <el-input-number v-model="temp.bfq_num" :max="30" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="判断题分数" prop="tfq_score">
          <el-input-number v-model="temp.tfq_score" :max="5" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="判断题数目" prop="tfq_num">
          <el-input-number v-model="temp.tfq_num" :max="30" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="代码阅读题分数" prop="crq_score">
          <el-tooltip content="代码阅读题固定为每小题 6 分" effect="dark" placement="top-start">
            <el-input-number v-model="temp.bfq_score" :max="6" :min="6" disabled></el-input-number>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="代码阅读题数目" prop="crq_num">
          <el-input-number v-model="temp.crq_num" :max="5" :min="1"></el-input-number>
        </el-form-item>
        <el-form-item label="编程题分数" prop="cq_score">
          <el-input-number v-model="temp.cq_score" :max="20" :min="3"></el-input-number>
        </el-form-item>
        <el-form-item label="编程题数目" prop="cq_num">
          <el-input-number v-model="temp.cq_num" :max="3" :min="1"></el-input-number>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          确定
        </el-button>
      </div>
    </el-dialog>

    <!--    <el-dialog :close-on-click-modal="false" :title="'确认删除以下 '+rowsToBeDeleted.length+' 条记录？'"-->
    <!--               :visible.sync="dialogDeleteVisible">-->
    <!--      <el-table :data="rowsToBeDeleted" max-height="800">-->
    <!--        <el-table-column align="center" label="ID" property="id" width="150"></el-table-column>-->
    <!--        <el-table-column align="center" label="发布者工号" property="publisher_teacher_id" width="150"></el-table-column>-->
    <!--        <el-table-column align="center" label="开始时刻" property="stem" width="200"></el-table-column>-->
    <!--      </el-table>-->
    <!--      <div slot="footer" class="dialog-footer">-->
    <!--        <el-button @click="dialogDeleteVisible = false">-->
    <!--          取消-->
    <!--        </el-button>-->
    <!--        <el-button type="danger" @click="deleteData">-->
    <!--          确定-->
    <!--        </el-button>-->
    <!--      </div>-->
    <!--    </el-dialog>-->
  </div>
</template>

<script>
import {createExams, deleteExams, getExams, updateExam} from '@/api/exam'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination'
import MarkdownEditor from "@/components/MarkdownEditor";
import _ from "lodash"
import {parseTime} from "@/utils/time";

export default {
  name: 'ExamList',
  components: {MarkdownEditor, Pagination},
  directives: {waves},
  filters: {
    parseTime
  },
  data() {
    return {
      tableKey: 0,

      list: null,
      total: 0,

      listLoading: true,
      listQuery: {
        publisher_teacher_id: '',
        page_index: 1,
        page_size: 20
      },

      temp: {
        id: undefined,
        start_time: new Date(),
        end_time: new Date(),// TODO
        time_allowed: 120,
        mcq_score: 2,
        mcq_num: 20,
        maq_score: 3,
        maq_num: 5,
        bfq_score: 3,
        bfq_num: 5,
        tfq_score: 2,
        tfq_num: 5,
        crq_score: 6,
        crq_num: 2,
        cq_score: 8,
        cq_num: 1
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },

      rowsToBeDeleted: [],
      dialogDeleteVisible: false,

      rules: {
        stem: [{required: true, message: '必须填写开始时刻', trigger: 'change'},
          {max: 200, message: '不得超过 200 个字符', trigger: 'change'}],
        right_answer: [{required: true, message: '必须给出正确答案', trigger: 'change'}]
      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    checkStatus(exam) {
      if (new Date(exam.end_time) <= new Date()) {
        return '已结束'
      } else if (new Date() < new Date(exam.start_time)) {
        return '未开始'
      } else {
        return '进行中'
      }
    },
    getList() {
      this.listLoading = true
      getExams(this.listQuery).then(body => {
        this.list = body.data
        this.total = body.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page_index = 1
      this.getList()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        start_time: new Date(),
        end_time: new Date(),
        time_allowed: 120,
        mcq_score: 2,
        mcq_num: 20,
        maq_score: 3,
        maq_num: 5,
        bfq_score: 3,
        bfq_num: 5,
        tfq_score: 2,
        tfq_num: 5,
        crq_score: 6,
        crq_num: 2,
        cq_score: 8,
        cq_num: 1
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          // TODO: RFC3339
          createExams([this.temp]).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: '创建成功',
              showClose: true,
              type: 'success'
            })
            this.getList()
          })
        }
      })
    },
    handleUpdate(row) {
      _.merge(this.temp, row)
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateExam('exam', tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.temp.id)
            this.list.splice(index, 1, this.temp)
            this.dialogFormVisible = false
            this.$message({
              message: '修改已保存',
              showClose: true,
              type: 'success'
            })
          })
        }
      })
    },
    handleMultiDelete() {
      let rows = this.$refs.examTable.selection;
      if (rows.length === 0) {
        this.$message({
          message: '没有任何一项被选中，勾选表格左侧以多选',
          showClose: true,
          type: 'warning'
        })
      } else {
        this.rowsToBeDeleted = Object.assign([], rows)
        this.dialogDeleteVisible = true
      }
    },
    handleDelete(row) {
      this.rowsToBeDeleted = []
      this.rowsToBeDeleted[0] = Object.assign({}, row)
      this.dialogDeleteVisible = true
    },
    deleteData() {
      console.log(this.rowsToBeDeleted)
      deleteExams('exam', this.rowsToBeDeleted.map(v => v.id)).then(() => {
        this.dialogDeleteVisible = false
        this.$message({
          message: '删除成功',
          showClose: true,
          type: 'success'
        })
        this.getList()
      })
    },
    handleGetStudentList(examId) {

    }
  }
}
</script>
