package verniy

type mediaResponse struct {
	Data struct {
		Media Media `json:"media"`
	} `json:"data"`
}

func (c *Client) mediaQuery(params QueryParam, fields ...MediaField) string {
	p := make([]string, len(fields))
	for i := range fields {
		p[i] = string(fields[i])
	}
	return FieldObject("Media", params, p...)
}

func (c *Client) getMedia(id int, mediaType MediaType, fields ...MediaField) (*Media, error) {
	query := FieldObject("query", QueryParam{
		"$id":   "Int",
		"$type": "MediaType",
	}, c.mediaQuery(QueryParam{
		"id":   "$id",
		"type": "$type",
	}, fields...))

	var d mediaResponse
	err := c.post(query, queryVariable{
		"id":   id,
		"type": mediaType,
	}, &d)
	if err != nil {
		return nil, err
	}

	return &d.Data.Media, nil
}

// GetAnime to get anime data.
func (c *Client) GetAnime(id int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldID,
			MediaFieldTitle(
				MediaTitleFieldRomaji,
				MediaTitleFieldEnglish,
				MediaTitleFieldNative),
			MediaFieldType,
			MediaFieldFormat,
			MediaFieldStatusV2,
			MediaFieldDescription,
			MediaFieldDescription,
			MediaFieldStartDate,
			MediaFieldEndDate,
			MediaFieldCountryOfOrigin,
			MediaFieldSourceV2,
			MediaFieldCoverImage(
				MediaCoverImageFieldExtraLarge,
				MediaCoverImageFieldLarge,
				MediaCoverImageFieldMedium,
				MediaCoverImageFieldColor),
			MediaFieldBannerImage,
			MediaFieldGenres,
			MediaFieldSynonyms,
			MediaFieldAverageScore,
			MediaFieldMeanScore,
			MediaFieldPopularity,
			MediaFieldFavourites,
			MediaFieldTags(
				MediaTagFieldName,
				MediaTagFieldDescription,
				MediaTagFieldCategory,
				MediaTagFieldRank,
				MediaTagFieldIsGeneralSpoiler,
				MediaTagFieldIsMediaSpoiler,
				MediaTagFieldIsAdult),
			MediaFieldRelations(
				MediaConnectionFieldPageInfo(
					PageInfoFieldTotal,
					PageInfoFieldPerPage,
					PageInfoFieldCurrentPage,
					PageInfoFieldLastPage,
					PageInfoFieldHasNextPage),
				MediaConnectionFieldEdges(
					MediaEdgeFieldRelationTypeV2,
					MediaEdgeFieldNode(
						MediaFieldID,
						MediaFieldTitle(
							MediaTitleFieldRomaji,
							MediaTitleFieldEnglish,
							MediaTitleFieldNative),
						MediaFieldFormat,
						MediaFieldType,
						MediaFieldStatusV2,
						MediaFieldCoverImage(
							MediaCoverImageFieldMedium)))),
			MediaFieldIsAdult,
			MediaFieldRankings(
				MediaRankFieldID,
				MediaRankFieldRank,
				MediaRankFieldType,
				MediaRankFieldFormat,
				MediaRankFieldYear,
				MediaRankFieldSeason,
				MediaRankFieldAllTime,
				MediaRankFieldContext,
			),
			MediaFieldSeason,
			MediaFieldSeasonYear,
			MediaFieldEpisodes,
			MediaFieldDuration,
			MediaFieldStudios(
				MediaParamStudios{},
				StudioConnectionFieldEdges(
					StudioEdgeFieldIsMain,
					StudioEdgeFieldNode(
						StudioFieldID,
						StudioFieldName,
						StudioFieldIsAnimationStudio))),
		}
	}
	return c.getMedia(id, MediaTypeAnime, fields...)
}

// GetAnimeCharacters to get list of characters in anime.
func (c *Client) GetAnimeCharacters(id int, page int, perPage int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldCharacters(
				MediaParamCharacters{
					Page:    page,
					PerPage: perPage,
					Sort:    []CharacterSort{CharacterSortRole, CharacterSortRelevance, CharacterSortID},
				},
				CharacterConnectionFieldPageInfo(
					PageInfoFieldTotal,
					PageInfoFieldPerPage,
					PageInfoFieldCurrentPage,
					PageInfoFieldLastPage,
					PageInfoFieldHasNextPage),
				CharacterConnectionFieldEdges(
					CharacterEdgeFieldRole,
					CharacterEdgeFieldNode(
						CharacterFieldID,
						CharacterFieldName(CharacterNameFieldFull),
						CharacterFieldImage(CharacterImageFieldMedium)),
					CharacterEdgeFieldVoiceActors(
						CharacterEdgeParamVoiceActors{},
						StaffFieldID,
						StaffFieldName(StaffNameFieldFull),
						StaffFieldImage(StaffImageFieldMedium),
						StaffFieldLanguage))),
		}
	}
	return c.getMedia(id, MediaTypeAnime, fields...)
}

