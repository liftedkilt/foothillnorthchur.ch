$.get("http://localhost:8888/playlist", function (response) {
    // store game board size
    song_one_title = response.data
});

var song1 = new Vue({
	el: '#button-2',
	data: {
		message: 'Song 2'
	}
})

var song2 = new Vue({
	el: '#button-3',
	data: {
		message: 'Song 3'
	}
})