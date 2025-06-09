<template>
    <div>
        <div class="card" style="margin-left:5px;">
             <el-select v-model="value" clearable placeholder="选择用户" style="width: 240px"  @change="contactLoad">
                <el-option v-for="item in data.tableData" :key="item.value" :label="item.label" :value="item.name" />
            </el-select> 
        </div>
        <div class="card" style="margin-bottom: 5px">

        </div>
        <div class="card">
            <el-table :data="data.contactList" @selection-change="" stripe>
                <el-table-column type="selection" width="55" />
                <el-table-column prop="telegram_id" label="电报ID" show-overflow-tooltip />
                <el-table-column prop="first_name" label="姓氏" show-overflow-tooltip />
                <el-table-column prop="last_name" label="名称" show-overflow-tooltip />
                <el-table-column prop="username" label="用户名" show-overflow-tooltip />
                
                <el-table-column prop="deptName" label="操作">
                    <template #default="scope">
                        <el-button @click="handleUpdata(scope.row)" type="primary" :icon="Edit" circle></el-button>
                        <el-button @click="del(scope.row.phone)" type="danger" :icon="Delete" circle></el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div style="margin-top: 15px">
                <el-pagination @change="contactLoad(value)" v-model:current-page="data.pageNum" v-model:page-size="data.pageSize"
                    :page-sizes="[5, 10, 15, 20]" background layout="total,sizes,prev,pager,next,jumper"
                    :total="data.total" />
            </div>
        </div>


    </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import request from "@/utils/request.js";
import { ElMessage } from "element-plus";
import { Edit, Delete, UploadFilled } from "@element-plus/icons-vue";
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
const value = ref(null)
const selectedItem = ref<any>(null)  
const data = reactive({
    name: null,
    tableData: [
    ],
    contactList:[],
    pageNum: 1,
    pageSize: 10,
    total: 0,
    form: {// 其他表单字段
        username: userStore.userInfo.nickName,
        // 存储上传的文件
        session_files: []
    },
    personal_details: {},
    individual_grouping: {},
    listGroupId: null,
    formGroup: false,
    Ids: [],
    rules: {
        username: [
            { required: true, message: '请输入账号', trigger: 'blur' },
        ],
        name: [
            { required: true, message: '请输入姓名', trigger: 'blur' },
            { min: 2, max: 5, message: '长度在 2 到 5 个字符', trigger: 'blur' }
        ],
        no: [
            { required: true, message: '请输入工号', trigger: 'blur' }
        ]
    },
    titleName: '添加账号',
    deptList: [],


}
)
const contactLoad = (val: string) => {
    const item = data.tableData.find(o => o.name === val)
    selectedItem.value = item
    request.post('/api/get_contacts', {
        page: data.pageNum,
        page_size: data.pageSize,
        phone: item.phone
    }).then(res => {

      
 
        data.contactList = res.data.data.contacts;
        data.total = res.data.data.count;
        console.log("列表数据", res);

    }).catch(err => {
        console.error('接口异常：', err);
        ElMessage.error('服务器错误，请稍后重试');
    });

}
const load = () => {
    console.log("请求参数", { page: data.pageNum, page_size: data.pageSize, master_username: data.form.username })

    request.post('/api/get_session', {
        page: 1,
        page_size: 1000000000,
        master_username: data.form.username
    }).then(res => {


        data.tableData = res.data.session;


        console.log("列表数据", res);

    }).catch(err => {
        console.error('接口异常：', err);
        ElMessage.error('服务器错误，请稍后重试');
    });
}
onMounted(() => {
    console.log('调用 load 前的用户名:', data.form.username);

    load()
})  
</script>

<style scoped>
.upload-demo :deep(.el-upload-dragger) {
    padding: 0 !important
}

.el-upload__text {
    font-size: 10px;
}
</style>