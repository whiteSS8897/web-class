<template>
    <div v-if="keep_log_in === false">
        <loginpage
        @check_login="handle_change_log" @get_login_user="handle_login_user" @get_login_nickname="handle_login_nickname"
        v-bind:_kli="keep_log_in"/>
    </div>
    <div v-if="keep_log_in === true">
        <navbar @update_light="set_new_light" @list_expand="click_list_expand" @log_out="_logout" @personal_setting="open_personal_setting"
        v-bind:choose="choose" v-bind:list_names="list_names"
        v-bind:now_nickname="now_nickname" v-bind:now_user="now_user" v-bind:avatar="received_avatar"/>
        <div class="box">
            <div class="list">
                <mainlist @update="list_choose_handle" v-bind:names="list_names"/>
            </div>
            
            <div class="main_page">
                <page1_1 v-if="choose === '1-1'"/>
                <page1_2 v-if="choose === '1-2'"/>
                <page1_3 v-if="choose === '1-3'" v-bind:user_name="now_user"/>
                <page2_1 v-if="choose === '2-1'"/>
                <page2_2 v-if="choose === '2-2'"/>
                <page3_1 v-if="choose === '3-1'"/>
                <page3_2 v-if="choose === '3-2'"/>
                <page3_3 v-if="choose === '3-3'"/>
                <page3_4 v-if="choose === '3-4'"/>
                <page3_5 v-if="choose === '3-5'"/>
                <page4 v-if="choose === '4'"/>
                <page5 v-if="choose === '5'"/>
                <page6 v-if="choose === '6'"/>
                <page7 v-if="choose === '7'"/>
                <page8 v-if="choose === '8'"/>

                <personal_setting v-if="choose === 'personal_setting'"
                v-bind:user_name="now_user" v-bind:nickname="now_nickname" v-bind:avatar="received_avatar"
                @re_check="_check"/>
            </div>
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onMounted} from "vue";
import page1_1 from "./components/page1_1.vue";
import page1_2 from "./components/page1_2.vue";
import page1_3 from "./components/page1_3.vue";
import page2_1 from "./components/page2_1.vue";
import page2_2 from "./components/page2_2.vue";
import page3_1 from "./components/page3_1.vue";
import page3_2 from "./components/page3_2.vue";
import page3_3 from "./components/page3_3.vue";
import page3_4 from "./components/page3_4.vue";
import page3_5 from "./components/page3_5.vue";
import page4 from "./components/page4.vue";
import page5 from "./components/page5.vue";
import page6 from "./components/page6.vue";
import page7 from "./components/page7.vue";
import page8 from "./components/page8.vue";


import data from "/src/assets/SettingConfig.json";
var _settings = reactive({});
var _frontend = reactive({});
onMounted(()=>{
    _settings = data;
    _frontend = _settings["FrontEnd"]
    _check();
});


import navbar from "./components/navbar.vue"
const light = ref(90)
const set_new_light = (new_light)=>{
    light.value = new_light.value
    document.querySelector(".main_page").style.opacity=light.value/100;
}


import mainlist from "./components/mainlist.vue"
const list_names = {"page1_1":"ARC 升級表",
                    "page1_2":"AUT 升級表",
                    "page1_3":"符文金額小算盤",
                    "page2_1":"頁面2-1",
                    "page2_2":"頁面2-2",
                    "page3_1":"頁面3-1",
                    "page3_2":"頁面3-2",
                    "page3_3":"頁面3-3",
                    "page3_4":"頁面3-4",
                    "page3_5":"頁面3-5",
                    "page4":"星力★",
                    "page5":"星火數據",
                    "page6":"各種小算盤",
                    "page7":"頁面7",
                    "page8":"頁面8"};

var list_expand = true  //是否已展開
const click_list_expand = ()=>{
    if(list_expand){
        document.querySelector(".list").style.width="0px";  //收縮(寬變為0px)
        list_expand = !list_expand
    }else{
        document.querySelector(".list").style.width="220px";  //展開(寬變回220px)
        list_expand = !list_expand
    }
}

//控制主畫面顯示第幾個
const choose = ref("1-1")  //表示選擇清單第幾個
    //接收從mainlist傳來的選擇
    const list_choose_handle = (new_list_choise)=>{
        choose.value = new_list_choise;
    }


