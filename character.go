package verniy

type characterResponse struct {
	Data struct {
		Character Character `json:"character"`
	} `json:"data"`
}

func (c *Client) characterQuery(params QueryParam, fields ...CharacterField) string {
	p := make([]string, len(fields))
	for i := range fields {
		p[i] = string(fields[i])
	}
	return FieldObject("Character", params, p...)
}

// GetCharacter to get character data.
func (c *Client) GetCharacter(id int, fields ...CharacterField) (*Character, error) {
	if len(fields) == 0 {
		fields = []CharacterField{
			CharacterFieldID,
			CharacterFieldName(
				CharacterNameFieldFirst,
				CharacterNameFieldMiddle,
				CharacterNameFieldLast,
				CharacterNameFieldFull,
				CharacterNameFieldNative,
				CharacterNameFieldAlternative,
				CharacterNameFieldAlternativeSpoiler),
			CharacterFieldImage(CharacterImageFieldLarge),
			CharacterFieldDescription,
			CharacterFieldGender,
			CharacterFieldDateOfBirth,
			CharacterFieldAge,
			CharacterFieldFavourite,
		}
	}

	query := FieldObject("query", QueryParam{
		"$id": "Int",
	}, c.characterQuery(QueryParam{
		"id": "$id",
	}, fields...))

	var d characterResponse
	err := c.post(query, queryVariable{
		"id": id,
	}, &d)
	if err != nil {
		return nil, err
	}

	return &d.Data.Character, nil
}

// GetCharacterAnime to get list of anime the character play.
func (c *Client) GetCharacterAnime(id int, page int, perPage int) (*Character, error) {
	return c.GetCharacter(id, CharacterFieldMedia(CharacterParamMedia{
		Type:    TypeAnime,
		Page:    page,
		PerPage: perPage,
	},
		MediaConnectionFieldPageInfo(
			PageInfoFieldTotal,
			PageInfoFieldPerPage,
			PageInfoFieldCurrentPage,
			PageInfoFieldLastPage,
			PageInfoFieldHasNextPage),
		MediaConnectionFieldEdges(
			MediaEdgeFieldCharacterRole,
			MediaEdgeFieldVoiceActors(
				MediaEdgeParamVoiceActors{},
				StaffFieldID,
				StaffFieldLanguage,
				StaffFieldName(StaffNameFieldFull),
				StaffFieldImage(StaffImageFieldMedium)),
			MediaEdgeFieldNode(
				MediaFieldID,
				MediaFieldType,
				MediaFieldTitle(
					MediaTitleFieldEnglish,
					MediaTitleFieldNative,
					MediaTitleFieldRomaji),
				MediaFieldCoverImage(
					MediaCoverImageFieldMedium)))))
}

// GetCharacterManga to get list of manga the character play.
func (c *Client) GetCharacterManga(id int, page int, perPage int) (*Character, error) {
	return c.GetCharacter(id, CharacterFieldMedia(CharacterParamMedia{
		Type:    TypeManga,
		Page:    page,
		PerPage: perPage,
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
				MediaFieldCoverImage(
					MediaCoverImageFieldMedium)))))
}
