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
  name: 'HelloWorld',
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

      if (!this.user)
        this.errors.push("User must be filled in!");

      if (this.questions.length !== Object.keys(this.answers).length)
        this.errors.push("Please choose an answer of all questions!");

      if (this.errors.length > 0) return;

      const self = this;
      Vue.axios
        .post(`${process.env.ROOT_API}quiz/answer`, {
            user: self.user,
            answers: self.answers,
        })
        .then(() => {
          this.$emit('submitted');
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
  margin-top: 20px;
}

.foot > *:not(:first-child) {
  margin-left: 10px;
}

button, input[type="text"] {
    background-color: transparent;
    border: 2px solid #357edd;
    color: #357edd;
    font-family: 'Open Sans', sans-serif;
    font-weight: 600;
    line-height: 1;
    border-radius: 0.5rem;
    font-size: 0.8em;
    min-height: 2rem;
    flex-grow: 1;
    text-overflow: ellipsis;
    padding-left: 6px;
    padding-right: 6px;
    outline: none;
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

.answers {
  display: inline-block;
  margin-bottom: 20px;
  text-align: left;
}

.answers label {
  display: block;
}
</style>
