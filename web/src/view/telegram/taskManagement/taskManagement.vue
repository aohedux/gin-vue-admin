<template>
    <div>
        <div class="card" style="margin-bottom: 5px">
            <el-input style="width: 240px; margin-right: 10px" v-model="data.name" placeholder="请输入名称搜索"
                prefix-icon="Search"></el-input>
            <el-select style="width: 15%" v-model="data.form.userName" placeholder="请选择分组名搜索"
                @change="handleGroupChange">
                <el-option v-for="group in data.tableDataGroup" :key="group.id" :label="group.label"
                    :value="group.groupName" />
            </el-select>
            <el-button type="primary" @click="load">查 询</el-button>
            <el-button type="warning" @click="reset">重 置</el-button>
        </div>
        <div class="card" style="margin-bottom: 5px">
            <el-button type="primary" @click="handleAdd">
                <ChatDotRound style="width: 1em; height: 1em; margin-right: 3px" />批量发送
            </el-button>
            <el-button type="primary" @click="delAll">
                <Link style="width: 1em; height: 1em; margin-right: 3px" />批量转发
            </el-button>
            <el-button type="primary" @click="groupAll">
                <ChatDotRound style="width: 1em; height: 1em; margin-right: 3px" />批量私聊
            </el-button>
            <el-button type="primary" @click="batchUpdateOnline(true)">
                <VideoPlay style="width: 1em; height: 1em; margin-right: 3px" />启动
            </el-button>
            <el-button type="primary" @click="batchUpdateOnline(false)">
                <VideoPause style="width: 1em; height: 1em; margin-right: 3px" />暂停
            </el-button>
            <el-button type="danger" plain @click="batchUpdateOnline(false)">
                <Delete style="width: 1em; height: 1em; margin-right: 3px" />删除
            </el-button>
            <el-button type="primary" plain @click="batchUpdateOnline(false)">重新筛选</el-button>
        </div>
        <div class="card">
            <el-table :data="data.tableData" @selection-change="handleSelectionChange" stripe>
                <el-table-column type="selection" width="55" />
                <el-table-column prop="name" label="账号" show-overflow-tooltip />
                <el-table-column prop="phone" label="任务名称" show-overflow-tooltip />
                <el-table-column prop="username" label="任务状态" show-overflow-tooltip />
                <el-table-column prop="groupNumber" label="任目标数量" show-overflow-tooltip />
                <el-table-column prop="tag" label="成功数量" show-overflow-tooltip>
                    <template #default="scope">
                        {{ scope.row.tag }}
                    </template>
                </el-table-column>
                <el-table-column prop="is_premium" label="单个间隔" show-overflow-tooltip>
                    <template #default="{ row }">
                        {{ row.is_premium ? '是' : '否' }}
                    </template>
                </el-table-column>
                <el-table-column prop="online" label="预计耗时">
                    <template #default="{ row }">
                        <el-switch v-model="row.online" inline-prompt :active-value="true" :inactive-value="false"
                            active-text="在线" inactive-text="离线" @change="handleOnlineChange(row)"
                            show-overflow-tooltip />
                    </template>
                </el-table-column>
                <el-table-column prop="is_blocked" label="循环间隔">
                    <template #default="{ row }">
                        {{ {
                            true: '🔴 异常',
                            false: '🟢 正常'
                        }[row.is_blocked] || '⚪ 状态异常' }}
                    </template>
                </el-table-column>
                <el-table-column prop="server" label="已循环" show-overflow-tooltip />
                <el-table-column prop="proxy" label="代理IP" show-overflow-tooltip />
                <el-table-column prop="loginMethod" label="创建时间" show-overflow-tooltip>
                    <template #default="{ row }">
                        协议号
                    </template>
                </el-table-column>
                <el-table-column prop="created_at" label="创建时间" show-overflow-tooltip />
                <el-table-column prop="deptName" label="操作">
                    <template #default="scope">
                        <el-button @click="handleUpdata(scope.row)" type="primary" :icon="Edit" circle></el-button>
                        <el-button @click="del(scope.row.phone)" type="danger" :icon="Delete" circle></el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div style="margin-top: 15px">
                <el-pagination @change="load" v-model:current-page="data.pageNum" v-model:page-size="data.pageSize"
                    :page-sizes="[5, 10, 15, 20,100]" background layout="total,sizes,prev,pager,next,jumper"
                    :total="data.total" />
            </div>
        </div>




        <!-- 批量发送  模态框 -->
        <el-drawer v-model="data.formVisible" :title="data.titleName" direction="rtl" size="50%">
            <el-form-item label="账号" required>
                <el-select v-model="data.form.userName" multiple placeholder="请选择账号">
                    <el-option v-for="group in data.tableDataGroup" :key="group.id" :label="group.label"
                        :value="group.groupName" />
                </el-select>
            </el-form-item>

            <el-form-item label="发送目标" required>
                <el-radio-group v-model="data.target">
                    <el-radio label="用户">用户</el-radio>
                    <el-radio label="群组">群组</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="消息图片">
                <el-upload class="upload-demo" action="#" :auto-upload="false" :show-file-list="true"
                    list-type="picture-card">
                    <el-icon>
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>

   

                <el-form-item label="消息内容" required>
                    <div v-for="(msg, index) in data.messages" :key="index"
                        style="display: flex; align-items: flex-start; margin-bottom: 12px;">
                        <el-input v-model="data.messages[index]" type="textarea" :rows="3" placeholder="请输入内容"
                            maxlength="4090" show-word-limit style="flex: 1;" />
                        <el-button type="danger" icon="Delete" size="small" @click="removeMessage(index)"
                            style="margin-left: 10px;" :disabled="data.messages.length === 1">
                            删除
                        </el-button>
                    </div>

                    <el-button type="primary" plain size="small" @click="addMessage">
                        添加一条消息
                    </el-button>
                </el-form-item>

            <el-form-item label="链接预览" required>
                <el-switch v-model="data.preview" />
            </el-form-item>

            <el-form-item label="消息尾缀" required>
                <el-radio-group v-model="data.suffix">
                    <el-radio label="表情">表情</el-radio>
                    <el-radio label="字母">字母</el-radio>
                    <el-radio label="关闭">关闭</el-radio>
                </el-radio-group>
            </el-form-item>

            <el-form-item label="发送方式" required>
                <el-radio-group v-model="data.sendType">
                    <el-radio label="单次发送">单次发送</el-radio>
                    <el-radio label="循环发送">循环发送</el-radio>
                </el-radio-group>
            </el-form-item>

            <el-form-item label="单个间隔">
                <el-input-number v-model="data.intervalMin" :min="0" /> ~
                <el-input-number v-model="data.intervalMax" :min="1" />
                <span style="margin-left: 10px;">秒</span>
            </el-form-item>

            <el-form-item label="循环间隔">
                <el-input-number v-model="data.loopMin" :min="1" /> ~
                <el-input-number v-model="data.loopMax" :min="data.loopMin + 1" />
                <span style="margin-left: 10px;">分钟</span>
            </el-form-item>

            <el-form-item label="休息时间段">
                <el-date-picker v-model="data.restTime" type="datetimerange" start-placeholder="开始时间"
                    end-placeholder="结束时间" style="width: 100%;" />
            </el-form-item>

            <el-form-item label="停止日期">
                <el-date-picker v-model="data.stopDate" type="date" placeholder="请选择日期" />
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="submit">提交</el-button>
                <el-button @click="cancel">取消</el-button>
            </el-form-item>
        </el-drawer>

        <!-- 批量发送  模态框 结束 -->

        <el-dialog v-model="data.personalData" width="500" destroy-on-close>

            <el-tabs v-model="activeTab">
                <!-- 页面1 -->
                <el-tab-pane label="编辑电报个人信息" name="page1">
                    <br>
                    <div style="display: flex; align-items: center; margin-bottom: 10px;">
                        <span style="width: 60px; text-align: right; margin-right: 10px;">用户名:</span>
                        <el-input v-model="data.personal_details" :value="data.form.username" placeholder="用户名"
                            style="width: 80%;" />
                    </div>
                    <div style="display: flex; align-items: center; margin-bottom: 10px;">
                        <span style="width: 60px; text-align: right; margin-right: 10px;">姓氏:</span>
                        <el-input v-model="data.personal_details" value="" placeholder="姓氏" style="width: 80%;" />
                    </div>
                    <div style="display: flex; align-items: center; margin-bottom: 10px;">
                        <span style="width: 60px; text-align: right; margin-right: 10px;">名称:</span>
                        <el-input v-model="data.personal_details" value="" placeholder="名称" style="width: 80%;" />
                    </div>
                    <div style="display: flex; align-items: center; margin-bottom: 10px;">
                        <span style="width: 60px; text-align: right; margin-right: 10px;">个人简介:</span>
                        <el-input v-model="data.personal_details" value="" placeholder="个人简介" style="width: 80%;" />
                    </div>

                </el-tab-pane>

                <!-- 页面2 -->
                <el-tab-pane label="修改电报分组" name="page2">
                    <br>
                    <el-input v-model="data.individual_grouping" value="" placeholder="年龄" />
                </el-tab-pane>

                <!-- 页面3 -->
                <el-tab-pane label="修改默认" name="page3">
                    <br>
                    <el-input v-model="data.form.address" placeholder="地址" />
                </el-tab-pane>
            </el-tabs>

            <template #footer>
                <div class="dialog-footer">
                    <el-button type="warning" @click="savePersonalInformation">确认</el-button>
                    <el-button type="primary" @click="data.personalData = false">取消</el-button>
                </div>
            </template>
        </el-dialog>
        <el-dialog :title="data.titleName" v-model="data.formGroup" width="500" destroy-on-close>

            <br>
            分组: <el-select style="width: 50%" v-model="data.form.groupName_list" placeholder="请选择分组名搜索"
                @change="handleGroupList">
                <el-option v-for="group in data.tableDataGroup" :key="group.id" :label="group.label"
                    :value="group.groupName" />
            </el-select>
            <template #footer>
                <div class="dialog-footer">
                    <el-button type="warning" @click="setAllGroup">确认</el-button>
                    <el-button type="primary" @click="data.formGroup = false">取消</el-button>
                </div>
            </template>
        </el-dialog>

    </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import request from "@/utils/request.js";
