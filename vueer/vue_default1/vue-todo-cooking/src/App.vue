<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
  <div id="App">
    <h1>To-Do List</h1>
    <ul>
      <li>
        <to-do-item-vue label="My ToDo Item" :done="false"></to-do-item-vue>
      </li>
    </ul>
  </div>
</template>

<script>
import toDoItemVue from './components/toDoItem.vue';

export default {
  name: 'App',
  components : {
    toDoItemVue
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
