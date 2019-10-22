package parser

import (
	"learngo/crawler/engine"
	"learngo/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<div class="des f-cl" data-v-3c42fade>.+\|\s*(\d+)岁\s*\|.+\|.+\|.+\|.+</div>`)
var heightRe = regexp.MustCompile(
	`<div class="des f-cl" data-v-3c42fade>.+\|.+\|.+\|.+\|.*(\d+)cm.*\|.+</div>`)
var incomeRe = regexp.MustCompile(
	`<div class="des f-cl" data-v-3c42fade>.+\|.+\|.+\|.+\|.+\|\s*([^<]+)</div>`)
var weightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>(\d+)kg</div>`)
var genderRe = regexp.MustCompile(
	`"genderString":"([.+]+)"`)
var xinzuoRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798>([^>]{2})座.*</div>`)
var marriageRe = regexp.MustCompile(
	`<div class="des f-cl" data-v-3c42fade>.+\|.+\|.+\|\s+([^\s]+)\s+\|.+\|.+</div>`)
var educationRe = regexp.MustCompile(
	`<div class="des f-cl" data-v-3c42fade>.+\|.+\|\s+([^\s]+)\s+\|.+\|.+\|.+</div>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hokouRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798>籍贯:([^>]+)</div>`)
var houseRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798>([^房]{2,4}房)</div>`)
var carRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798>([^车]{2,4}车)</div>`)
var guessRe = regexp.MustCompile(
	`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
var idUrlRe = regexp.MustCompile(
	`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(
		contents, incomeRe)
	profile.Gender = extractString(
		contents, genderRe)
	profile.Car = extractString(
		contents, carRe)
	profile.Education = extractString(
		contents, educationRe)
	profile.Hokou = extractString(
		contents, hokouRe)
	profile.House = extractString(
		contents, houseRe)
	profile.Marriage = extractString(
		contents, marriageRe)
	profile.Occupation = extractString(
		contents, occupationRe)
	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}

	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
