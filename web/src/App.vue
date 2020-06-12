<template>
  <div id="app">
    <b-tabs fill>
      <b-tab title="List all domains consulted" active><ServerList/></b-tab>
      <b-tab title="Search a new domain"><ServerInfo /></b-tab>
  </b-tabs>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import ServerInfo from './components/ServerInfo.vue'
import ServerList from './components/ServerList.vue'

@Component({
  components: {
    ServerInfo,
    ServerList
  }
})
export default class App extends Vue {}
</script>
