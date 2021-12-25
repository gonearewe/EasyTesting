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

    <el-form>
      <el-form-item label="要查看的考试（ID）">
        <el-checkbox-group v-model="selectedExamIds" @change="updateListAndChart">
          <el-tooltip v-for="exam in allExams" effect="dark" placement="top-start">
            <div slot="content">
              发布者工号：{{ exam.publisher_teacher_id }}<br/>
              开始时刻：{{ exam.start_time | parseTime }}<br/>
              结束时刻：{{ exam.end_time | parseTime }}<br/>
            </div>
            <el-checkbox :label="exam.id" border></el-checkbox>
          </el-tooltip>
        </el-checkbox-group>
      </el-form-item>
    </el-form>

    <el-form :inline="true" :model="statistics">
      <el-form-item label="平均分">
        <span>{{ statistics.average }}</span>
      </el-form-item>
      <el-form-item label="最高分">
        <span>{{ statistics.max }}</span>
      </el-form-item>
      <el-form-item label="最低分">
        <span>{{ statistics.min }}</span>
      </el-form-item>
    </el-form>

    <el-tabs @tab-click="refreshChart">
      <el-tab-pane label="表格">
        <el-table
          ref="examineeTable"
          v-loading="listLoading"
          :data="list"
          border
          fit
          max-height="1000"
          show-header
          stripe
          style="width: 100%;"
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
                {{ row.score * 1.0 / 10 }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>

      <!--      use `lazy` to avoid echarts displaying in a small size-->
      <el-tab-pane label="图示" lazy>
        <div class="chart-container">
          <chart ref="chart" height="100%" width="100%"/>
        </div>
      </el-tab-pane>
    </el-tabs>

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
      selectedExamIds: [],
      allExams: [],
      examinees: {},
      chartData: {},

      statistics: {
        average: 0,
        max: 0,
        min: 0
      },

      list: null,
      listQuery: {
        student_id: undefined,
        student_name: ''
      },
      listLoading: false,

      downloadLoading: false,

      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },

      rowsToBeDeleted: [],
      dialogDeleteVisible: false,

      dialogExamineeVisible: false,

    }
  },
  created() {
    getEndedExams().then(body => {
      this.allExams = body
    })
  },
  watch: {
    list(newList, oldList) {
      this.statistics.average = newList.reduce((p, c) => p + c.score, 0) / (10 * newList.length)
      this.statistics.max = Math.max(...newList.map(e => e.score)) / 10
      this.statistics.min = Math.min(...newList.map(e => e.score)) / 10
    },
    chartData(newData, oldData) {
      this.refreshChart(newData)
    }
  },
  methods: {
    getScoreColor(score) {
      if (score < 600) {
        return 'danger'
      } else if (score < 800) {
        return 'warning'
      } else if (score < 900) {
        return 'info'
      } else {
        return 'success'
      }
    },
    updateListAndChart() {
      this.listLoading = true
      this.list = []
      this.chartData = {}
      this.selectedExamIds.forEach(exam_id => {
        if (this.examinees.hasOwnProperty(exam_id)) {
          this.list.push(...this.examinees[exam_id])
          this.chartData[exam_id] = this.examinees[exam_id].map(e => e.score / 10)
        } else {
          getExaminees({"exam_id": exam_id}).then(body => {
              this.examinees[exam_id] = body
              this.list.push(...body)
              this.chartData[exam_id] = body.map(e => e.score / 10)
            }
          )
        }
      })
      this.$nextTick(() => {
        this.listLoading = false
      })
    },
    refreshChart(data) {
      if (this.$refs.chart) {
        let list = []
        let data = data || this.chartData
        // console.log(data,data.entries)
        for (const [key, value] of Object.entries(data)) {
          // map everybody's scores into buckets at the interval of 10 points
          let buckets = new Array(10).fill(0)
          for (const score of value) {
            if (score === 100) {
              buckets[9] += 1
            } else {
              buckets[Math.floor(score / 10)] += 1
            }
          }
          list.push({name: 'ID-' + key, data: buckets})
        }
        this.$refs.chart.setData(list)
      }
    },
    handleDownload() {
      // this.downloadLoading = true
      // import('@/utils/Export2Excel').then(excel => {
      //   const tHeader = ['student_id', 'name', 'class_id']
      //   let queryAll = Object.assign({}, this.listQuery, {page_index: 1, page_size: Math.pow(2, 32)})
      //   getStudents(queryAll).then(body => {
      //     excel.export_json_to_excel({
      //       header: tHeader,
      //       data: body.data.map(student => tHeader.map(k => student[k])),
      //       filename: 'students'
      //     })
      //     this.downloadLoading = false
      //   })
      // })
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
