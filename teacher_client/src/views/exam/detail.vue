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
    <div>
      <el-checkbox-group v-model="checkboxGroup1">
        <el-checkbox-button v-for="city in cities" :key="city" :label="city">{{ city }}</el-checkbox-button>
      </el-checkbox-group>
    </div>

    <el-table
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
          <router-link v-if="checkStatus(row,currentDatetime)==='已结束'" :to="'/example/edit/'+row.id">
            <el-button icon="el-icon-info" size="small" type="success">
              查看考生作答情况
            </el-button>
          </router-link>
          <span v-else-if="checkStatus(row,currentDatetime)==='进行中'" class="link-type"
                @click="handleGetExamineeList(row.id)">
            查看考生名单
          </span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="发布者工号">
        <template slot-scope="{row}">
          <span :class="checkStatus(row,currentDatetime)==='未开始'?'link-type':''"
                @click="checkStatus(row,currentDatetime)==='未开始'?handleUpdate(row):undefined">
            {{ row.publisher_teacher_id }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="开始时刻">
        <template slot-scope="{row}">
          <span :class="checkStatus(row,currentDatetime)==='未开始'?'link-type':''"
                @click="checkStatus(row,currentDatetime)==='未开始'?handleUpdate(row):undefined">
          {{ row.start_time | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="结束时刻">
        <template slot-scope="{row}">
          <span :class="checkStatus(row,currentDatetime)==='未开始'?'link-type':''"
                @click="checkStatus(row,currentDatetime)==='未开始'?handleUpdate(row):undefined">
          {{ row.end_time | parseTime }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="状态" width="100">
        <template slot-scope="{row}">
          <el-tag
            :type="new Map([['已结束', 'info'], ['进行中', 'success'], ['未开始', 'primary']]).get(checkStatus(row,currentDatetime))">
            {{ checkStatus(row, currentDatetime) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="230">
        <template slot-scope="{row}">
          <el-button :disabled="checkStatus(row,currentDatetime)!=='未开始'" size="mini" type="primary"
                     @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button :disabled="checkStatus(row,currentDatetime)==='进行中'" size="mini" type="danger"
                     @click="handleDelete(row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {createExams, deleteExams, getExaminees, getExams, updateExam} from '@/api/exam'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination'
import MarkdownEditor from "@/components/MarkdownEditor";
import _ from "lodash"
import {parseTime} from "@/utils/time";

export default {
  name: 'ResultList',
  components: {MarkdownEditor, Pagination},
  directives: {waves},
  filters: {
    parseTime
  },
  computed: {},
  data() {
    return {
      selectedExams: [],
      allExams: [],

      list: null,

      listLoading: true,
      listQuery: {
        publisher_teacher_id: '',
        page_index: 1,
        page_size: 20
      },

      temp: {
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
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },

      rowsToBeDeleted: [],
      dialogDeleteVisible: false,

      examineeList: [],
      dialogExamineeVisible: false,

      rules: {
        start_time: [{required: true, message: '必须填写开始时刻', trigger: 'change'},
          {validator: validateStartTime, trigger: 'change'}],
        end_time: [{required: true, message: '必须填写结束时刻', trigger: 'change'},
          {validator: validateEndTime, trigger: 'change'}],
        total_score: [{validator: validateTotalScore, trigger: 'change'}],
      }
    }
  },
  created() {
    this.getList()
    setInterval(() => this.currentDatetime = new Date(), 1000)
  },
  methods: {
    checkStatus(exam, currentDatetime) {
      if (new Date(exam.end_time) <= currentDatetime) {
        return '已结束'
      } else if (currentDatetime < new Date(exam.start_time)) {
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
      let st = this.currentDatetime
      st.setHours(14)
      st.setMinutes(0)
      st.setSeconds(0)
      let et = this.currentDatetime
      et.setHours(17)
      et.setMinutes(0)
      et.setSeconds(0)
      this.temp = {
        id: undefined,
        start_time: st,
        end_time: et,
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
    anyExamActive(exams) {
      return exams.some(exam => this.checkStatus(exam, this.currentDatetime) === '进行中')
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
      } else if (this.anyExamActive(rows)) {
        this.$message({
          message: '其中有考试正在进行中，无法删除',
          showClose: true,
          type: 'error'
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
      deleteExams(this.rowsToBeDeleted.map(v => v.id)).then(() => {
        this.dialogDeleteVisible = false
        this.$message({
          message: '删除成功',
          showClose: true,
          type: 'success'
        })
        this.getList()
      })
    },
    handleGetExamineeList(examId) {
      getExaminees({"exam_id": examId}).then(body => {
        this.examineeList = body
        this.dialogExamineeVisible = true
      })
    }
  }
}
</script>
