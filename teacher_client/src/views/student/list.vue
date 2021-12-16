<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.student_id" class="filter-item" placeholder="学号" style="width: 200px;"
                @keyup.enter.native="handleFilter"/>
      <el-input v-model="listQuery.name" class="filter-item" placeholder="姓名" style="width: 200px;"
                @keyup.enter.native="handleFilter"/>
      <el-input v-model="listQuery.class_id" class="filter-item" placeholder="班号" style="width: 200px;"
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
      ref="studentTable"
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
      <el-table-column align="center" label="学号">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.student_id }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="姓名">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="班号">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.class_id }}</span>
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
                @pagination="getList"/>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="dataForm" :model="temp" :rules="rules" label-position="left" label-width="100px"
               style="width: 400px; margin-left:50px;">
        <el-form-item label="学号" prop="student_id">
          <el-input v-model="temp.student_id"/>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item label="班号" prop="class_id">
          <el-input v-model="temp.class_id"/>
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
        <el-table-column label="ID" property="id" width="150"></el-table-column>
        <el-table-column label="学号" property="student_id" width="150"></el-table-column>
        <el-table-column label="姓名" property="name" width="200"></el-table-column>
        <el-table-column label="班号" property="class_id" width="150"></el-table-column>
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
        <el-table-column label="学号" property="student_id" width="150"></el-table-column>
        <el-table-column label="姓名" property="name" width="200"></el-table-column>
        <el-table-column label="班号" property="class_id" width="150"></el-table-column>
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
import {createStudents, deleteStudents, getStudents, updateStudent} from '@/api/student'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination'

export default {
  name: 'StudentList',
  components: {Pagination},
  directives: {waves},
  data() {
    return {
      tableKey: 0,

      list: null,
      total: 0,

      listLoading: true,
      listQuery: {
        student_id: '',
        name: '',
        class_id: '',
        page_index: 1,
        page_size: 20
      },

      temp: {
        id: undefined,
        student_id: '',
        name: '',
        class_id: ''
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
        student_id: [{required: true, message: '必须填写学号', trigger: 'change'}],
        name: [{required: true, message: '必须填写姓名', trigger: 'change'}],
        class_id: [{required: true, message: '必须填写班号', trigger: 'change'}]
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
      getStudents(this.listQuery).then(body => {
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
        student_id: '',
        name: '',
        class_id: ''
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
          createStudents([this.temp]).then(() => {
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
          updateStudent(tempData).then(() => {
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
      let rows = this.$refs.studentTable.selection;
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
      deleteStudents(this.rowsToBeDeleted.map(v => v.id)).then(() => {
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
        const tHeader = ['student_id', 'name', 'class_id']
        let queryAll = Object.assign({}, this.listQuery, {page_index: 1, page_size: Math.pow(2, 32)})
        getStudents(queryAll).then(body => {
          excel.export_json_to_excel({
            header: tHeader,
            data: body.data.map(student => tHeader.map(k => student[k])),
            filename: 'students'
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
        xlsx.extract_from_excel(file).then(students => {
          this.rowsToBeAdded = students
          this.dialogImportVisible = true
        })
      })
    },
    uploadData() {
      console.log(this.rowsToBeAdded)
      createStudents(this.rowsToBeAdded).then(() => {
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
