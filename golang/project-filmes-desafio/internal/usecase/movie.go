package usecase

import (
	"bufio"
	"encoding/csv"
	"os"
	"project-filmes-desafio/internal/schemas"
	"strconv"
)

func readMoviesFromFile(filePath string) ([]schemas.Movie, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	customReader := csv.NewReader(bufio.NewReader(file))
	customReader.Comma = ','

	_, err = customReader.Read()
	if err != nil {
		return nil, err
	}

	var movies []schemas.Movie

	for {
		record, err := customReader.Read()
		if err != nil {
			break
		}

		movieID := record[0]
		title := record[1]
		genres := record[2]

		movieIDInt, _ := strconv.Atoi(movieID)

		movie := schemas.Movie{
			MovieID: movieIDInt,
			Title:   title,
			Genres:  genres,
		}

		movies = append(movies, movie)
	}

	return movies, nil
}
