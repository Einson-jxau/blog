<template>
  <v-footer app padless absolute v-if="!this.isMessage">
    <div class="footer-wrap">
      <div>©2019 - {{ new Date().getFullYear() }} By Einson</div>
      <a href="https://beian.miit.gov.cn/" target="_blank">
        赣ICP备20008771号
      </a>
    </div>
  </v-footer>
</template>

<script>
  export default {
    computed: {
      isMessage() {
        return this.$route.path === "/message";
      }
    }
  };
</script>

<style scoped>
  .footer-wrap {
    width: 100%;
    line-height: 2;
    position: relative;
    padding: 20px 20px;
    color: #eee;
    font-size: 14px;
    text-align: center;
    /*background: linear-gradient(-45deg, rgb(20, 20, 226), #17db8f, #23a6d5, #23d5ab);*/
    background: #121212;
    background-size: 400% 400%;
    animation: Gradient 10s ease infinite;
  }

  .footer-wrap a {
    color: #eee !important;
  }

  @keyframes Gradient {
    0% {
      background-position: 0 50%;
    }
    50% {
      background-position: 100% 50%;
    }
    100% {
      background-position: 0 50%;
    }
  }
</style>
