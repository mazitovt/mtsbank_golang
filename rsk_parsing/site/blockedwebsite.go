package site

import (
	"fmt"
	"time"
)

type BlockedWebsite struct {
	date           time.Time
	url            string
	ip             string
	blockInitiator string
	blocksCount    int
}

func (b *BlockedWebsite) Date() time.Time {
	return b.date
}

func NewBlockedWebsite(date time.Time, url string, ip string, blockInitiator string, blocksCount int) *BlockedWebsite {
	return &BlockedWebsite{date: date, url: url, ip: ip, blockInitiator: blockInitiator, blocksCount: blocksCount}
}

func (b *BlockedWebsite) SetDate(date time.Time) *BlockedWebsite {
	b.date = date
	return b
}

func (b *BlockedWebsite) SetUrl(url string) *BlockedWebsite {
	b.url = url
	return b
}

func (b *BlockedWebsite) SetIp(ip string) *BlockedWebsite {
	b.ip = ip
	return b
}

func (b *BlockedWebsite) SetBlockInitiator(blockInitiator string) *BlockedWebsite {
	b.blockInitiator = blockInitiator
	return b
}

func (b *BlockedWebsite) SetBlocksCount(blocksCount int) *BlockedWebsite {
	b.blocksCount = blocksCount
	return b
}

func (b BlockedWebsite) String() string {
	return fmt.Sprintf("[%v] url: %v ip: %v", b.date.Format(time.RFC3339), b.url, b.ip)
}