// GetAnimeStaff to get list of staff in anime.
func (c *Client) GetAnimeStaff(id int, page int, perPage int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldStaff(
				MediaParamStaff{
					Page:    page,
					PerPage: perPage,
					Sort:    []StaffSort{StaffSortRelevance, StaffSortID},
				},
				StaffConnectionFieldEdges(
					StaffEdgeFieldRole,
					StaffEdgeFieldNode(
						StaffFieldID,
						StaffFieldName(StaffNameFieldFull),
						StaffFieldImage(StaffImageFieldMedium)))),
		}
	}
	return c.getMedia(id, MediaTypeAnime, fields...)
}

// GetAnimeStats to get anime stats.
func (c *Client) GetAnimeStats(id int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldStats(
				MediaStatsFieldScoreDistribution,
				MediaStatsFieldStatusDistribution),
		}
	}
	return c.getMedia(id, MediaTypeAnime, fields...)
}

// GetManga to get manga data.
func (c *Client) GetManga(id int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldID,
			MediaFieldTitle(
				MediaTitleFieldRomaji,
				MediaTitleFieldEnglish,
				MediaTitleFieldNative),
			MediaFieldType,
			MediaFieldFormat,
			MediaFieldStatusV2,
			MediaFieldDescription,
			MediaFieldDescription,
			MediaFieldStartDate,
			MediaFieldEndDate,
			MediaFieldCountryOfOrigin,
			MediaFieldSourceV2,
			MediaFieldCoverImage(
				MediaCoverImageFieldExtraLarge,
				MediaCoverImageFieldLarge,
				MediaCoverImageFieldMedium,
				MediaCoverImageFieldColor),
			MediaFieldBannerImage,
			MediaFieldGenres,
			MediaFieldSynonyms,
			MediaFieldAverageScore,
			MediaFieldMeanScore,
			MediaFieldPopularity,
			MediaFieldFavourites,
			MediaFieldTags(
				MediaTagFieldName,
				MediaTagFieldDescription,
				MediaTagFieldCategory,
				MediaTagFieldRank,
				MediaTagFieldIsGeneralSpoiler,
				MediaTagFieldIsMediaSpoiler,
				MediaTagFieldIsAdult),
			MediaFieldRelations(
				MediaConnectionFieldPageInfo(
					PageInfoFieldTotal,
					PageInfoFieldPerPage,
					PageInfoFieldCurrentPage,
					PageInfoFieldLastPage,
					PageInfoFieldHasNextPage),
				MediaConnectionFieldEdges(
					MediaEdgeFieldRelationTypeV2,
					MediaEdgeFieldNode(
						MediaFieldID,
						MediaFieldTitle(
							MediaTitleFieldRomaji,
							MediaTitleFieldEnglish,
							MediaTitleFieldNative),
						MediaFieldFormat,
						MediaFieldType,
						MediaFieldStatusV2,
						MediaFieldCoverImage(
							MediaCoverImageFieldMedium)))),
			MediaFieldIsAdult,
			MediaFieldRankings(
				MediaRankFieldID,
				MediaRankFieldRank,
				MediaRankFieldType,
				MediaRankFieldFormat,
				MediaRankFieldYear,
				MediaRankFieldSeason,
				MediaRankFieldAllTime,
				MediaRankFieldContext,
			),
			MediaFieldChapters,
			MediaFieldVolumes,
		}
	}
	return c.getMedia(id, MediaTypeManga, fields...)
}

// GetMangaCharacters to get list of characters in manga.
func (c *Client) GetMangaCharacters(id int, page int, perPage int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldCharacters(
				MediaParamCharacters{
					Page:    page,
					PerPage: perPage,
					Sort:    []CharacterSort{CharacterSortRole, CharacterSortRelevance, CharacterSortID},
				},
				CharacterConnectionFieldPageInfo(
					PageInfoFieldTotal,
					PageInfoFieldPerPage,
					PageInfoFieldCurrentPage,
					PageInfoFieldLastPage,
					PageInfoFieldHasNextPage),
				CharacterConnectionFieldEdges(
					CharacterEdgeFieldRole,
					CharacterEdgeFieldNode(
						CharacterFieldID,
						CharacterFieldName(CharacterNameFieldFull),
						CharacterFieldImage(CharacterImageFieldMedium)))),
		}
	}
	return c.getMedia(id, MediaTypeManga, fields...)
}

// GetMangaStaff to get list of staff in manga.
func (c *Client) GetMangaStaff(id int, page int, perPage int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldStaff(
				MediaParamStaff{
					Page:    page,
					PerPage: perPage,
					Sort:    []StaffSort{StaffSortRelevance, StaffSortID},
				},
				StaffConnectionFieldEdges(
					StaffEdgeFieldRole,
					StaffEdgeFieldNode(
						StaffFieldID,
						StaffFieldName(StaffNameFieldFull),
						StaffFieldImage(StaffImageFieldMedium)))),
		}
	}
	return c.getMedia(id, MediaTypeManga, fields...)
}

// GetMangaStats to get manga stats.
func (c *Client) GetMangaStats(id int, fields ...MediaField) (*Media, error) {
	if len(fields) == 0 {
		fields = []MediaField{
			MediaFieldStats(
				MediaStatsFieldScoreDistribution,
				MediaStatsFieldStatusDistribution),
		}
	}
	return c.getMedia(id, MediaTypeManga, fields...)
}
