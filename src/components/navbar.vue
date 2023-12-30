<template>
    <!-- 上方橫條 -->
    <div class="navbar">

        
        <!-- 橫條左邊 -->
        <div style="display:flex; align-items:center;">
            
            <!-- 展開清單按鈕 -->
            <div class="button list_image header_button" @click="click_list_expand">
                <img src="https://api.iconify.design/prime:list.svg" width="36" height="36">
            </div>
            
            <!-- Netflow字樣與後面副標 -->
            <div style="display:flex; align-items:end;">
                <div class="title1"><img src="/src/assets/imgs/maplestory.png" width=30 style="margin-right:2px;">楓谷資料站</div>
                <div class="title2" v-if="choose === '1-1'">---{{list_names["page1_1"]}}</div>
                <div class="title2" v-if="choose === '1-2'">---{{list_names["page1_2"]}}</div>
                <div class="title2" v-if="choose === '1-3'">---{{list_names["page1_3"]}}</div>
                <div class="title2" v-if="choose === '2-1'">---{{list_names["page2_1"]}}</div>
                <div class="title2" v-if="choose === '2-2'">---{{list_names["page2_2"]}}</div>
                <div class="title2" v-if="choose === '3-1'">---{{list_names["page3_1"]}}</div>
                <div class="title2" v-if="choose === '3-2'">---{{list_names["page3_2"]}}</div>
                <div class="title2" v-if="choose === '3-3'">---{{list_names["page3_3"]}}</div>
                <div class="title2" v-if="choose === '3-4'">---{{list_names["page3_4"]}}</div>
                <div class="title2" v-if="choose === '3-5'">---{{list_names["page3_5"]}}</div>
                <div class="title2" v-if="choose === '4'">---{{list_names["page4"]}}</div>
                <div class="title2" v-if="choose === '5'">---{{list_names["page5"]}}</div>
                <div class="title2" v-if="choose === '6'">---{{list_names["page6"]}}</div>
                <div class="title2" v-if="choose === '7'">---{{list_names["page7"]}}</div>
                <div class="title2" v-if="choose === '8'">---{{list_names["page8"]}}</div>

                <div class="title2" v-if="choose === 'personal_setting'">---個人設定</div>
            </div>
            
        </div>
        
        <!-- 橫條右邊 -->
        <div style="flex:auto; display:flex; align-items:center; justify-content:flex-end; padding-right:20px;">
            
            <!-- 亮度滑桿 -->
            <div style="width:max(15%,180px);">
                <v-slider
                v-model="light"
                :min="0" :max="100" :step="1"
                prepend-icon="mdi-lightbulb-on-10"
                hide-details
                thumb-label thumb-size="10"/>
            </div>
            
            <!-- 亮度數值框 -->
            <div style="width:58px; margin-right:50px;">
                <v-text-field
                v-model="light"
                hide-details
                density="compact"/>
            </div>
            
            <!-- 使用者(整體) -->
            <div class="button user navbar_button" @click="open_personal_setting">
                    <!-- 使用者icon -->
                <div class="user_icon">
                    <img src="https://api.iconify.design/svg-spinners:6-dots-rotate.svg?color=%23aaaaaa" width="26" height="26">
                </div>
                    {{ now_nickname }}
            </div>

            <!-- 使用者與登出間的分隔線 -->
            <div style="height:30px; border-right:2px #777777 solid; margin:0px 15px 0px 15px;"></div>

            <!-- 登出按鈕 -->
            <div class="button logout navbar_button" @click="_logout">登出</div>
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";

const light = ref(90)
const emit = defineEmits(["update_light","list_expand","log_out","personal_setting"])
watch(light,(new_light)=>{
    if(new_light > 100){light.value=100;}
    emit("update_light",light)
})

const click_list_expand = ()=>{emit("list_expand")}

const props = defineProps({
    choose:{},
    list_names:{},
    now_nickname:{},
    now_user:{},
})

const _logout = ()=>{emit("log_out")}
const open_personal_setting = ()=>{emit("personal_setting")}
</script>


<style scoped>
.button{  /* 按鈕(所有應該要可以點擊的東西) */
    cursor:pointer;  /* 滑鼠會變點擊圖樣(手指) */
    transition-duration:0.2s;  /* 因為按鈕都有做滑鼠移入會變色 */
}
.navbar{  /* 上方橫條 */
    display:flex;
    background-color:#aaaaaa;
    height:60px;
    margin-bottom:10px;
    align-items:center;
    white-space:nowrap;
}
.list_image{  /* 展開選單的按鈕 */
    width:50px;
    height:50px;
    display:flex;
    justify-content:center;
    align-items:center;
    margin-left:10px;
    border:0px #000000 solid;
    border-radius:100%;
}
.title1{  /* Netflow字樣 */
    font-size:30px;
    padding-left:25px;
    display:flex;
    align-items:center;
}
.title2{  /* 後面副標(當前頁面) */
    font-size:16px;
    padding-left:10px;
    padding-bottom:4px;
}
.navbar_button{
    background-color:#aaaaaa;
}
.navbar_button:hover{
    background-color: #666666;
}
.user{  /* "使用者"按鈕(整體) */
    display:flex;
    align-items:center;
    padding:5px;
    border:0px #000000 solid;
    border-radius:10px;
}
.user_icon{  /* 使用者icon */
    width:26px;
    height:26px;
    background-color:#444444;
    border-radius:100%;
    margin-right:5px;
    display:flex;
    align-items:center;
    justify-content:center;
    overflow:hidden;
}
.logout{  /* 登出按鈕 */
    display:flex;
    align-items:center;
    padding:5px;
    border:0px #000000 solid;
    border-radius:10px;
}
</style>
