<template>
    <div style="display:flex;">
        <div>
            <div class="search_area">
                <div class="inside_search_area">
                    <div class="button change_search_way_button" @click="change_search_way()">
                        <span v-if="search_as_scope">以範圍搜尋</span>
                        <span v-if="!search_as_scope">以等級搜尋</span>
                    </div>
                    <select class="scope_select" v-model="scope_choosed" v-if="search_as_scope">
                        <option v-for="(op,index) in scope_option_begins" v-bind:value="index">{{ op }}～{{ scope_option_ends[index] }}</option>
                    </select>
                    <input type=number v-model="level" v-if="!search_as_scope" class="level_input">
                </div>
                <div v-if="out_of_range" class="warning_out_of_range">等級需在0~275之間</div>
            </div>

            <div style="background-color:#ffff88; text-align:center; padding:5px; margin-top:40px;">武器星火攻擊力小算盤</div>
            <div style="background-color:#ffffbb66; padding:5px; text-align:center;">
                武器白字：<input type=number class="attack_input" v-model="input_attack">
                <br>
                <select class="tier_select" v-model="input_tier">
                    <option value="1">T1</option>
                    <option value="2">T2</option>
                    <option value="3">T3</option>
                    <option value="4">T4</option>
                    <option value="5">T5</option>
                    <option value="6">T6</option>
                    <option value="7">T7</option>
                </select>
            </div>
            <div style="display:flex; justify-content:center; font-size:30px;">⇓</div>
            <div style="display:flex; justify-content:center;">
                <span style="background-color:#ffff88; padding:5px 10px;">星火數值</span>
                <span style="background-color:#ffffbb66; padding:5px 10px;">
                    +{{ Round_to(input_attack*(Math.floor(level/40)+1)*weapon_bonus[input_tier-1],6) }}攻
                </span>
            </div>
        </div>

        <div>
            <table style="margin-left:30px;">
                <tr style="position:relative;" v-if="search_as_scope"><th colspan="8" class="table_title">
                    {{ scope_option_begins[scope_choosed] }}～{{ scope_option_ends[scope_choosed] }}等
                </th></tr>
                <tr style="position:relative;" v-if="!search_as_scope"><th colspan="8" class="table_title">
                    {{ level }}等
                </th></tr>
                <th>星火能力</th>
                <th v-for="i in Array.from(Array(7).keys())">T{{ i+1 }}</th>
                <tr>
                    <td>單屬</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1)*(Math.floor(level/20)+1) }}</td>
                </tr>
                <tr>
                    <td>雙屬</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1)*(Math.floor(level/40)+1) }}</td>
                </tr>
                <tr>
                    <td>攻擊力</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1) }}</td>
                </tr>
                <tr>
                    <td>防禦力</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1)*(Math.floor(level/20)+1) }}</td>
                </tr>
                <tr>
                    <td>HP／MP</td>
                    <td v-for="i in Array.from(Array(7).keys())" v-if="search_as_scope && scope_choosed!==0">{{ 30*(i+1)*(scope_choosed*2) }} || {{ 30*(i+1)*(scope_choosed*2+1) }}</td>
                    <td v-for="i in Array.from(Array(7).keys())" v-if="!search_as_scope && level>9">{{ 30*(i+1)*Math.floor(level/10) }}</td>
                    <td v-for="i in Array.from(Array(7).keys())" v-if="search_as_scope && scope_choosed===0">{{ 30*(i+1) }}</td>
                    <td v-for="i in Array.from(Array(7).keys())" v-if="!search_as_scope && level<=9">{{ 30*(i+1) }}</td>
                </tr>
                <tr>
                    <td>移動速度</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1) }}</td>
                </tr>
                <tr>
                    <td>跳躍力</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1) }}</td>
                </tr>
                <tr>
                    <td>全屬%</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1) }}%</td>
                </tr>
                <tr>
                    <td style="text-wrap:nowrap;">裝備等級減少</td>
                    <td v-for="i in Array.from(Array(7).keys())">-{{ (i+1)*5 }}</td>
                </tr>
                <tr><th colspan="8" style="padding:4px 0;">武器專用</th></tr>
                <tr>
                    <td>攻擊力</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ Round_to((Math.floor(level/40)+1)*weapon_bonus[i]*100,6) }}%</td>
                </tr>
                <tr>
                    <td>總傷害</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1) }}%</td>
                </tr>
                <tr>
                    <td style="text-wrap:nowrap;">BOSS傷害</td>
                    <td v-for="i in Array.from(Array(7).keys())">{{ (i+1)*2 }}%</td>
                </tr>
            </table>
            <div v-if="240<=level && level<=259" style="color:#ff0000; margin-left:30px;">
                韓版為了抑制永恆裝的強度，單屬跟防禦力需扣掉T數（12 , 24 , 36 , 48 , 60 , 72 , 84）
            </div>
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";
const scope_option_begins = Array.from(Array(14).keys(),x=>x*20);
const scope_option_ends = Array.from(Array(13).keys(),x=>x*20+19); scope_option_ends.push(275);
const scope_choosed = ref(0);
const level = ref(0);
watch(scope_choosed,(new_num)=>{level.value = new_num*20;})
const out_of_range = ref(false)
watch(level,(new_num,old_num)=>{
    if(new_num<0 || new_num>275){
        out_of_range.value = true;
        level.value = old_num;
        let search_area = document.querySelector(".search_area");
        search_area.classList.add("wrong_shake");
        setTimeout(() =>{search_area.classList.remove("wrong_shake");},450);
        setTimeout(() =>{out_of_range.value = false;},1500);
    }
})

