package verniy

type genreResponse struct {
	Data struct {
		Genres []string `json:"genreCollection"`
	} `json:"data"`
}

// GetGenres to get all genre list.
func (c *Client) GetGenres() ([]string, error) {
	query := FieldObject("query", nil, "GenreCollection")

	var d genreResponse
	err := c.post(query, nil, &d)
	if err != nil {
		return nil, err
	}

	return d.Data.Genres, nil
}