import { ElMessage, ElMessageBox } from "element-plus";
import { Edit, Delete, UploadFilled } from "@element-plus/icons-vue";
import type { UploadInstance } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

const activeTab = ref('page1')
// 文件列表（Element Plus 用来显示）
const userStore = useUserStore()
const uploadFileList = ref([]);

const uploadRef = ref<UploadInstance>()
const fileRawList = ref<File[]>([]);

const data = reactive({
    name: null,
    tableData: [
    ],
    tableDataGroup: [

    ],
    selectedGroupId: null,
    pageNum: 1,
    pageSize: 10,
    total: 0,
    formVisible: false,
    personalData: false,
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

    //测试
    accounts: [],
    target: '群组',
    messages: [''], // 默认有一条
    preview: true,
    suffix: '字母',
    sendType: '单次发送',
    intervalMin: 0,
    intervalMax: 60,
    loopMin: 5,
    loopMax: 10,
    restTime: [],
    stopDate: null

}
)
//
const addMessage = () => {
  data.messages.push('')
}

const removeMessage = (index) => {
  if (data.messages.length > 1) {
    data.messages.splice(index, 1)
  }
}

const submit = () => {
    console.log('提交表单:', data.value)
}


const cancel = () => {
    console.log('取消操作')
}
//
function savePersonalInformation() {
    switch (activeTab.value) {
        case 'page1':
            console.log('保存个人信息:', data.personal_details)
            // 执行页面1的保存逻辑
            break

        case 'page2':
            console.log('保存分组信息:', data.individual_grouping)
            // 执行页面2的保存逻辑
            break

        case 'page3':
            console.log('保存地址信息:', data.form.address)
            // 执行页面3的保存逻辑
            break

        default:
            console.warn('未知页面')
    }

    // 可选：关闭对话框
    data.personalData = false
}
function handleGroupChange(groupName) {
    const group = data.tableDataGroup.find(g => g.groupName === groupName)
    data.selectedGroupId = group ? group.id : null
    console.log('当前分组ID：', data.selectedGroupId)
}

