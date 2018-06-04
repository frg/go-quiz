<template>
  <div class="leaderboard">
    <h2>Leaderboard</h2>
    <ol>
      <li v-for="leader in leaderboard">
        <div v-if="leader !== null" v-bind:class="{ isUser: leader.user == user }"><span class="user">{{ leader.user }}</span> - {{ leader.score }}pts <time :datetime="leader.createdAt">{{ leader.createdAtHuman }}</time></div>
      </li>
    </ol>

    <ul>
      <li v-for="fact in facts">{{ fact }}</li>
    </ul>

    <button v-on:click="goBack">Submit Another!</button>
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
  props: ['user'],
  data() {
    return {
      leaderboard: [],
      facts: [],
    };
  },
  methods: {
    goBack: function() {
      this.$emit('goBack');
    }
  },
  mounted() {
    Vue.axios
      .get(`${process.env.ROOT_API}quiz/leaderboard`)
      .then((response) => {
        let leaderboard = response.data.map(element => {
          element.createdAtHuman = Moment.utc(element.createdAt).fromNow();
          return element;
        });

        for (var i = leaderboard.length; i < 10; i++) {
          leaderboard.push(null);
        }

        this.leaderboard = leaderboard;

        Vue.axios
            .get(`${process.env.ROOT_API}quiz/leaderboard/relative/${this.user}`)
            .then((response) => {
              this.facts = response.data;
            });
      });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.isUser {
  color: #357edd;
}

.user {
  font-weight: bold;
}

ol {
  position: relative;
}

li time {
  font-size: 14px;
  margin-left: 12px;
}

.leaderboard {
  font-size: 20px;
}
</style>
