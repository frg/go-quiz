<template>
  <div id="app">
    <h1>This Glorious Go-Quiz!</h1>
    <Quiz/>
  </div>
</template>

<script>
import Quiz from './components/Quiz';

export default {
  name: 'App',
  components: {
    Quiz,
  },
};
</script>

<style>
body {
    margin:40px auto;
    max-width:650px;
    line-height:1.6;
    font-size:18px;
    color:#444;
    padding:0 10px
}

h1, h2, h3 {
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
</style>
