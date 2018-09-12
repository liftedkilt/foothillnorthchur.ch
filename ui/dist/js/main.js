const vm = new Vue({
    el: '#info',
    data: {
        playlist: []
    },
    mounted() {
        axios.get("https://api.liftedkilt.xyz/playlist")
            .then(response => { this.playlist = response.data.videos })
    }
});
