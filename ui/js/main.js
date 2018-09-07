const vm = new Vue({
    el: '#info',
    data: {
        info: []
    },
    mounted() {
        axios.get("http://192.168.86.92:8889/playlist")
            .then(response => { this.info = response.data[0] })
    }
});
