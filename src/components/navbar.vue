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
                <div class="title1">楓谷資料站</div>
                <div class="title2" v-if="choose === '1-1'">--{{list_names["page1_1"]}}</div>
                <div class="title2" v-if="choose === '1-2'">--{{list_names["page1_2"]}}</div>
                <div class="title2" v-if="choose === '2-1'">--{{list_names["page2_1"]}}</div>
                <div class="title2" v-if="choose === '2-2'">--{{list_names["page2_2"]}}</div>
                <div class="title2" v-if="choose === '3-1'">--{{list_names["page3_1"]}}</div>
                <div class="title2" v-if="choose === '3-2'">--{{list_names["page3_2"]}}</div>
                <div class="title2" v-if="choose === '3-3'">--{{list_names["page3_3"]}}</div>
                <div class="title2" v-if="choose === '3-4'">--{{list_names["page3_4"]}}</div>
                <div class="title2" v-if="choose === '3-5'">--{{list_names["page3_5"]}}</div>
                <div class="title2" v-if="choose === '4'">--{{list_names["page4"]}}</div>
                <div class="title2" v-if="choose === '5'">--{{list_names["page5"]}}</div>
                <div class="title2" v-if="choose === '6'">--{{list_names["page6"]}}</div>
                <div class="title2" v-if="choose === '7'">--{{list_names["page7"]}}</div>
                <div class="title2" v-if="choose === '8'">--{{list_names["page8"]}}</div>
            </div>
            
        </div>
        
        <!-- 橫條右邊 -->
        <div style="flex:auto; display:flex; align-items:center; justify-content:flex-end;">
            
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
            <div style="width:58px;">
                <v-text-field
                v-model="light"
                hide-details
                density="compact"/>
            </div>
            
            <div style="padding:0 20px 0 50px;">
                by 小白SSS
            </div>
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onMounted} from "vue";

const light = ref(90)
const emit = defineEmits(["update_light","list_expand"])
watch(light,(new_light)=>{
    if(new_light > 100){light.value=100}
    emit("update_light",light)
})

const click_list_expand = ()=>{emit("list_expand")}

const props = defineProps({
    choose:{},
    list_names:{}
})
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
}
.title2{  /* 後面副標(當前頁面) */
    font-size:16px;
    padding-left:10px;
    padding-bottom:4px;
}
.header_button{
    background-color:#aaaaaa;
}
.header_button:hover{
    background-color: #666666;
}
</style>
