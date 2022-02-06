<template>
  <div class="app-wrapper">
    <app-main/>
  </div>
</template>

<script>
import {AppMain} from './components'
import ResizeMixin from './mixin/ResizeHandler'

export default {
  name: 'Layout',
  components: {
    AppMain
  },
  mixins: [ResizeMixin],
  computed: {
    device() {
      return 'desktop'
    },
  },
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";
@import "~@/styles/variables.scss";

.app-wrapper {
  @include clearfix;
  position: relative;
  height: 100%;
  width: 100%;
}

</style>
