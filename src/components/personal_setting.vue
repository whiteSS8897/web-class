<template>
    <dialog id="change_avatar">
        <div class="inner_change_avatar">
            <img v-if="avatar" v-bind:src="avatar_url" style="max-width:100%; max-height:100%; margin:auto;"/>
        </div>
        <div style="display:flex; justify-content:center; position:absolute; top:70%; left:0; right:0;">
            <div class="button change_avatar_button" @click="update_avatar">更改</div>
            <div class="button change_avatar_button close_dialog_button" @click="close_dialog">關閉</div>
        </div>
    </dialog>
    <div class="profile">
        <div style="display:flex;align-items:center;">
            <div class="user_icon">
                <img v-bind:src="now_avatar" width="30" height="30">
            </div>
            <div style="font-size:22px; display:flex;"> {{ now_user }}
                <div style="margin-left:20px;" class="change_avatar button" @click="$refs.avatarInput.click()">更改頭像</div>
                <input type="file" style="display:none;" ref="avatarInput" @change="upload_avatar" name="image">

                
            </div>
        </div>
        <details>
            <summary class="item button" style="display:flex;align-items:center;">
                    <div style="padding:0px 20px 0px 30px;">暱稱</div>
                    {{ now_nickname }}
            </summary>
            <div style="padding:0px 50px; display:flex; align-items:center;">
                <input id="nickname" placeholder="請輸入新暱稱" v-model="_new_nickname" @keyup.enter.prevent="" class="input_box">
                <div class="button confirm_button" @click="update_nickname(_new_nickname)">更改</div>
            </div>
        </details>
    </div>
    <!-- <div style="display:flex; margin-top:20px;">
        <div class="button confirm_button" @click="get_image()">測試接收頭像</div>
    </div>
    {{ received_avatar }}
    <img v-bind:src="received_avatar"/> -->
</template>


<script setup>
import {computed, reactive, ref, onBeforeMount, onMounted, onBeforeUpdate, onUpdated, watch} from "vue";
import data from "../../public/json/SettingConfig.json";

    var _settings = reactive({});
    var _frontend = reactive({});
    onBeforeMount(()=>{
            _settings = data;
            _frontend = _settings["FrontEnd"]
            // console.log("※check receive")
            // console.log("username:",now_user.value)
            // console.log("nickname:",now_nickname.value)
            // console.log("check receive※")
        });

    const props = defineProps({
        user_name:{},
        nickname:{},
        avatar:{}
    })
    const now_user = computed(()=>{
        return props.user_name
    })
    const now_nickname = computed(()=>{
        return props.nickname
    })
    const now_avatar = computed(()=>{
        return props.avatar
    })

    const _new_nickname = ref("")
    const emit = defineEmits(["re_check","update_avatar"])
    const update_nickname = async (new_nickname) =>{
        const _body = {"userName":now_user.value,"newNickname":new_nickname}
        console.log(JSON.stringify(_body))
        const requestOptions = {
                                method:"POST",
                                headers:{
                                    "Content-Type": "application/json"
                                },
                                body:JSON.stringify(_body)
        }
        await fetch("http://"+_frontend["HOSTNAMe"]+":"+_frontend["LoginBackEndPort"]+"/updateNickname/",requestOptions)
        .then(res =>{
            return res.text();
        })
        .then(res =>{
            return JSON.parse(res);
        })
        .then(res =>{
            console.log(res);
            if(res["message"] === "Nickname updated successfully"){
                emit("re_check");
            }
        })
    }

    const avatar = ref()
    const avatar_url = ref()
    const upload_avatar = async (e) =>{
        avatar.value = e.target.files[0];
        console.log(avatar.value);
        avatar_url.value = window.URL.createObjectURL(avatar.value);
        console.log(avatar_url.value)
        let _dialog = document.querySelector("#change_avatar");
        _dialog.showModal();
        e.target.value = "";
    }
    /*const update_avatar = async () =>{
        const _body = {"image":avatar_url.value};
        const requestOptions = {
                                method:"POST",
                                headers:{
                                    "Content-Type": "application/json"
                                },
                                body:JSON.stringify(_body)
        }
    */
    const update_avatar = async () => {
        var formData = new FormData();
        formData.append("image",avatar.value);
        formData.append("userName",now_user.value)
        const requestOptions = {
            method: "POST",
            body: formData
        };
        console.log(formData.getAll("image"))
        console.log(formData.getAll("userName"))
        // const _body = new FormData();
        // _body.append("image",avatar.value)
        // const requestOptions = {
        //                         method:"POST",
        //                         body:_body
        // }



        await fetch("http://"+_frontend["HOSTNAMe"]+":"+_frontend["LoginBackEndPort"]+"/getImage/",requestOptions)
        .then(res =>{
            return res.text();
        })
        .then(res =>{
            return JSON.parse(res);
        })
        .then(res =>{
            console.log(res);
        })
        close_dialog();
        emit("update_avatar");
    }

    const close_dialog = ()=>{
        let _dialog = document.querySelector("#change_avatar");
        _dialog.close();
    }
    

