<template>
  <div class="ServerInfo">
    <p> Enter the domain that you want to consult </p>
    <input v-model="domain">
    <button @click="searchDomain">Send</button>
    <transition name="bounce">
    <h1 v-if="show">{{ info }}</h1>
    </transition>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'

@Component
export default class ServerInfo extends Vue {
  domain = ''
  show = false
  info = null
  searchDomain () {
    axios
      .get('http://localhost:4000/servers_info/' + this.domain)
      .then(response => {
        this.info = response.data
        this.show = true
        console.log(this.info)
      })
  }
}

</script>

<style>
.bounce-enter-active {
  animation: bounce-in .5s;
}
.bounce-leave-active {
  animation: bounce-in .5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>
