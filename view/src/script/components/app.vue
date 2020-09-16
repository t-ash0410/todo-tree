<template>
  <div id="app-2">
    <signup></signup>
  </div>
</template>
<script lang="ts">
  import Vue from "vue"
  export default class App extends Vue { }

  import signup from "./signup.vue"
  Vue.component('signup', signup)
</script>
