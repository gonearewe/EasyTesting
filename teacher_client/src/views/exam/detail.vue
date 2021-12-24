<template>
  <div class="app-container">
    <div class="filter-container">
      <el-tooltip class="item" content="导出满足当前筛选条件的所有数据至一个 Excel(*.xlsx) 文件" effect="dark" placement="top-start">
        <el-button v-waves :loading="downloadLoading" class="filter-item" icon="el-icon-download" type="success"
                   @click="handleDownload">
          导出
        </el-button>
      </el-tooltip>
    </div>
    <div>
      <el-checkbox-group v-model="selectedExams">
        <el-checkbox-button v-for="exam in allExams" :key="exam" :label="exam.id"></el-checkbox-button>
      </el-checkbox-group>
    </div>

    <el-table
      ref="examineeTable"
      v-loading="listLoading"
      :data="list"
      border
      fit
      show-header
      stripe
      style="width: 100%;"
      max-height="1000"
    >
      <el-table-column align="center" label="学号" property="student_id"></el-table-column>
      <el-table-column align="center" label="姓名" property="student_name"></el-table-column>
      <el-table-column align="center" label="进入考试时刻">
        <template slot-scope="{row}">
          {{ row.start_time | parseTime }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="交卷时刻">
        <template slot-scope="{row}">
          {{ row.end_time | parseTime }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="得分" width="100">
        <template slot-scope="{row}">
          <el-tag
            :type="getScoreColor(row.score)">
            {{ row.score / 10 }}
          </el-tag>
        </template>
      </el-table-column>
    </el-table>

    <div class="chart-container">
      <chart height="100%" width="100%"/>
    </div>
  </div>
</template>

<script>
import {getEndedExams, getExaminees} from '@/api/exam'
import waves from '@/directive/waves' // waves directive
import Chart from '@/views/exam/Chart'
import {parseTime} from "@/utils/time";

export default {
  name: 'ResultList',
  components: {Chart},
  directives: {waves},
  filters: {
    parseTime
  },
  data() {
    return {
      selectedExams: [],
      allExams: [],

      list: null,
      listQuery: {
        student_id: undefined,
        student_name: ''
      },
      listLoading: true,

      downloadLoading: false,

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

    }
  },
  created() {
    this.getAllExams()
  },
  methods: {
    getScoreColor(score) {
      if (score < 600) {
        return 'danger'
      } else if (score < 80) {
        return 'warning'
      } else if (score < 90) {
        return 'info'
      } else {
        return 'success'
      }
    },
    getAllExams() {
      getEndedExams().then(body => {
        this.allExams = body
      })
    },
    handleFilter() {
      this.listQuery.page_index = 1
      this.getAllExams()
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

<style scoped>
.chart-container {
  position: relative;
  width: 100%;
  height: calc(100vh - 84px);
}
</style>
