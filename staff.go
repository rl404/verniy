package verniy

type staffResponse struct {
	Data struct {
		Staff Staff `json:"staff"`
	} `json:"data"`
}

func (c *Client) staffQuery(params QueryParam, fields ...StaffField) string {
	p := make([]string, len(fields))
	for i := range fields {
		p[i] = string(fields[i])
	}
	return FieldObject("Staff", params, p...)
}

// GetStaff to get staff & voice actor data.
func (c *Client) GetStaff(id int, fields ...StaffField) (*Staff, error) {
	if len(fields) == 0 {
		fields = []StaffField{
			StaffFieldID,
			StaffFieldName(
				StaffNameFieldFirst,
				StaffNameFieldMiddle,
				StaffNameFieldLast,
				StaffNameFieldFull,
				StaffNameFieldNative,
				StaffNameFieldAlternative),
			StaffFieldImage(StaffImageFieldLarge),
			StaffFieldDescription,
			StaffFieldFavourites,
			StaffFieldAge,
			StaffFieldHomeTown,
			StaffFieldDateOfBirth,
			StaffFieldDateOfDeath,
		}
	}

	query := FieldObject("query", QueryParam{
		"$id": "Int",
	}, c.staffQuery(QueryParam{
		"id": "$id",
	}, fields...))

	var d staffResponse
	err := c.post(query, queryVariable{
		"id": id,
	}, &d)
	if err != nil {
		return nil, err
	}

	return &d.Data.Staff, nil
}

// GetStaffCharacters to get list of character the staff play.
func (c *Client) GetStaffCharacters(id int, page int, perPage int) (*Staff, error) {
	return c.GetStaff(id, StaffFieldCharacterMedia(StaffParamCharacterMedia{
		Page:    page,
		PerPage: perPage,
		Sort:    []MediaSort{MediaSortStartDateDesc},
	},
		MediaConnectionFieldPageInfo(
			PageInfoFieldTotal,
			PageInfoFieldPerPage,
			PageInfoFieldCurrentPage,
			PageInfoFieldLastPage,
			PageInfoFieldHasNextPage),
		MediaConnectionFieldEdges(
			MediaEdgeFieldCharacterRole,
			MediaEdgeFieldNode(
				MediaFieldID,
				MediaFieldType,
				MediaFieldTitle(
					MediaTitleFieldEnglish,
					MediaTitleFieldNative,
					MediaTitleFieldRomaji),
				MediaFieldCoverImage(MediaCoverImageFieldMedium)),
			MediaEdgeFieldCharacters(
				CharacterFieldID,
				CharacterFieldName(CharacterNameFieldFull),
				CharacterFieldImage(CharacterImageFieldMedium)))))
}

func (c *Client) getStaffMedia(id int, mediaType MediaType, page int, perPage int) (*Staff, error) {
	return c.GetStaff(id, StaffFieldStaffMedia(StaffParamStaffMedia{
		Page:    page,
		PerPage: perPage,
		Type:    mediaType,
		Sort:    []MediaSort{MediaSortStartDateDesc},
	},
		MediaConnectionFieldPageInfo(
			PageInfoFieldTotal,
			PageInfoFieldPerPage,
			PageInfoFieldCurrentPage,
			PageInfoFieldLastPage,
			PageInfoFieldHasNextPage),
		MediaConnectionFieldEdges(
			MediaEdgeFieldStaffRole,
			MediaEdgeFieldNode(
				MediaFieldID,
				MediaFieldType,
				MediaFieldTitle(
					MediaTitleFieldEnglish,
					MediaTitleFieldNative,
					MediaTitleFieldRomaji),
				MediaFieldCoverImage(MediaCoverImageFieldMedium)))))
}

// GetStaffAnime to get staff anime list.
func (c *Client) GetStaffAnime(id int, page int, perPage int) (*Staff, error) {
	return c.getStaffMedia(id, TypeAnime, page, perPage)
}

// GetStaffManga to get staff manga list.
func (c *Client) GetStaffManga(id int, page int, perPage int) (*Staff, error) {
	return c.getStaffMedia(id, TypeManga, page, perPage)
}
