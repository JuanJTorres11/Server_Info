<template>
  <div class="ServerList">
    {{ info }}
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import axios from 'axios'

@Component
export default class ServerList extends Vue {
  info = null
  mounted () {
    axios
      .get('http://localhost:4000/servers')
      .then(response => {
        this.info = response.data
        console.log('Lista es ' + this.info)
      })
  }
}

</script>
