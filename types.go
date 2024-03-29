// Copyright (c) 2019 Teodor Spæren
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

type RedditSubmissionRaw struct {
	AllAwardings               []interface{} `json:"all_awardings"`
	AllowLiveComments          bool          `json:"allow_live_comments"`
	ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
	ApprovedBy                 interface{}   `json:"approved_by"`
	Archived                   bool          `json:"archived"`
	Author                     string        `json:"author"`
	AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
	AuthorFlairCSSClass        interface{}   `json:"author_flair_css_class"`
	AuthorFlairRichtext        []interface{} `json:"author_flair_richtext"`
	AuthorFlairTemplateID      interface{}   `json:"author_flair_template_id"`
	AuthorFlairText            interface{}   `json:"author_flair_text"`
	AuthorFlairTextColor       interface{}   `json:"author_flair_text_color"`
	AuthorFlairType            string        `json:"author_flair_type"`
	AuthorFullname             string        `json:"author_fullname"`
	AuthorPatreonFlair         bool          `json:"author_patreon_flair"`
	BannedAtUtc                interface{}   `json:"banned_at_utc"`
	BannedBy                   interface{}   `json:"banned_by"`
	CanGild                    bool          `json:"can_gild"`
	CanModPost                 bool          `json:"can_mod_post"`
	Category                   interface{}   `json:"category"`
	Clicked                    bool          `json:"clicked"`
	ContentCategories          interface{}   `json:"content_categories"`
	ContestMode                bool          `json:"contest_mode"`
	Created                    int           `json:"created"`
	CreatedUtc                 int           `json:"created_utc"`
	DiscussionType             interface{}   `json:"discussion_type"`
	Distinguished              interface{}   `json:"distinguished"`
	Domain                     string        `json:"domain"`
	Downs                      int           `json:"downs"`
	Edited                     bool          `json:"edited"`
	Gilded                     int           `json:"gilded"`
	Gildings                   struct{}      `json:"gildings"`
	Hidden                     bool          `json:"hidden"`
	HideScore                  bool          `json:"hide_score"`
	ID                         string        `json:"id"`
	IsCrosspostable            bool          `json:"is_crosspostable"`
	IsMeta                     bool          `json:"is_meta"`
	IsOriginalContent          bool          `json:"is_original_content"`
	IsRedditMediaDomain        bool          `json:"is_reddit_media_domain"`
	IsRobotIndexable           bool          `json:"is_robot_indexable"`
	IsSelf                     bool          `json:"is_self"`
	IsVideo                    bool          `json:"is_video"`
	Likes                      interface{}   `json:"likes"`
	LinkFlairBackgroundColor   string        `json:"link_flair_background_color"`
	LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
	LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
	LinkFlairText              interface{}   `json:"link_flair_text"`
	LinkFlairTextColor         string        `json:"link_flair_text_color"`
	LinkFlairType              string        `json:"link_flair_type"`
	Locked                     bool          `json:"locked"`
	Media                      interface{}   `json:"media"`
	MediaEmbed                 struct{}      `json:"media_embed"`
	MediaOnly                  bool          `json:"media_only"`
	ModNote                    interface{}   `json:"mod_note"`
	ModReasonBy                interface{}   `json:"mod_reason_by"`
	ModReasonTitle             interface{}   `json:"mod_reason_title"`
	ModReports                 []interface{} `json:"mod_reports"`
	Name                       string        `json:"name"`
	NoFollow                   bool          `json:"no_follow"`
	NumComments                int           `json:"num_comments"`
	NumCrossposts              int           `json:"num_crossposts"`
	NumReports                 interface{}   `json:"num_reports"`
	Over18                     bool          `json:"over_18"`
	ParentWhitelistStatus      string        `json:"parent_whitelist_status"`
	Permalink                  string        `json:"permalink"`
	Pinned                     bool          `json:"pinned"`
	Pwls                       int           `json:"pwls"`
	Quarantine                 bool          `json:"quarantine"`
	RemovalReason              interface{}   `json:"removal_reason"`
	ReportReasons              interface{}   `json:"report_reasons"`
	RetrievedOn                int           `json:"retrieved_on"`
	Saved                      bool          `json:"saved"`
	Score                      int           `json:"score"`
	SecureMedia                interface{}   `json:"secure_media"`
	SecureMediaEmbed           struct{}      `json:"secure_media_embed"`
	Selftext                   string        `json:"selftext"`
	SelftextHTML               interface{}   `json:"selftext_html"`
	SendReplies                bool          `json:"send_replies"`
	Spoiler                    bool          `json:"spoiler"`
	Stickied                   bool          `json:"stickied"`
	Subreddit                  string        `json:"subreddit"`
	SubredditID                string        `json:"subreddit_id"`
	SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
	SubredditSubscribers       int           `json:"subreddit_subscribers"`
	SubredditType              string        `json:"subreddit_type"`
	SuggestedSort              interface{}   `json:"suggested_sort"`
	Thumbnail                  string        `json:"thumbnail"`
	ThumbnailHeight            interface{}   `json:"thumbnail_height"`
	ThumbnailWidth             interface{}   `json:"thumbnail_width"`
	Title                      string        `json:"title"`
	TotalAwardsReceived        int           `json:"total_awards_received"`
	Ups                        int           `json:"ups"`
	URL                        string        `json:"url"`
	UserReports                []interface{} `json:"user_reports"`
	ViewCount                  interface{}   `json:"view_count"`
	Visited                    bool          `json:"visited"`
	WhitelistStatus            string        `json:"whitelist_status"`
	Wls                        int           `json:"wls"`
}

type RedditSubmissionRefined struct {
	ApprovedAtUtc        interface{} `json:"approved_at_utc"`
	ApprovedBy           interface{} `json:"approved_by"`
	Author               string      `json:"author"`
	AuthorFullname       string      `json:"author_fullname"`
	CreatedUtc           int         `json:"created_utc"`
	Domain               string      `json:"domain"`
	Gilded               int         `json:"gilded"`
	ID                   string      `json:"id"`
	IsCrosspostable      bool        `json:"is_crosspostable"`
	IsMeta               bool        `json:"is_meta"`
	IsOriginalContent    bool        `json:"is_original_content"`
	IsRedditMediaDomain  bool        `json:"is_reddit_media_domain"`
	IsRobotIndexable     bool        `json:"is_robot_indexable"`
	IsSelf               bool        `json:"is_self"`
	IsVideo              bool        `json:"is_video"`
	Likes                interface{} `json:"likes"`
	Name                 string      `json:"name"`
	NumComments          int         `json:"num_comments"`
	NumCrossposts        int         `json:"num_crossposts"`
	Over18               bool        `json:"over_18"`
	Permalink            string      `json:"permalink"`
	Pwls                 int         `json:"pwls"`
	RetrievedOn          int         `json:"retrieved_on"`
	Score                int         `json:"score"`
	Selftext             string      `json:"selftext"`
	SelftextHTML         string      `json:"selftext_html"`
	Subreddit            string      `json:"subreddit"`
	SubredditID          string      `json:"subreddit_id"`
	SubredditSubscribers int         `json:"subreddit_subscribers"`
	Thumbnail            string      `json:"thumbnail"`
	ThumbnailHeight      int         `json:"thumbnail_height"`
	ThumbnailWidth       int         `json:"thumbnail_width"`
	Title                string      `json:"title"`
	URL                  string      `json:"url"`
	WhitelistStatus      string      `json:"whitelist_status"`
	Wls                  int         `json:"wls"`
}
