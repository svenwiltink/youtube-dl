package youtubedl

import (
	"context"
	"encoding/json"
	"os/exec"
	"time"
)

var timeout = time.Second * 5

type MetaData struct {
	Title       string      `json:"title"`
	Thumbnail   string      `json:"thumbnail"`
	Thumbnails  []Thumbnail `json:"thumbnails"`
	Duration    float64     `json:"duration"`
	Description string      `json:"description"`
	Formats     []Format    `json:"formats"`
	Uploader    string      `json:"uploader"`
}

type Thumbnail struct {
	Id  int    `json:"id,string"`
	Url string `json:"url"`
}

type Format struct {
	Protocol string  `json:"protocol"`
	Ext      string  `json:"ext"`
	Format   string  `json:"format"`
	Url      string  `json:"url"`
	Abr      float64 `json:"abr"`
	Vcodec   string  `json:"vcodec"`
}

func GetMetaData(url string) (*MetaData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "youtube-dl", "-j", url)
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	if ctx.Err() != nil {
		return nil, err
	}

	metaData := &MetaData{}
	if err = json.Unmarshal(output, metaData); err != nil {
		return nil, err
	}

	return metaData, err
}

func SetTimeout(newtimeout time.Duration) {
	timeout = newtimeout
}
