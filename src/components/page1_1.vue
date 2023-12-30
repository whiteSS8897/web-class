<template>
    <table>
        <th style="min-width:80px;"></th>
        <th style="min-width:80px;"><div class="header">顆數</div></th>
        <th style="min-width:100px;"><div class="header">累積顆數</div></th>
        <th><div class="header">消逝的旅途<img src="/src/assets/imgs/消逝的旅途.png" class="symbol_img"></div></th>
        <th><div class="header">啾啾愛爾蘭<img src="/src/assets/imgs/啾啾愛爾蘭.png" class="symbol_img"></div></th>
        <th><div class="header">拉契爾恩<img src="/src/assets/imgs/拉契爾恩.png" class="symbol_img"></div></th>
        <th><div class="header">阿爾卡娜<img src="/src/assets/imgs/阿爾卡娜.png" class="symbol_img"></div></th>
        <th><div class="header">魔菈斯<img src="/src/assets/imgs/魔菈斯.png" class="symbol_img"></div></th>
        <th><div class="header">艾斯佩拉<img src="/src/assets/imgs/艾斯佩拉.png" class="symbol_img"></div></th>
        <tr v-for="n in Object.values(arc_prices['Vanishing_Journey']).length" v-bind:key="n">
            <td class="r_head">Lv.{{ n }}</td>
            <td style="text-align:center;">{{ arc_demand['Demand'][n] }}</td>
            <td style="text-align:center;">{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(0,n)) }}</td>
            <td data-cell="消逝的旅途"> {{ formatNumber(arc_prices['Vanishing_Journey'][n]) }} </td>
            <td data-cell="啾啾愛爾蘭"> {{ formatNumber(arc_prices['Chu_Chu_Island'][n]) }} </td>
            <td data-cell="拉契爾恩"> {{ formatNumber(arc_prices['Lachelein'][n]) }} </td>
            <td data-cell="阿爾卡娜"> {{ formatNumber(arc_prices['Arcana'][n]) }} </td>
            <td data-cell="魔菈斯"> {{ formatNumber(arc_prices['Morass'][n]) }} </td>
            <td data-cell="艾斯佩拉"> {{ formatNumber(arc_prices['Esfera'][n]) }} </td>
        </tr>
        <tr>
            <td class="total_cost total_cost_head" colspan="3">總金額</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Vanishing_Journey']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Chu_Chu_Island']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Lachelein']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Arcana']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Morass']))) }}</td>
            <td class="total_cost">{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Esfera']))) }}</td>
        </tr>
    </table>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";

import _arc_prices from "/src/assets/ARC_prices.json";
import _arc_demand from "/src/assets/ARC_demand.json";
var arc_prices = reactive({});
var arc_demand = reactive({});
onBeforeMount(()=>{
    arc_prices = _arc_prices;
    arc_demand = _arc_demand;
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
    text-align:center;
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
