<template>
  <div class="ServerInfo" id="ServerInfo">
    <b-row>
      <b-col cols=6>
        <input v-model="domain" placeholder="Enter the domain that you want to consult" class="form-control">
      </b-col>
      <b-col cols=2>
        <b-button variant="info" @click="searchDomain">Send</b-button>
        <b-button v-if="show" variant="danger" @click="show = false">Close</b-button>
      </b-col>
    </b-row>
    <transition name="bounce">
    <div v-if="show">
      <b-card v-if="validLogo" :title="domain" bg-variant="info" text-variant="white" footer="Servers" id="img">
        <b-row no-gutters>
          <b-col cols=3>
            <b-card-img :src="info.logo" alt="Image" style="max-height:150px; max-width:150px" class="rounded-0" ></b-card-img>
          </b-col>
          <b-col>
            <b-card-text>
              <b>The servers have changed: </b> {{info.servers_changed}} <br>
              <b>SSL Grade: </b> {{info.ssl_grade}} <br>
              <b>Previoud SSL Grade: </b> {{info.previous_ssl_grade}} <br>
              <b>Title: </b> {{info.title}} <br>
              <b>Is the domain down: </b> {{info.is_down}} <br>
            </b-card-text>
          </b-col>
        </b-row>
      </b-card>
      <b-card v-else :title="domain" bg-variant="info" text-variant="white" footer="Servers">
        <b-card-text>
          <b>The servers have changed: </b> {{info.servers_changed}} <br>
          <b>SSL Grade: </b> {{info.ssl_grade}} <br>
          <b>Previoud SSL Grade: </b> {{info.previous_ssl_grade}} <br>
          <b>Title: </b> {{info.title}} <br>
          <b>Is the domain down: </b> {{info.is_down}} <br>
        </b-card-text>
      </b-card>
       <b-card-group deck>
         <div v-for="server in info.endpoints" v-bind:key="server.address">
            <b-card :title="server.address" border-variant="success">
              <b-card-text>
                <b>SSL Grade: </b> {{server.ssl_grade}} <br>
                <b>Country: </b> {{server.country}} <br>
                <b>Owner: </b> {{server.owner}} <br>
              </b-card-text>
            </b-card>
         </div>
       </b-card-group>
    </div>
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
  info: any = {}
  searchDomain () {
    axios
      .get('http://localhost:4000/servers_info/' + this.domain)
      .then(response => {
        this.info = response.data
        this.show = true
        console.log(this.info)
      })
  }

  validLogo () {
    if (this.info!.logo === '') {
      return false
    } else {
      return true
    }
  }
}

</script>

<style>
#ServerInfo {
  margin-left: 30px
}
#img {
  width: 500px;
}
.card {
  width: 400px;
  margin:20px
}
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
