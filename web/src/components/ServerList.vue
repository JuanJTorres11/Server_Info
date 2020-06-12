<template>
  <div class="ServerList">
    <b-tabs pills card vertical>
      <div v-for="domain in info.items" v-bind:key="domain.domain">
        <b-tab :title="domain.domain">
            <b-card v-if="validLogo(domain.logo)" :title="domain.domain" body-class="max-width:500px" footer="Servers">
              <b-row no-gutters>
                <b-col cols=3>
                  <b-card-img :src="domain.logo" alt="Image" style="max-height:150px; max-width:150px" class="rounded-0" ></b-card-img>
                </b-col>
                <b-col>
                  <b-card-text>
                    <b>SSL Grade: </b> {{domain.ssl_grade}} <br>
                    <b>Title: </b> {{domain.title}} <br>
                    <b>Time: </b> {{domain.time_consulted}} <br>
                  </b-card-text>
                </b-col>
              </b-row>
            </b-card>
            <b-card v-else :title="domain.domain" body-class="max-width:500px" footer="Servers">
              <b-card-text>
                <b>SSL Grade: </b> {{domain.ssl_grade}} <br>
                <b>Title: </b> {{domain.title}} <br>
                <b>Time: </b> {{domain.time_consulted}} <br>
              </b-card-text>
            </b-card>
            <b-card-group deck>
              <div v-for="server in domain.endpoints" v-bind:key="server.address">
                  <b-card :title="server.address">
                    <b-card-text>
                      <b>SSL Grade: </b> {{server.ssl_grade}} <br>
                      <b>Country: </b> {{server.country}} <br>
                      <b>Owner: </b> {{server.owner}} <br>
                    </b-card-text>
                  </b-card>
              </div>
            </b-card-group>
        </b-tab>
      </div>
    </b-tabs>
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

  validLogo (logo: string) {
    if (logo === '') {
      return false
    } else {
      return true
    }
  }
}

</script>
