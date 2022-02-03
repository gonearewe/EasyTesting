<template>
  <div class="app-container">
    <div class="dashboard-container">
      <el-row :gutter="10" class="panel-group">
        <el-col :lg="6" :sm="12" :xs="15" class="card-panel-col">
          <router-link to="/exam">
            <el-card class="card-panel" shadow="hover">
              <img class="image" src="@/assets/dashboard/exam.png">
              <div class="description">
                <span>管理考试信息，查看成绩分析</span>
              </div>
            </el-card>
          </router-link>
        </el-col>
        <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
          <router-link to="/student">
            <el-card class="card-panel" shadow="hover">
              <img class="image" src="@/assets/dashboard/student.png">
              <div class="description">
                <span>管理学生名单，批量地处理学生信息</span>
              </div>
            </el-card>
          </router-link>
        </el-col>
        <el-col v-permission="['admin']" :lg="6" :sm="12" :xs="12" class="card-panel-col">
          <router-link to="/teacher">
            <el-card class="card-panel" shadow="hover">
              <img class="image" src="@/assets/dashboard/teacher.png">
              <div class="description">
                <span>作为超级管理员，管理教师名单</span>
              </div>
            </el-card>
          </router-link>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import permission from '@/directive/permission/index.js'

const sources = ['exam', 'student', 'question'].map(e => {
  return '@/assets/dashboard/' + e + '.png'
})

export default {
  directives: {permission},
  name: 'DashboardAdmin'
}
</script>

<style lang="scss" scoped>
.dashboard-container {
  .panel-group {
    margin-top: 18px;

    .card-panel-col {
      margin: 22px;

      .card-panel {
        .image {
          float: left;
          width: 30%;
          display: block;
          margin: 20px;
        }

        .description {
          padding: 30px 20px 30px 50px;
        }

        margin: 18px;
        border-color: rgba(0, 0, 0, .05);
        background: #ffe;
        font-size: 20px;
        position: relative;
        overflow: hidden;
        color: #666;
      }
    }
  }
}
</style>
