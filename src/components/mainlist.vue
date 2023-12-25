<template>
    <details class="detail_1" open>
        <summary class="button summary_1" style="height:25px;" @click="all_other_hide($event)">
        分類ㄧ
        </summary>

            <div style="border-left:16px #222222 solid;">
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('1-1') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)"
                    style="border-left:16px #555555 solid; background-color:#99bbff;"
                    >
                {{names["page1_1"]}}
                </div>
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('1-2') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page1_2"]}}
                </div>
            </div>
    </details>

    <details class="detail_1">
        <summary class="button summary_1" style="height:25px;" @click="all_other_hide($event)">
        分類二
        </summary>

            <div style="border-left:16px #222222 solid;">
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('2-1') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page2_1"]}}
                </div>
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('2-2') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page2_2"]}}
                </div>
            </div>
    </details>

    <details class="detail_1">
        <summary class="button summary_1" style="height:25px;" @click="all_other_hide($event)">
        分類三
        </summary>

            <div style="border-left:16px #222222 solid;">
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('3-1') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page3_1"]}}
                </div>
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('3-2') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page3_2"]}}
                </div>
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('3-3') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page3_3"]}}
                </div>
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('3-4') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page3_4"]}}
                </div>
                <div class="button detail_1_choice final_choice"
                    @click="choose_list('3-5') ,choosed_list($event)"
                    @mouseenter="touch_list($event)"
                    @mouseleave="untouch_list($event)">
                {{names["page3_5"]}}
                </div>
            </div>
    </details>
    
    <div class="button list_choice final_choice"
        @click="choose_list('4') ,all_other_hide($event) ,choosed_list($event)"
        @mouseenter="touch_list($event)"
        @mouseleave="untouch_list($event)"
        style="background-color:#dddddd;">
        {{names["page4"]}}
    </div>
    <div class="button list_choice final_choice"
        @click="choose_list('5') ,all_other_hide($event) ,choosed_list($event)"
        @mouseenter="touch_list($event)"
        @mouseleave="untouch_list($event)"
        style="background-color:#dddddd;">
        {{names["page5"]}}
    </div>
    <div class="button list_choice final_choice"
        @click="choose_list('6') ,all_other_hide($event) ,choosed_list($event)"
        @mouseenter="touch_list($event)"
        @mouseleave="untouch_list($event)"
        style="background-color:#dddddd;">
        {{names["page6"]}}
    </div>
    <div class="button list_choice final_choice"
        @click="choose_list('7') ,all_other_hide($event) ,choosed_list($event)"
        @mouseenter="touch_list($event)"
        @mouseleave="untouch_list($event)"
        style="background-color:#dddddd;">
        {{names["page7"]}}
    </div>
    <div class="button list_choice final_choice"
        @click="choose_list('8') ,all_other_hide($event) ,choosed_list($event)"
        @mouseenter="touch_list($event)"
        @mouseleave="untouch_list($event)"
        style="background-color:#dddddd;">
        {{names["page8"]}}
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onMounted} from "vue";

//接收mainpage.vue丟過來的list_names
const props = defineProps({
    names:{
        type:Object
    }
})

    
//處理選擇
    const choose = ref("1")  //表示選擇清單哪一個的變數
    const choose_list = (n) => {choose.value = n}  //控制選擇清單哪一個(@click)
    //將選擇上傳至App
    const emit = defineEmits(["update"])
    watch(choose,(newchoice)=>{
        emit("update",newchoice)
    })

    const all_other_hide = (n)=>{
        var details = document.querySelectorAll(".detail_1");
        for(let i of details){
            i.open = false;
        }
        n.open = true;
    }


//處理變色
    //滑鼠移入選項
    const touch_list = (n)=>{
            //原色的狀況
            if(n.target.style.backgroundColor==="rgb(221, 221, 221)"){
                n.target.style.backgroundColor="#888888";
                n.target.style.borderLeft="16px #555555 solid";
            }
            //藍色的狀況(被選的)
            else if(n.target.style.backgroundColor==="rgb(153, 187, 255)"){
                n.target.style.backgroundColor="#7799cc";
                n.target.style.borderLeft="16px #555555 solid";
            }
        }
    //滑鼠移出選項
        const untouch_list = (n)=>{
            //原色的狀況
            if(n.target.style.backgroundColor==="rgb(136, 136, 136)"){
                n.target.style.backgroundColor="#dddddd";
                n.target.style.borderLeft="4px #888888 solid";
            }
            //藍色的狀況(被選的)
            else if(n.target.style.backgroundColor==="rgb(119, 153, 204)"){
                n.target.style.backgroundColor="#99bbff";
            }
        }
    //選擇的選項要變為藍色
        const choosed_list = (n)=>{
            var list_choices = document.querySelectorAll(".final_choice");
            //將所有選項改為原色
            for(let i of list_choices){
                i.style.backgroundColor="#dddddd";
                i.style.borderLeft="4px #888888 solid";
            }
            n.target.style.backgroundColor="#99bbff";  //選擇的選項設為藍色
            n.target.style.borderLeft="16px #555555 solid";  //選擇的選項左邊框加粗
        }


</script>


<style scoped>
.button{  /* 按鈕(所有應該要可以點擊的東西) */
    cursor:pointer;  /* 滑鼠會變點擊圖樣(手指) */
    transition-duration:0.2s;  /* 因為按鈕都有做滑鼠移入會變色 */
}
.list_choice ,.summary_1{  /* 清單內一個個的選項 */
    height:32px;
    padding:3px 0 3px 8px;
    border-bottom:1px #bbbbbb dotted;
    border-left:4px #999999 solid;
    background-color:#dddddd;
    display:flex;
    align-items:center;
    white-space:nowrap;
}
.list_choice:hover ,.detail_1_choice:hover{
    border-left:16px #555555 solid;
    background-color:#888888;
}
.summary_1:hover{
    border-left:16px #555555 solid;
    background-color:#7755aa;
}
.detail_1[open] > summary{
    background-color:#ccaaff;
    border-left:16px #222222 solid;
}
.detail_1_choice{
    padding:3px 0 3px 8px;
    background-color:#dddddd;
    border-bottom:1px #bbbbbb dotted;
    border-left:4px #999999 solid;
}
</style>