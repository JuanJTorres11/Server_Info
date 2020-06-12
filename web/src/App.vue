<template>
  <div id="app">
    <button @click="show = !show">Search for a new domain!</button>
    <ServerInfo v-if="show" />
    <ServerList/>
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
export default class App extends Vue {
  show = false
}
</script>