</script>


<style scoped>
.button{  /* 按鈕(所有應該要可以點擊的東西) */
    cursor:pointer;  /* 滑鼠會變點擊圖樣(手指) */
    transition-duration:0.2s;  /* 因為按鈕都有做滑鼠移入會變色 */
}
.profile{
    background-color:#c0c0c0;
    padding:16px;
    border-radius:24px;
}
.user_icon{  /* 使用者icon */
    width:30px;
    height:30px;
    background-color:#444444;
    border-radius:100%;
    margin-right:5px;
    display:flex;
    align-items:center;
    justify-content:center;
    overflow:hidden;
}
.item{
    display:flex;
    align-items:center;
    background-color:#aaaaaa;
    margin-top:16px;
    padding:8px 8px 8px 20px;
    border-radius:10px;
}
.item:hover{
    background-color:#888888;
}
details[open] > summary:before {
    transform: rotate(90deg);
}
summary:before {
    content:"➢";
    font-size:24px;
    position: absolute;
    transform-origin:50%;
    transition: .35s transform ease;
}
.input_box{
    padding:6px 10px;
    border-radius:6px;
    box-shadow:-1px 1px 4px 1px #999999 inset;
}
.confirm_button{
    background-color:#0066ff;
    color:#ffffff;
    padding:4px 10px;
    font-size:14px;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000066 inset,-1px 1px 4px 1px #ffffff88 inset;
    margin-left:6px;
}
.confirm_button:hover{
    background-color:#0000dd;
}
.change_avatar{
    font-size:16px;
    background-color:#777777;
    color:#ffffff;
    padding:4px 10px;
    font-size:14px;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000066 inset,-1px 1px 4px 1px #ffffff88 inset;
    margin-left:6px;
}
.change_avatar:hover{
    background-color:#444444;
}
#change_avatar{
    background-color:#ffffff00;
    width:100%;
    height:100%;
    margin:auto;
    outline:none;
    border:none;
    overflow:hidden;
}
#change_avatar::backdrop{
    background-color:#000000cc;
}
.inner_change_avatar{
    background-color:#ffffff;
    margin:auto;
    padding:16px;
    border-radius:16px;
    padding:16px;
    width:30%;
    height:40%;
    position:absolute;
    top:0;
    right:0;
    bottom:0;
    left:0;
}
.change_avatar_button{
    background-color:#0066ff;
    color:#ffffff;
    padding:6px 16px;
    font-size:16px;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000066 inset,-1px 1px 4px 1px #ffffff88 inset;
    margin-top:16px;
}
.change_avatar_button:hover{
    background-color:#0000dd;
}
.close_dialog_button{
    background-color:#888888;
    color:#ffffff;
    margin-left:16px;
}
.close_dialog_button:hover{
    background-color:#555555;
}
</style>