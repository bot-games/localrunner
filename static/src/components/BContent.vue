<template>
  <MDBRow>
    <MDBCol v-if="loading || error">
      <div class="alert alert-danger" v-if="error">
        <strong>Something went wrong.</strong>
        <br><br>
        {{ error }}
      </div>

      <div class="d-flex justify-content-center m-5" v-if="loading && !error">
        <MDBSpinner color="primary" size="5rem"/>
      </div>
    </MDBCol>

    <slot v-if="!loading && !error"></slot>
  </MDBRow>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {MDBCol, MDBRow, MDBSpinner} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'BContent',
  props: {
    loading: Boolean,
    error: String,
  },
  components: {
    MDBCol, MDBRow, MDBSpinner
  }
});
</script>
