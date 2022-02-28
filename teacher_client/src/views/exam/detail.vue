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
        <span>{{ statistics.average.toFixed(1) }}</span>
      </el-form-item>
      <el-form-item label="最高分">
        <span>{{ statistics.max.toFixed(1) }}</span>
      </el-form-item>
      <el-form-item label="最低分">
        <span>{{ statistics.min.toFixed(1) }}</span>
      </el-form-item>
    </el-form>

    <el-tabs>
      <!--      put chart on the first tab to prevent chart not showing data-->
      <!--      use `lazy` to avoid echarts displaying in a small size-->
      <el-tab-pane label="图示" lazy>
        <div class="chart-container">
          <chart ref="chart" height="100%" width="100%"/>
        </div>
      </el-tab-pane>

      <el-tab-pane label="表格">
        <el-table
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
    </el-tabs>

  </div>
</template>

<script>
import {getEndedExams, getExaminees} from '@/api/exam'
import waves from '@/directive/waves' // waves directive
import chart from '@/views/exam/chart'
import {parseTime} from "@/utils/time"

export default {
  name: 'ResultList',
  components: {Chart: chart},
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
      listLoading: false,

      downloadLoading: false,
    }
  },
  created() {
    getEndedExams().then(body => {
      this.allExams = body
    })

    if (this.$route.params && this.$route.params.id) {
      this.selectedExamIds.push(parseInt(this.$route.params.id))
      this.updateListAndChart()
    }
  },
  watch: {
    list(newList, oldList) {
      this.statistics.average = newList.reduce((p, c) => p + c.score, 0) / (10 * newList.length)
      this.statistics.max = Math.max(...newList.map(e => e.score)) / 10
      this.statistics.min = Math.min(...newList.map(e => e.score)) / 10
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
    async updateListAndChart() {
      this.listLoading = true
      this.list = []
      this.chartData = {}
      for (const exam_id of this.selectedExamIds) {
        if (!this.examinees.hasOwnProperty(exam_id)) {
          // lazy, so as to only load necessary Examinees
          await getExaminees({"exam_id": exam_id}).then(body => {
              this.examinees[exam_id] = body
            }
          )
        }
        this.list.push(...this.examinees[exam_id])
        this.chartData[exam_id] = this.examinees[exam_id].map(e => e.score / 10)
      }
      this.refreshChart()
      this.$nextTick(() => {
        this.listLoading = false
      })
    },
    refreshChart() {
      console.log(this.$refs)
      if (this.$refs.chart) {
        let list = []
        for (const [key, value] of Object.entries(this.chartData)) {
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
        console.log(list)
        this.$refs.chart.setData(list)
      }
    },
    handleDownload() {
      this.downloadLoading = true
      import('@/utils/Export2Excel').then(excel => {
        const tHeader = ['考试 ID', '学号', '姓名', '得分', '进入考试时刻', '交卷时刻']
        let rows = []
        for (let examId of this.selectedExamIds) {
          for (let examinee of this.examinees[examId]) {
            rows.push([examId, examinee.student_id, examinee.student_name,
              examinee.score / 10, parseTime(examinee.start_time), parseTime(examinee.end_time)])
          }
        }
        excel.export_json_to_excel({
          header: tHeader,
          data: rows,
          filename: 'results'
        })
        this.downloadLoading = false
      })
    }
  }
}
</script>

<style scoped>
.chart-container {
  margin-top: 40px;
  margin-bottom: 30px;
  position: relative;
  width: 100%;
  height: calc(100vh - 84px);
}
</style>
