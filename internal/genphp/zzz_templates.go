// Code generated by "esc -o zzz_templates.go -pkg genphp templates"; DO NOT EDIT.

package genphp

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/load.twig": {
		local:   "templates/load.twig",
		size:    208,
		modtime: 1526223692,
		compressed: `
H4sIAAAAAAAC/4zMvQqDMBRA4T1PcZcQhaIPYKFIsSrYVvozBzE3eCFNg6lTyLuX1rVDpzOdb7tzk2NM
oSaLCYjrvqmO5flSy7671+1J9uWtERtQNNvhgQlIeWi7SkpIIS0YIzuaRSH8/CADkZMd8+VFJnOTEwUL
HPRzBk0Ggey3Hnj8AxrN4D36PIT1jnElAwe06oPyyN4BAAD//9B/+5nQAAAA
`,
	},

	"/templates/structtypes.twig": {
		local:   "templates/structtypes.twig",
		size:    715,
		modtime: 1526223323,
		compressed: `
H4sIAAAAAAAC/3yQQa+UMBzEz+2nGBMI7EYfd1eenr16NIb0wR+pgbZpi3lrw3c3pSzirvFG+v/NMDMf
PprBcN6OwjmEAH811CgxEZYl5JA9jLCkfPPnkC+gV0+qWwUP51VHqpN9JOVkRppIeYfPTqsvZKUY5S/x
MhIC58zML6Ns4bzwskXWauXp1aNGMXhv3leVaweaxJO234vLAx7/Gtm74JEMOXpt0UsaO0iVPhzyhbPq
fN6urZ5itmaUinbo6fiaFDjHqn/Rt5rRJiHVHu+nsMhC2Oy2UCnTQbHB/axaL7XCj8M+VJ4QOGOZnuMY
wlpxLTljrPi0TVSgfv7XSG8TFedIyP02nLHThbP/DcQyd8vRoYbTTczW7I8lMj9I9+75viNWZ9mjxBvQ
ZPy1xNHrhFRr7fW12OWr/S3gN9RHUXRcUt7DeMySn61CNLpwtvCF/w4AAP//08/bRssCAAA=
`,
	},

	"/templates/util.twig": {
		local:   "templates/util.twig",
		size:    539,
		modtime: 1526222360,
		compressed: `
H4sIAAAAAAAC/5yQQW6rMBCG155TzAIJI+WdwC9p79BlVVkuGYQjsJE9rkgr7l5hCCpN1UW34/+b+X7/
fxjaAaBJrmbrHUavL9E7HSlY09l3kli8mS4RVvgBwjYokfqBr1/m+UUE4hScAjEBiMInxiOaEMxVVgpW
0ka9jL7DjQ9k6nabm4iFdWca8XjCgjrqyfGaXVb1xK0/axpt5Ci3zAHL2f/ppl9uJ4TIp/WQ4nzHJz5s
1L/TDpIVVmpGJhC5jkguEssZei4fa++YRi5fcmjpnReu5W9V/euFar7r+qN+jvwiv37okrvTVX/y3O1U
sE9M8BkAAP//+0ZubxsCAAA=
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/templates": {
		isDir: true,
		local: "templates",
	},
}
