package models

import (
    "github.com/jinzhu/gorm"
)

type ComicHistory struct {
    ID          uint `gorm:"id"`
}

type Trending struct {
    ID                                  int `gorm:"id" json:"id"`
    Title                               string `gorm:"title" json:"title"`
    Slug                                string `gorm:"slug" json:"slug"`
    Image                               string `gorm:"image" json:"image"`
    Tooltip                             string `gorm:"tooltip" json:"tooltip"`
    RedditUpvotes                       int `gorm:"reddit_upvotes" json:"reddit_upvotes"`
    YesterdayRedditUpvotes              int `gorm:"yesterday_reddit_upvotes" json:"yesterday_reddit_upvotes"`
    RedditChange                        int `gorm:"reddit_change" json:"reddit_change"`
    NineGagUpvotes                      int `gorm:"9gag_upvotes" json:"9gag_upvotes"`
    Yesterday9gagUpvotes                int `gorm:"yesterday_9gag_upvotes" json:"yesterday_9gag_upvotes"`
    NineGagChange                       int `gorm:"9gag_change" json:"9gag_change"`
    CheezUpvotes                        int `gorm:"cheez_upvotes" json:"cheez_upvotes"`
    YesterdayCheezUpvotes               int `gorm:"yesterday_cheez_upvotes" json:"yesterday_cheez_upvotes"`
    CheezChange                         int `gorm:"cheez_change" json:"cheez_change"`
    FacebookLikes                       int `gorm:"facebook_likes" json:"facebook_likes"`
    YesterdayFacebookLikes              int `gorm:"yesterday_facebook_likes" json:"yesterday_facebook_likes"`
    FacebookChange                      int `gorm:"facebook_change" json:"facebook_change"`
    TumblrNotes                         int `gorm:"tumblr_notes" json:"tumblr_notes"`
    YesterdayTumblrNotes                int `gorm:"yesterday_tumblr_notes" json:"yesterday_tumblr_notes"`
    TumblrChange                        int `gorm:"tumblr_change" json:"tumblr_change"`
    TwitterShits                        int `gorm:"twitter_shits" json:"twitter_shits"`
    YesterdayTwitterShits               int `gorm:"yesterday_twitter_shits" json:"yesterday_twitter_shits"`
    TwitterChange                       int `gorm:"twitter_change" json:"twitter_change"`
    RedditMetricsUpdatedOn              int `gorm:"reddit_metrics_updated_on" json:"reddit_metrics_updated_on"`
    YesterdayRedditMetricsUpdatedOn     int `gorm:"yesterday_reddit_metrics_updated_on" json:"yesterday_reddit_metrics_updated_on"`
    NineGagMetricsUpdatedOn             int `gorm:"9gag_metrics_updated_on" json:"9gag_metrics_updated_on"`
    YesterdayNineGagMetricsUpdatedOn    int `gorm:"yesterday_9gag_metrics_updated_on" json:"yesterday_9gag_metrics_updated_on"`
    CheezMetricsUpdatedOn               int `gorm:"cheez_metrics_updated_on" json:"cheez_metrics_updated_on"`
    YesterdayCheezMetricsUpdatedOn      int `gorm:"yesterday_cheez_metrics_updated_on" json:"yesterday_cheez_metrics_updated_on"`
    FacebookMetricsUpdatedOn            int `gorm:"facebook_metrics_updated_on" json:"facebook_metrics_updated_on"`
    YesterdayFacebookMetricsUpdatedOn   int `gorm:"yesterday_facebook_metrics_updated_on" json:"yesterday_facebook_metrics_updated_on"`
    TumblrMetricsUpdatedOn              int `gorm:"tumblr_metrics_updated_on" json:"tumblr_metrics_updated_on"`
    YesterdayTumblrMetricsUpdatedOn     int `gorm:"yesterday_tumblr_metrics_updated_on" json:"yesterday_tumblr_metrics_updated_on"`
    TwitterMetricsUpdatedOn             int `gorm:"twitter_metrics_updated_on" json:"twitter_metrics_updated_on"`
    YesterdayTwitterMetricsUpdatedOn    int `gorm:"yesterday_twitter_metrics_updated_on" json:"yesterday_twitter_metrics_updated_on"`
    Reddit                              int `gorm:"reddit" json:"reddit"`
    Cheezburger                         int `gorm:"cheezburger" json:"cheezburger"`
    NineGag                             int `gorm:"9gag" json:"9gag"`
    Facebook                            int `gorm:"facebook" json:"facebook"`
    Tumblr                              int `gorm:"tumblr" json:"tumblr"`
    Twitter                             int `gorm:"twitte" json:"twitter"`
}


