const vm = new Vue({
    el: '#info',
    data: {
        playlist: []
    },
    mounted() {
        axios.get("http://api.liftedkilt.xyz:8888/playlist")
            .then(response => { this.playlist = response.data.videos })
    }
});