const search_as_scope = ref(false);
const change_search_way = ()=>{
    if(!search_as_scope.value){scope_choosed.value = Math.floor(level.value/20);}
    search_as_scope.value = !search_as_scope.value;
}
const input_attack = ref(0);
const input_tier = ref(1);
const weapon_bonus = reactive([0.01 ,0.02 ,0.03 ,0.044 ,0.0605 ,0.07986 ,0.102487]);
watch(input_tier,(new_val)=>{input_tier.value = Number(new_val);})
const Round_to = (num,f)=>{return Math.round(num*10**f)/10**f;}
</script>


<style scoped>
select{
    background-color:#ffffff;
    text-align:end;
    border:#000000 solid 1px;
    border-radius:5px;
    appearance:auto;
}
.scope_select{
    width:130px;
    padding-right:6px;
    font-weight:900;
    font-size:18px;
}
input{
    background-color:#ffffff;
    border:#888888 solid 1px;
    border-radius:6px;
    text-align:center;
}
.level_input{
    width:80px;
    height:34px;
    padding:2px 0;
    font-size:18px;
}
.button{  /* 按鈕(所有應該要可以點擊的東西) */
    cursor:pointer;  /* 滑鼠會變點擊圖樣(手指) */
    transition-duration:0.2s;  /* 因為按鈕都有做滑鼠移入會變色 */
}
.change_search_way_button{
    font-size:14px;
    background-color:#ff99ff88;
    width:100px;
    padding-top:6px;
    padding-bottom:6px;
    margin:0 10px;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000044 inset,-1px 1px 4px 1px #ffffff88 inset;
    white-space:nowrap;
}
.change_search_way_button:hover{
    background-color:#ee11ee88;
}
.search_area{
    border:#cccccc dashed 3px;
    padding:8px 20px;
}
.inside_search_area{
    display:flex;
    align-items:center;
}
@keyframes shake{
    0%{
        margin-left:0;
    }25%{
        margin-left:5px;
    }50%{
        margin-left:0;
    }75%{
        margin-left:-5px;
    }100%{
        margin-left:0;
    }
}
.wrong_shake{
    animation:shake .15s ease-in-out 3;
}
.warning_out_of_range{
    color:#ff0000;
    font-size:14px;
    padding-left:30px;
    white-space:nowrap;
}
table{
    border-collapse:collapse;
}
th{
    background-color:#888888;
    padding:10px 0;
    min-width:120px;
    max-width:120px;
    vertical-align:top;
}
th,td{
    border-right:#dddddd solid 1px;
}
tr:nth-child(2n){
    background-color:#cccccc;
}
tr:nth-child(2n+1){
    background-color:#c0c0c0;
}
td{
    padding:6px 20px;
    text-align:center;
}
.table_title::after{
    content:"";
    position:absolute;
    width:100%;
    height:100%;
    top:0;
    left:0;
    border:#000000 double 3px;
}
.tier_select{
    width:60px;
    font-size:22px;
    margin-top:10px;
}
.attack_input{
    width:70px;
    font-size:18px;
}
</style>
