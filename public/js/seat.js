var thisFilm = $(this).prev('span');

$.ajax({
	type: "GET",
	url: "/film_info/:film_id",
	data: filmData,
	success: function(data) {
			thisFilm,.html(data);
	}