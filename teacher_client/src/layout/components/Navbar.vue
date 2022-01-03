<template>
  <div class="navbar">
    <hamburger :is-active="sidebar.opened" class="hamburger-container" @toggleClick="toggleSideBar"/>

    <breadcrumb class="breadcrumb-container"/>

    <div class="right-menu">
      <el-dropdown class="avatar-container" trigger="click">
        <div class="avatar-wrapper">
          <img class="user-avatar" src="../../icons/avatar.gif">
          <i class="el-icon-caret-bottom"/>
        </div>
        <el-dropdown-menu slot="dropdown" class="user-dropdown">
          <router-link to="/">
            <el-dropdown-item> 首页</el-dropdown-item>
          </router-link>
          <el-dropdown-item @click.native="handleUpdate"> 个人信息</el-dropdown-item>
          <a href="https://github.com/gonearewe/EasyTesting" target="_blank">
            <el-dropdown-item>文档</el-dropdown-item>
          </a>
          <el-dropdown-item @click.native="openAboutMsgBox">关于</el-dropdown-item>
          <el-dropdown-item divided @click.native="logout">
            <span style="display:block;">退出登录</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>

    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" title="编辑个人信息"
               width="30%">
      <el-form ref="dataForm" :model="temp" :rules="rules" label-position="left" label-width="100px"
               style="margin-left:50px;">
        <el-form-item label="工号" prop="teacher_id">
          <el-input v-model="temp.teacher_id"/>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="temp.password" placeholder="若不修改密码，请留空"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false"> 取消</el-button>
        <el-button type="danger" @click="updateData()"> 修改</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {mapGetters} from 'vuex'
import Breadcrumb from '@/components/Breadcrumb'
import Hamburger from '@/components/Hamburger'
import {updateTeacherProfile} from "@/api/teacher";
import _ from "lodash"
import {sha256} from "js-sha256"

export default {
  components: {
    Breadcrumb,
    Hamburger
  },
  computed: {
    ...mapGetters([
      'sidebar',
      'name',
      'teacher_id',
      'id'
    ])
  },
  data() {
    return {
      dialogFormVisible: false,
      temp: {
        name: '',
        teacher_id: '',
        password: ''
      },
      rules: {
        teacher_id: [{required: true, message: '必须填写工号', trigger: 'change'},
          {max: 10, message: '不得超过 10 个字符', trigger: 'change'}],
        name: [{required: true, message: '必须填写姓名', trigger: 'change'},
          {max: 50, message: '不得超过 50 个字符', trigger: 'change'}]
      },
    }
  },
  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    async logout() {
      await this.$store.dispatch('user/logout')
      this.$router.push(`/login?redirect=${this.$route.fullPath}`)
    },
    openAboutMsgBox() {
      this.$alert('Easy Testing 系统是一个用于 Python 课程考核的在线考试系统。这是系统的教师端。详情请查看文档。',
        '关于 E-Testing')
    },
    resetTemp() {
      this.temp = {
        id: this.id,
        teacher_id: this.teacher_id,
        name: this.name,
        password: ''
      }
    },
    handleUpdate() {
      this.resetTemp()
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          let tmp = _.assign({}, this.temp)
          // modify the copy instead because `this.temp` is the `model` and any change to it will be shown to user
          if (tmp.password) { // `if` is needed because empty string can also be encoded by sha256
            tmp.password = sha256(tmp.password)
          }
          updateTeacherProfile(tmp).then(() => {
            this.$message({
              message: '修改成功，请重新登录',
              showClose: true,
              type: 'success'
            })
            this.logout()
          })
        }
      })
    },
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, .08);

  .hamburger-container {
    line-height: 46px;
    height: 100%;
    float: left;
    cursor: pointer;
    transition: background .3s;
    -webkit-tap-highlight-color: transparent;

    &:hover {
      background: rgba(0, 0, 0, .025)
    }
  }

  .breadcrumb-container {
    float: left;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background .3s;

        &:hover {
          background: rgba(0, 0, 0, .025)
        }
      }
    }

    .avatar-container {
      margin-right: 30px;

      .avatar-wrapper {
        margin-top: 5px;
        position: relative;

        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          position: absolute;
          right: -20px;
          top: 25px;
          font-size: 12px;
        }
      }
    }
  }
}
</style>
