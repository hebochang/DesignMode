package main

import (
	"bytes"
	"encoding/gob"
)

type CPU struct {
	Name string
}

type ROM struct {
	Name string
}

type Disk struct {
	Name string
}

type Computer struct {
	Cpu  CPU
	Rom  ROM
	Disk Disk
}

func (s *Computer) Clone() *Computer {
	resume := *s
	return &resume
}

func (s *Computer) BackUp() *Computer {
	pc := new(Computer)
	if err := deepCopy(pc, s); err != nil {
		panic(err.Error())
	}
	return pc
}

//深拷贝结构体
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
