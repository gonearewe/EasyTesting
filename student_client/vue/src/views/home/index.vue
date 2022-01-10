<template>
  <div class="app-container">
    <el-tabs type="border-card">
      <el-tab-pane label="单选题">
        <el-card v-for="(mcq,i) in questions.mcq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown>{{ mcq.stem }}</vue-markdown>
          </div>
          <el-radio-group v-model="answers.mcq[i]">
            <div v-for="j in 4">
              <el-radio :label="j">
                <vue-markdown>{{ mcq.choices[j - 1] }}</vue-markdown>
              </el-radio>
            </div>
          </el-radio-group>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="多选题">
        <el-card v-for="(maq,i) in questions.maq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown>{{ maq.stem }}</vue-markdown>
          </div>
          <el-checkbox-group v-model="answers.maq[i]">
            <div v-for="(c,j) in maq.choices">
              <el-checkbox :label="j+1">
                <vue-markdown>{{ c }}</vue-markdown>
              </el-checkbox>
            </div>
          </el-checkbox-group>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="填空题">
        <el-card v-for="(bfq,i) in questions.bfq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown>{{ bfq.stem }}</vue-markdown>
          </div>
          <el-input v-for="j in bfq.blank_num" v-model="answers.bfq[i][j-1]" autosize style="margin: 10px"
                    type="textarea"></el-input>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="判断题">
        <el-card v-for="(tfq,i) in questions.tfq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown>{{ tfq.stem }}</vue-markdown>
          </div>
          <el-radio-group v-model="answers.tfq[i]">
            <el-radio-button v-for="j in 2" :label="[true,false][j-1]"></el-radio-button>
          </el-radio-group>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="代码阅读题">
        <el-card v-for="(crq,i) in questions.crq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown>{{ crq.stem }}</vue-markdown>
          </div>
          <el-input v-for="j in crq.blank_num" v-model="answers.crq[i][j-1]" autosize style="margin: 10px"
                    type="textarea"></el-input>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="编程题"></el-tab-pane>
    </el-tabs>
    <back-to-top/>
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
      for (const [questionName, questionArray] of Object.entries(questions)) {
        _.shuffle(questionArray)
      }
      this.questions = questions
    })
  },
  data() {
    return {
      questions: {},
      answers: {
        // HACK: for we can't bind v-model to attributes later added in `created` hook
        "mcq": new Array(100),
        "maq": Array.from(new Array(100), e => []),
        "bfq": Array.from(new Array(100), e => new Array(3)),
        "tfq": new Array(100),
        "crq": Array.from(new Array(100), e => new Array(6)),
        "cq": new Array(100)
      },
    }
  }
}
</script>

<style lang="scss" scoped>
.box-card {
  margin: 30px;
}
</style>
