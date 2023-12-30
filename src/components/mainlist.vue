<template>
    <div>
        <details class="detail_1" open>
            <summary class="button summary_1" style="height:25px;" @click="all_other_hide($event)">
            符文相關
            </summary>

                <div style="border-left:16px #8800ff solid;">
                    <div class="button detail_1_choice final_choice choosed_final_choice"
                        @click="choose_list('1-1') ,choosed_list($event)">
                    {{names["page1_1"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('1-2') ,choosed_list($event)">
                    {{names["page1_2"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('1-3') ,choosed_list($event)">
                    {{names["page1_3"]}}
                    </div>
                </div>
        </details>

        <details class="detail_1">
            <summary class="button summary_1" style="height:25px;" @click="all_other_hide($event)">
            class 2
            </summary>

                <div style="border-left:16px #8800ff solid;">
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('2-1') ,choosed_list($event)">
                    {{names["page2_1"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('2-2') ,choosed_list($event)">
                    {{names["page2_2"]}}
                    </div>
                </div>
        </details>

        <details class="detail_1">
            <summary class="button summary_1" style="height:25px;" @click="all_other_hide($event)">
            class 3
            </summary>

                <div style="border-left:16px #8800ff solid;">
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('3-1') ,choosed_list($event)">
                    {{names["page3_1"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('3-2') ,choosed_list($event)">
                    {{names["page3_2"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('3-3') ,choosed_list($event)">
                    {{names["page3_3"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('3-4') ,choosed_list($event)">
                    {{names["page3_4"]}}
                    </div>
                    <div class="button detail_1_choice final_choice"
                        @click="choose_list('3-5') ,choosed_list($event)">
                    {{names["page3_5"]}}
                    </div>
                </div>
        </details>
        
        <div class="button list_choice final_choice"
            @click="choose_list('4') ,all_other_hide($event) ,choosed_list($event)">
            {{names["page4"]}}
        </div>
        <div class="button list_choice final_choice"
            @click="choose_list('5') ,all_other_hide($event) ,choosed_list($event)">
            {{names["page5"]}}
        </div>
        <div class="button list_choice final_choice"
            @click="choose_list('6') ,all_other_hide($event) ,choosed_list($event)">
            {{names["page6"]}}
        </div>
        <div class="button list_choice final_choice"
            @click="choose_list('7') ,all_other_hide($event) ,choosed_list($event)">
            {{names["page7"]}}
        </div>
        <div class="button list_choice final_choice"
            @click="choose_list('8') ,all_other_hide($event) ,choosed_list($event)">
            {{names["page8"]}}
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";

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

    //選擇的選項要變為藍色
        const choosed_list = (n)=>{
            var list_choices = document.querySelectorAll(".final_choice");
            //將所有選項改為原色
            for(let i of list_choices){
                i.classList.remove("choosed_final_choice")
            }
            n.target.classList.add("choosed_final_choice")
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
    white-space:nowrap;
}
.list_choice:hover ,.detail_1_choice:hover{
    background-color:#888888;
    border-left:16px #555555 solid;
}
.summary_1:hover{
    background-color:#7755aa;
    border-left:16px #555555 solid;
}
.detail_1[open] > summary{
    background-color:#ccaaff;
    border-left:16px #8800ff solid;
}
.detail_1_choice{
    padding:3px 0 3px 8px;
    background-color:#dddddd;
    border-bottom:1px #bbbbbb dotted;
    border-left:4px #999999 solid;
    white-space:nowrap;
}
.choosed_final_choice{
    background-color:#99bbff;
    border-left:16px #555555 solid;
}
.choosed_final_choice:hover{
    background-color:#6688bb;
}
</style>