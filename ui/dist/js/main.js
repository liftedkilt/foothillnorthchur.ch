const vm = new Vue({
    el: '#info',
    data: {
        playlist: []
    },
    mounted() {
        axios.get("http://192.168.86.92:2015/playlist")
            .then(response => { this.playlist = response.data.videos })
    }
});
