package verniy

// MediaType is main media type.
type MediaType string

const (
	TypeAnime MediaType = "ANIME"
	TypeManga MediaType = "MANGA"
)

// MediaFormat is anime & manga format.
type MediaFormat string

const (
	FormatTv      MediaFormat = "TV"
	FormatTvShort MediaFormat = "TV_SHORT"
	FormatMovie   MediaFormat = "MOVIE"
	FormatSpecial MediaFormat = "SPECIAL"
	FormatOVA     MediaFormat = "OVA"
	FormatONA     MediaFormat = "ONA"
	FormatMusic   MediaFormat = "MUSIC"
	FormatManga   MediaFormat = "MANGA"
	FormatNovel   MediaFormat = "NOVEL"
	FormatOneShot MediaFormat = "ONE_SHOT"
)

// MediaStatus is anime & manga status.
type MediaStatus string

const (
	StatusFinished       MediaStatus = "FINISHED"
	StatusReleasing      MediaStatus = "RELEASING"
	StatusNotYetReleased MediaStatus = "NOT_YET_RELEASED"
	StatusCancelled      MediaStatus = "CANCELLED"
	StatusHiatus         MediaStatus = "HIATUS"
)

// MediaSeason is anime season.
type MediaSeason string

const (
	SeasonWinter MediaSeason = "WINTER" // 12-2
	SeasonSpring MediaSeason = "SPRING" // 3-5
	SeasonSummer MediaSeason = "SUMMER" // 6-8
	SeasonFall   MediaSeason = "FALL"   // 9-11
)

// MediaSource is anime & manga source.
type MediaSource string

const (
	SourceOriginal    MediaSource = "ORIGINAL"
	SourceManga       MediaSource = "MANGA"
	SourceLightNovel  MediaSource = "LIGHT_NOVEL"
	SourceVisualNovel MediaSource = "VISUAL_NOVEL"
	SourceVideoGame   MediaSource = "VIDEO_GAME"
	SourceOther       MediaSource = "OTHER"
	SourceNovel       MediaSource = "NOVEL"
	SourceDoujinshi   MediaSource = "DOUJINSHI"
	SourceAnime       MediaSource = "ANIME"
)

// MediaRankType is type of ranking.
type MediaRankType string

const (
	RankRated   MediaRankType = "RATED"
	RankPopular MediaRankType = "POPULAR"
)

// MediaListStatus is user's anime & manga status.
type MediaListStatus string

const (
	ListCurrent   MediaListStatus = "CURRENT"
	ListPlanning  MediaListStatus = "PLANNING"
	ListCompleted MediaListStatus = "COMPLETED"
	ListDropped   MediaListStatus = "DROPPED"
	ListPaused    MediaListStatus = "PAUSED"
	ListRepeating MediaListStatus = "REPEATING"
)

// MediaRelation is anime & manga relation.
type MediaRelation string

const (
	RelationAdaptation  MediaRelation = "ADAPTATION"
	RelationPrequel     MediaRelation = "PREQUEL"
	RelationSequel      MediaRelation = "SEQUEL"
	RelationParent      MediaRelation = "PARENT"
	RelationSideStory   MediaRelation = "SIDE_STORY"
	RelationCharacter   MediaRelation = "CHARACTER"
	RelationSummary     MediaRelation = "SUMMARY"
	RelationAlternative MediaRelation = "ALTERNATIVE"
	RelationSpinOff     MediaRelation = "SPIN_OFF"
	RelationOther       MediaRelation = "OTHER"
	RelationSource      MediaRelation = "SOURCE"
	RelationCompilation MediaRelation = "COMPILATION"
	RelationContains    MediaRelation = "CONTAINS"
)

// CharacterRole is type of character role.
type CharacterRole string

const (
	RoleMain       CharacterRole = "MAIN"
	RoleSupporting CharacterRole = "SUPPORTING"
	RoleBackground CharacterRole = "BACKGROUND"
)

// UserTitleLanguage is default user anime & manga title language.
type UserTitleLanguage string

const (
	LanguageRomaji          UserTitleLanguage = "ROMAJI"
	LanguageEnglish         UserTitleLanguage = "ENGLISH"
	LanguageNative          UserTitleLanguage = "NATIVE"
	LanguageRomajiStylised  UserTitleLanguage = "ROMAJI_STYLISED"
	LanguageEnglishStylised UserTitleLanguage = "ENGLISH_STYLISED"
	LanguageNativeStylised  UserTitleLanguage = "NATIVE_STYLISED"
)

