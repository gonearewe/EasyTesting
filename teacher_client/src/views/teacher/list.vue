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
      <el-button v-waves :loading="downloadLoading" class="filter-item" type="primary" icon="el-icon-download"
                 @click="handleDownload">
        导出
      </el-button>
    </div>

    <el-table
      ref = "teacherTable"
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      stripe
      fit
      show-header
      style="width: 100%;"
      highlight-current-row
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
      <el-table-column label="Actions" align="center" width="230" class-name="small-padding fixed-width">
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

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
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

    <el-dialog title="确认删除" :visible.sync="dialogDeleteVisible">
      <el-table :data="rowsToBeDeleted">
        <el-table-column property="id" label="ID" width="150"></el-table-column>
        <el-table-column property="teacher_id" label="工号" width="150"></el-table-column>
        <el-table-column property="name" label="姓名" width="200"></el-table-column>
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
import {getTeachers, createTeachers, updateTeacher, deleteTeachers} from '@/api/teacher'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

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
            this.getList()
            this.dialogFormVisible = false
            this.$message({
              message: '创建成功',
              showClose: true,
              type: 'success'
            })
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
    deleteData(){
      // TODO
      console.log(this.rowsToBeDeleted)
      this.dialogDeleteVisible = false
    },
    handleDownload() {
      this.downloadLoading = true
      import('@/vendor/Export2Excel').then(excel => {
        const tHeader = ['timestamp', 'title', 'type', 'importance', 'status']
        const filterVal = ['timestamp', 'title', 'type', 'importance', 'status']
        const data = this.formatJson(filterVal)
        excel.export_json_to_excel({
          header: tHeader,
          data,
          filename: 'table-list'
        })
        this.downloadLoading = false
      })
    },
    formatJson(filterVal) {
      return this.list.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    }
  }
}
</script>


<!--<template>-->
<!--  <div class="app-container">-->
<!--    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">-->
<!--      <el-table-column align="center" label="ID" width="80">-->
<!--        <template slot-scope="scope">-->
<!--          <span>{{ scope.row.id }}</span>-->
<!--        </template>-->
<!--      </el-table-column>-->

<!--      <el-table-column width="180px" align="center" label="工号">-->
<!--        <template slot-scope="scope">-->
<!--          <span>{{ scope.row.teacher_id }}</span>-->
<!--        </template>-->
<!--      </el-table-column>-->

<!--      <el-table-column width="120px" align="center" label="姓名">-->
<!--        <template slot-scope="scope">-->
<!--          <span>{{ scope.row.name }}</span>-->
<!--        </template>-->
<!--      </el-table-column>-->

<!--      <el-table-column width="100px" label="是管理员">-->
<!--        <template slot-scope="scope">-->
<!--          <el-button v-if="scope.row.is_admin" type="success" round>是</el-button>-->
<!--          <el-button v-else type="info" round>否</el-button>-->
<!--        </template>-->
<!--      </el-table-column>-->

<!--      <el-table-column align="center" label="操作" width="120">-->
<!--        <template slot-scope="{row}">-->
<!--          <el-row>-->
<!--            <router-link :to="'/teacher/edit/'+row.id">-->
<!--              <el-button type="primary" size="small" icon="el-icon-edit">-->
<!--                编辑-->
<!--              </el-button>-->
<!--            </router-link>-->
<!--            <el-button type="danger" size="small" icon="el-icon-delete" @click="delete(row.id)">-->
<!--              删除-->
<!--            </el-button>-->
<!--          </el-row>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--    </el-table>-->

<!--    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page_index" :limit.sync="listQuery.page_size"-->
<!--                @pagination="getList"/>-->
<!--  </div>-->
<!--</template>-->

<!--<script>-->
<!--import {getTeachers} from '@/api/teacher'-->
<!--import Pagination from '@/components/Pagination' // Secondary package based on el-pagination-->

<!--export default {-->
<!--  name: 'TeacherList',-->
<!--  components: {Pagination},-->
<!--  filters: {-->
<!--    statusFilter(status) {-->
<!--      const statusMap = {-->
<!--        published: 'success',-->
<!--        draft: 'info',-->
<!--        deleted: 'danger'-->
<!--      }-->
<!--      return statusMap[status]-->
<!--    }-->
<!--  },-->
<!--  data() {-->
<!--    return {-->
<!--      list: null,-->
<!--      total: 0,-->
<!--      listLoading: true,-->
<!--      listQuery: {-->
<!--        page_index: 1,-->
<!--        page_size: 20-->
<!--      }-->
<!--    }-->
<!--  },-->
<!--  created() {-->
<!--    this.getList()-->
<!--  },-->
<!--  methods: {-->
<!--    getList() {-->
<!--      this.listLoading = true-->
<!--      getTeachers(this.listQuery).then(body => {-->
<!--        this.list = body.data-->
<!--        this.total = body.total-->
<!--        this.listLoading = false-->
<!--      })-->
<!--    },-->
<!--    delete(ids) {-->
<!--      console.log(ids)-->
<!--      console.log(this)-->
<!--      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {-->
<!--        confirmButtonText: '确定',-->
<!--        cancelButtonText: '取消',-->
<!--        type: 'warning'-->
<!--      }).then(() => {-->
<!--        this.$message({-->
<!--          type: 'success',-->
<!--          message: '删除成功!'-->
<!--        });-->
<!--      }).catch(() => {-->
<!--        this.$message({-->
<!--          type: 'info',-->
<!--          message: '已取消删除'-->
<!--        });-->
<!--      })-->
<!--    }-->
<!--  }-->
<!--}-->
<!--</script>-->

<!--<style scoped>-->
<!--.edit-input {-->
<!--  padding-right: 100px;-->
<!--}-->

<!--.cancel-btn {-->
<!--  position: absolute;-->
<!--  right: 15px;-->
<!--  top: 10px;-->
<!--}-->
<!--</style>-->
