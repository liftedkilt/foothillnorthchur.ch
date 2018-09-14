<template>
  <div class="playlistbox">
    <!-- <div v-for="song in playlist"> -->
    <a v-for="song in playlist" :key="song.videoID" :href="'https://liftedkilt.github.io/' + song.videoID"><div class="v-thumb"><img v-bind:src="'https://images.liftedkilt.xyz/q100/https://i.ytimg.com/vi/' + song.videoID + '/mqdefault.jpg'"></div><div class="v-title">{{ song.videoTitle }}</div></a>
  <!-- </div> -->
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'PlaylistBox',
  props: {
    msg: String,
  },
  data() {
    return {
      playlist: [],
    };
  },
  created() {
    axios.get('https://api.liftedkilt.xyz/playlist')
      .then((response) => {
        // JSON responses are automatically parsed.
        this.playlist = response.data.videos;
      });
  },
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.playlistbox {
  margin-top: 5%;
}
img {
  max-width: 75%;
  height: auto;
}
a {
  text-decoration: none;
  color: #050;
}

.playlistbox {
  display: grid;
  margin-left: 3%;
  margin-right: 5%;
  margin-top: 100px;
  grid-template-columns: 4fr 6fr;

  a {
    display: grid;
    grid-template-columns: 4fr 6fr;
    grid-column: 1 / 3;
  }

  .v-image {
    grid-column: 1 / 2;
  }

  .v-title {
    grid-column: 2 / 3 ;
    align-self: center;
    text-align: left;
    color: #050;
    padding-left: 10px;
  }
}
</style>
