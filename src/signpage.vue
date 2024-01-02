<template>
    <div style="height:100vh; display:flex; align-items:center; justify-content:center;">
        <div id="login_page" class="login_page">
            <div style="display:flex; justify-content:center; align-items:center; font-size:1.5rem; margin-top:30px;">
                <img src="/src/assets/imgs/maplestory.png" style="width:2rem;">
                楓谷資料站登入介面
            </div>
            <form style="margin-top:30px;">
                <fieldset style="border:none; width:calc(100% - 38px); padding:15px;">
                    <div style="margin-left:15px;">
                        <div style="font-size:1.2rem;">&#10146;帳號</div>
                        <input id="account" name="loginpage" placeholder="請輸入帳號" class="log_input" v-model="_account" @keyup.enter.prevent="_log_in"/>
                        <br/>
                        <div style="height:20px;"></div>
                        
                        <div style="font-size:1.2rem;">&#10146;密碼</div>
                        <input id="password1" name="loginpage" type="password" placeholder="請輸入密碼" class="log_input" v-model="_password" @keyup.enter.prevent="_log_in"/>
                        <br/>
                        <div style="margin-bottom:0px;"></div>
                        
                        <input id="show_password1" name="loginpage" type="checkbox" v-model="show_password1"/>
                        <label for="show_password1">顯示密碼</label>
                        <br/>

                        <div style="font-size:1.2rem; margin-top:16px;">&#10146;確認密碼</div>
                        <input id="password2" name="loginpage" type="password" placeholder="請再次輸入密碼" class="log_input" v-model="_password_check" @keyup.enter.prevent="_log_in"/>
                        <br/>
                        <div style="margin-bottom:0px;"></div>
                        
                        <input id="show_password2" name="loginpage" type="checkbox" v-model="show_password2"/>
                        <label for="show_password2">顯示密碼</label>
                        <br/>
                        <div style="height:60px;"><div style="padding:10px 0 0 10px; color:#dd0000;">
                            <span v-if="wrong_check_password">※兩次密碼不同</span>
                        </div></div>
                    </div>
                </fieldset>
            </form>
            <div style="display:flex; align-items:center;">
                <div class="button login_button" @click="_sign_up">註冊</div>
                <div class="button sign_in_button" @click="open_log_in_page">登入頁面</div>
            </div>
            <!-- <div class="button login_button" @click="_check">檢查</div> -->
            <!-- <div>{{ _token }}</div> -->
        </div>
        <!-- <div class="button login_button" @click="_href">連結</div> -->
    </div>
</template>

<script setup>
import {computed, reactive, ref, onMounted, watch} from "vue";

import data from "/src/assets/SettingConfig.json";
    var _settings = reactive({});
    var _frontend = reactive({});
    onMounted(()=>{
        _settings = data;
        _frontend = _settings["FrontEnd"]
    });
    // const _href= ()=>{document.location.href="http://localhost:5173/";}
    
    const show_password1 = ref(false);
    watch(show_password1,(ifshow)=>{
        let p = document.querySelector("#password1");
        if(ifshow){p.type = "text";}
        else{p.type = "password";}
    })
    const show_password2 = ref(false);
    watch(show_password2,(ifshow)=>{
        let p = document.querySelector("#password2");
        if(ifshow){p.type = "text";}
        else{p.type = "password";}
    })

    const _account = ref('');
    const _password = ref('');
    const _password_check = ref('');

    const wrong_check_password = ref(false)

    const _sign_up = async () =>{
        if(_password.value !== _password_check.value){
            const login_page = document.querySelector("#login_page");
            login_page.classList.add("wrong_shake");
            wrong_check_password.value = true;
            setTimeout(() =>{login_page.classList.remove("wrong_shake");},450)
            return
        }
        const _body = {"password":_password.value,"username":_account.value}
        const requestOptions = {
                                method:"POST",
                                headers:{
                                    "Content-Type": "application/json"
                                },
                                body:JSON.stringify(_body)

        }
        await fetch("http://"+_frontend["Hostname"]+":"+_frontend["Backend_port"]+"/signup/",requestOptions)
        .then(res =>{
            return res.json();
        })
        .then(res =>{
            console.log(res)
        })
        .catch(err =>{
            console.log(err);
        })
    };

    const emit = defineEmits(["open_log_in"])

const open_log_in_page = ()=>{emit("open_log_in");}
</script>

<style scoped>
.login_page{
    background-color:#bbbbbbee;
    width:400px;
    height:60%;
    overflow:auto;
    border-radius:10px;
    box-shadow:0 0 30px 10px #ffffff88,0 0 20px 20px #ffffff88 inset;
}
.button{  /* 按鈕(所有應該要可以點擊的東西) */
    cursor:pointer;  /* 滑鼠會變點擊圖樣(手指) */
    transition-duration:0.2s;  /* 因為按鈕都有做滑鼠移入會變色 */
}
.login_button{
    background-color:#0066ff;
    color:#ffffff;
    width:70px;
    padding-top:6px;
    padding-bottom:6px;
    margin-left:32px;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000066 inset,-1px 1px 4px 1px #ffffff88 inset;
}
.login_button:hover{
    background-color:#0000dd;
}
.log_input{
    font-size:1rem;
    width:100%;
    padding:10px;
    background-color:#ffffff;
    border-radius:6px;
}
@keyframes shake{
    0%{
        margin-left:0;
    }25%{
        margin-left:15px;
    }50%{
        margin-left:0;
    }75%{
        margin-left:-15px;
    }100%{
        margin-left:0;
    }
}
.wrong_shake{
    animation:shake .15s ease-in-out 3;
}
.sign_in_button{
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
    margin-left:16px;
}
.sign_in_button:hover{
    background-color:#444444;
}
</style>
