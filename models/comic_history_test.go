package models

import (
    _"fmt"
    "testing"
    "consolia-api/utils"

    "github.com/stretchr/testify/assert"
    "github.com/jinzhu/gorm"
)

func TestGetTrending(t *testing.T) {

    comic := Comic{}
    conf, err := utils.GetConf()
    assert.NoError(t, err)

    trendingComics := comic.GetTrending(conf.DB)

    shouldBeTrendingComics := getComicWithSocialMediaUpdateLastTenHours(conf.DB)

    if len(shouldBeTrendingComics) > 0 {

        assert.Equal(t, len(trendingComics), len(shouldBeTrendingComics), "Number of trending comics should be equal to estimated number")

        for _, trendingComic := range trendingComics {

            assert.NotNil(t, trendingComic, "comic should be there")

            // Count the percentages of change in social media platforms. It should be > 0 since len(shouldBeTrendingComics) > 0
            assert.NotEqual(t,  trendingComic.RedditChange +
                                trendingComic.NineGagChange +
                                trendingComic.CheezChange +
                                trendingComic.TumblrChange +
                                trendingComic.TwitterChange +
                                trendingComic.FacebookChange, 0, "There must be a social media change somewhere")

            // Count the amount of social media updates. It should be > 0 since len(shouldBeTrendingComics) > 0
            assert.NotEqual(t,  trendingComic.Reddit +
                                trendingComic.NineGag +
                                trendingComic.Cheezburger +
                                trendingComic.Tumblr +
                                trendingComic.Twitter +
                                trendingComic.Facebook, 0, "There must be at least 1 social media platform update")


            if trendingComic.Reddit > 0 {
                assert.NotEqual(t, trendingComic.RedditChange, 0, "If there is a reddit update, there should be a percentage change")
                assert.NotEqual(t, trendingComic.RedditUpvotes, 0, "If there is a reddit update, there should be reddit upvotes")
                assert.True(t, trendingComic.RedditUpvotes > trendingComic.YesterdayRedditUpvotes, 0, "If there is a reddit change, there should be fewer upvotes for yesterday")
            }
            if trendingComic.NineGag > 0 {
                assert.NotEqual(t, trendingComic.NineGagChange, 0, "If there is a 9gag update, there should be a percentage change")
                assert.NotEqual(t, trendingComic.NineGagUpvotes, 0, "If there is a 9gag change, there should be 9gag upvotes")
                assert.True(t, trendingComic.NineGagUpvotes > trendingComic.YesterdayNineGagUpvotes, 0, "If there is a 9gag change, there should be fewer upvotes for yesterday")
            }
            if trendingComic.Cheezburger > 0 {
                assert.NotEqual(t, trendingComic.CheezChange, 0, "If there is a cheezburger update, there should be a percentage change")
                assert.NotEqual(t, trendingComic.CheezUpvotes, 0, "If there is a cheezburger change, there should be cheezburger upvotes")
                assert.True(t, trendingComic.CheezUpvotes > trendingComic.YesterdayCheezUpvotes, 0, "If there is a cheezburger change, there should be fewer upvotes for yesterday")
            }
            if trendingComic.Tumblr > 0 {
                assert.NotEqual(t, trendingComic.TumblrChange, 0, "If there is a tumblr update, there should be a percentage change")
                assert.NotEqual(t, trendingComic.TumblrNotes, 0, "If there is a tumblr change, there should be reddit notes")
                assert.True(t, trendingComic.TumblrNotes > trendingComic.YesterdayTumblrNotes, 0, "If there is a tumblr change, there should be fewer notes for yesterday")
            }
            if trendingComic.Twitter > 0 {
                assert.NotEqual(t, trendingComic.TwitterChange, 0, "If there is a twitter update, there should be a percentage change")
                assert.NotEqual(t, trendingComic.TwitterShits, 0, "If there is a twitter change, there should be twitter shits")
                assert.True(t, trendingComic.TwitterShits > trendingComic.YesterdayTwitterShits, 0, "If there is a twitter change, there should be fewer shits for yesterday")
            }
            if trendingComic.Facebook > 0 {
                assert.NotEqual(t, trendingComic.FacebookChange, 0, "If there is a facebook update, there should be a percentage change")
                assert.NotEqual(t, trendingComic.FacebookLikes, 0, "If there is a facebook change, there should be facebook likes")
                assert.True(t, trendingComic.FacebookLikes > trendingComic.YesterdayFacebookLikes, 0, "If there is a facebook change, there should be fewer likes for yesterday")
            }

        }
    }
}

func getComicWithSocialMediaUpdateLastTenHours (db *gorm.DB) []Comic {

    comics := []Comic{}

    db.Raw(`
            SELECT
                *
            FROM
                comics C
            where
                C.reddit_metrics_updated_on > now() - interval 10 hour
            or
                C.nine_gag_metrics_updated_on > now() - interval 10 hour
            or
                C.cheez_metrics_updated_on > now() - interval 10 hour
            or
                C.facebook_metrics_updated_on > now() - interval 10 hour
            or
                C.tumblr_metrics_updated_on > now() - interval 10 hour
            or
                C.twitter_metrics_updated_on > now() - interval 10 hour
    `).Scan(&comics)

    return comics
}
