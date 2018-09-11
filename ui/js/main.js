const vm = new Vue({
    el: '#info',
    data: {
        playlist: [
            {
              "videoID": "xJxK7xWTqYg",
              "videoTitle": "Our Great God - Fernando Ortega"
            },
            {
              "videoID": "GfAZhvMG2Ew",
              "videoTitle": "His Love Can Never Fail - Kenwood Music"
            },
            {
              "videoID": "i01ZER00X70",
              "videoTitle": "O Sing My Soul"
            },
            {
              "videoID": "gkvQsz7_Ui8",
              "videoTitle": "Creation Sings The Father's Song"
            },
            {
              "videoID": "SuGDU3hAZRk",
              "videoTitle": "Escaping the Darkness"
            },
            {
              "videoID": "iia_owSfv-8",
              "videoTitle": "We Look to You [Official Lyric Video]"
            }
          ]

    // },
    // mounted() {
    //     axios.get("http://192.168.86.92:8889/playlist")
    //         .then(response => { this.playlist = response.data[0] })
    }
});
