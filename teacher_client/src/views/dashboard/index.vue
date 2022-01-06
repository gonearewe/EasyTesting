<template>
  <div class="app-container">
    <div class="dashboard-container">
      <el-row :gutter="40" class="panel-group">
        <el-col :lg="6" :sm="12" :xs="12" class="card-panel-col">
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
                <span>管理学生名单</span>
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

    <!--    TODO: a div showing github repo releases-->
    <div v-show="false" class="dashboard-container">
      <github-corner class="github-corner"/>
      <el-timeline>
        <a v-for="(activity, index) in activities" href="https://github.com/gonearewe/EasyTesting" target="_blank">
          <el-timeline-item
            :key="index"
            :timestamp="activity.timestamp">
            {{ activity.content }}
          </el-timeline-item>
        </a>
      </el-timeline>
    </div>
  </div>
</template>

<script>
import GithubCorner from '@/components/GithubCorner'
import permission from '@/directive/permission/index.js'

const sources = ['exam', 'student', 'question'].map(e => {
  return '@/assets/dashboard/' + e + '.png'
})

export default {
  directives: {permission},
  name: 'DashboardAdmin',
  components: {
    GithubCorner,
  },
  created() {
    // getRepoReleases().then((res) => {
    //     // this.activities
    //     console.log(res)
    //   }
    // )
  },
  data() {
    return {
      activities: [{
        content: '活动按期开始',
        timestamp: '2018-04-15'
      }, {
        content: '通过审核',
        timestamp: '2018-04-13'
      }, {
        content: '创建成功',
        timestamp: '2018-04-11'
      }]
    }
  }
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
          width: 150px;
          display: block;
          margin: 20px;
        }

        .description {
          padding: 30px 20px 30px 50px;
        }

        margin: 20px;
        border-color: rgba(0, 0, 0, .05);
        background: #ffe;
        font-size: 20px;
        position: relative;
        overflow: hidden;
        color: #666;
      }
    }
  }

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }
}
</style>
