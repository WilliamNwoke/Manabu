<template>
  <div id="app" class="container-xl p-8">
    <FeedReader />
  </div>
</template>

<script>
import FeedReader from "./components/feed-reader/index.vue";

export default {
  name: "app",

  components: { FeedReader },
};
</script>
