<template>
    <table style="display:contents;">
        <th>符文</th>
        <th>幾等到幾等</th>
        <th>顆數</th>
        <th>金額</th>
        <tr>
            <td><div class="choice_label">消逝的旅途<img src="/src/assets/imgs/消逝的旅途.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Vanishing_Journey_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Vanishing_Journey_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(Vanishing_Journey_begin-1,Vanishing_Journey_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Vanishing_Journey']).slice(Vanishing_Journey_begin-1,Vanishing_Journey_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">啾啾愛爾蘭<img src="/src/assets/imgs/啾啾愛爾蘭.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Chu_Chu_Island_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Chu_Chu_Island_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(Chu_Chu_Island_begin-1,Chu_Chu_Island_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Chu_Chu_Island']).slice(Chu_Chu_Island_begin-1,Chu_Chu_Island_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">拉契爾恩<img src="/src/assets/imgs/拉契爾恩.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Lachelein_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Lachelein_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(Lachelein_begin-1,Lachelein_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Lachelein']).slice(Lachelein_begin-1,Lachelein_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">阿爾卡娜<img src="/src/assets/imgs/阿爾卡娜.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Arcana_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Arcana_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(Arcana_begin-1,Arcana_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Arcana']).slice(Arcana_begin-1,Arcana_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">魔菈斯<img src="/src/assets/imgs/魔菈斯.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Morass_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Morass_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(Morass_begin-1,Morass_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Morass']).slice(Morass_begin-1,Morass_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">艾斯佩拉<img src="/src/assets/imgs/艾斯佩拉.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Esfera_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Esfera_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(arc_demand['Demand']).slice(Esfera_begin-1,Esfera_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(arc_prices['Esfera']).slice(Esfera_begin-1,Esfera_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">賽爾尼溫<img src="/src/assets/imgs/賽爾尼溫.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Cernium_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Cernium_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(Cernium_begin-1,Cernium_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Cernium']).slice(Cernium_begin-1,Cernium_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">阿爾克斯<img src="/src/assets/imgs/阿爾克斯.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Arcus_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Arcus_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(Arcus_begin-1,Arcus_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Arcus']).slice(Arcus_begin-1,Arcus_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">奧迪溫<img src="/src/assets/imgs/奧迪溫.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Odium_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Odium_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(Odium_begin-1,Odium_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Odium']).slice(Odium_begin-1,Odium_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">桃園境<img src="/src/assets/imgs/桃園境.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Shangri_La_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Shangri_La_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(Shangri_La_begin-1,Shangri_La_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Shangri_La']).slice(Shangri_La_begin-1,Shangri_La_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">阿爾特利亞<img src="/src/assets/imgs/阿爾特利亞.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Arteria_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Arteria_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(Arteria_begin-1,Arteria_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Arteria']).slice(Arteria_begin-1,Arteria_end-1))) }}</td>
        </tr>
        <tr>
            <td><div class="choice_label">卡爾西溫<img src="/src/assets/imgs/卡爾西溫.png" class="symbol_img"></div></td>
            <td>
                <div class="input_level_area">
                    <input type=number class="input_level" v-model="Carcion_begin">
                    <div style="font-size:24px;">～</div>
                    <input type=number class="input_level" v-model="Carcion_end">
                </div>
            </td>
            <td>{{ Sum_of_Nums(Object.values(aut_demand['Demand']).slice(Carcion_begin-1,Carcion_end-1)) }}</td>
            <td>{{ formatNumber(Sum_of_Nums(Object.values(aut_prices['Carcion']).slice(Carcion_begin-1,Carcion_end-1))) }}</td>
        </tr>
        <tr>
            <td colspan="3" class="total_cost total_cost_head">總金額</td>
            <td style="text-align:end;" class="total_cost">{{ formatNumber(
                Sum_of_Nums(Object.values(arc_prices['Vanishing_Journey']).slice(Vanishing_Journey_begin-1,Vanishing_Journey_end-1))+
                Sum_of_Nums(Object.values(arc_prices['Chu_Chu_Island']).slice(Chu_Chu_Island_begin-1,Chu_Chu_Island_end-1))+
                Sum_of_Nums(Object.values(arc_prices['Lachelein']).slice(Lachelein_begin-1,Lachelein_end-1))+
                Sum_of_Nums(Object.values(arc_prices['Arcana']).slice(Arcana_begin-1,Arcana_end-1))+
                Sum_of_Nums(Object.values(arc_prices['Morass']).slice(Morass_begin-1,Morass_end-1))+
                Sum_of_Nums(Object.values(arc_prices['Esfera']).slice(Esfera_begin-1,Esfera_end-1))+
                Sum_of_Nums(Object.values(aut_prices['Cernium']).slice(Cernium_begin-1,Cernium_end-1))+
                Sum_of_Nums(Object.values(aut_prices['Arcus']).slice(Arcus_begin-1,Arcus_end-1))+
                Sum_of_Nums(Object.values(aut_prices['Odium']).slice(Odium_begin-1,Odium_end-1))+
                Sum_of_Nums(Object.values(aut_prices['Shangri_La']).slice(Shangri_La_begin-1,Shangri_La_end-1))+
                Sum_of_Nums(Object.values(aut_prices['Arteria']).slice(Arteria_begin-1,Arteria_end-1))+
                Sum_of_Nums(Object.values(aut_prices['Carcion']).slice(Carcion_begin-1,Carcion_end-1))
                ) }}</td>
        </tr>
    </table>
