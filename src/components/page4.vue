<template>
    <div style="display:flex;">
        <div>
            <div style="border:#cccccc dashed 3px;">
                <div style="text-align:center; font-size:20px;">
                    查詢等級範圍
                </div>
                <div style="text-align:center; font-size:20px;">
                    <select class="scope_select" v-model="scope_choosed">
                        <option v-for="scope in Object.keys(star).sort()" v-bind:value="scope">{{ scope }}</option>
                    </select>
                </div>
            </div>
            <div style="background-color:#ffff88; text-align:center; padding:5px; margin-top:40px; ">星力小算盤</div>
            <div style="background-color:#ffffbb66; display:flex; align-items:center; padding:5px;">
                
                <select class="star_select" v-model="begin">
                    <option v-for="s in Array.from(Array(10).keys(),x=>x+15)" v-bind:value="s">{{ s }}★</option>
                </select>
                ～
                <select class="star_select" v-model="end">
                    <option v-for="s in Array.from(Array(10).keys(),x=>x+16)" v-bind:value="s">{{ s }}★</option>
                </select>
                <div class="change_if_weapon button" @click="change_if_weapon">
                    <span v-if="is_weapon">武器</span>
                    <span v-if="!is_weapon">非武器</span>
                </div>
            </div>
            <div style="display:flex; justify-content:center; font-size:30px;">⇓</div>
            <div style="display:flex; justify-content:center;">
                <span style="background-color:#ffff88; padding:5px 10px;">總共差</span>
                <span style="background-color:#ffffbb66; padding:5px 10px;">{{ calculated_attack }}攻</span>
                <span style="background-color:#ffffbb66; padding:5px 10px;">{{ calculated_all_stat }}全</span>
            </div>
        </div>
        <div>
            <table style="margin-left:30px;">
                <tr style="position:relative;"><th colspan="4" class="table_title">{{ scope_choosed }}等</th></tr>
                <th>星數</th>
                <th>攻擊(武器)</th>
                <th>攻擊(非武器)</th>
                <th>全屬</th>

                <tr v-for="s in Object.keys(star[scope_choosed]['weapon'])">
                    <td>{{ s }}★
                        <span class="high_light" v-if="begin<s && s<=end"></span>
                        <span class="high_light_all_stat" v-if="begin<s && s<=end"></span>
                    </td>
                    <td><span v-if="star[scope_choosed]['weapon'][s] !== '-'">+</span>{{ star[scope_choosed]['weapon'][s] }}
                        <span class="high_light" v-if="begin<s && s<=end && is_weapon"></span>
                    </td>
                    <td>+{{ star[scope_choosed]['non_weapon'][s] }}
                        <span class="high_light" v-if="begin<s && s<=end && !is_weapon"></span>
                    </td>

                    <td rowspan="7" v-if="s==16">+{{ star[scope_choosed]['all_stat'] }}全</td>
                    <td v-if="s>22"></td>
                </tr>
            </table>
        </div>
    </div>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";

import star from "/src/assets/star.json";
const scope_choosed = ref(Object.keys(star).sort()[0]);
const begin = ref(15);
const end = ref(25);
const is_weapon = ref(true);
const change_if_weapon = ()=>{is_weapon.value = !is_weapon.value;}

const calculated_attack = ref(37);
const calculated_all_stat = ref(35);
watch([scope_choosed,begin,end,is_weapon],(newarr)=>{
    let search = star[newarr[0]];
    let temp
    calculated_attack.value = 0;
    calculated_all_stat.value = 0;
    if(newarr[3]){temp = search['weapon']}
    else{temp = search['non_weapon']}
    for(let i=newarr[1]+1;i<=newarr[2];i++){
        if(typeof(temp[i]) === "number"){calculated_attack.value += temp[i];}
        if(typeof(temp[i]) === "number"){calculated_all_stat.value += search['all_stat']}
    }
})
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
}
table{
    border-collapse:collapse;

}
th{
    background-color:#888888;
    padding:10px 0;
    min-width:130px;
    max-width:130px;
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
    position: relative;
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
.star_select{
    width:60px;
}
.high_light::after{
    content:"";
    position:absolute;
    width:100%;
    height:100%;
    background-color: #ffff0018;
    top:0;
    left:0;
}
.high_light_all_stat::after{
    content:"";
    position:absolute;
    width:100%;
    height:100%;
    background-color: #ffff0018;
    top:0;
    left:390px;
}
.button{  /* 按鈕(所有應該要可以點擊的東西) */
    cursor:pointer;  /* 滑鼠會變點擊圖樣(手指) */
    transition-duration:0.2s;  /* 因為按鈕都有做滑鼠移入會變色 */
}
.change_if_weapon{
    font-weight:900;
    background-color:#ff99ff88;
    padding:2px 0;
    border-radius:10px;
    display:flex;
    justify-content:center;
    align-items:center;
    box-shadow: 2px -2px 4px 1px #00000044 inset,-1px 1px 4px 1px #ffffff88 inset;
    width:70px;
    margin-left:10px;
}
.change_if_weapon:hover{
    background-color:#ee11ee88;
}
</style>