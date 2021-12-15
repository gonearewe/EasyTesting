<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.teacher_id" placeholder="工号" style="width: 200px;" class="filter-item"
                @keyup.enter.native="handleFilter"/>
      <el-input v-model="listQuery.name" placeholder="姓名" style="width: 200px;" class="filter-item"
                @keyup.enter.native="handleFilter"/>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        搜索
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit"
                 @click="handleCreate">
        添加
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="danger" icon="el-icon-edit"
                 @click="handleMultiDelete">
        删除
      </el-button>

      <el-button v-waves class="filter-item" icon="el-icon-upload" type="primary"
                 @click="$refs.file_picker.click()">
        导入
      </el-button>


      <el-button v-waves :loading="downloadLoading" class="filter-item" icon="el-icon-download" type="success"
                 @click="handleDownload">
        导出
      </el-button>
    </div>

    <input ref="file_picker" accept="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" hidden
           type='file' @change="handleUpload"/>

    <el-table
      ref="teacherTable"
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      stripe
      fit
      show-header
      style="width: 100%;"
    >

      <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column label="ID" prop="id" align="center" width="80">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="工号" align="center">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.teacher_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="姓名" align="center">
        <template slot-scope="{row}">
          <span class="link-type" @click="handleUpdate(row)">{{ row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="是管理员？" class-name="status-col" width="100" align="center">
        <template slot-scope="{row}">
          <el-tag :type="row.is_admin | boolColorFilter">
            {{ row.is_admin | boolFilter }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" class-name="small-padding fixed-width" label="操作" width="230">
        <template slot-scope="{row}">
          <el-button type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button size="mini" type="danger" @click="handleDelete(row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page_index" :limit.sync="listQuery.page_size"
                @pagination="getList"/>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="30%">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="100px"
               style="width: 400px; margin-left:50px;">
        <el-form-item label="工号" prop="teacher_id">
          <el-input v-model="temp.teacher_id"/>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input v-model="temp.name"/>
        </el-form-item>
        <el-form-item label="管理员权限" prop="is_admin">
          <el-checkbox v-model="temp.is_admin" border label="授予"/>
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
      <el-table :data="rowsToBeDeleted">
        <el-table-column property="id" label="ID" width="150"></el-table-column>
        <el-table-column property="teacher_id" label="工号" width="150"></el-table-column>
        <el-table-column property="name" label="姓名" width="200"></el-table-column>
        <el-table-column class-name="status-col" label="是管理员？" width="100">
          <template slot-scope="{row}">
            <el-tag :type="row.is_admin | boolColorFilter">
              {{ row.is_admin | boolFilter }}
            </el-tag>
          </template>
        </el-table-column>
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

    <el-dialog :visible.sync="dialogImportVisible" width="30%">
      <el-upload
        :auto-upload="false"
        :http-request="uploadData"
        :limit="1"
        :show-file-list="false"
        accept=".xlsx"
        action=""
        class="upload-demo"
        drag
      >
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div slot="tip" class="el-upload__tip">只能上传 xlsx 文件</div>
      </el-upload>
    </el-dialog>

    <el-dialog :visible.sync="dialogPvVisible" title="Reading statistics">
      <el-table :data="pvData" border fit highlight-current-row style="width: 100%">
        <el-table-column prop="key" label="Channel"/>
        <el-table-column prop="pv" label="Pv"/>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogPvVisible = false">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {createTeachers, deleteTeachers, getTeachers, updateTeacher} from '@/api/teacher'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination'

export default {
  name: 'TeacherList',
  components: {Pagination},
  directives: {waves},
  filters: {
    boolFilter(bool) {
      return {true: '是', false: '否'}[bool]
    },
    boolColorFilter(bool) {
      return {true: 'success', false: 'info'}[bool]
    }
  },
  data() {
    return {
      tableKey: 0,

      list: null,
      total: 0,

      listLoading: true,
      listQuery: {
        teacher_id: '',
        name: '',
        page_index: 1,
        page_size: 20
      },

      temp: {
        id: undefined,
        teacher_id: '',
        name: '',
        is_admin: false
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: '编辑',
        create: '创建'
      },

      rowsToBeDeleted: [],
      dialogDeleteVisible: false,

      dialogImportVisible: false,

      dialogPvVisible: false,
      pvData: [],
      rules: {
        type: [{required: true, message: 'type is required', trigger: 'change'}],
        timestamp: [{type: 'date', required: true, message: 'timestamp is required', trigger: 'change'}],
        title: [{required: true, message: 'title is required', trigger: 'blur'}]
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
      getTeachers(this.listQuery).then(body => {
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
        teacher_id: '',
        name: '',
        is_admin: false
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
          createTeachers([this.temp]).then(() => {
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
          updateTeacher(tempData).then(() => {
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
      let rows = this.$refs.teacherTable.selection;
      this.rowsToBeDeleted = Object.assign([], rows)
      this.dialogDeleteVisible = true
    },
    handleDelete(row) {
      this.rowsToBeDeleted = []
      this.rowsToBeDeleted[0] = Object.assign({}, row)
      this.dialogDeleteVisible = true
    },
    handleFetchPv(pv) {
      fetchPv(pv).then(response => {
        this.pvData = response.data.pvData
        this.dialogPvVisible = true
      })
    },
    deleteData() {
      console.log(this.rowsToBeDeleted)
      deleteTeachers(this.rowsToBeDeleted.map(v => v.id)).then(() => {
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
        const tHeader = ['teacher_id', 'name', 'is_admin']
        let queryAll = Object.assign({}, this.listQuery, {page_index: 1, page_size: Math.pow(2, 32)})
        getTeachers(queryAll).then(body => {
          excel.export_json_to_excel({
            header: tHeader,
            data: body.data.map(teacher => tHeader.map(k => teacher[k])),
            filename: 'teachers'
          })
          this.downloadLoading = false
        })
      })
    },
    handleUpload(e) {
      e.stopPropagation()
      e.preventDefault()
      let file = e.target.files[0]
      import("@/utils/ImportFromExcel").then(xlsx => {
        xlsx.extract_from_excel(file).then(teachers => {
          console.log(teachers)
          createTeachers(teachers)
        })
      })
    },
    uploadData(files) {
      console.log(files)
    }

  }
}
</script>
