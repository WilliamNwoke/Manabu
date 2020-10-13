<template>
  <div>
    <h1 class="mb-6">Feed Reader</h1>
    <div class="feed-container">
      <FeedList />>
      <ArticleList />
      <SingleArticle />
    </div>
  </div>
</template>

<script>
import FeedList from "./FeedList.vue";
import ArticleList from "./ArticleList.vue";
import SingleArticle from "./SingleArticle.vue";
export default {
  components: { FeedList, ArticleList, SingleArticle },
};
</script>

<style>
.feed-container {
  display: grid;
  grid-template-columns: 15% 15% auto;
  grid-gap: 20px;
}
</style>