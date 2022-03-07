<template>
  <div class="login-container">
    <particles-bg :bg="true" type="circle"/>

    <el-form ref="loginForm" :model="loginForm" :rules="loginRules" auto-complete="on" class="login-form"
             label-position="left">

      <div class="title-container">
        <h3 class="title">教师登录</h3>
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

      <p class="label">工号</p>
      <el-form-item prop="teacher_id">
        <span class="svg-container">
          <svg-icon icon-class="user"/>
        </span>
        <el-input
          ref="username"
          v-model="loginForm.teacher_id"
          auto-complete="on"
          name="工号"
          placeholder="例如：2010204520"
          tabindex="2"
          type="text"
        />
      </el-form-item>

      <p class="label">密码</p>
      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon icon-class="password"/>
        </span>
        <el-input
          :key="passwordType"
          ref="password"
          v-model="loginForm.password"
          :type="passwordType"
          auto-complete="on"
          name="密码"
          tabindex="3"
          @keyup.enter.native="handleLogin"
        />
        <span class="show-pwd" @click="showPwd">
          <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'"/>
        </span>
      </el-form-item>

      <el-button :disabled="loading" :loading="loading" type="warning" @click.native.prevent="handleLogin">
        登录
      </el-button>

    </el-form>
  </div>
</template>

<script>
import {validTeacherId} from '@/utils/validate'
import {ParticlesBg} from "particles-bg-vue";

export default {
  name: 'Login',
  components: {
    ParticlesBg
  },
  data() {
    const validateTeacherId = (rule, value, callback) => {
      if (!validTeacherId(value)) {
        callback(new Error('Please enter the correct user name'))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      if (value.length < 5) {
        // TODO: demand strong password
        callback(new Error('The password can not be less than 5 characters'))
      } else {
        callback()
      }
    }
    return {
      loginForm: {
        server_addr: 'http://localhost:9000',
        teacher_id: '0',
        password: 'ET000'
      },
      loginRules: {
        teacher_id: [{required: true, trigger: 'blur', validator: validateTeacherId}],
        password: [{required: true, trigger: 'blur', validator: validatePassword}]
      },
      loading: false,
      passwordType: 'password',
      redirect: undefined
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
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
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
    top: 150px;
    border-radius: 20px;
    max-width: 100%;
    padding: 50px 35px 0;
    margin: 0 auto;
    overflow: hidden;
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

  .label {
    color: #f0c78a;
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
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

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }
}
</style>
