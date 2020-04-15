package annotate

import (
	"github.com/Pantani/replacer/pkg/replacer"
	"testing"
)

var (
	names = replacer.Names{
		replacer.Name{
			Name:    "alex",
			Url:     "http://alex.com",
			Snippet: `<a href="http://alex.com">{{.Name}}</a>`,
		},
		replacer.Name{
			Name:    "bo",
			Url:     "http://bo.com",
			Snippet: `<a href="http://alex.com">{{.Name}}</a>`,
		},
		replacer.Name{
			Name:    "casey",
			Url:     "http://casey.com",
			Snippet: `<a href="http://alex.com">{{.Name}}</a>`,
		},
	}
)

func TestScrapeHtml(t *testing.T) {
	type args struct {
		text    string
		replace replacer.Names
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple text 1", args{`my name is alex`, names}, `my name is <a href="http://alex.com">alex</a>`},
		{"simple text 2", args{`my name is todd`, names}, `my name is todd`},
		{"simple text 3", args{`alex, bo, and casey went to the park.`, names}, `<a href="http://alex.com">alex</a>, <a href="http://bo.com">bo</a>, and <a href="http://casey.com">casey</a> went to the park.`},
		{"name", args{`alex alexander alexandria alexbocasey`, names}, `<a href="http://alex.com">alex</a> alexander alexandria alexbocasey`},
		{"html div", args{`<div data-alex="alex">alex</div>`, names}, `<div data-alex="alex"><a href="http://alex.com">alex</a></div>`},
		{"html links", args{`<a href="http://foo.com">alex is already linked</a> but alex is not`, names}, `<a href="http://foo.com">alex is already linked</a> but <a href="http://alex.com">alex</a> is not`},
		{"html paragraphs", args{`<div><p>this is paragraph 1 about alex.</p><p>alex's paragraph number 2.</p><p>and some closing remarks about alex</p></div>`, names}, `<div><p>this is paragraph 1 about <a href="http://alex.com">alex</a>.</p><p><a href="http://alex.com">alex</a>'s paragraph number 2.</p><p>and some closing remarks about <a href="http://alex.com">alex</a></p></div>`},
		{"html data", args{`<div data-alex="alex">alex</div>`, names}, `<div data-alex="alex"><a href="http://alex.com">alex</a></div>`},
		{"html", args{`<div><p>this is paragraph 1 about alex.</p><p>alex's paragraph number 2.</p><p>and some closing remarks about alex</p></div>`, names}, `<div><p>this is paragraph 1 about <a href="http://alex.com">alex</a>.</p><p><a href="http://alex.com">alex</a>'s paragraph number 2.</p><p>and some closing remarks about <a href="http://alex.com">alex</a></p></div>`},
		{"html table", args{`<div><ul><li>alex</li><li>bo</li><li>bob</li><li>casey</li></ul></div><div><p>this is paragraph 1 about alex.</p><p>alex's paragraph number 2.</p><p>and some closing remarks about alex</p></div>`, names}, `<div><ul><li><a href="http://alex.com">alex</a></li><li><a href="http://bo.com">bo</a></li><li>bob</li><li><a href="http://casey.com">casey</a></li></ul></div><div><p>this is paragraph 1 about <a href="http://alex.com">alex</a>.</p><p><a href="http://alex.com">alex</a>'s paragraph number 2.</p><p>and some closing remarks about <a href="http://alex.com">alex</a></p></div>`},
		{"div class 1", args{`<div class=\"row\"><div class=\"col-md-6\"><h2> Sourcegraph makes programming <strong>delightful.</strong></h2><p>We want to make you even better at what you do best: building software to solve real problems.</p><p>Sourcegraph makes it easier to find the information you need: documentation, examples, usage statistics, answers, and more.</p><p>We're just getting started, and we'd love to hear from you. <a href=\"/contact\" ui-sref=\"help.contact\">Get in touch with us.</a></p></div><div class=\"col-md-4 team\"><h3>Team</h3><ul><li><img src=\"https://secure.gravatar.com/avatar/c728a3085fc16da7c594903ea8e8858f?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Beyang Liu</strong><br><a target=\"_blank\" href=\"http://github.com/beyang\">github.com/beyang</a><a href=\"mailto:beyang@sourcegraph.com\">beyang@sourcegraph.com</a></div></li><li><img src=\"https://secure.gravatar.com/avatar/d491971c742b8249341e495cf53045ea?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Quinn Slack</strong><br><a target=\"_blank\" href=\"http://github.com/sqs\">github.com/sqs</a><a href=\"mailto:sqs@sourcegraph.com\">sqs@sourcegraph.com</a></div></li><li><img src=\"https://s3-us-west-2.amazonaws.com/public-dev/milton.png\" class=\"pull-left\"><div class=\"bio\"><strong>Milton</strong> the Australian Shepherd </div></li></ul><p><a href=\"/contact\" ui-sref=\"help.contact\">Want to join us?</a></p></div></div>`, names}, `<div class=\"row\"><div class=\"col-md-6\"><h2> Sourcegraph makes programming <strong>delightful.</strong></h2><p>We want to make you even better at what you do best: building software to solve real problems.</p><p>Sourcegraph makes it easier to find the information you need: documentation, examples, usage statistics, answers, and more.</p><p>We're just getting started, and we'd love to hear from you. <a href=\"/contact\" ui-sref=\"help.contact\">Get in touch with us.</a></p></div><div class=\"col-md-4 team\"><h3>Team</h3><ul><li><img src=\"https://secure.gravatar.com/avatar/c728a3085fc16da7c594903ea8e8858f?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Beyang Liu</strong><br><a target=\"_blank\" href=\"http://github.com/beyang\">github.com/beyang</a><a href=\"mailto:beyang@sourcegraph.com\">beyang@sourcegraph.com</a></div></li><li><img src=\"https://secure.gravatar.com/avatar/d491971c742b8249341e495cf53045ea?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Quinn Slack</strong><br><a target=\"_blank\" href=\"http://github.com/sqs\">github.com/sqs</a><a href=\"mailto:sqs@sourcegraph.com\">sqs@sourcegraph.com</a></div></li><li><img src=\"https://s3-us-west-2.amazonaws.com/public-dev/milton.png\" class=\"pull-left\"><div class=\"bio\"><strong>Milton</strong> the Australian Shepherd </div></li></ul><p><a href=\"/contact\" ui-sref=\"help.contact\">Want to join us?</a></p></div></div>`},
		{"div class 2", args{`<div class='<div class=\"name\">name</a>'>name</div>`, names}, `<div class='<div class=\"name\">name</a>'>name</div>`},
		{"empty", args{"", names}, ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScrapeHtml(tt.args.text, tt.args.replace); got != tt.want {
				t.Errorf("ScrapeHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceStrings(t *testing.T) {
	type args struct {
		text    string
		replace replacer.Names
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simple text 1", args{`my name is alex`, names}, `my name is <a href="http://alex.com">alex</a>`},
		{"simple text 2", args{`my name is todd`, names}, `my name is todd`},
		{"simple text 3", args{`alex, bo, and casey went to the park.`, names}, `<a href="http://alex.com">alex</a>, <a href="http://bo.com">bo</a>, and <a href="http://casey.com">casey</a> went to the park.`},
		{"name", args{`alex alexander alexandria alexbocasey`, names}, `<a href="http://alex.com">alex</a> alexander alexandria alexbocasey`},
		{"html div", args{`<div data-alex="alex">alex</div>`, names}, `<div data-<a href="http://alex.com">alex</a>="<a href="http://alex.com">alex</a>"><a href="http://alex.com">alex</a></div>`},
		{"html links", args{`<a href="http://foo.com">alex is already linked</a> but alex is not`, names}, `<a href="http://foo.com"><a href="http://alex.com">alex</a> is already linked</a> but <a href="http://alex.com">alex</a> is not`},
		{"html paragraphs", args{`<div><p>this is paragraph 1 about alex.</p><p>alex's paragraph number 2.</p><p>and some closing remarks about alex</p></div>`, names}, `<div><p>this is paragraph 1 about <a href="http://alex.com">alex</a>.</p><p><a href="http://alex.com">alex</a>'s paragraph number 2.</p><p>and some closing remarks about <a href="http://alex.com">alex</a></p></div>`},
		{"html data", args{`<div data-alex="alex">alex</div>`, names}, `<div data-<a href="http://alex.com">alex</a>="<a href="http://alex.com">alex</a>"><a href="http://alex.com">alex</a></div>`},
		{"html", args{`<div><p>this is paragraph 1 about alex.</p><p>alex's paragraph number 2.</p><p>and some closing remarks about alex</p></div>`, names}, `<div><p>this is paragraph 1 about <a href="http://alex.com">alex</a>.</p><p><a href="http://alex.com">alex</a>'s paragraph number 2.</p><p>and some closing remarks about <a href="http://alex.com">alex</a></p></div>`},
		{"html table", args{`<div><ul><li>alex</li><li>bo</li><li>bob</li><li>casey</li></ul></div><div><p>this is paragraph 1 about alex.</p><p>alex's paragraph number 2.</p><p>and some closing remarks about alex</p></div>`, names}, `<div><ul><li><a href="http://alex.com">alex</a></li><li><a href="http://bo.com">bo</a></li><li>bob</li><li><a href="http://casey.com">casey</a></li></ul></div><div><p>this is paragraph 1 about <a href="http://alex.com">alex</a>.</p><p><a href="http://alex.com">alex</a>'s paragraph number 2.</p><p>and some closing remarks about <a href="http://alex.com">alex</a></p></div>`},
		{"div class 1", args{`<div class=\"row\"><div class=\"col-md-6\"><h2> Sourcegraph makes programming <strong>delightful.</strong></h2><p>We want to make you even better at what you do best: building software to solve real problems.</p><p>Sourcegraph makes it easier to find the information you need: documentation, examples, usage statistics, answers, and more.</p><p>We're just getting started, and we'd love to hear from you. <a href=\"/contact\" ui-sref=\"help.contact\">Get in touch with us.</a></p></div><div class=\"col-md-4 team\"><h3>Team</h3><ul><li><img src=\"https://secure.gravatar.com/avatar/c728a3085fc16da7c594903ea8e8858f?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Beyang Liu</strong><br><a target=\"_blank\" href=\"http://github.com/beyang\">github.com/beyang</a><a href=\"mailto:beyang@sourcegraph.com\">beyang@sourcegraph.com</a></div></li><li><img src=\"https://secure.gravatar.com/avatar/d491971c742b8249341e495cf53045ea?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Quinn Slack</strong><br><a target=\"_blank\" href=\"http://github.com/sqs\">github.com/sqs</a><a href=\"mailto:sqs@sourcegraph.com\">sqs@sourcegraph.com</a></div></li><li><img src=\"https://s3-us-west-2.amazonaws.com/public-dev/milton.png\" class=\"pull-left\"><div class=\"bio\"><strong>Milton</strong> the Australian Shepherd </div></li></ul><p><a href=\"/contact\" ui-sref=\"help.contact\">Want to join us?</a></p></div></div>`, names}, `<div class=\"row\"><div class=\"col-md-6\"><h2> Sourcegraph makes programming <strong>delightful.</strong></h2><p>We want to make you even better at what you do best: building software to solve real problems.</p><p>Sourcegraph makes it easier to find the information you need: documentation, examples, usage statistics, answers, and more.</p><p>We're just getting started, and we'd love to hear from you. <a href=\"/contact\" ui-sref=\"help.contact\">Get in touch with us.</a></p></div><div class=\"col-md-4 team\"><h3>Team</h3><ul><li><img src=\"https://secure.gravatar.com/avatar/c728a3085fc16da7c594903ea8e8858f?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Beyang Liu</strong><br><a target=\"_blank\" href=\"http://github.com/beyang\">github.com/beyang</a><a href=\"mailto:beyang@sourcegraph.com\">beyang@sourcegraph.com</a></div></li><li><img src=\"https://secure.gravatar.com/avatar/d491971c742b8249341e495cf53045ea?s=64\" class=\"pull-left\"><div class=\"bio\"><strong>Quinn Slack</strong><br><a target=\"_blank\" href=\"http://github.com/sqs\">github.com/sqs</a><a href=\"mailto:sqs@sourcegraph.com\">sqs@sourcegraph.com</a></div></li><li><img src=\"https://s3-us-west-2.amazonaws.com/public-dev/milton.png\" class=\"pull-left\"><div class=\"bio\"><strong>Milton</strong> the Australian Shepherd </div></li></ul><p><a href=\"/contact\" ui-sref=\"help.contact\">Want to join us?</a></p></div></div>`},
		{"div class 2", args{`<div class='<div class=\"name\">name</a>'>name</div>`, names}, `<div class='<div class=\"name\">name</a>'>name</div>`},
		{"empty", args{"", names}, ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceStrings(tt.args.text, tt.args.replace); got != tt.want {
				t.Errorf("replaceStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