</template>


<script setup>
import {computed, reactive, ref, watch, onBeforeMount} from "vue";

import _arc_prices from "/src/assets/ARC_prices.json";
import _arc_demand from "/src/assets/ARC_demand.json";
import _aut_prices from "/src/assets/AUT_prices.json";
import _aut_demand from "/src/assets/AUT_demand.json";
var arc_prices = reactive({});
var arc_demand = reactive({});
var aut_prices = reactive({});
var aut_demand = reactive({});
onBeforeMount(()=>{
    arc_prices = _arc_prices;
    arc_demand = _arc_demand;
    aut_prices = _aut_prices;
    aut_demand = _aut_demand;
})

const Sum_of_Nums = (arr)=>{
    if(arr.length === 0){return 0;}
    return arr.reduce((a,b)=>a+b);
}
const formatNumber = (num)=>{return num.toString().replace(/(\d)(?=(?:\d{3})+$)/g,"$1,")}

const Vanishing_Journey_begin = ref(1);
const Vanishing_Journey_end = ref(1);
const Chu_Chu_Island_begin = ref(1);
const Chu_Chu_Island_end = ref(1);
const Lachelein_begin = ref(1);
const Lachelein_end = ref(1);
const Arcana_begin = ref(1);
const Arcana_end = ref(1);
const Morass_begin = ref(1);
const Morass_end = ref(1);
const Esfera_begin = ref(1);
const Esfera_end = ref(1);
const Cernium_begin = ref(1);
const Cernium_end = ref(1);
const Arcus_begin = ref(1);
const Arcus_end = ref(1);
const Odium_begin = ref(1);
const Odium_end = ref(1);
const Shangri_La_begin = ref(1);
const Shangri_La_end = ref(1);
const Arteria_begin = ref(1);
const Arteria_end = ref(1);
const Carcion_begin = ref(1);
const Carcion_end = ref(1);

watch(Vanishing_Journey_begin,(new_num)=>{
    if(new_num < 1){Vanishing_Journey_begin.value=1;}
    if(new_num > 19){Vanishing_Journey_begin.value=19;}
    if(new_num > Vanishing_Journey_end.value){Vanishing_Journey_end.value=Vanishing_Journey_begin.value;}
})
watch(Vanishing_Journey_end,(new_num)=>{
    if(new_num < 1){Vanishing_Journey_end.value=1;}
    if(new_num > 20){Vanishing_Journey_end.value=20;}
    if(new_num < Vanishing_Journey_begin.value){Vanishing_Journey_begin.value=Vanishing_Journey_end.value;}
})

watch(Chu_Chu_Island_begin,(new_num)=>{
    if(new_num < 1){Chu_Chu_Island_begin.value=1;}
    if(new_num > 19){Chu_Chu_Island_begin.value=19;}
    if(new_num > Chu_Chu_Island_end.value){Chu_Chu_Island_end.value=Chu_Chu_Island_begin.value;}
})
watch(Chu_Chu_Island_end,(new_num)=>{
    if(new_num < 1){Chu_Chu_Island_end.value=1;}
    if(new_num > 20){Chu_Chu_Island_end.value=20;}
    if(new_num < Chu_Chu_Island_begin.value){Chu_Chu_Island_begin.value=Chu_Chu_Island_end.value;}
})

watch(Lachelein_begin,(new_num)=>{
    if(new_num < 1){Lachelein_begin.value=1;}
    if(new_num > 19){Lachelein_begin.value=19;}
    if(new_num > Lachelein_end.value){Lachelein_end.value=Lachelein_begin.value;}
})
watch(Lachelein_end,(new_num)=>{
    if(new_num < 1){Lachelein_end.value=1;}
    if(new_num > 20){Lachelein_end.value=20;}
    if(new_num < Lachelein_begin.value){Lachelein_begin.value=Lachelein_end.value;}
})

watch(Arcana_begin,(new_num)=>{
    if(new_num < 1){Arcana_begin.value=1;}
    if(new_num > 19){Arcana_begin.value=19;}
    if(new_num > Arcana_end.value){Arcana_end.value=Arcana_begin.value;}
})
watch(Arcana_end,(new_num)=>{
    if(new_num < 1){Arcana_end.value=1;}
    if(new_num > 20){Arcana_end.value=20;}
    if(new_num < Arcana_begin.value){Arcana_begin.value=Arcana_end.value;}
})

watch(Morass_begin,(new_num)=>{
    if(new_num < 1){Morass_begin.value=1;}
    if(new_num > 19){Morass_begin.value=19;}
    if(new_num > Morass_end.value){Morass_end.value=Morass_begin.value;}
})
watch(Morass_end,(new_num)=>{
    if(new_num < 1){Morass_end.value=1;}
    if(new_num > 20){Morass_end.value=20;}
    if(new_num < Morass_begin.value){Morass_begin.value=Morass_end.value;}
})

