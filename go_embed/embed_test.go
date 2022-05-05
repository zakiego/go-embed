package go_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed version.txt
var version string

func TestStringEmbed(t *testing.T) {
	assert.Equal(t, version, "3.21.0")
}

//go:embed spongebob.jpg
var image []byte

func TestByteEmbed(t *testing.T) {
	err := ioutil.WriteFile("new.jpg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed new.jpg
var new []byte

func TestEqualByte(t *testing.T) {
	assert.Equal(t, image, new)
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedFileSystem(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	assert.Equal(t, string(a), "AAA")

	b, _ := files.ReadFile("files/b.txt")
	assert.Equal(t, string(b), "BBB")

	c, _ := files.ReadFile("files/c.txt")
	assert.Equal(t, string(c), "CCC")
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
