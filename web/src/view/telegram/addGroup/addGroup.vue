<template>
    <div>
        <div class="card" style="margin-bottom: 5px">
            <el-input style="width: 240px; margin-right: 10px" v-model="data.groupName" placeholder="请输入名称搜索"
                prefix-icon="Search"></el-input>
            <el-button type="primary" @click="load">查 询</el-button>
            <el-button type="warning" @click="reset">重 置</el-button>
        </div>
        <div class="card" style="margin-bottom: 5px">
            <el-button type="danger" @click="handleAdd">
                <Plus style="width: 1em; height: 1em; margin-right: 2px" />添加链接
            </el-button>
            <el-button type="info" @click="delAll">
                <Share style="width: 1em; height: 1em; margin-right: 2px" />极速加群
            </el-button>
        </div>
        <div class="card">
            <el-table :data="data.tableData" @selection-change="handleSelectionChange" stripe>
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" label="文件夹标题" show-overflow-tooltip />
                <el-table-column prop="id" label="邀请链接" show-overflow-tooltip />
                <el-table-column prop="id" label="群组数量" show-overflow-tooltip />
                <el-table-column prop="id" label="分组ID" show-overflow-tooltip />
                <el-table-column prop="groupName" label="分组名" show-overflow-tooltip />
                <el-table-column label="操作">
                    <template #default="scope">
                        <el-button @click="handleUpdata(scope.row)" type="primary" :icon="Edit" circle></el-button>
                        <el-button @click="delList(scope.row.id)" type="danger" :icon="Delete" circle></el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div style="margin-top: 15px">
                <el-pagination @change="load" v-model:current-page="data.pageNum" v-model:page-size="data.pageSize"
                    :page-sizes="[5, 10, 15, 20]" background layout="total,sizes,prev,pager,next,jumper"
                    :key="data.total" :total="data.total" />
            </div>
        </div>
        <el-dialog :title="data.titleName" v-model="data.formVisible" width="500" destroy-on-close>
            <br>
            <el-form ref="formRef" :model="data.form" :rules="formRules">
                <el-form-item label="分组名" prop="group">
                    <el-input v-model="data.form.group" autocomplete="off" placeholder="请输入分组名(必填)"
                        style="width:400px" />
                </el-form-item>
            </el-form>
            <br>
            <template #footer>
                <div class="dialog-footer">
                    <el-button type="warning" @click="validateForm">确认</el-button>
                    <el-button type="primary" @click="data.formVisible = false">取消</el-button>
                </div>
            </template>
        </el-dialog>
        <el-dialog :title="data.titleName" v-model="data.personalData" width="500" destroy-on-close>
            <el-form-item label="分组名:" prop="groupName">
                <el-input v-model="data.form.group" autocomplete="off" placeholder="请输入分组名" style="width:400px" />
            </el-form-item>
            <!-- <el-form-item label="年龄">
                <el-input-number v-model="data.form.age" :min="18" autocomplete="off" style="width:150px" />
            </el-form-item> -->
            <template #footer>
                <div class="dialog-footer">
                    <el-button type="warning" @click="updateGroup">确认</el-button>
                    <el-button type="primary" @click="data.personalData = false">取消</el-button>
                </div>
            </template>
        </el-dialog>


    </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import { Edit, Delete, UploadFilled, Plus } from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { useUserStore } from '@/pinia/modules/user'
import request from "@/utils/request.js";
// 文件列表（Element Plus 用来显示）
const userStore = useUserStore()

const formRef = ref(null);
const data = reactive({
    name: null,
    tableData: [
    ],
    pageNum: 1,
    pageSize: 10,
    total: 0,
    formVisible: false,
    personalData: false,
    form: {

    },
    Ids: [],
    titleName: '添加分组',
    deptList: [],
    groupName: '',
    userName: '',
})
//多选所有用户id
const handleSelectionChange = (rows) => {
    console.log(rows.map(item => item.id))
    data.Ids = rows.map(item => item.id)
}
const delAll = () => {
    del()
}
const delList = (id) => {
    data.Ids.push(id)
    del()
}
const reset = () => {
  data.groupName = null
  load()
}
//删除分组
const del = () => {
    ElMessageBox.confirm('删除后无法恢复，确定删除吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        request.post('/api/delete_telegram_group',
            { ids: data.Ids }).then(res => {
                data.Ids = []
                ElMessage.success('删除成功')
                load()
            }).catch(() => {
                data.Ids = []
                ElMessage.error('删除失败')
            })
    }).catch(() => {
        data.Ids = []
        ElMessage({
            type: 'info',
            message: '已取消删除'
        });
    });

}
// 验证规则
const formRules = {
    group: [
        { required: true, message: '分组名不能为空', trigger: 'blur' }
    ]
};
// 验证方法
const validateForm = () => {
    formRef.value.validate((valid) => {
        if (valid) {
            save(); // 验证通过后执行保存
        }
    });
};
const handleAdd = () => {
    data.formVisible = true
}
const save = () => {
    add()
}
const add = () => {
    const groupName = data.form.group;
    const userName = userStore.userInfo.nickName;
    request.post('/api/create_telegram_group', {
        groupName, userName
    }).then(res => {



        data.formVisible = false;
        data.form = {}
        load()


    }).catch(err => {

    })

}
const load = () => {
    console.log("请求参数", { page: data.pageNum, page_size: data.pageSize, user_name: userStore.userInfo.nickName })

   request.post('/api/list_telegram_group', {
        page: data.pageNum,
        page_size: data.pageSize,
        user_name: userStore.userInfo.nickName,
        group_name: data.groupName
    }).then(res => {


        data.tableData = res.data.list;
        data.total = res.data.total;

        // console.log("列表数据", res);

    }).catch(err => {
        console.error('接口异常：', err);
        ElMessage.error('服务器错误，请稍后重试');
    });
}
const handleUpdata = (row) => {
    data.titleName = '修改分组信息'
    data.form = JSON.parse(JSON.stringify(row))
    data.personalData = true

}
const updateGroup = () => {
    const groupName = data.form.groupName;
    const id = data.form.id;
    const username = userStore.userInfo.nickName;
    console.log("修改前数据", groupName, id, username);
    request.post('/api/update_telegram_group', {
        groupName, id, username
    }).then(res => {
        data.personalData = false;
        data.form = {}
        load();
    }).catch(err => {
        console.log(err)
    })
}
onMounted(() => {
    console.log("加载数据");
    load();
});
</script>

<style scoped>
.upload-demo :deep(.el-upload-dragger) {
    padding: 0 !important
}

.el-upload__text {
    font-size: 10px;
}
</style>