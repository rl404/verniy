package verniy

type pageResponse struct {
	Data struct {
		Page Page `json:"Page"`
	} `json:"data"`
}

// Page is pagination response from anilist.
type Page struct {
	PageInfo   PageInfo    `json:"pageInfo"`
	Media      []Media     `json:"media"`
	Characters []Character `json:"characters"`
	Staff      []Staff     `json:"staff"`
	Studios    []Studio    `json:"studios"`
}

func (c *Client) pageQuery(params QueryParam, fields ...PageField) string {
	p := make([]string, len(fields))
	for i := range fields {
		p[i] = string(fields[i])
	}
	return FieldObject("Page", params, p...)
}

func (c *Client) page(page int, perPage int, fields ...PageField) (*Page, error) {
	fields = append(fields, PageFieldPageInfo(
		PageInfoFieldTotal,
		PageInfoFieldPerPage,
		PageInfoFieldCurrentPage,
		PageInfoFieldLastPage,
		PageInfoFieldHasNextPage,
	))

	query := FieldObject("query", QueryParam{
		"$page":    "Int",
		"$perPage": "Int",
	}, c.pageQuery(QueryParam{
		"page":    "$page",
		"perPage": "$perPage",
	}, fields...))

	var d pageResponse
	err := c.post(query, queryVariable{
		"page":    page,
		"perPage": perPage,
	}, &d)
	if err != nil {
		return nil, err
	}

	return &d.Data.Page, nil
}
