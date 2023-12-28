<template>
    <div class="login_page">
        <div style="display:flex; justify-content:center; align-items:center; font-size:1.5rem; margin-top:30px;">
            <img src="../../public/favicon.ico" style="width:2rem;">
            Netflow登入介面
        </div>
        <form style="margin-top:30px;">
            <fieldset style="border:none; width:calc(100% - 38px); padding:15px;">
                <div style="margin-left:15px;">
                    <div style="font-size:1.2rem;">&#10146;帳號</div>
                    <input id="account" name="loginpage" placeholder="請輸入帳號" class="log_input" v-model="_account" @keyup.enter.prevent="_log_in"/>
                    <br/>
                    <div style="height:20px;"></div>
                    
                    <div style="font-size:1.2rem;">&#10146;密碼</div>
                    <input id="password" name="loginpage" type="password" placeholder="請輸入密碼" class="log_input" v-model="_password" @keyup.enter.prevent="_log_in"/>
                    <br/>
                    <div style="margin-bottom:0px;"></div>
                    
                    <input id="show_password" name="loginpage" type="checkbox" v-model="show_password"/>
                    <label for="show_password">顯示密碼</label>
                    <br/>
                    <div style="height:60px;"></div>
                </div>
            </fieldset>
        </form>
        <div class="button login_button" @click="_log_in">登入</div>
        <div class="button login_button" @click="_check">檢查</div>
        <!-- <div>{{ _token }}</div> -->
    </div>
    <!-- <div class="button login_button" @click="_href">連結</div> -->
</template>

<script setup>
import {computed, reactive, ref, onBeforeMount, onMounted, onBeforeUpdate, onUpdated, watch} from "vue";
    // const _href= ()=>{document.location.href="http://localhost:5173/";}
    
    const show_password = ref(false);
    watch(show_password,(ifshow)=>{
        let p = document.querySelector("#password");
        if(ifshow){p.type = "text";}
        else{p.type = "password";}
    })

    const _account = ref('');
    const _password = ref('');
    const _token = ref('');

    function getCookie(name) {
        const temp = `; ${document.cookie}`;
        const parts = temp.split(`; ${name}=`);
        if (parts.length === 2) return parts.pop().split(';').shift();
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
        await fetch("http://localhost:9805/checkToken/",requestOptions)
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
        })
    }
    
    const _log_in = async () =>{
        const _body = {"password":_password.value,"username":_account.value}
        const requestOptions = {
                                method:"POST",
                                headers:{
                                    "Content-Type": "application/json"
                                },
                                body:JSON.stringify(_body)

        }
        await fetch("http://localhost:9805/login/",requestOptions)
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
            _token.value = res["token"]
            console.log(res);
            console.log(_token.value)
            document.cookie = `token=${_token.value}`;
        })
        .catch(err =>{
            console.log(err);
        })
    };

    
    

</script>

<style scoped>
.login_page{
    background-color:#bbbbbbee;
    width:400px;
    height:50%;
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
    margin-right:40px;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000066 inset,-1px 1px 4px 1px #ffffff88 inset;
    margin-left:32px;
}
.login_button:hover{
    background-color:#0000dd;
}
.log_input{
    font-size:1rem;
    width:calc(100% - 24px - 20px);
    padding:10px;
}
</style>
