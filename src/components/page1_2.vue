<template>
    <table>
        <th style="min-width:80px;"></th>
        <th style="min-width:80px;"><div class="header">顆數</div></th>
        <th style="min-width:100px;"><div class="header">累積顆數</div></th>
        <th><div class="header">賽爾尼溫<img src="/src/assets/imgs/賽爾尼溫.png" class="symbol_img"></div></th>
        <th><div class="header">阿爾克斯<img src="/src/assets/imgs/阿爾克斯.png" class="symbol_img"></div></th>
        <th><div class="header">奧迪溫<img src="/src/assets/imgs/奧迪溫.png" class="symbol_img"></div></th>
        <th><div class="header">桃園境<img src="/src/assets/imgs/桃園境.png" class="symbol_img"></div></th>
        <th><div class="header">阿爾特利亞<img src="/src/assets/imgs/阿爾特利亞.png" class="symbol_img"></div></th>
        <th><div class="header">卡爾西溫<img src="/src/assets/imgs/卡爾西溫.png" class="symbol_img"></div></th>
        <tr v-for="n in Object.values(aut_prices['Cernium']).length" v-bind:key="n">
            <td class="r_head">Lv.{{ n }}</td>
            <td style="text-align:center;">{{ aut_demand['Demand'][n] }}</td>
            <td style="text-align:center;">{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(0,n)) }}</td>
            <td data-cell="消逝的旅途"> {{ formatNumber(aut_prices['Cernium'][n]) }} </td>
            <td data-cell="啾啾愛爾蘭"> {{ formatNumber(aut_prices['Arcus'][n]) }} </td>
            <td data-cell="拉契爾恩"> {{ formatNumber(aut_prices['Odium'][n]) }} </td>
            <td data-cell="阿爾卡娜"> {{ formatNumber(aut_prices['Shangri_La'][n]) }} </td>
            <td data-cell="魔菈斯"> {{ formatNumber(aut_prices['Arteria'][n]) }} </td>
            <td data-cell="艾斯佩拉"> {{ formatNumber(aut_prices['Carcion'][n]) }} </td>
        </tr>
        <tr>
            <td class="total_cost total_cost_head" colspan="3">總金額</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Cernium']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Arcus']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Odium']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Shangri_La']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Arteria']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Carcion']))) }}</td>
        </tr>
    </table>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";

import _aut_prices from "/src/assets/AUT_prices.json";
import _aut_demand from "/src/assets/AUT_demand.json";
var aut_prices = reactive({});
var aut_demand = reactive({});
onBeforeMount(()=>{
    aut_prices = _aut_prices;
    aut_demand = _aut_demand;
})

const Sum_of_Nums = (arr)=>{return arr.reduce((a,b)=>a+b);}
const formatNumber = (num)=>{return num.toString().replace(/(\d)(?=(?:\d{3})+$)/g,"$1,")}
</script>


<style scoped>
table{
    border-collapse:collapse;

}
th{
    background-color:#888888;
    padding:10px 0;
    min-width:130px;
    max-width:130px;
    height:80px;
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
    text-align:end;
}
.symbol_img{
    width:30px;
    height:30px;
    margin:0 300px;
}
.header{
    display:flex;
    justify-content:center;
    align-content:center;
    flex-wrap:wrap;
    height:100%;
}
.r_head{
    background-color:#999999;
    border-top:#dddddd solid 1px;
    text-align: center;
    font-weight:900;
}
.total_cost{
    border-top:#dddddd solid 5px;
    background-color:#aaaaaa;
}
.total_cost_head{
    text-align:center;
    font-weight:900;
    background-color:#888888;
}
</style>
