package site

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type BlockedWebsitesParser interface {
	Parse(string) ([]BlockedWebsite, error)
}

type RskParser struct{}

func NewRskParser() *RskParser {
	return &RskParser{}
}

func (r *RskParser) Parse(url string) (res []BlockedWebsite, err error) {

	cl := http.Client{Timeout: 5 * time.Second}
	resp, err := cl.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	res, err = r.parse(resp.Body)

	return
}

func (r *RskParser) parse(resp io.Reader) (res []BlockedWebsite, err error) {
	doc, err := goquery.NewDocumentFromReader(resp)

	if err != nil {
		return
	}

	doc.Find(".t431__table tbody tr.t431__evenrow").Each(func(i int, selection *goquery.Selection) {
		bws := BlockedWebsite{}

		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 1:
				d, err := time.Parse("02.01.2006", strings.TrimSpace(selection.Text()))
				if err != nil {
					return
				}
				bws.SetDate(d)
			case 2:
				bws.SetUrl(strings.TrimSpace(selection.Text()))
			case 3:
				bws.SetIp(strings.TrimSpace(selection.Text()))
			case 4:
				bws.SetBlockInitiator(strings.TrimSpace(selection.Text()))
			case 5:
				parseInt, err := strconv.ParseInt(strings.TrimSpace(selection.Text()), 0, 64)
				if err != nil {
					return
				}
				bws.SetBlocksCount(int(parseInt))
			}
		})

		res = append(res, bws)
	})

	return
}