function handleGroupList(groupName) {
    const group = data.tableDataGroup.find(g => g.groupName === groupName)
    data.listGroupId = group ? group.id : null
    console.log('当前分组ID：', data.selectedGroupId)
} const handleOnlineChange = (row) => {
    console.log('当前行数据：', row)
    console.log('当前 online 状态：', row.online)
    const phone = [row.phone]


    if (row.online == 1) {

        console.log('登入')

        request.post("/api/sign_in", { phone: phone, proxy: "", server: row.server })
            .then(res => {
                load()

            })
        return
    }
    if (row.online == 0) {
        console.log('登出')
        console.log(row.phone)
        request.post("/api/sign_out", { phone: phone }).then(res => {
            load()

        })
    }

    //   // 如果你要发送请求更新这个状态：
    //   request.post('/api/update_online_status', { 
    //     id: row.id,
    //     online: row.online
    //   }).then(res => {
    //     ElMessage.success('更新成功')
    //   }).catch(err => {
    //     ElMessage.error('更新失败')
    //   })
}
// 每次添加文件时触发
const clearUploadFiles = () => {
    uploadRef.value?.clearFiles();      // 清空 el-upload 显示的文件
    fileRawList.value = [];             // 清空你保存的文件变量
    uploadFileList.value = [];          // 清空 UI 文件列表（显示用）
};
const handleFileChange = (uploadFile, uploadFiles) => {
    fileRawList.value = uploadFiles.map(f => f.raw); // 这里存储原始文件
    uploadFileList.value = uploadFiles; // 显示在列表中
    // 自动过滤无效文件
    console.log('上传文件', uploadFile, uploadFiles)

    uploadFileList.value = uploadFiles.filter(f => {
        const { isSession, isSizeValid } = validateSessionFile(f.raw)
        return isSession && isSizeValid
    })
    if (uploadFileList.value.length == 0) {
        ElMessage.error(uploadFile.name + '文件不符合要求')

    }

};
const loadGroupList = () => {
    console.log("请求参数", { page: data.pageNum, page_size: data.pageSize, user_name: userStore.userInfo.nickName })

    request.post('/api/list_telegram_group', {
        page: data.pageNum,
        page_size: 1000000,
        user_name: userStore.userInfo.nickName,
        group_name: data.groupName
    }).then(res => {


        data.tableDataGroup = res.data.list;
        data.total = res.data.total;

        // console.log("列表数据", res);

    }).catch(err => {
        console.error('接口异常：', err);
        ElMessage.error('服务器错误，请稍后重试');
    });
}

