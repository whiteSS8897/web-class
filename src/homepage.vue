<template>
    <navbar @update_light="set_new_light" @list_expand="click_list_expand"
        v-bind:choose="choose" v-bind:list_names="list_names"/>
    <div class="box">
        <div class="list">
            <mainlist @update="list_choose_handle" v-bind:names="list_names"/>
        </div>

        <div class="main_page">
            <div v-if="choose === '1-1'"><page1_1/></div>
            <div v-if="choose === '1-2'"><page1_2/></div>
            <div v-if="choose === '2-1'"><page2_1/></div>
            <div v-if="choose === '2-2'"><page2_2/></div>
            <div v-if="choose === '3-1'"><page3_1/></div>
            <div v-if="choose === '3-2'"><page3_2/></div>
            <div v-if="choose === '3-3'"><page3_3/></div>
            <div v-if="choose === '3-4'"><page3_4/></div>
            <div v-if="choose === '3-5'"><page3_5/></div>
            <div v-if="choose === '4'"><page4/></div>
            <div v-if="choose === '5'"><page5/></div>
            <div v-if="choose === '6'"><page6/></div>
            <div v-if="choose === '7'"><page7/></div>
            <div v-if="choose === '8'"><page8/></div>
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onMounted} from "vue";
import page1_1 from "./components/page1_1.vue";
import page1_2 from "./components/page1_2.vue";
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


import navbar from "./components/navbar.vue"
const light = ref(90)
const set_new_light = (new_light)=>{
    light.value = new_light.value
    document.querySelector(".main_page").style.opacity=light.value/100;
}


import mainlist from "./components/mainlist.vue"
const list_names = {"page1_1":"頁面1-1",
                    "page1_2":"頁面1-2",
                    "page2_1":"頁面2-1",
                    "page2_2":"頁面2-2",
                    "page3_1":"頁面3-1",
                    "page3_2":"頁面3-2",
                    "page3_3":"頁面3-3",
                    "page3_4":"頁面3-4",
                    "page3_5":"頁面3-5",
                    "page4":"頁面4",
                    "page5":"頁面5",
                    "page6":"頁面6",
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
