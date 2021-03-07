package request

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"
)

type Upload struct {
	Filename         string `r:"filename" json:"filename" form:"filename"`
	TotalSize        string `r:"totalSize" json:"totalSize" form:"totalSize"`
	Identifier       string `r:"identifier" json:"identifier" form:"identifier"`
	TotalChunks      string `r:"totalChunks" json:"totalChunks" form:"totalChunks"`
	ChunkNumber      string `r:"chunkNumber" json:"chunkNumber" form:"chunkNumber"`
	CurrentChunkSize string `r:"currentChunkSize" json:"currentChunkSize" form:"currentChunkSize"`
}

func (u *Upload) Create(path string) model.SimpleUploader {
	return model.SimpleUploader{
		Filename:         u.Filename,
		TotalSize:        u.TotalSize,
		Identifier:       u.Identifier,
		TotalChunks:      u.TotalChunks,
		ChunkNumber:      u.ChunkNumber,
		CurrentChunkSize: u.CurrentChunkSize,
		CurrentChunkPath: path,
	}
}

type CheckFileMd5 struct {
	Md5 string `json:"md5" form:"md5"`
}

type MergeFileMd5 struct {
	Md5      string `json:"md5" form:"md5"`
	Filename string `json:"fileName" form:"fileName"`
}
