<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.publisher_teacher_id" class="filter-item" placeholder="出题人工号" style="width: 200px;"
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

      <el-tooltip class="item" content="从 Excel(*.xlsx) 文件中导入数据，文件需遵守文档中的规范" effect="dark" placement="top-start">
        <el-button v-waves class="filter-item" icon="el-icon-upload" type="primary"
                   @click="$refs.file_picker.click()">
          导入
        </el-button>
      </el-tooltip>

      <el-tooltip class="item" content="导出满足当前筛选条件的所有数据至一个 Excel(*.xlsx) 文件" effect="dark" placement="top-start">
        <el-button v-waves :loading="downloadLoading" class="filter-item" icon="el-icon-download" type="success"
                   @click="handleDownload">
          导出
        </el-button>
      </el-tooltip>
    </div>

    <input ref="file_picker" accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" hidden
           type='file' @change="handleUpload"/>

    <el-table
      :key="tableKey"
      ref="mcqTable"
      v-loading="listLoading"
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
      <el-table-column align="center" label="出题人工号">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.publisher_teacher_id }}</span>
        </template>
      </el-table-column>
      <el-table-column align="left" header-align="center" label="题干" show-overflow-tooltip>
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.stem }}</span>
        </template>
      </el-table-column>
      <!--      <el-table-column align="center" label="选项">-->
      <el-table-column v-for="i in 4" :label="'选项 '+i" align="left" header-align="center" show-overflow-tooltip>
        <template slot-scope="{row}">
          <el-tag v-show="row.right_answer===i" style="margin-right:10px;" type="success">正解</el-tag>
          <span class="link-type" @click="handleUpdate(row)">{{ row.choices[i - 1] }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="230">
        <template slot-scope="{row}">
          <el-button size="mini" type="primary" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="handleDelete(row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :limit.sync="listQuery.page_size" :page.sync="listQuery.page_index" :total="total"
                align="center" @pagination="getList"/>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="60%">
      <el-form ref="dataForm" :model="temp" :rules="rules" label-position="left" label-width="100px"
               style="margin-left:50px;">
        <el-form-item label="题干" prop="stem">
          <markdown-editor v-model="temp.stem"/>
        </el-form-item>
        <el-form-item label="选项" prop="choices">
          <markdown-editor v-for="i in 4" v-model="temp.choices[i-1]"/>
        </el-form-item>
        <el-form-item label="正确答案" prop="right_answer">
          <el-radio-group v-model="temp.right_answer">
            <el-radio-button v-for="i in 4" :label="i" :value="i">
              {{ '第' + ['一', '二', '三', '四'][i - 1] + '个选项' }}
            </el-radio-button>
          </el-radio-group>
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

    <el-dialog :title="'确认删除以下 '+rowsToBeDeleted.length+' 条记录？'"
               :visible.sync="dialogDeleteVisible">
      <el-table :data="rowsToBeDeleted" max-height="800">
        <el-table-column align="center" label="ID" property="id" width="150"></el-table-column>
        <el-table-column align="center" label="出题人工号" property="publisher_teacher_id" width="150"></el-table-column>
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

    <el-dialog :title="'确认添加以下 '+rowsToBeAdded.length+' 条记录？'"
               :visible.sync="dialogImportVisible">
      <el-table :data="rowsToBeAdded" max-height="800">
        <el-table-column align="center" label="出题人工号" property="publisher_teacher_id" width="150"></el-table-column>
        <el-table-column align="center" label="题干" property="stem" width="200"></el-table-column>
        <el-table-column label="选项" property="choices" width="150"></el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogImportVisible = false">
          取消
        </el-button>
        <el-button type="primary" @click="uploadData">
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

export default {
  name: 'StudentList',
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
        choices: [],
        right_answer: 1
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },

      rowsToBeDeleted: [],
      dialogDeleteVisible: false,

      rowsToBeAdded: [],
      dialogImportVisible: false,

      rules: {
        publisher_teacher_id: [{required: true, message: '必须填写出题人工号', trigger: 'change'}],
        stem: [{required: true, message: '必须填写题干', trigger: 'change'}],
        choices: [{required: true, message: '必须填写选项', trigger: 'change'}]
      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getQuestions('mcq', this.listQuery).then(body => {
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
        publisher_teacher_id: '',
        stem: '',
        choices: ''
      }
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
          createQuestions('mcq', [this.temp]).then(() => {
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
      this.temp = Object.assign({}, row) // copy obj
      this.dialogStatus = 'update'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          updateQuestion('mcq', tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.temp.id)
            this.list.splice(index, 1, this.temp)
            this.dialogFormVisible = false
            this.$message({
              message: '修改已保存',
              showClose: true,
              type: 'success'
            })
          })
        }
      })
    },
    handleMultiDelete() {
      let rows = this.$refs.mcqTable.selection;
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
      this.dialogDeleteVisible = true
    },
    deleteData() {
      console.log(this.rowsToBeDeleted)
      deleteQuestions('mcq', this.rowsToBeDeleted.map(v => v.id)).then(() => {
        this.dialogDeleteVisible = false
        this.$message({
          message: '删除成功',
          showClose: true,
          type: 'success'
        })
        this.getList()
      })
    },
    handleDownload() {
      this.downloadLoading = true
      import('@/utils/Export2Excel').then(excel => {
        const tHeader = ['publisher_teacher_id', 'stem', 'choices']
        let queryAll = Object.assign({}, this.listQuery, {page_index: 1, page_size: Math.pow(2, 32)})
        getQuestions(queryAll).then(body => {
          excel.export_json_to_excel({
            header: tHeader,
            data: body.data.map(mcq => tHeader.map(k => mcq[k])),
            filename: 'mcqs'
          })
          this.downloadLoading = false
        })
      })
    },
    handleUpload(e) {
      e.stopPropagation()
      e.preventDefault()
      let file = e.target.files[0]
      // clear input file so that this event can be triggered for multiple times
      e.target.value = ''
      import("@/utils/ImportFromExcel").then(xlsx => {
        xlsx.extract_from_excel(file).then(mcqs => {
          this.rowsToBeAdded = mcqs
          this.dialogImportVisible = true
        })
      })
    },
    uploadData() {
      console.log(this.rowsToBeAdded)
      createQuestions('mcq', this.rowsToBeAdded).then(() => {
        this.dialogImportVisible = false
        this.$message({
          message: '批量创建成功',
          showClose: true,
          type: 'success'
        })
        this.getList()
      })
    }
  }
}
</script>
