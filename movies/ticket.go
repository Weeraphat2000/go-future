package movies

func CheckTicketAvailability(movieId string) string {
	switch movieId {
	case "1":
		return "Tickets are available for The Shawshank Redemption."
	case "2":
		return "Tickets are sold out for The Godfather."
	case "3":
		return "Tickets are available for The Dark Knight."
	default:
		return "Movie not found with ID: " + movieId
	}
}
