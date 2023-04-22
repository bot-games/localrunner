<template>
  <MDBRow class="mt-3">
    <MDBCol>
      <h1>Battles</h1>
    </MDBCol>
  </MDBRow>

  <b-content :loading="gamesLoading" :error="gamesError">
    <MDBCol>
      <MDBTable sm striped v-if="games.length">
        <tbody>
        <tr v-for="game in games" :key="game.id" :class="{'table-warning': game.debug}">
          <td>
            {{ $filters.formatDateTime(game.ts) }}
            <MDBIcon icon="bug" iconStyle="fas" v-if="game.debug" title="Debug game" style="cursor: help"/>
          </td>

          <template v-for="(participant, i) in game.participants" :key="i">
            <td class="text-center" v-if="i>0">
              <MDBIcon icon="times" iconStyle="fas"/>
            </td>

            <td
              :class="{'table-info': game.winner === 0, 'table-success': game.winner === i+1, 'table-danger': game.winner !==0 && game.winner !== i+1}">
              {{ participant.name }}
            </td>

            <td class="text-end"
                :class="{'table-info': game.winner === 0, 'table-success': game.winner === i+1, 'table-danger': game.winner !==0 && game.winner !== i+1}">
              <MDBIcon icon="stopwatch" iconStyle="fas" class="ms-1" style="color: red"
                       v-if="game.finished && (game.timeout & (2**i))"/>
            </td>
          </template>

          <td class="text-center">
            <MDBIcon icon="hourglass-half" iconStyle="fas" size="lg" v-if="!game.finished"/>
            <router-link :to="`/battle/${game.id}`" v-if="game.finished">
              <MDBIcon icon="eye" iconStyle="far" size="lg"/>
            </router-link>
          </td>
        </tr>
        </tbody>
      </MDBTable>
      <div class="alert alert-warning" v-else>
        No battles found!
      </div>
    </MDBCol>
  </b-content>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {MDBCol, MDBIcon, MDBRow, MDBTable} from 'mdb-vue-ui-kit'
import BContent from "@/components/BContent.vue"
import api, {GamesListGameV1} from '@/api'

export default defineComponent({
  name: 'BattlesView',
  components: {
    BContent,
    MDBCol, MDBIcon, MDBRow, MDBTable
  },
  data() {
    return {
      gamesLoading: true,
      gamesError: '',
      games: [] as Array<GamesListGameV1>
    }
  },
  mounted() {
    this.load()
  },
  methods: {
    load(): void {
      this.gamesLoading = true
      this.gamesError = ''

      api.GamesListV1({}).then(games => {
        this.games = games
      }).catch(err => {
        this.gamesError = err
      }).finally(() => {
        this.gamesLoading = false
      })
    }
  }
})
</script>
