package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Comic struct {
	ID        	uint       `gorm:"id" json:"id"`
	Slug 		string     `json:"slug"`
	Title		string     `json:"title"`
	Sublog		string     `json:"sublog"`
	Tooltip		string     `json:"tooltip"`
	Image		string     `json:"image"`
	Thumbnail	string     `json:"thumbnail"`
	ComicWidth	int        `json:"comic_width"`
	ComicHeight	int        `json:"comic_height"`
	PostOn 		time.Time  `json:"post_on"`
	PostedOn 	time.Time  `json:"posted_on"`
	Enabled		bool       `json:"enabled"`
	RedditUrl	string     `json:"reddit_url"`
	RedditUpvotes	int    `json:"reddit_upvotes"`
	TumblrUrl	string     `json:"tumblr_url"`
	TumblrNotes	int        `json:"tumblr_notes"`
	NineGagUrl	string     `gorm:"nine_gag_url" json:"nine_gag_url"`
	NineGagUpvotes	int    `gorm:"nine_gag_upvotes" json:"nine_gag_upvotes"`
	FacebookLikes	int    `json:"facebook_likes"`
	CheezUrl	string     `gorm:"cheez_url" json:"cheez_url"`
	CheezUpvotes	int    `gorm:"cheez_upvotes" json:"cheez_upvotes"`
}

type PopularComics struct {
    Reddit      []Comic     `json:"reddit"`
    NineGag    []Comic     `json:"nine_gag"`
    Cheezburger []Comic     `json:"cheezburger"`
}

type AwesomeStats struct {
    Stat string             `json:"stat"`
}

type Notification struct {
    Message string             `json:"message"`
    visibility string       `json:"visibility"`
    Type string             `json:"type"`
}

type ComicActivity struct {
    Activity string         `json:"activity"`
    Comic Comic             `gorm:"embedded"`
}

// top 10 reddit, cheezburger & 9gag
func (c *Comic) GetPopularSocialMedia (db *gorm.DB) PopularComics {

    comics := PopularComics{
        Reddit: []Comic{},
        NineGag: []Comic{},
        Cheezburger: []Comic{},
    }

	// db.Find(&comics)
    db.Raw("SELECT * FROM `comics` where `enabled` = 1 order by `reddit_upvotes` desc limit 10").Scan(&comics.Reddit)
    db.Raw("SELECT * FROM `comics` where `enabled` = 1 order by `nine_gag_upvotes` desc limit 10").Scan(&comics.NineGag)
    db.Raw("SELECT * FROM `comics` where `enabled` = 1 order by `cheez_upvotes` desc limit 10").Scan(&comics.Cheezburger)

	return comics
}

// x days old, x upvotes, etc
func (c *Comic) GetAwesomeStats (db *gorm.DB) []AwesomeStats {

    awesomeStats := []AwesomeStats{}

    db.Raw(`
            select arr as "stat" from (
                    SELECT concat(count(*), " comics online") as "arr" FROM comics where enabled = 1
                            union
                    SELECT concat(format(sum(comic_width * comic_height), 0), " pixels&#178;") as "arr" FROM comics where enabled = 1
                            union
                    SELECT concat(sum(ga_sessions), " GA sessions") as "arr" FROM comics
                            union
                    SELECT concat(sum(reddit_upvotes) + sum(nine_gag_upvotes) + sum(cheez_upvotes) + sum(tumblr_notes) + sum(facebook_likes), " upvotes") as "arr" FROM comics
                            union
                    select concat(round((count(*)-1) / (timestampdiff(day, '2015-05-13', now()) / 7), 1), " avg comics per week") as "arr" FROM comics where enabled = 1
                            union
                    select concat(timestampdiff(day, '2015-05-13', now()), " days old") as "arr"
                            union
                    select concat(timestampdiff(month, '2015-05-13', now()), " months old") as "arr"
                            union
                    select concat(count(*), " updates from social media") as "arr" from history_of_things
            ) woot;
    `).Scan(&awesomeStats)

    return awesomeStats
}


func (c *Comic) GetUpcomingComic (db *gorm.DB) Comic {

    comic := Comic{}

    db.Raw(`
            SELECT
                *
            FROM
                comics
            where
                enabled = 0
            order by
                id asc
            limit 1
    `).Scan(&comic)

    return comic
}


func (c *Comic) GetNotifications (db *gorm.DB) []Notification {

    notifications := []Notification{}

    db.Raw(`
            select * from (
                            (
                    SELECT
                            concat("Current comic has the highest rating in ", datediff((select posted_on from consolia.comics where enabled = 1 order by id desc limit 1), posted_on), " days") as "message"
                            ,if (datediff((select posted_on from consolia.comics where enabled = 1 order by id desc limit 1), posted_on) > 13, 1, 0) as "visibility"
                            ,"info" as "type"
                    FROM
                            consolia.comics
                    where
                            enabled = 1
                    and
                            id < (select id from consolia.comics where enabled = 1 order by id desc limit 1)
                    and
                            (reddit_upvotes + cheez_upvotes + nine_gag_upvotes) >= (select (reddit_upvotes + cheez_upvotes + nine_gag_upvotes) from consolia.comics where enabled = 1 order by id desc limit 1)
                    order by
                            posted_on desc
                    limit 1
            )
            union all
            (
                    SELECT
                             concat(count(*), " comic(s) have no dimensions specified") as "message"
                            ,if (count(*) > 0, 1, 0) as "visibility"
                            ,"alert" as "type"
                    FROM
                            consolia.comics
                    where
                            enabled = 0
                    and (
                            comic_width is null or
                            comic_height is null
                    )
                    limit 1
            )
            union all
            (
                    SELECT
                             concat(count(*), " comic(s) don't have a tumblr title set up") as "message"
                            ,if (count(*) > 0, 1, 0) as "visibility"
                            ,"warning" as "type"
                    FROM
                            consolia.comics
                    where
                            enabled = 1
                    and
                            id > 49
                    and (
                                tumblr_title is null
                            and tumblr_notes = 0
                            and tumblr_url is null
                    )
                    limit 1
            )
            ) asd
            where visibility = 1
    `).Scan(&notifications)

    return notifications
}

func (c *Comic) GetPublishActivity (db *gorm.DB) []ComicActivity {

    activities := []ComicActivity{}

    db.Raw(`
        select * from (
        (
            SELECT
                     *
                    ,"will_publish_tomorrow" as "activity"
            FROM
                    comics
            where
                    enabled = 0
            and
                    date(post_on) = date_add(current_date(), interval 1 day)
            order by
                    post_on asc
        )
        union
        (
            SELECT
                     *
                    ,"published_today" as "activity"
            FROM
                    comics
            where
                    enabled = 1
            and
                    date(posted_on) = current_date()
            order by
                    posted_on asc
        )) asd
    `).Scan(&activities)

    return activities
}
