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
      <el-tab-pane label="编程题">
        <el-card v-for="(cq,i) in questions.cq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown>{{ cq.stem }}</vue-markdown>
            <div>
              <el-tag v-if="cq.input === ''" type="info">无输入</el-tag>
              <el-tag v-else :type="cq.is_input_from_file?'success':'warning'" style="margin-right:10px;">
                {{ cq.is_input_from_file ? '输入来自文件' : '输入来自命令行' }}
              </el-tag>
              <el-tag :type="cq.is_output_to_file?'success':'warning'" style="margin-right:10px;">
                {{ cq.is_output_to_file ? '输出至文件' : '输出至命令行' }}
              </el-tag>
              <el-button size="small" style="float: right" type="primary" @click="runStudentCode(cq,answers.cq[i])">
                运行代码
              </el-button>
            </div>
          </div>
          <VueCodeEditor
            v-model="answers.cq[i].code"
            :options="{
        enableBasicAutocompletion: true,
        enableLiveAutocompletion: true,
        fontSize: 14,
        highlightActiveLine: true,
        // enableSnippets: true,
        showLineNumbers: true,
        tabSize: 2,
        showPrintMargin: false,
        showGutter: true,
    }"
            height="400px"
            lang="python"
            style="border-radius: 20px; margin-bottom: 50px"
            theme="tomorrow_night_eighties"
            @init="editorInit"
          />
          <el-input
            v-model="answers.cq[i].console_output"
            :rows="8"
            placeholder="终端的输出都会显示在这里"
            readonly
            type="textarea"
          >
          </el-input>
        </el-card>
      </el-tab-pane>
    </el-tabs>
    <div class="fixed-box">
      <flip-countdown :deadline="deadline" :showDays="false" countdownSize="x-large" labelSize="small"
                      @timeElapsed="saveAnswers"></flip-countdown>
      <el-button size="medium" type="warning" @click="saveAnswers">提交答卷</el-button>
    </div>
    <back-to-top/>
  </div>
</template>

<script>
import _ from 'lodash'
import {getMyQuestions, loadMyAnswerModels, runCode, saveMyAnswerModels, submitMyAnswers} from "@/api"
import VueMarkdown from 'vue-markdown'
import BackToTop from '@/components/BackToTop'
import VueCodeEditor from 'vue2-code-editor'
import {sha256} from "js-sha256"
import FlipCountdown from 'vue2-flip-countdown'

export default {
  name: 'HOME',
  components: {
    VueMarkdown,
    BackToTop,
    VueCodeEditor,
    FlipCountdown
  },
  created() {
    this.$nextTick(() => {
      // 禁用右键
      document.oncontextmenu = new Function("event.returnValue=false");
      // 禁用选择
      document.onselectstart = new Function("event.returnValue=false");
    })

    getMyQuestions().then(questions => {
      for (const [questionName, questionArray] of Object.entries(questions)) {
        _.shuffle(questionArray)
      }
      this.questions = questions
      this.answers.cq = this.questions.cq.map(e => {
        return {code: e.template, console_output: '', right: false}
      })

      // try loading models, may not succeed if we're starting the client for the first time
      loadMyAnswerModels().then(body=>{
        this.$message({
          message: '已恢复你的作答至上一次保存时的状态',
          showClose: true,
          type: 'success'
        })
        this.answers = body
        }
      )
    })

    setInterval(() => this.saveAnswers(), 2 * 60 * 1000) // auto-save every 2 minutes
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
        "cq": Array.from(new Array(100), e => {
          return {code: '', console_output: '', right: false}
        }),
      },
      deadline: this.$store.getters.exam_deadline || '2089-01-01 08:30:00'
    }
  },
  methods: {
    editorInit() {
      // vue2-code-editor/node_modules/
      require('brace/ext/language_tools') //language extension prerequsite...
      require('brace/mode/python') //language
      require('brace/theme/tomorrow_night_eighties')
      require('brace/snippets/python') //snippet
    },
    runStudentCode(cq, answer) {
      let tmp = _.merge({}, cq)
      tmp.code = answer.code
      runCode(tmp).then(body => {
        answer.console_output = body.console_output
        answer.right = body.pass
      })
    },
    saveAnswers() {
      const questions = _.merge({}, this.questions)
      const answers = _.merge({}, this.answers)
      saveMyAnswerModels(answers)
      let req = {}
      req.mcq = questions.mcq.map((e, i) => {
        return {id: e.id, answer: answers.mcq[i]}
      })
      req.maq = questions.maq.map((e, i) => {
        return {id: e, answer: answers.maq[i]}
      })
      req.bfq = questions.bfq.map((e, i) => {
        return {id: e.id, answer: answers.bfq[i].slice(0, e.blank_num)}
      })
      req.tfq = questions.tfq.map((e, i) => {
        return {id: e.id, answer: answers.tfq[i]}
      })
      req.crq = questions.crq.map((e, i) => {
        return {id: e.id, answer: answers.crq[i].slice(0, e.blank_num)}
      })
      req.cq = questions.cq.map((e, i) => {
        return {id: e.id, answer: answers.cq[i].code, right: answers.cq[i].right}
      })
      submitMyAnswers(req).then(() => {
        this.$message({
          message: '你的作答已自动保存',
          showClose: true,
          type: 'success'
        })
      })
    }
  },
}
</script>

<style lang="scss" scoped>
.box-card {
  margin: 30px;
}

.fixed-box {
  position: fixed;
  top: 50px;
  right: 10px;

  .el-button {
    float: right;
    margin: 10px;
  }
}
</style>
