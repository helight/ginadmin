<template>
  <van-nav-bar
    title="个人介绍"
    left-text="首页"
    right-text="返回"
    @click-left="onClickLeft"
    @click-right="onClickRight"
  />
</template>
<script>
import Vue from 'vue'
import { NavBar } from 'vant'
// import { Toast } from 'vant';
Vue.use(NavBar)

export default {
  methods: {
    onClickLeft() {
      // Toast('返回');
      this.$router.push({
        path: `/home`
      })
    },
    onClickRight() {
      history.back(-1)
    }
  }
}
</script>
