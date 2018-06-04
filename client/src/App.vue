<template>
  <div id="app">
    <h1>This Glorious Go-Quiz!</h1>
    <Quiz v-if="step == 1" @goToNext="goToNext" v-model="user"/>
    <Leaderboard v-if="step == 2" :user="user" @goBack="goBack" />
  </div>
</template>

<script>
import Quiz from './components/Quiz';
import Leaderboard from './components/Leaderboard';

export default {
  name: 'App',
  components: {
    Quiz,
    Leaderboard,
  },
  data: function() {
    return {
      step: 1,
      user: null
    }
  },
  methods: {
    goToNext(user) {
      this.step++;
      this.user = user;
    },
    goBack() {
      this.step--;
    }
  }
};
</script>

<style>
body {
  margin: 40px auto;
  max-width: 650px;
  line-height: 1.6;
  font-size: 18px;
  color: #444;
  padding: 0 10px
}

h1,
h2,
h3 {
  line-height:1.2
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

button, input[type="text"] {
  background-color: transparent;
  border: 2px solid #357edd;
  color: #357edd;
  font-family: 'Open Sans', sans-serif;
  font-weight: 600;
  line-height: 2;
  border-radius: 8px;
  font-size: 0.9em;
  min-height: 36px;
  -webkit-box-flex: 1;
  -ms-flex-positive: 1;
  flex-grow: 1;
  text-overflow: ellipsis;
  padding: 4px 16px;
  outline: none;
  margin-top: 20px;
}

button {
  cursor: pointer;
  background-color: #357edd;
  color: white;
}

button:hover {
  opacity: 0.8;
}

::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
    color: #357edd;
    opacity: 1; /* Firefox */
}
</style>
