<template>
  <div class="app-container">
    <el-tabs tabPosition="left" type="border-card">
      <el-tab-pane label="单选题">
        <el-card v-for="(mcq,i) in questions.mcq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + i + ' 题' }}</el-tag>
            <vue-markdown>{{ mcq.stem }}</vue-markdown>
          </div>
          <el-radio-group v-model="mcq.choices">
            <div v-for="i in 4">
              <el-radio :label="i">
                <vue-markdown>{{ mcq.choices[i - 1] }}</vue-markdown>
              </el-radio>
            </div>
          </el-radio-group>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="多选题">配置管理</el-tab-pane>
      <el-tab-pane label="填空题">角色管理</el-tab-pane>
      <el-tab-pane label="判断题">定时任务补偿</el-tab-pane>
      <el-tab-pane label="代码阅读题">定时任务补偿</el-tab-pane>
      <el-tab-pane label="编程题">定时任务补偿</el-tab-pane>
    </el-tabs>
    <back-to-top :back-position="50" transition-name="fade"/>
  </div>
</template>

<script>
import _ from 'lodash'
import {getMyQuestions} from "@/api"
import VueMarkdown from 'vue-markdown'
import BackToTop from '@/components/BackToTop'

export default {
  name: 'HOME',
  components: {
    VueMarkdown,
    BackToTop
  },
  created() {
    getMyQuestions().then(questions => {
      for (const v of Object.values(questions)) {
        _.shuffle(v)
      }
      this.questions = questions
    })
  },
  data() {
    return {
      questions: []
    }
  }
}
</script>

<style lang="scss" scoped>
.box-card {
  margin: 30px;
}
</style>
