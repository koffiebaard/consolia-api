package models

import (
    _"fmt"
    "testing"
    "consolia-api/utils"

    "github.com/stretchr/testify/assert"
)


func TestGetPopularSocialMedia(t *testing.T) {

    comic := Comic{}
    conf, err := utils.GetConf()
    assert.NoError(t, err)

    comics := comic.GetPopularSocialMedia(conf.DB)

    assert.NotNil(t, comics.Reddit, "Reddit not empty")
    assert.NotNil(t, comics.NineGag, "9Gag not empty")
    assert.NotNil(t, comics.Cheezburger, "Cheezburger not empty")

    assert.Equal(t, len(comics.Reddit), 10, "Reddit has 10 comics")
    assert.Equal(t, len(comics.NineGag), 10, "9Gag has 10 comics")
    assert.Equal(t, len(comics.Cheezburger), 10, "Cheezburger has 10 comics")

    for _, redditComic := range comics.Reddit {
        isValidComic(t, redditComic)
    }
    for _, nineGagComic := range comics.NineGag {
        isValidComic(t, nineGagComic)
    }
    for _, cheezburgerComic := range comics.Cheezburger {
        isValidComic(t, cheezburgerComic)
    }
}

func isValidComic (t *testing.T, comic Comic) {
    assert.NotNil(t, comic, "No comic can be nil")
    assert.NotNil(t, comic.ID, "ID must be present")
    assert.NotNil(t, comic.Slug, "Slug must be present")
    assert.Equal(t, len(comic.Slug) > 2, true, "Slug must be larger than 2 characters")
    assert.NotNil(t, comic.Image, "Image must be present")
    assert.Equal(t, len(comic.Image) > 4, true, "Image must be larger than 4 characters")
    assert.NotNil(t, comic.Enabled, "Enabled cannot be nil")
    assert.NotNil(t, comic.ComicWidth, "Width cannot be nil")
    assert.NotNil(t, comic.ComicHeight, "Height cannot be nil")
}

func TestGetAwesomeStats(t *testing.T) {

    comic := Comic{}
    conf, err := utils.GetConf()
    assert.NoError(t, err)

    stats := comic.GetAwesomeStats(conf.DB)

    assert.Equal(t, len(stats), 8, "There are currently 8 stats implemented")

    for _, stat := range stats {
        assert.NotNil(t, stat.Stat, "No stat can be nil")
    }
}


func TestGetUpcomingComic(t *testing.T) {

    comic := Comic{}
    conf, err := utils.GetConf()
    assert.NoError(t, err)

    upcomingComic := comic.GetUpcomingComic(conf.DB)

    assert.NotNil(t, upcomingComic, "There should always be an upcoming comic")
    assert.NotNil(t, upcomingComic.ID, "ID should be present")
    assert.Equal(t, upcomingComic.Enabled, false, "Comic should not be enabled")
}


func TestGetNotifications(t *testing.T) {

    comic := Comic{}
    conf, err := utils.GetConf()
    assert.NoError(t, err)

    notifications := comic.GetNotifications(conf.DB)

    assert.NotNil(t, notifications, "There should always be at least 1 notification")

    for _, notification := range notifications {

        assert.NotNil(t, notification, "Notification should be there")
        assert.NotNil(t, notification.Message, "Notification message should be there")
        assert.Equal(t, IsValidNotificationType(notification.Type), true, "Notification type should be alert, warning or info")
    }
}

func IsValidNotificationType(notificationType string) bool {
    switch notificationType {
    case
        "warning",
        "alert",
        "info":
        return true
    }
    return false
}


func TestGetPublishActivity(t *testing.T) {

    comic := Comic{}
    conf, err := utils.GetConf()
    assert.NoError(t, err)

    activities := comic.GetPublishActivity(conf.DB)

    assert.NotNil(t, activities, "Activities cannot be nil")

    for _, activity := range activities {

        assert.NotNil(t, activity, "activity should be there")
        assert.NotNil(t, activity.Comic, "Comic should be attached to activity")
        assert.Equal(t, IsValidActivity(activity.Activity), true, "Activity should be will_publish_tomorrow or published_today")

        isValidComic(t, activity.Comic)
    }
}

func IsValidActivity(activity string) bool {
    switch activity {
    case
        "will_publish_tomorrow",
        "published_today":
        return true
    }
    return false
}