// request.get('/dept/selectAll').then(res => {
//     data.deptList = res.data
// })
// 文件类型验证方法
const validateSessionFile = (file) => {
    // 双重验证：文件名后缀 + MIME类型（可选）
    const isSession = [
        file.name.split('.').pop().toLowerCase() === 'session',      // 验证扩展名
        // file.type === 'application/octet-stream'                  // 可选验证MIME类型
    ].every(Boolean)

    // 文件大小验证（500KB限制）
    const isSizeValid = file.size / 1024 <= 500

    return { isSession, isSizeValid }
}
const beforeUpload = (file) => {
    const { isSession, isSizeValid } = validateSessionFile(file)

    console.log('文件类型', isSession, file)
    if (!isSession) {
        ElMessage.error(`文件 [${file.name}] 类型不符合要求，必须为 .session 文件`)
        console.log('文件类型不符合要求')
        return false
    }

    if (!isSizeValid) {
        ElMessage.error(`文件 [${file.name}] 大小超过500KB限制`)
        console.log('文件类型不符合要求1')
        return false
    }

    return true
};

const load = () => {
    console.log("请求参数", { page: data.pageNum, page_size: data.pageSize, master_username: userStore.userInfo.nickName })

    request.post('/api/get_session', {
        tag: data.selectedGroupId,
        name: data.name,
        page: data.pageNum,
        page_size: data.pageSize,
        master_username: userStore.userInfo.nickName

    }).then(res => {


        data.tableData = res.data.session;
        data.total = res.data.count;

        console.log("列表数据", res);

    }).catch(err => {
        console.error('接口异常：', err);
        ElMessage.error('服务器错误，请稍后重试');
    });
}
onMounted(() => {
    console.log('调用 load 前的用户名:', data.form.username);

    loadGroupList()
    load()
})
const reset = () => {
    ``
    data.form.userName = null
    data.name = null
    data.selectedGroupId = null
    load()
}
const handleAdd = () => {
    data.formVisible = true
    data.titleName = '发送消息'

}

