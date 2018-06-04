<template>
  <div class="leaderboard">
    <h3>Leaderboard</h3>
    <ol>
      <li v-for="leader in leaderboard">{{ leader.user }} - {{ leader.score }} ({{ leader.createdAt }})</li>
    </ol>
  </div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import Moment from 'moment';

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
        this.leaderboard = response.data.map(element => {
          element.createdAt = Moment.utc(element.createdAt).fromNow();
          return element;
        });
      });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
