<template>
  <div id="mymain">
    <!-- swiper -->
    <section class="swiper-container">
      <div class="swiper-wrapper t-c">
        <div class="swiper-slide"><img src="../../assets/logo1.png"></div>
      </div>
    </section>
    <div class="t-l f24 pl1 pt2 pb2 pr1 bg-color9 color8 clearfix">
      <div class="fl" />
      <div id="demo" class="o-h fl" style="height:22px;line-height:22px;" />
    </div>
    <!-- 导航 -->
    <section class="demo-nav f40 clearfix">
      <div class="col-12 t-c pt2 pb2" style="border-left:none;">
        <router-link to="/home">
          <i class="color-success icon iconfont icon-xiangqingjieshao" />
          个人介绍
        </router-link>
      </div>
      <div class="col-12 t-c pt2 pb2" style="border-left:none;">
        <router-link to="/home">
          <i class="color-success icon iconfont icon-yuyuezhongxin" />
          项目介绍
        </router-link>
      </div>
      <div class="col-12 t-c pt2 pb2" style="border-left:none;">
        <router-link to="/intro">
          <i class="color-success icon iconfont icon-daohang" />
          人员介绍
        </router-link>
      </div>
      <div class="col-12 t-c pt2 pb2" style="border-left:none;">
        <router-link to="/intro">
          <i class="color-success icon iconfont icon-daohang" />
          人员介绍
        </router-link>
      </div>
    </section>
    <FootBody />
  </div>
</template>
<style scoped>
    .demo-nav {
        line-height: 1.8em;
    }

    .demo-nav div {
        border-left: 1px solid #ccc;
        border-bottom: 1px solid #ccc;
    }

    .demo-nav a {
        display: block;
        width: 100%;
        height: 100%;
    }

    .demo-nav div:nth-child(4n+1) {
        border-left: none;
    }
</style>
<style scoped>
@import '~@/assets/css/iconfont.css';
@import '~@/assets/css/miniMoblie.css';

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
<script>
import FootBody from '../common/foot.vue'

export default {
  name: 'Mymain',
  components: {
    FootBody
  }
}
</script>