const add = () => {
    data.titleName = '添加账号'
    const formData = new FormData()

    // 添加多个文件，字段名统一为 'session'（后端接收多文件字段名一致）
    fileRawList.value.forEach(file => {
        formData.append('session_files[]', file);
    });



    formData.append('master_username', data.form.username);
    console.log(data.form.username);
    // 打印 formData 的键值对
    formData.forEach((value, key) => {
        if (value instanceof File) {
            console.log(`key: ${key}, filename: ${value.name}, size: ${value.size}`);
        } else {
            console.log(`key: ${key}, value: ${value}`);
        }
    });
    console.log(formData);
    console.log(data.form);
    try {
        request.post('/api/import_session', formData, {
            headers: {
                'Content-Type': 'multipart/form-data'  // 必须设置请求头
            }
        }).then(res => {


            ElMessage.success('上传成功');
            clearUploadFiles();
            data.formVisible = false;

            load()

        })
    } catch (error) {
        ElMessage.error('上传失败');
        console.error('错误详情:', error);
        clearUploadFiles();
        data.formVisible = false
        // load()
        load()
    }

    data.formVisible = false
    // load()

}

const save = () => {

    add()
    // formRef.value.validate(valid => {
    //     if (valid)
    //         data.form.id ? updata() : add()
    // })

}

const del = (phone) => {
    const phonenumber = []
    phonenumber.push(phone)
    ElMessageBox.confirm('删除后无法恢复，确定删除吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        request.post('/api/delete_session',
            { phone: phonenumber }).then(res => {
                if (res.code == 0) {
                    ElMessage.success('删除成功')
                    load()
                } else {
                    ElMessage.error(res.msg)
                }
            })
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '已取消删除'
        });
    });

}
const handleUpdata = (row) => {
    data.form = JSON.parse(JSON.stringify(row))
    if (row.online == false) {
        ElMessage.error('账号未登入无法编辑')

        return
    }
    data.personalData = true
}
const handleSelectionChange = (rows) => {
    console.log(rows.map(item => item.phone))
    data.Ids = rows.map(item => item.phone)
}

const setAllGroup = () => {
    if (data.listGroupId == null) {
        ElMessage.error('请选择分组')
        return
    }

    request.post('/api/update_group', { phone: data.Ids, group_id: data.listGroupId }).then(res => {
        if (res.code == 0) {
            ElMessage.success('分组设置成功')
            load()
        } else {
            ElMessage.error(res.msg)
        }
    })



}
//批量分组
const groupAll = () => {
    if (data.Ids.length == 0) {
        ElMessage.error('请选择要分组的账号')
        return
    }
    data.form.groupName_list = null
    data.listGroupId = null
    data.formGroup = true



}
const delAll = () => {
    if (data.Ids.length == 0) {
        ElMessage.error('请选择要删除的数据')
        return
    }
    ElMessageBox.confirm('删除后无法恢复，确定删除吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
        request.post('/api/delete_session', { phone: data.Ids }).then(res => {
            if (res.code == 0) {
                ElMessage.success('删除成功')
                load()
            } else {
                ElMessage.error(res.msg)
            }
        })
        console.log("jinru")
    }).catch()
}


const batchUpdateOnline = (newStatus) => {
    if (data.Ids.length == 0) {
        ElMessage.error('请选择要上线的账号')
        return
    }
    // 前端先更新 UI 表格中的 online 状态

    const phonenumber = []
    phonenumber.push(...data.Ids)
    console.log("批量上线", phonenumber)
    if (newStatus == true) {
        request.post("/api/sign_in", { phone: phonenumber, proxy: "", server: 0 })
            .then(res => {

                load()

            }).catch(err => {
                console.error('接口异常：', err);
                ElMessage.error('服务器错误，请稍后重试');
            });
        return
    }
    if (newStatus == false) {
        console.log("批量下线", phonenumber)
        request.post("/api/sign_out", { phone: phonenumber }).then(res => {
            load()
        })
    }



}
const importSuccess = (res) => {
    if (res.code == 200) {
        ElMessage.success('导入成功')
        load()
    } else {
        ElMessage.error(res.msg)
    }
}
</script>

<style scoped>
.upload-demo :deep(.el-upload-dragger) {
    padding: 0 !important
}

.el-upload__text {
    font-size: 10px;
}
</style>