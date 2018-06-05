<template>
  <div id="quiz" >
    <ul v-if="errors.length" class="errors">
      <li v-for="error in errors">{{ error }}</li>
    </ul>

    <form @submit.prevent="handleSubmit">
      <div v-for="(question, qkey) in questions" :key="qkey">
        <h3 class="question">{{ question.question }}</h3>
        <div class="answers">
          <label v-for="(answer, akey) in question.answers" :key="akey">
            <input type="radio" v-model="answers[question.id]" :name="question.id" :value="akey">
            {{ String.fromCharCode(97 + akey) }} : {{ answer }}
          </label>
        </div>
      </div>
      <div class="foot">
        <input type="text" name="user" placeholder="Insert your name" v-model="user">
        <button>Submit Quiz</button>
      </div>
    </form>
  </div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';

Vue.use(VueAxios, axios);

export default {
  name: 'Quiz',
  data() {
    return {
      questions: [],
      user: null,
      answers: {},
      errors: [],
    };
  },
  methods: {
    handleSubmit: function handleSubmit () {
      this.errors = [];

      // Sanitize the user
      this.user = this.user.trim().split(" ").join("_").replace(/[^a-zA-Z0-9-_]/g, '')

      if (!this.user)
        this.errors.push("User must be filled in!");

      if (this.questions.length !== Object.keys(this.answers).length)
        this.errors.push("Please choose an answer of all questions!");

      if (this.errors.length > 0) return;

      const self = this;
      Vue.axios
        .post(`${process.env.ROOT_API}quiz/answer`, {
            user: this.user,
            answers: self.answers,
        })
        .then(() => {
          this.$emit('goToNext', this.user);
        });
    },
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
.errors {
  margin: 0;
  list-style: none;
  border: 2px solid #f64a4a;
  border-radius: 6px;
  padding: 10px 20px;
}

.foot {
  display: flex;
}

.foot > *:not(:first-child) {
  margin-left: 10px;
}

.answers {
  display: inline-block;
  margin-bottom: 20px;
  text-align: left;
}

.answers label {
  display: block;
}
</style>
