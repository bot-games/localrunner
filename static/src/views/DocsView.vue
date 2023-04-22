<template>
  <div class="swagger" id="swagger"/>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import SwaggerUI from 'swagger-ui'

const HideCurlPlugin = () => {
  return {
    wrapComponents: {
      curl: () => () => null
    }
  }
}

export default defineComponent({
  name: 'DocsView',
  mounted() {
    SwaggerUI({
      dom_id: '#swagger',
      url: '/swagger.json',
      docExpansion: "list",
      defaultModelRendering: 'model',
      defaultModelsExpandDepth: -1,
      deepLinking: true,
      plugins: [HideCurlPlugin]
    })
  }
})
</script>

<style lang="scss">
@import "~swagger-ui/dist/swagger-ui.css";
</style>
