<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.publisher_teacher_id" class="filter-item" placeholder="出题者工号" style="width: 200px;"
                @keyup.enter.native="handleFilter"/>
      <el-button v-waves class="filter-item" icon="el-icon-search" type="primary" @click="handleFilter">
        搜索
      </el-button>
      <el-button v-waves class="filter-item" icon="el-icon-edit" type="primary"
                 @click="handleCreate">
        添加
      </el-button>
      <el-tooltip class="item" content="勾选表格左侧以多选，删除所有选中项" effect="dark" placement="top-start">
        <el-button v-waves class="filter-item" icon="el-icon-delete" type="danger"
                   @click="handleMultiDelete">
          删除
        </el-button>
      </el-tooltip>
    </div>

    <el-skeleton v-if="listLoading" :rows="6" animated/>
    <el-table
      :key="tableKey"
      ref="maqTable"
      v-else
      :data="list"
      border
      fit
      show-header
      stripe
      style="width: 100%;"
    >

      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column align="center" label="ID" prop="id" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="出题者工号" width="150">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.publisher_teacher_id }}</span>
        </template>
      </el-table-column>
      <el-table-column align="left" header-align="center" label="题干" show-overflow-tooltip>
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.stem }}</span>
        </template>
      </el-table-column>
      <el-table-column v-for="i in 7" :label="'选项 '+i" align="left" header-align="center" show-overflow-tooltip>
        <template slot-scope="{row}">
          <el-tag v-show="row.right_answer.includes(i)" style="margin-right:10px;" type="success">正解</el-tag>
          <span v-if="i <= row.choices.length" class="link-type" @click="handleUpdate(row)">
            {{ row.choices[i - 1] }}
          </span>
          <el-tag v-else type="info">无</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" header-align="center" label="难度" width="60">
        <template slot-scope="{row}">
          <el-tooltip v-if="row.overall_score < 100" content="数据过少，无法分析" effect="dark" placement="top">
            <el-tag type="info">无</el-tag>
          </el-tooltip>
          <el-tag v-else :type="getDifficultyColor(1-row.overall_correct_score/row.overall_score)">
            {{ (1 - row.overall_correct_score / row.overall_score).toFixed(1) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="200">
        <template slot-scope="{row}">
          <el-button size="mini" type="primary" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-popconfirm
            confirm-button-type="danger"
            icon-color="red"
            style="margin-left: 10px"
            title="确定删除吗？"
            @confirm="handleDelete(row)"
          >
            <el-button slot="reference" size="mini" type="danger">
              删除
            </el-button>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :limit.sync="listQuery.page_size" :page.sync="listQuery.page_index" :total="total"
                align="center" @pagination="getList"/>

    <el-dialog :close-on-click-modal="false" :title="textMap[dialogStatus]"
               :visible.sync="dialogFormVisible" width="60%">
      <el-form ref="dataForm" :model="temp" :rules="rules" label-position="left" label-width="100px"
               style="margin-left:50px;">
        <el-form-item label="题干" prop="stem">
          <markdown-editor v-model="temp.stem"/>
        </el-form-item>
        <el-form-item label="选项数" prop="blank_num">
          <el-radio-group v-model="temp.choice_num" @change="updateTemp">
            <el-radio-button v-for="i in 4" :label="i + 3">
              {{ ['4', '5', '6', '7'][i - 1] + ' 个' }}
            </el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="选项" prop="choices">
          <markdown-editor v-for="i in temp.choice_num" v-model="temp.choices[i-1]"/>
        </el-form-item>
        <el-form-item label="正确答案" prop="right_answer">
          <el-checkbox-group v-model="temp.right_answer">
            <el-checkbox v-for="i in temp.choice_num" :label="i" border>
              {{ '第' + ['一', '二', '三', '四', '五', '六', '七'][i - 1] + '个选项' }}
            </el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          确定
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :close-on-click-modal="false" :title="'确认删除以下 '+rowsToBeDeleted.length+' 条记录？'"
               :visible.sync="dialogDeleteVisible">
      <el-table :data="rowsToBeDeleted" max-height="800">
        <el-table-column align="center" label="ID" property="id" width="100"></el-table-column>
        <el-table-column align="center" label="出题者工号" property="publisher_teacher_id" width="150"></el-table-column>
        <el-table-column header-align="center" label="题干" property="stem" show-overflow-tooltip
                         width="200"></el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogDeleteVisible = false">
          取消
        </el-button>
        <el-button type="danger" @click="deleteData">
          确定
        </el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {createQuestions, deleteQuestions, getQuestions, updateQuestion} from '@/api/question'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination'
import MarkdownEditor from "@/components/MarkdownEditor";
import _ from "lodash"
import {getDifficultyColor} from "@/views/question/common";

export default {
  name: 'MaqList',
  components: {MarkdownEditor, Pagination},
  directives: {waves},
  data() {
    return {
      tableKey: 0,

      list: null,
      total: 0,

      listLoading: true,
      listQuery: {
        publisher_teacher_id: '',
        page_index: 1,
        page_size: 20
      },

      temp: {
        id: undefined,
        stem: '',
        choice_num: 4,
        choices: [],
        right_answer: []
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },

      rowsToBeDeleted: [],
      dialogDeleteVisible: false,

      rules: {
        stem: [{required: true, message: '必须填写题干', trigger: 'change'},
          {max: 200, message: '不得超过 200 个字符', trigger: 'change'}],
        right_answer: [{required: true, message: '必须给出正确答案', trigger: 'change'}]
      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getDifficultyColor,
    getList() {
      this.listLoading = true
      getQuestions('maq', this.listQuery).then(body => {
        this.list = body.data
        this.total = body.total
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page_index = 1
      this.getList()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        stem: '',
        blank_num: 4,
        choices: [],
        right_answer: []
      }
    },
    updateTemp() {
      this.temp.choices.splice(this.temp.choice_num)
      this.temp.right_answer = []
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          let req = _.merge({}, this.temp)
          delete req.blank_num
          createQuestions('maq', [req]).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: '创建成功',
              showClose: true,
              type: 'success'
            })
            this.getList()
          })
        }
      })
    },
    handleUpdate(row) {
      _.merge(this.temp, row)
      // this.temp = Object.assign({}, row) // NOTICE: shadow copy will cause problems on array
      this.temp.choice_num = row.choices.length
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const req = _.merge({}, this.temp)
          delete req.blank_num
          updateQuestion('maq', req).then(() => {
            this.dialogFormVisible = false
            this.$message({
              message: '修改已保存',
              showClose: true,
              type: 'success'
            })
            this.getList()
          })
        }
      })
    },
    handleMultiDelete() {
      let rows = this.$refs.maqTable.selection;
      if (rows.length === 0) {
        this.$message({
          message: '没有任何一项被选中，勾选表格左侧以多选',
          showClose: true,
          type: 'warning'
        })
      } else {
        this.rowsToBeDeleted = Object.assign([], rows)
        this.dialogDeleteVisible = true
      }
    },
    handleDelete(row) {
      this.rowsToBeDeleted = []
      this.rowsToBeDeleted[0] = Object.assign({}, row)
      this.deleteData()
    },
    deleteData() {
      deleteQuestions('maq', this.rowsToBeDeleted.map(v => v.id)).then(() => {
        this.dialogDeleteVisible = false
        this.$message({
          message: '删除成功',
          showClose: true,
          type: 'success'
        })
        this.getList()
      })
    }
  }
}
</script>
