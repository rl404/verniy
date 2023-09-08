package main

import (
	"encoding/json"
	"os"

	"github.com/rl404/verniy"
)

func main() {
	c := verniy.New()

	d, err := c.GetAnime(21)
	// d, err := c.GetAnimeCharacters(1,1,10)
	// d, err := c.GetAnimeStaff(1,1,10)
	// d, err := c.GetAnimeStats(1)

	// d, err := c.GetManga(105778)
	// d, err := c.GetMangaCharacters(105778, 1, 10)
	// d, err := c.GetMangaStaff(105778, 1, 10)
	// d, err := c.GetMangaStats(105778)

	// d, err := c.GetCharacter(40)
	// d, err := c.GetCharacterAnime(40, 1, 5)
	// d, err := c.GetCharacterManga(40, 1, 5)

	// d, err := c.GetStaff(96868)
	// d, err := c.GetStaffCharacters(106184, 1, 10)
	// d, err := c.GetStaffAnime(106184, 1, 10)
	// d, err := c.GetStaffManga(96868, 1, 10)

	// d, err := c.GetUser("rl404")
	// d, err := c.GetUserFavouriteAnime("Josh", 1, 10)
	// d, err := c.GetUserFavouriteManga("Josh", 1, 10)
	// d, err := c.GetUserFavouriteCharacters("ZaeWarudo", 1, 10)
	// d, err := c.GetUserFavouriteStaff("ZaeWarudo", 1, 10)
	// d, err := c.GetUserFavouriteStudios("ZaeWarudo", 1, 10)
	// d, err := c.GetUserAnimeList("rl404")
	// d, err := c.GetUserMangaList("rl404")

	// d, err := c.GetStudio(803, 1, 10)
	// d, err := c.GetStudios(1, 10)
	// d, err := c.GetGenres()
	// d, err := c.GetTags()

	// d, err := c.SearchAnime(verniy.PageParamMedia{
	// Season:     verniy.SeasonWinter,
	// SeasonYear: 2021,
	// 	TagIn: []string{"Action"},
	// }, 1, 10)
	// d, err := c.SearchManga(verniy.PageParamMedia{
	// 	Search: "chainsaw",
	// }, 1, 10)
	// d, err := c.SearchCharacter(verniy.PageParamCharacters{
	// 	Search: "shinobu",
	// }, 1, 10)
	// d, err := c.SearchStaff(verniy.PageParamStaff{
	// 	Search: "kana",
	// }, 1, 10)

	if err != nil {
		panic(err)
	}

	dd, _ := json.MarshalIndent(d, "", "  ")
	// fmt.Println(string(dd))

	f, _ := os.Create("../data.json")
	defer f.Close()
	_, _ = f.Write(dd)
}