// NotificationType is user notification type.
type NotificationType string

const (
	NotificationActivityMessage         NotificationType = "ACTIVITY_MESSAGE"
	NotificationActivityReply           NotificationType = "ACTIVITY_REPLY"
	NotificationActivityFollowing       NotificationType = "FOLLOWING"
	NotificationActivityMention         NotificationType = "ACTIVITY_MENTION"
	NotificationThreadCommentMention    NotificationType = "THREAD_COMMENT_MENTION"
	NotificationThreadSubscribed        NotificationType = "THREAD_SUBSCRIBED"
	NotificationThreadCommentReply      NotificationType = "THREAD_COMMENT_REPLY"
	NotificationAiring                  NotificationType = "AIRING"
	NotificationActivityLike            NotificationType = "ACTIVITY_LIKE"
	NotificationActivityReplyLike       NotificationType = "ACTIVITY_REPLY_LIKE"
	NotificationThreadLike              NotificationType = "THREAD_LIKE"
	NotificationThreadCommentLike       NotificationType = "THREAD_COMMENT_LIKE"
	NotificationActivityReplySubscribed NotificationType = "ACTIVITY_REPLY_SUBSCRIBED"
	NotificationRelatedMediaAddition    NotificationType = "RELATED_MEDIA_ADDITION"
)

// UserStaffNameLanguage is default user staff naming language.
type UserStaffNameLanguage string

const (
	UserStaffLanguageRomajiWestern UserStaffNameLanguage = "ROMAJI_WESTERN"
	UserStaffLanguageRomaji        UserStaffNameLanguage = "ROMAJI"
	UserStaffLanguageNative        UserStaffNameLanguage = "NATIVE"
)

// ReviewRating is type of review rating.
type ReviewRating string

const (
	RatingNoVote   ReviewRating = "NO_VOTE"
	RatingUpVote   ReviewRating = "UP_VOTE"
	RatingDownVote ReviewRating = "DOWN_VOTE"
)

// RecommendationRating is type of recommendation rating.
type RecommendationRating string

const (
	RatingNoRating RecommendationRating = "NO_RATING"
	RatingRateUp   RecommendationRating = "RATE_UP"
	RatingRateDown RecommendationRating = "RATE_DOWN"
)

// ScoreFormat is scoring format.
type ScoreFormat string

const (
	ScorePoint100        ScoreFormat = "POINT_100"
	ScorePoint100Decimal ScoreFormat = "POINT_10_DECIMAL"
	ScorePoint10         ScoreFormat = "POINT_10"
	ScorePoint5          ScoreFormat = "POINT_5"
	ScorePoint3          ScoreFormat = "POINT_3"
)

// ModRole is mod role.
type ModRole string

const (
	RoleAdmin            ModRole = "ADMIN"
	RoleLeadDeveloper    ModRole = "LEAD_DEVELOPER"
	RoleDeveloper        ModRole = "DEVELOPER"
	RoleLeadCommunity    ModRole = "LEAD_COMMUNITY"
	RoleCommunity        ModRole = "COMMUNITY"
	RoleDiscordCommunity ModRole = "DISCORD_COMMUNITY"
	RoleLeadAnimeData    ModRole = "LEAD_ANIME_DATA"
	RoleAnimeData        ModRole = "ANIME_DATA"
	RoleLeadMangaData    ModRole = "LEAD_MANGA_DATA"
	RoleMangaData        ModRole = "MANGA_DATA"
	RoleLeadSocialMedia  ModRole = "LEAD_SOCIAL_MEDIA"
	RoleSocialMedia      ModRole = "SOCIAL_MEDIA"
	RoleRetired          ModRole = "RETIRED"
)

// CharacterSort is sorting option for character list.
type CharacterSort string

const (
	CharacterSortID             CharacterSort = "ID"
	CharacterSortIDDesc         CharacterSort = "ID_DESC"
	CharacterSortRole           CharacterSort = "ROLE"
	CharacterSortRoleDesc       CharacterSort = "ROLE_DESC"
	CharacterSortSearchMatch    CharacterSort = "SEARCH_MATCH"
	CharacterSortFavourites     CharacterSort = "FAVOURITES"
	CharacterSortFavouritesDesc CharacterSort = "FAVOURITES_DESC"
	CharacterSortRelevance      CharacterSort = "RELEVANCE"
)

