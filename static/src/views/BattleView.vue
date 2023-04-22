<template>
  <MDBRow class="mt-3">
    <MDBCol>
      <h1>The battle
        <MDBBadge color="secondary" pill>{{ battleId }}</MDBBadge>
      </h1>
      <b-content :loading="loading" :error="error" class="mt-3">
        <MDBCol ref="player"/>
      </b-content>
    </MDBCol>
  </MDBRow>
</template>

<script lang="ts">
import api, {GamesDataGameV1} from '@/api'
import {defineComponent} from 'vue'
import {MDBBadge, MDBCol, MDBRow} from "mdb-vue-ui-kit";
import BContent from "@/components/BContent.vue";

// eslint-disable-next-line @typescript-eslint/no-namespace
declare namespace player {
  export class Player {
    constructor(container: HTMLElement, data: GamesDataGameV1);
  }
}

declare global {
  interface Window {
    player: (p: typeof player) => void;
  }
}

export default defineComponent({
  name: 'BattleView',
  components: {MDBBadge, BContent, MDBCol, MDBRow},
  data() {
    return {
      loading: true,
      error: ''
    }
  },
  computed: {
    battleId() {
      return this.$route.params.battleId
    }
  },
  mounted() {
    api.GamesDataV1({
      id: this.$route.params.battleId as string
    }).then(data => {
      window.player = (p: typeof player) => {
        this.loading = false
        this.$nextTick(() => {
          new p.Player((this.$refs.player as { $el: HTMLElement }).$el, data)
        })
      }
      const script = document.createElement('script')
      script.src = '/player/bundle.js'
      document.body.appendChild(script)
    }).catch(err => {
      this.error = err
    })
  }
})
</script>
