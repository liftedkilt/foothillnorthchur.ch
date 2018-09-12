const vm = new Vue({
    el: '#info',
    data: {
        playlist: []
    },
    mounted() {
        axios.get("http://localhost:9999/playlist")
            .then(response => { this.playlist = response.data.videos })
    }
});