// StaffSort is sorting option for staff list.
type StaffSort string

const (
	StaffSortID             StaffSort = "ID"
	StaffSortIDDesc         StaffSort = "ID_DESC"
	StaffSortRole           StaffSort = "ROLE"
	StaffSortRoleDesc       StaffSort = "ROLE_DESC"
	StaffSortLanguage       StaffSort = "LANGUAGE"
	StaffSortLanguageDesc   StaffSort = "LANGUAGE_DESC"
	StaffSortSearchMatch    StaffSort = "SEARCH_MATCH"
	StaffSortFavourites     StaffSort = "FAVOURITES"
	StaffSortFavouritesDesc StaffSort = "FAVOURITES_DESC"
	StaffSortRelevance      StaffSort = "RELEVANCE"
)

// MediaSort is sorting option for anime & manga list.
type MediaSort string

const (
	MediaSortID               MediaSort = "ID"
	MediaSortIDDesc           MediaSort = "ID_DESC"
	MediaSortTitleRomaji      MediaSort = "TITLE_ROMAJI"
	MediaSortTitleRomajiDesc  MediaSort = "TITLE_ROMAJI_DESC"
	MediaSortTitleEnglish     MediaSort = "TITLE_ENGLISH"
	MediaSortTitleEnglishDesc MediaSort = "TITLE_ENGLISH_DESC"
	MediaSortTitleNative      MediaSort = "TITLE_NATIVE"
	MediaSortTitleNativeDesc  MediaSort = "TITLE_NATIVE_DESC"
	MediaSortType             MediaSort = "TYPE"
	MediaSortTypeDesc         MediaSort = "TYPE_DESC"
	MediaSortFormat           MediaSort = "FORMAT"
	MediaSortFormatDesc       MediaSort = "FORMAT_DESC"
	MediaSortStartDate        MediaSort = "START_DATE"
	MediaSortStartDateDesc    MediaSort = "START_DATE_DESC"
	MediaSortEndDate          MediaSort = "END_DATE"
	MediaSortEndDateDesc      MediaSort = "END_DATE_DESC"
	MediaSortScore            MediaSort = "SCORE"
	MediaSortScoreDesc        MediaSort = "SCORE_DESC"
	MediaSortPopularity       MediaSort = "POPULARITY"
	MediaSortPopularityDesc   MediaSort = "POPULARITY_DESC"
	MediaSortTrending         MediaSort = "TRENDING"
	MediaSortTrendingDesc     MediaSort = "TRENDING_DESC"
	MediaSortEpisodes         MediaSort = "EPISODES"
	MediaSortEpisodesDesc     MediaSort = "EPISODES_DESC"
	MediaSortDuration         MediaSort = "DURATION"
	MediaSortDurationDesc     MediaSort = "DURATION_DESC"
	MediaSortStatus           MediaSort = "STATUS"
	MediaSortStatusDesc       MediaSort = "STATUS_DESC"
	MediaSortChapters         MediaSort = "CHAPTERS"
	MediaSortChaptersDesc     MediaSort = "CHAPTERS_DESC"
	MediaSortVolumes          MediaSort = "VOLUMES"
	MediaSortVolumesDesc      MediaSort = "VOLUMES_DESC"
	MediaSortUpdatedAt        MediaSort = "UPDATED_AT"
	MediaSortUpdatedAtDesc    MediaSort = "UPDATED_AT_DESC"
	MediaSortSearchMatch      MediaSort = "SEARCH_MATCH"
	MediaSortFavourites       MediaSort = "FAVOURITES"
	MediaSortFavouritesDesc   MediaSort = "FAVOURITES_DESC"
)

// StaffLanguage is staff language.
type StaffLanguage string