import loginpage from "./loginpage.vue"

const keep_log_in = ref(false);
const now_user = ref("unknown")
const now_nickname = ref("unknown")
const handle_change_log = (new_state)=>{
    keep_log_in.value = new_state
}
const handle_login_user = (login_user)=>{
    now_user.value = login_user
}
const handle_login_nickname = (login_nickname)=>{
    now_nickname.value = login_nickname
}
const _logout = () =>{
        let d = new Date()
        keep_log_in.value = false;
        document.cookie = "token='';expires=" + d.toGMTString();
}

import personal_setting from "./components/personal_setting.vue";
const open_personal_setting = ()=>{choose.value = "personal_setting"}


function getCookie(name) {
        const temp = `; ${document.cookie}`;
        console.log(temp)
        const parts = temp.split(`; ${name}=`);
        console.log(parts)
        if (parts.length === 2) return parts.pop().split(';').shift();
        else return ''
    }
const _check = async () =>{
        let check_token = getCookie("token")
        console.log(check_token)
        const _body = {"token":check_token}
        const requestOptions = {
                                method:"POST",
                                headers:{
                                    "Content-Type": "application/json"
                                },
                                body:JSON.stringify(_body)
        }
        await fetch("http://"+_frontend["Hostname"]+":"+_frontend["Backend_port"]+"/checkToken/",requestOptions)
        .then(res =>{
            return res.text();
        })
        .then(res =>{
            return res.substring(JSON.stringify(_body).length);
        })
        .then(res =>{
            return JSON.parse(res);
        })
        .then(res =>{
            console.log(res);
            if (res["message"] === "Token is valid"){
                //console.log(keep_log_in.value)
                keep_log_in.value = true
                now_user.value = res["username"]
                now_nickname.value = res["nickname"]
                get_image();
                console.log("username:",now_user.value)
                console.log("nickname:",now_nickname.value)

                //console.log(keep_log_in.value)
            }
            else{
                //console.log(keep_log_in.value)
                keep_log_in.value = false
                //console.log(keep_log_in.value)
            }
        })
    }

const received_avatar = ref();
const get_image = async () =>{
    const _body = {"userName":now_user.value}
    console.log(JSON.stringify(_body))
    const requestOptions = {
                            method:"POST",
                            headers:{
                                "Content-Type": "application/json"
                            },
                            body:JSON.stringify(_body)
    }
    await fetch("http://"+_frontend["Hostname"]+":"+_frontend["Backend_port"]+"/getImage/",requestOptions)
    .then(res =>{
        return res.json();
    })
    // .then(res =>{
    //     return JSON.stringify(res);
    // })
    .then(res =>{
        received_avatar.value = res[""];
        const imageBase64 = res;
        // console.log(res["profile"]);
        // console.log(imageBase64.profile.Imagebase64);
        
        // console.log(objectURL);
        if (res["profile"] === "Null image"){
            received_avatar.value = "https://api.iconify.design/svg-spinners:6-dots-rotate.svg?color=%23aaaaaa"
            console.log("Null image")
        }
        else {
            const byteString = atob(imageBase64.profile.Imagebase64);

            const ab = new ArrayBuffer(byteString.length);
            const ia = new Uint8Array(ab);
            for (let i = 0; i < byteString.length; i++) {
                ia[i] = byteString.charCodeAt(i);
            }
            const blob = new Blob([ab]);
            const objectURL = URL.createObjectURL(blob);
            received_avatar.value = objectURL
        }
        console.log(received_avatar.value)
    })
    .catch(()=>{
        received_avatar.value = "https://api.iconify.design/svg-spinners:6-dots-rotate.svg?color=%23aaaaaa"
        console.log("fail to get image")
    })

}
</script>


<style scoped>
.box{  /* 橫條下方(清單+主畫面) */
    display:flex;
    height:90vh;
}
.list{  /* 清單 */
    flex:none;
    width:220px;
    margin-left:10px;
    overflow-x:hidden;
    overflow-y:auto;
    transition-duration:0.8s;
}
.main_page{  /* 主畫面 */
    flex:auto;
    margin-left:10px;
    margin-right:10px;
    border:0px #000000 solid;
    opacity:90%;
    padding:10px;
    background-color:#dddddd;
    overflow:auto;
}
</style>
