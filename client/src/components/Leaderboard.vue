<template>
  <div class="leaderboard">
    <h3>Leaderboard</h3>
    <ul v-for="leader in leaderboard">
      <li>{{ leader.user }} - {{ leader.score }} ({{ leader.createdAt }})</li>
    </ul>
  </div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';

Vue.use(VueAxios, axios);

export default {
  name: 'Leaderboard',
  data() {
    return {
      leaderboard: [],
    };
  },
  mounted() {
    Vue.axios
      .get(`${process.env.ROOT_API}quiz/leaderboard`)
      .then((response) => {
        this.leaderboard = response.data;
      });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
