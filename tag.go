package verniy

type tagResponse struct {
	Data struct {
		Tags []MediaTag `json:"mediaTagCollection"`
	} `json:"data"`
}

func (c *Client) tagQuery(fields ...MediaTagField) string {
	p := make([]string, len(fields))
	for i := range fields {
		p[i] = string(fields[i])
	}
	return FieldObject("MediaTagCollection", nil, p...)
}

// GetTags to get all tag list.
func (c *Client) GetTags(fields ...MediaTagField) ([]MediaTag, error) {
	if len(fields) == 0 {
		fields = []MediaTagField{
			MediaTagFieldName,
			MediaTagFieldDescription,
			MediaTagFieldCategory,
			MediaTagFieldIsAdult,
		}
	}

	query := FieldObject("query", nil, c.tagQuery(fields...))

	var d tagResponse
	err := c.post(query, nil, &d)
	if err != nil {
		return nil, err
	}

	return d.Data.Tags, nil
}
