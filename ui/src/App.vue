<template>
<div id="app">
  <nav>
    <ul class="navbar">
      <li id="home">Foothill North</li>
    </ul>
  </nav>

<PlaylistBox />

</div>
</template>

<script>
import PlaylistBox from './components/PlaylistBox.vue';

export default {
  name: 'app',
  components: {
    PlaylistBox,
  },
};
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

a {
  text-decoration: none;
}

.navbar {
  color: #2b7554;
  display: grid;
  justify-content: left;
  border-bottom: 1px solid black;
  overflow: hidden;
  position: fixed;
  top: 0;
  width: 100%;
  background-color: white;
  margin-top: 0px;
  height: 40px;
  padding-top: 10px;

  li {
    list-style-type: none;
    text-align: left;
  }

  #home {
    font-size: 24px;
  }
}

</style>
