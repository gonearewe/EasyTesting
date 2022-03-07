<template>
  <div class="login-container">
    <particles-bg :bg="true" type="square"/>

    <el-form ref="loginForm" :model="loginForm" auto-complete="on" class="login-form"
             label-position="left">

      <div class="title-container">
        <h3 class="title">学生登录</h3>
      </div>

      <p class="label">服务器地址</p>
      <el-form-item prop="server_addr">
        <span class="svg-container">
          <svg-icon icon-class="server"/>
        </span>
        <el-input
          v-model="loginForm.server_addr"
          auto-complete="on"
          name="服务器地址"
          placeholder="例如：http://192.168.1.3:443"
          tabindex="1"
          type="text"
        />
      </el-form-item>

      <p class="label">学号</p>
      <el-form-item prop="student_id">
        <span class="svg-container">
          <svg-icon icon-class="id"/>
        </span>
        <el-input
          v-model="loginForm.student_id"
          auto-complete="on"
          name="学号"
          placeholder="例如：2020204520"
          tabindex="2"
          type="text"
        />
      </el-form-item>

      <p class="label">姓名</p>
      <el-form-item prop="name">
        <span class="svg-container">
          <svg-icon icon-class="user"/>
        </span>
        <el-input
          v-model="loginForm.name"
          name="姓名"
          placeholder="例如：王小明"
          tabindex="3"
          type="text"
        />
      </el-form-item>

      <p class="label">考试号</p>
      <el-form-item prop="exam_id">
        <span class="svg-container">
          <svg-icon icon-class="exam"/>
        </span>
        <el-input
          ref="exam_id"
          v-model="loginForm.exam_id"
          auto-complete="on"
          name="考试号"
          tabindex="4"
          @keyup.enter.native="dialogVisible=true"
        />
      </el-form-item>

      <el-button :disabled="loading" type="warning" @click.native.prevent="dialogVisible=true">
        登录
      </el-button>
    </el-form>

    <el-dialog :close-on-click-modal="false" title="注意事项" :visible.sync="dialogVisible">
      <span>关于软件</span>
      <vue-markdown>{{ notice.about }}</vue-markdown>
      <el-checkbox v-model="aboutChecked">我已阅读软件使用说明</el-checkbox>
      <br>
      <el-checkbox v-model="promiseChecked">
        <vue-markdown>{{ notice.promise }}</vue-markdown>
      </el-checkbox>
      <div slot="footer" class="dialog-footer">
        <el-button :disabled="!aboutChecked||!promiseChecked||loading" :loading="loading"
                   type="primary" @click="handleLogin">
          开始考试
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import VueMarkdown from 'vue-markdown'
import {ParticlesBg} from "particles-bg-vue";

export default {
  name: 'Login',
  components: {
    VueMarkdown,
    ParticlesBg
  },
  created() {
    this.$nextTick(() => {
      // 禁用右键
      document.oncontextmenu = new Function("event.returnValue=false");
      // 禁用选择
      document.onselectstart = new Function("event.returnValue=false");
    })
  },
  data() {
    return {
      loginForm: {
        server_addr: 'http://localhost:9000',
        student_id: '2020501880',
        name: '小明',
        exam_id: '4'
      },
      loading: false,
      redirect: undefined,

      dialogVisible: false,
      aboutChecked: false,
      promiseChecked: false,
      notice: {
        about:
        // multiline string mustn't contain leading indents, or markdown rendering will fail
          `
**Easy Testing** 是一个用于《Python 编程语言》课程的在线考试系统。
这是它的考生客户端。你可以在此答题。

考试题目分为**单项选择题、多项选择题、填空题、判断对错题、读程序题、编程题**五个部分。
你可以通过点击对应的选项卡跳转到任意题型进行作答。你可以以**任意顺序进行作答**。
编程题要求你编写程序解决问题，软件允许你**运行程序获得终端输出**；运行出错或超时同样会有输出。

右侧悬浮的倒计时会**显示你的剩余作答时间**，时间归零后软件会自动提交答卷，考试同时结束，作答终止。
如要提前交卷，请点击右侧“保存答卷”按钮，待**出现成功提示后**关闭软件窗口。

软件会每隔一段时间将你的作答情况**自动保存**至服务端，并给予成功提示。
你也可以通过界面右侧的“保存答卷”按钮手动保存，但是**请勿频繁提交**以免增加服务器压力。

考试过程中请勿关闭软件。如遇到**软件卡死、电脑意外关机**等情况，请尽快**重新打开软件并登录**，
软件会恢复你的试卷与上次保存的作答情况，但是中途消耗的时间无法补偿。
你也可以使用其他电脑重新打开软件并登录，但是同一时刻不允许在多个电脑上登录。
        `,
        promise: "**本人知晓我校考场规则和违纪处分条例的有关规定，保证遵守考场规则，诚实做人。**"
      }
    }
  },
  watch: {
    $route: {
      handler: function (route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    handleLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/login', this.loginForm).then(() => {
            this.$router.push({path: this.redirect || '/'})
          }).finally(() => {
            this.loading = false
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss">
/* 修复input 背景不协调 和光标变色 */
/* Detail see https://github.com/PanJiaChen/vue-element-admin/pull/927 */

$bg: #283443;
$light_gray: #fff;
$cursor: #fff;

@supports (-webkit-mask: none) and (not (cater-color: $cursor)) {
  .login-container .el-input input {
    color: $cursor;
  }
}

/* reset element-ui css */
.login-container {
  .el-button {
    width: 30%;
    margin-bottom: 30px;
    display: block;
    margin-left: auto;
    margin-right: auto;
  }

  .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;

    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      color: $light_gray;
      height: 47px;
      caret-color: $cursor;

      &:-webkit-autofill {
        box-shadow: 0 0 0px 1000px $bg inset !important;
        -webkit-text-fill-color: $cursor !important;
      }
    }
  }

  .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: rgba(0, 0, 0, 0.1);
    border-radius: 5px;
    color: #454545;
  }

  .label {
    color: #f0c78a;
  }
}
</style>

<style lang="scss" scoped>
$bg: gray;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100%;
  width: 100%;
  overflow: hidden;

  .login-form {
    position: relative;
    background-color: $bg;
    width: 520px;
    top: 50px;
    border-radius: 20px;
    max-width: 100%;
    padding: 50px 35px 0;
    margin: 0 auto;
    overflow: hidden;
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .tips {
    font-size: 14px;
    color: #fff;
    margin-bottom: 10px;

    span {
      &:first-of-type {
        margin-right: 16px;
      }
    }
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }
}
</style>
