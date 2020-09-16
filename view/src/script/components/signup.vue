<template>
  <div>
    <h1>Welcome to Todo Tree!</h1>
    <input v-model="name" placeholder="name">
    <input v-model="email" placeholder="email">
    <input v-model="confirm_email" placeholder="confirm email">
    <button @click='signUp'>Sign Up</button>
  </div>
</template>
<script lang="ts">
  import Vue from "vue"
  import Component from "vue-class-component"
  import Axios from "axios"

  @Component
  export default class Hello extends Vue {
    name: string = ""
    email: string = ""
    confirm_email: string = ""
    
    signUp(): void {
      if(this.email !== this.confirm_email){
        console.log("invalid email");
        return;
      }
      Axios.post('/users', {
          params: {
              Name: this.name,
              Email: this.email
          }
      })
      .then(response => {
          if (response.status != 200) {
              throw new Error('error')
          }
      })
    }
  }
</script>
