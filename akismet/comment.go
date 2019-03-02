package akismet

import (
	"errors"
	"net/url"
	"time"
)

type Comment struct {
	Blog                string
	UserIP              string
	UserAgent           string
	Referrer            string
	Permalink           string
	CommentType         string
	CommentAuthor       string
	CommentAuthorEmail  string
	CommentAuthorURL    string
	CommentContent      string
	CommentDate         time.Time
	CommentPostModified time.Time
	BlogLang            string
	BlogCharset         string
	UserRole            string
	Test                bool
}

func addOptionalString(v url.Values, key, value string) {
	if value != "" {
		v.Add(key, value)
	}
}

func addOptionalTime(v url.Values, key string, t time.Time) {
	if !t.IsZero() {
		v.Add(key, t.UTC().Format(time.RFC3339))
	}
}

func addOptionalBool(v url.Values, key string, value bool) {
	if value {
		v.Add(key, "true")
	}
}

func (c Comment) encode() (string, error) {
	if c.Blog == "" {
		return "", errors.New("Blog field is empty but required")
	}
	if c.UserIP == "" {
		return "", errors.New("UserIP field is empty but required")
	}
	if c.UserAgent == "" {
		return "", errors.New("UserAgent field is empty but required")
	}

	v := url.Values{}
	v.Add("blog", c.Blog)
	v.Add("user_ip", c.UserIP)
	v.Add("user_agent", c.UserAgent)
	addOptionalString(v, "referrer", c.Referrer)
	addOptionalString(v, "permalink", c.Permalink)
	addOptionalString(v, "comment_type", c.CommentType)
	addOptionalString(v, "comment_author", c.CommentAuthor)
	addOptionalString(v, "comment_author_email", c.CommentAuthorEmail)
	addOptionalString(v, "comment_author_url", c.CommentAuthorURL)
	addOptionalString(v, "comment_content", c.CommentContent)
	addOptionalTime(v, "comment_date_gmt", c.CommentDate)
	addOptionalTime(v, "comment_post_modified_gmt", c.CommentPostModified)
	addOptionalString(v, "blog_lang", c.BlogLang)
	addOptionalString(v, "blog_charset", c.BlogCharset)
	addOptionalString(v, "user_role", c.UserRole)
	addOptionalBool(v, "is_test", c.Test)

	return v.Encode(), nil
}