watch(Esfera_begin,(new_num)=>{
    if(new_num < 1){Esfera_begin.value=1;}
    if(new_num > 19){Esfera_begin.value=19;}
    if(new_num > Esfera_end.value){Esfera_end.value=Esfera_begin.value;}
})
watch(Esfera_end,(new_num)=>{
    if(new_num < 1){Esfera_end.value=1;}
    if(new_num > 20){Esfera_end.value=20;}
    if(new_num < Esfera_begin.value){Esfera_begin.value=Esfera_end.value;}
})

watch(Cernium_begin,(new_num)=>{
    if(new_num < 1){Cernium_begin.value=1;}
    if(new_num > 10){Cernium_begin.value=10;}
    if(new_num > Cernium_end.value){Cernium_end.value=Cernium_begin.value;}
})
watch(Cernium_end,(new_num)=>{
    if(new_num < 1){Cernium_end.value=1;}
    if(new_num > 11){Cernium_end.value=11;}
    if(new_num < Cernium_begin.value){Cernium_begin.value=Cernium_end.value;}
})

watch(Arcus_begin,(new_num)=>{
    if(new_num < 1){Arcus_begin.value=1;}
    if(new_num > 10){Arcus_begin.value=10;}
    if(new_num > Arcus_end.value){Arcus_end.value=Arcus_begin.value;}
})
watch(Arcus_end,(new_num)=>{
    if(new_num < 1){Arcus_end.value=1;}
    if(new_num > 11){Arcus_end.value=11;}
    if(new_num < Arcus_begin.value){Arcus_begin.value=Arcus_end.value;}
})

watch(Odium_begin,(new_num)=>{
    if(new_num < 1){Odium_begin.value=1;}
    if(new_num > 10){Odium_begin.value=10;}
    if(new_num > Odium_end.value){Odium_end.value=Odium_begin.value;}
})
watch(Odium_end,(new_num)=>{
    if(new_num < 1){Odium_end.value=1;}
    if(new_num > 11){Odium_end.value=11;}
    if(new_num < Odium_begin.value){Odium_begin.value=Odium_end.value;}
})

watch(Shangri_La_begin,(new_num)=>{
    if(new_num < 1){Shangri_La_begin.value=1;}
    if(new_num > 10){Shangri_La_begin.value=10;}
    if(new_num > Shangri_La_end.value){Shangri_La_end.value=Shangri_La_begin.value;}
})
watch(Shangri_La_end,(new_num)=>{
    if(new_num < 1){Shangri_La_end.value=1;}
    if(new_num > 11){Shangri_La_end.value=11;}
    if(new_num < Shangri_La_begin.value){Shangri_La_begin.value=Shangri_La_end.value;}
})

watch(Arteria_begin,(new_num)=>{
    if(new_num < 1){Arteria_begin.value=1;}
    if(new_num > 10){Arteria_begin.value=10;}
    if(new_num > Arteria_end.value){Arteria_end.value=Arteria_begin.value;}
})
watch(Arteria_end,(new_num)=>{
    if(new_num < 1){Arteria_end.value=1;}
    if(new_num > 11){Arteria_end.value=11;}
    if(new_num < Arteria_begin.value){Arteria_begin.value=Arteria_end.value;}
})

watch(Carcion_begin,(new_num)=>{
    if(new_num < 1){Carcion_begin.value=1;}
    if(new_num > 10){Carcion_begin.value=10;}
    if(new_num > Carcion_end.value){Carcion_end.value=Carcion_begin.value;}
})
watch(Carcion_end,(new_num)=>{
    if(new_num < 1){Carcion_end.value=1;}
    if(new_num > 11){Carcion_end.value=11;}
    if(new_num < Carcion_begin.value){Carcion_begin.value=Carcion_end.value;}
})
</script>


<style scoped>
.symbol_img{
    width:40px;
    height:40px;
}
.choice_label{
    display:flex;
    align-items:center;
    justify-content:end;
    padding:5px 0;
    font-size:18px;
    font-weight:700;
}
.input_level{
    background-color:#ffffff;
    border:#888888 solid 1px;
    border-radius:6px;
    width:50px;
    height:34px;
    padding:2px 0;
    font-size:18px;
    text-align:center;
}
.input_level_area{
    display:flex;
    align-items:center;
    height:50px;
}

table{
    border-collapse:collapse;

}
th{
    background-color:#888888;
    padding:10px 0;
    min-width:130px;
    max-width:130px;
    vertical-align:middle;
}
th:nth-child(1){
    min-width:200px;
    max-width:200px;
}
th:nth-child(3){
    min-width:80px;
    max-width:80px;
}
th,td{
    border-right:#dddddd solid 1px;
}
tr:nth-child(2n+1){
    background-color:#c0c0c0;
}
tr:nth-child(2n){
    background-color:#cccccc;
}
td{
    padding:6px 20px;
    text-align:center;
}
td:nth-child(4){
    text-align:end;
    width:160px;
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