const (
	StaffLanguageJapanese   StaffLanguage = "JAPANESE"
	StaffLanguageEnglish    StaffLanguage = "ENGLISH"
	StaffLanguageKorean     StaffLanguage = "KOREAN"
	StaffLanguageItalian    StaffLanguage = "ITALIAN"
	StaffLanguageSpanish    StaffLanguage = "SPANISH"
	StaffLanguagePortuguese StaffLanguage = "PORTUGUESE"
	StaffLanguageFrench     StaffLanguage = "FRENCH"
	StaffLanguageGerman     StaffLanguage = "GERMAN"
	StaffLanguageHebrew     StaffLanguage = "HEBREW"
	StaffLanguageHungarian  StaffLanguage = "HUNGARIAN"
)

// StudioSort is sorting option for studio list.
type StudioSort string

const (
	StudioSortID             StudioSort = "ID"
	StudioSortIDDesc         StudioSort = "ID_DESC"
	StudioSortName           StudioSort = "NAME"
	StudioSortNameDesc       StudioSort = "NAME_DESC"
	StudioSortSearchMatch    StudioSort = "SEARCH_MATCH"
	StudioSortFavourites     StudioSort = "FAVOURITES"
	StudioSortFavouritesDesc StudioSort = "FAVOURITES_DESC"
)

// MediaTrendSort is sorting option for media trend list.
type MediaTrendSort string

const (
	MediaTrendSortID             MediaTrendSort = "ID"
	MediaTrendSortIDDesc         MediaTrendSort = "ID_DESC"
	MediaTrendSortMediaID        MediaTrendSort = "MEDIA_ID"
	MediaTrendSortMediaIDDesc    MediaTrendSort = "MEDIA_ID_DESC"
	MediaTrendSortDate           MediaTrendSort = "DATE"
	MediaTrendSortDateDesc       MediaTrendSort = "DATE_DESC"
	MediaTrendSortScore          MediaTrendSort = "SCORE"
	MediaTrendSortScoreDesc      MediaTrendSort = "SCORE_DESC"
	MediaTrendSortPopularity     MediaTrendSort = "POPULARITY"
	MediaTrendSortPopularityDesc MediaTrendSort = "POPULARITY_DESC"
	MediaTrendSortTrending       MediaTrendSort = "TRENDING"
	MediaTrendSortTrendingDesc   MediaTrendSort = "TRENDING_DESC"
	MediaTrendSortEpisode        MediaTrendSort = "EPISODE"
	MediaTrendSortEpisodeDesc    MediaTrendSort = "EPISODE_DESC"
)

// ReviewSort is sorting option for review list.
type ReviewSort string

const (
	ReviewSortID            ReviewSort = "ID"
	ReviewSortIDDesc        ReviewSort = "ID_DESC"
	ReviewSortScore         ReviewSort = "SCORE"
	ReviewSortScoreDesc     ReviewSort = "SCORE_DESC"
	ReviewSortRating        ReviewSort = "RATING"
	ReviewSortRatingDesc    ReviewSort = "RATING_DESC"
	ReviewSortCreatedAt     ReviewSort = "CREATED_AT"
	ReviewSortCreatedAtDesc ReviewSort = "CREATED_AT_DESC"
	ReviewSortUpdatedAt     ReviewSort = "UPDATED_AT"
	ReviewSortUpdatedAtDesc ReviewSort = "UPDATED_AT_DESC"
)

// RecommendationSort is sorting option for recommendation list.
type RecommendationSort string

const (
	RecommendationSortID         RecommendationSort = "ID"
	RecommendationSortIDDesc     RecommendationSort = "ID_DESC"
	RecommendationSortRating     RecommendationSort = "RATING"
	RecommendationSortRatingDesc RecommendationSort = "RATING_DESC"
)

// UserStatisticsSort is sorting option for user statistics list.
type UserStatisticsSort string

const (
	UserStatisticsSortID            UserStatisticsSort = "ID"
	UserStatisticsSortIDDesc        UserStatisticsSort = "ID_DESC"
	UserStatisticsSortCount         UserStatisticsSort = "COUNT"
	UserStatisticsSortCountDesc     UserStatisticsSort = "COUNT_DESC"
	UserStatisticsSortProgress      UserStatisticsSort = "PROGRESS"
	UserStatisticsSortProgessDesc   UserStatisticsSort = "PROGRESS_DESC"
	UserStatisticsSortMeanScore     UserStatisticsSort = "MEAN_SCORE"
	UserStatisticsSortMeanScoreDesc UserStatisticsSort = "MEAN_SCORE_DESC"
)
