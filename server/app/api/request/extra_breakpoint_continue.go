package request

import (
	model "github.com/flipped-aurora/gf-vue-admin/server/app/model/extra"
	"mime/multipart"
	"strconv"
)

type BaseBreakpointContinueChunk struct {
	FileChunkPath   string `json:"fileChunkPath" form:"fileChunkPath"`
	FileChunkNumber int    `json:"fileChunkNumber" form:"fileChunkNumber"`
}

type BaseBreakpointContinue struct {
	FileMd5     string `json:"fileMd5" form:"fileMd5"`
	FileName    string `json:"fileName" form:"fileName"`
	FilePath    string `json:"filePath" form:"filePath"`
	ChunkMd5    string `json:"chunkMd5" form:"chunkMd5"`
	ChunkNumber string `json:"chunkNumber" form:"chunkNumber"`
	ChunkTotal  string `json:"chunkTotal" form:"chunkTotal"`

	ChunkNumberInt int
}

type FindFile struct {
	Header *multipart.FileHeader
	BaseBreakpointContinue
}

func (f *FindFile) Find() *model.BreakpointContinue {
	var chunkTotal, _ = strconv.Atoi(f.ChunkTotal)
	return &model.BreakpointContinue{
		FileMd5:    f.FileMd5,
		FileName:   f.FileName,
		ChunkTotal: chunkTotal,
	}
}

type BreakpointContinueFinish struct {
	BaseBreakpointContinue
}

type BreakpointContinue struct {
	BaseBreakpointContinue
}

func (b *BreakpointContinue) Create() *model.BreakpointContinue {
	var chunkTotal, _ = strconv.Atoi(b.ChunkTotal)
	return &model.BreakpointContinue{
		FileMd5:    b.FileMd5,
		FileName:   b.FileName,
		ChunkTotal: chunkTotal,
	}
}

type RemoveChunk struct {
	BaseBreakpointContinue
}

type CreateFileChunk struct {
	GetById
	BaseBreakpointContinueChunk
}

func (c *CreateFileChunk) Create() *model.BreakpointContinueChunk {
	return &model.BreakpointContinueChunk{
		FileID:          c.Id,
		FileChunkPath:   c.FileChunkPath,
		FileChunkNumber: c.FileChunkNumber,
	}
}
