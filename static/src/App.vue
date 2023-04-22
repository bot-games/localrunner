<template>
  <MDBNavbar expand="lg" light bg="light">
    <MDBContainer>
      <MDBNavbarBrand href="/">LocalRunner</MDBNavbarBrand>

      <MDBNavbarToggler
        @click="navbarCollapse = !navbarCollapse"
        target="#navbarSupportedContent"
      />

      <MDBCollapse v-model="navbarCollapse" id="navbarSupportedContent">
        <MDBNavbarNav>
          <MDBNavbarItem to="/" :active="$route.path == '/'">Documentation</MDBNavbarItem>
          <MDBNavbarItem to="/battles" :active="$route.path.startsWith(`/battles`)">Battles</MDBNavbarItem>
        </MDBNavbarNav>
      </MDBCollapse>
    </MDBContainer>
  </MDBNavbar>

  <MDBContainer>
    <router-view/>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {
  MDBCollapse,
  MDBContainer,
  MDBNavbar,
  MDBNavbarBrand,
  MDBNavbarItem,
  MDBNavbarNav,
  MDBNavbarToggler
} from 'mdb-vue-ui-kit'

export default defineComponent({
  name: 'App',
  components: {
    MDBCollapse,
    MDBContainer,
    MDBNavbar,
    MDBNavbarBrand,
    MDBNavbarItem,
    MDBNavbarNav,
    MDBNavbarToggler
  },
  data() {
    return {
      navbarCollapse: false
    }
  }
})
</script>

<style lang="scss">
@import '~@/../mdb/scss/index.free.scss';

#app {
  font-family: Roboto, Helvetica, Arial, sans-serif;
}
</style>
