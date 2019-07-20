<template>
  <van-tabbar v-model="active">
    <van-tabbar-item icon="home-o" to="/home">首页</van-tabbar-item>
    <van-tabbar-item icon="search" to="/intro">介绍</van-tabbar-item>
    <van-tabbar-item icon="friends-o" to="/home">预约</van-tabbar-item>
    <van-tabbar-item icon="setting-o" to="/home">我的</van-tabbar-item>
  </van-tabbar>
</template>
<script>
import Vue from 'vue'
import 'vant/lib/index.css'
import { Tabbar, TabbarItem } from 'vant'
Vue.use(Tabbar).use(TabbarItem)

export default {
  data() {
    return {
      active: 0
    }
  }
}
</script>
