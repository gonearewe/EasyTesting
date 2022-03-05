<template>
  <div class="app-container">
    <el-dialog :close-on-click-modal="false" :visible.sync="tipVisible" title="小贴士">
      <div style="text-align: center">
        <img src="@/assets/tip_zoom.png" style="width: 100px;height: 80px;"/>
        <p><strong>Ctrl + 鼠标滚轮 缩放页面</strong></p>
      </div>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click.native="tipVisible=false">
          确定
        </el-button>
      </div>
    </el-dialog>

    <el-skeleton v-if="tabsLoading" :rows="6" animated style="top: 30px"/>
    <el-tabs v-else type="border-card">
      <el-tab-pane label="单选题">
        <el-card v-for="(mcq,i) in questions.mcq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <!--            NOTICE: Instead of code below, moving model into `source` can prevent-->
            <!--            an event called `rendered` triggered every time user presses a key which results-->
            <!--            in a significant performance loss.-->
            <!--            <vue-markdown>{{mcq.stem}}</vue-markdown>-->
            <vue-markdown :source="mcq.stem"></vue-markdown>
          </div>
          <el-radio-group v-model="answers.mcq[i]">
            <div v-for="j in 4">
              <el-radio :label="j">
                <vue-markdown :source="mcq.choices[j - 1]"></vue-markdown>
              </el-radio>
            </div>
          </el-radio-group>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="多选题">
        <el-card v-for="(maq,i) in questions.maq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown :source="maq.stem"></vue-markdown>
          </div>
          <el-checkbox-group v-model="answers.maq[i]">
            <div v-for="(c,j) in maq.choices">
              <el-checkbox :label="j+1">
                <vue-markdown :source="c"></vue-markdown>
              </el-checkbox>
            </div>
          </el-checkbox-group>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="填空题">
        <el-card v-for="(bfq,i) in questions.bfq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown :source="bfq.stem"></vue-markdown>
          </div>
          <el-input v-for="j in bfq.blank_num" v-model="answers.bfq[i][j-1]" autosize style="margin: 10px"
                    type="textarea"></el-input>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="判断题">
        <el-card v-for="(tfq,i) in questions.tfq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown :source="tfq.stem"></vue-markdown>
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
            <vue-markdown :source="crq.stem"></vue-markdown>
          </div>
          <el-input v-for="j in crq.blank_num" v-model="answers.crq[i][j-1]" autosize style="margin: 10px"
                    type="textarea"></el-input>
        </el-card>
      </el-tab-pane>
      <el-tab-pane label="编程题">
        <el-card class="box-card" shadow="hover">
          <p>
            <i class="el-icon-warning-outline"></i>
            请根据题意编写代码，从文件或命令行中读取输入，向文件或命令行打印输出。
            假如涉及文件，输入文件为工作目录下的 input.txt，输出文件则为 output.txt。
            示例：f = open("output.txt", "rw")
          </p>
          <p>
            <i class="el-icon-warning-outline"></i>
            你可通过 print 函数进行调试。系统自动保存或提交答卷时只会保存最近一次运行的代码，
            所以<strong>在完成或修改代码后请务必运行一次</strong>。
          </p>
        </el-card>
        <el-card v-for="(cq,i) in questions.cq" class="box-card" shadow="hover">
          <div slot="header">
            <el-tag style="margin-right: 10px">{{ '第 ' + (i + 1) + ' 题' }}</el-tag>
            <vue-markdown :source="cq.stem"></vue-markdown>
            <div>
              <el-tag v-if="cq.input === ''" type="info">无输入</el-tag>
              <el-tag v-else :type="cq.is_input_from_file?'success':'warning'" style="margin-right:10px;">
                {{ cq.is_input_from_file ? '输入来自文件' : '输入来自命令行' }}
              </el-tag>
              <el-tag :type="cq.is_output_to_file?'success':'warning'" style="margin-right:10px;">
                {{ cq.is_output_to_file ? '输出至文件' : '输出至命令行' }}
              </el-tag>
              <!--      disable the button when `runningCode` to achieve ratelimit -->
              <el-button :loading="runningCode" size="small" style="float: right" type="primary"
                         :disabled="runningCode" icon="el-icon-time" @click="runStudentCode(cq,answers.cq[i])">
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
      <!--      disable the button when `submitting` to achieve client-side ratelimit -->
      <el-button :disabled="submitting" :loading="submitting" icon="el-icon-upload"
                 size="medium" type="warning" @click="saveAnswers">
        保存答卷
      </el-button>
      <br>
      <el-button :disabled="submitting" icon="el-icon-s-claim"
                 size="medium" type="info" @click="notifyProgress">
        检查进度
      </el-button>
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
import FlipCountdown from 'vue2-flip-countdown'
import {shuffle} from "@/utils/random"

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
      for (const questionName in questions) {
        // shuffle question array, but with `exam_session_id` as the random seed
        // so that the question array is in the same order every time the student enters this page
        shuffle(questions[questionName], this.$store.getters.exam_session_id)
        // TODO: shuffle choice array of mcq and maq for anti-cheating
      }
      this.questions = questions
      // trim the length of answers
      Object.entries(this.questions).forEach(([key,arr])=>{
        this.answers[key].splice(arr.length)
      })
      this.answers.cq = this.questions.cq.map(e => {
        return {code: e.template, console_output: '', right: false}
      })

      // try loading models, may not succeed if we're starting the client for the first time
      loadMyAnswerModels().then(body => {
          // body will be empty if we're starting the client for the first time
          if (body && body.mcq && body.maq && body.tfq && body.bfq && body.crq && body.cq) {
            this.$message({
              message: '已恢复你的作答至上一次保存时的状态',
              duration: 5000,
              showClose: true,
              type: 'success'
            })
            this.answers = body
          }
        }
      ).finally(() => {
        this.tabsLoading = false
      })
    })

    // auto-save every x seconds, x <- [120,180), interval includes random to avoid flush to server
    setInterval(() => this.saveAnswers(),
      (2 * 60 + Math.floor(Math.random()*60)) * 1000 )
  },
  data() {
    return {
      tipVisible: true,
      tabsLoading: true,
      runningCode: false,
      submitting: false,
      questions: {},
      answers: {
        // HACK: for we can't bind v-model to attributes later added in `created` hook
        "mcq": new Array(80),
        "maq": Array.from(new Array(60), e => []),
        "bfq": Array.from(new Array(60), e => new Array(3)),
        "tfq": new Array(60),
        "crq": Array.from(new Array(15), e => new Array(6)),
        "cq": Array.from(new Array(10), e => {
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
      this.runningCode = true
      let tmp = _.merge({}, cq)
      tmp.code = answer.code
      runCode(tmp).then(body => {
        answer.console_output = body.console_output
        answer.right = body.pass
      }).finally(() => {
        this.runningCode = false
      })
    },
    freezeAnswers() {
      const questions = _.merge({}, this.questions)
      const answers = _.merge({}, this.answers)
      let ret = {}
      ret.mcq = questions.mcq.map((e, i) => {
        return {id: e.id, answer: answers.mcq[i]}
      })
      ret.maq = questions.maq.map((e, i) => {
        return {id: e.id, answer: answers.maq[i]}
      })
      ret.bfq = questions.bfq.map((e, i) => {
        return {id: e.id, answer: answers.bfq[i].slice(0, e.blank_num).map(e => e === null ? '' : e)}
      })
      ret.tfq = questions.tfq.map((e, i) => {
        return {id: e.id, answer: answers.tfq[i]}
      })
      ret.crq = questions.crq.map((e, i) => {
        return {id: e.id, answer: answers.crq[i].slice(0, e.blank_num).map(e => e || '')}
      })
      ret.cq = questions.cq.map((e, i) => {
        return {id: e.id, answer: answers.cq[i].code || '', right: answers.cq[i].right}
      })
      return ret
    },
    saveAnswers() {
      this.submitting = true
      saveMyAnswerModels(_.merge({}, this.answers))
      let req = this.freezeAnswers()
      submitMyAnswers(req).then(() => {
        this.$message({
          message: '你的作答已保存',
          showClose: true,
          type: 'success'
        })
      }).finally(() => {
        this.submitting = false
      })
    },
    notifyProgress() {
      let answers = this.freezeAnswers()
      let mcqCompletedCnt = answers.mcq.filter(e => e.answer).length
      let maqCompletedCnt = answers.maq.filter(e => e.answer.length > 0).length
      let bfqCompletedCnt = answers.bfq.filter(e => e.answer.every(a => a)).length
      let tfqCompletedCnt = answers.tfq.filter(e => typeof e.answer == "boolean").length
      let crqCompletedCnt = answers.crq.filter(e => e.answer.every(a => a)).length
      let uncompletedCnt = Object.values(answers).map(e => e.length).reduce((a, b) => a + b) -
        mcqCompletedCnt - maqCompletedCnt - bfqCompletedCnt - tfqCompletedCnt - crqCompletedCnt - answers.cq.length
      this.$notify({
        title: uncompletedCnt === 0 ? '你已完成所有试题' : '你还有 ' + uncompletedCnt + ' 道小题未完成',
        type: uncompletedCnt === 0 ? 'success' : 'warning',
        dangerouslyUseHTMLString: true,
        message: [
          '已完成 <strong>' + mcqCompletedCnt + '/' + answers.mcq.length + ' 的单选题</strong>',
          '已完成 <strong>' + maqCompletedCnt + '/' + answers.maq.length + ' 的多选题</strong>',
          '已完成 <strong>' + bfqCompletedCnt + '/' + answers.bfq.length + ' 的填空题</strong>',
          '已完成 <strong>' + tfqCompletedCnt + '/' + answers.tfq.length + ' 的判断题</strong>',
          '已完成 <strong>' + crqCompletedCnt + '/' + answers.crq.length + ' 的代码阅读题</strong>',
          '<strong>注意：' + answers.cq.length + ' 道编程题的进度无法统计</strong>',
        ].join('<hr/>'),
        position: 'top-left',
        duration: 8000
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
    margin: 10px 10px 0;
  }
}
</style>
