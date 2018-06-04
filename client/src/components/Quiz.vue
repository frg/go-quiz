<template>
  <div id="quiz">
    <div v-for="(question, qkey) in questions" :key="qkey">
      <div class="question">{{ question.question }}</div>
      <div class="answers">
        <label v-for="(answer, akey) in question.answers" :key="akey">
          <input type="radio" name="question0" value="a">
          {{ String.fromCharCode(97 + akey) }} : {{ answer }}
        </label>
      </div>
    </div>
    <input type="text" placeholder="Insert your name">
    <button>Submit Quiz</button>
  </div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';

Vue.use(VueAxios, axios);

export default {
  name: 'HelloWorld',
  data() {
    return {
      questions: [],
    };
  },
  methods() {
    // submitQuiz: function() {

    // }
  },
  mounted() {
    Vue.axios
      .get(`${process.env.ROOT_API}quiz`)
      .then((response) => {
        this.questions = response.data.questions;
      });
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.question {
  font-weight: 600;
}

.answers {
  margin-bottom: 20px;
}

.answers label {
  display: block;
}
</style>
