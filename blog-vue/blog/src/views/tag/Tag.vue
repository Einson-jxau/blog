<template>
  <div>
    <!-- banner -->
    <div class="tag-banner banner">
      <h1 class="banner-title">标签</h1>
    </div>
    <!-- 标签列表 -->
    <v-card class="blog-container">
      <div class="tag-cloud-title">标签<span style="font-size: medium;color: #6c9d8f"><{{ count }}个></span></div>
      <div class="tag-cloud">
        <router-link
          v-for="item of tagList"
          :key="item.id"
          :style="tagStyle"
          :to="'/tags/' + item.id"
        >
          {{ item.tagName }}
        </router-link>
      </div>
    </v-card>
  </div>
</template>

<script>
  import { getTags } from "../../api/api";

  export default {
    created() {
      this.listTags();
    },
    data: function() {
      return {
        tagStyle: {
          fontSize: Math.floor(Math.random() * 10) + 18 + 'px',
          color: 'rgb(' + Math.round(Math.random() * 255) + ',' + Math.round(Math.random() * 255) + ',' + Math.round(Math.random() * 255) + ')'
        },
        tagList: [],
        count: 0
      };
    },
    methods: {
      listTags() {
        getTags().then((data) => {
          console.log(data);
          this.tagList = data.tags;
          this.count = data.tags.length;
        });
      }
    }
  };
</script>

<style scoped>
  .tag-banner {
    background: #49b1f5 url(https://www.static.talkxj.com/73lleo.png) no-repeat center center;
  }

  .tag-cloud-title {
    line-height: 2;
    font-size: 36px;
    text-align: center;
  }

  @media (max-width: 759px) {
    .tag-cloud-title {
      font-size: 25px;
    }
  }

  .tag-cloud {
    text-align: center;
  }

  .tag-cloud a {
    color: #616161;
    display: inline-block;
    text-decoration: none;
    padding: 0 8px;
    line-height: 2;
    transition: all 0.3s;
  }

  .tag-cloud a:hover {
    color: #03a9f4 !important;
    transform: scale(1.1);
  }
</style>