func (c *Comic) GetComicHistory (db gorm.DB) []ComicHistory {

    comicHistories := []ComicHistory{}

    db.Raw(`

    `).Scan(&comicHistories)

    return comicHistories
}

func (c *Comic) GetTrending (db gorm.DB) []Trending {

    activities := []Trending{}

    db.Raw(`
            SELECT
            C.id,
            C.title,
            C.slug,
            C.image,
            C.tooltip,
            C.reddit_upvotes,
            H.value as 'yesterday_reddit_upvotes',
            round((C.reddit_upvotes - H.value) / C.reddit_upvotes * 100, 0) as 'reddit_change',
            C.9gag_upvotes,
            HG.value as 'yesterday_9gag_upvotes',
            round((C.9gag_upvotes - HG.value) / C.9gag_upvotes * 100, 0) as '9gag_change',
            C.cheez_upvotes,
            HC.value as 'yesterday_cheez_upvotes',
            round((C.cheez_upvotes - HC.value) / C.cheez_upvotes * 100, 0) as 'cheez_change',
            C.facebook_likes,
            HF.value as 'yesterday_facebook_likes',
            round((C.facebook_likes - HF.value) / C.facebook_likes * 100, 0) as 'facebook_change',
            C.tumblr_notes,
            HT.value as 'yesterday_tumblr_notes',
            round((C.tumblr_notes - HT.value) / C.tumblr_notes * 100, 0) as 'tumblr_change',
            C.twitter_shits,
            HTW.value as 'yesterday_twitter_shits',
            round((C.twitter_shits - HTW.value) / C.twitter_shits * 100, 0) as 'twitter_change',
            C.reddit_metrics_updated_on,
            C.reddit_metrics_updated_on,
            H.added as 'yesterday_reddit_metrics_updated_on',
            C.9gag_metrics_updated_on,
            HG.added as 'yesterday_9gag_metrics_updated_on',
            C.cheez_metrics_updated_on,
            HC.added as 'yesterday_cheez_metrics_updated_on',
            C.facebook_metrics_updated_on,
            HF.added as 'yesterday_facebook_metrics_updated_on',
            C.tumblr_metrics_updated_on,
            HT.added as 'yesterday_tumblr_metrics_updated_on',
            C.twitter_metrics_updated_on,
            HTW.added as 'yesterday_twitter_metrics_updated_on',
            C.reddit_metrics_updated_on > now() - interval 10 hour as reddit,
            C.cheez_metrics_updated_on > now() - interval 10 hour as cheezburger,
            C.9gag_metrics_updated_on > now() - interval 10 hour as 9gag,
            C.facebook_metrics_updated_on > now() - interval 10 hour as facebook,
            C.tumblr_metrics_updated_on > now() - interval 10 hour as tumblr,
            C.twitter_metrics_updated_on > now() - interval 10 hour as twitter
    FROM
            consolia.comics C
left join (select * from (select * from consolia.history_of_things H where H.collection = 'reddit_upvotes' and H.added < now() - interval 24 hour order by H.added desc) ZZ group by ZZ.thing_id) H on (H.thing_id = C.id)
    left join (select * from (select * from consolia.history_of_things H where H.collection = '9gag_upvotes' and H.added < now() - interval 24 hour order by H.added desc) ZZZ group by ZZZ.thing_id) HG on (HG.thing_id = C.id)
    left join (select * from (select * from consolia.history_of_things H where H.collection = 'cheez_upvotes' and H.added < now() - interval 24 hour order by H.added desc) ZZZZ group by ZZZZ.thing_id) HC on (HC.thing_id = C.id)
    left join (select * from (select * from consolia.history_of_things H where H.collection = 'facebook_likes' and H.added < now() - interval 24 hour order by H.added desc) ZZZZZ group by ZZZZZ.thing_id) HF on (HF.thing_id = C.id)
    left join (select * from (select * from consolia.history_of_things H where H.collection = 'tumblr_notes' and H.added < now() - interval 24 hour order by H.added desc) ZZZZZZ group by ZZZZZZ.thing_id) HT on (HT.thing_id = C.id)
    left join (select * from (select * from consolia.history_of_things H where H.collection = 'twitter_shits' and H.added < now() - interval 24 hour order by H.added desc) ZZZZZZZ group by ZZZZZZZ.thing_id) HTW on (HTW.thing_id = C.id)
    where
            C.reddit_metrics_updated_on > now() - interval 10 hour
    or
            C.9gag_metrics_updated_on > now() - interval 10 hour
    or
            C.cheez_metrics_updated_on > now() - interval 10 hour
    or
            C.facebook_metrics_updated_on > now() - interval 10 hour
    or
            C.tumblr_metrics_updated_on > now() - interval 10 hour
    or
            C.twitter_metrics_updated_on > now() - interval 10 hour
    `).Scan(&activities)

    return activities
}
