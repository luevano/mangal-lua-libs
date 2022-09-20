package http_test

import (
	"crypto/subtle"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	luahttp "github.com/metafates/mangal-lua-libs/http"
	"github.com/metafates/mangal-lua-libs/inspect"
	"github.com/metafates/mangal-lua-libs/plugin"
	luatime "github.com/metafates/mangal-lua-libs/time"
	lua "github.com/yuin/gopher-lua"
)

func basicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func httpCheckHeaders(w http.ResponseWriter, r *http.Request) {
	if header := r.Header.Get(`simple_header`); header != `check_header` {
		w.WriteHeader(412)
		return
	}
	w.Write([]byte(`OK`))
}

func httpCheckUserAgent(w http.ResponseWriter, r *http.Request) {
	if ua := r.UserAgent(); ua != `check_ua` {
		w.WriteHeader(412)
		return
	}
	w.Write([]byte(`OK`))
}

func getFormFile(r *http.Request, key, filename string) (err error) {
	file, header, err := r.FormFile(key)
	if err != nil {
		return err
	}
	defer file.Close()
	if header.Filename != filename {
		return fmt.Errorf("bad filename, get: %s except: %s\n", header.Filename, filename)
	}
	return nil
}

func httpUploadFile(w http.ResponseWriter, r *http.Request) {
	err := getFormFile(r, "file", "test.txt")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(`OK`))
}

func httpUploadFileWithFields(w http.ResponseWriter, r *http.Request) {
	err := getFormFile(r, "file", "test.txt")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if r.FormValue("foo") != "bar" {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(`OK`))
}

func httpUploadMultipleFile(w http.ResponseWriter, r *http.Request) {
	err := getFormFile(r, "file", "test.txt")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = getFormFile(r, "file1", "test1.txt")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(`OK`))
}

func httpUploadMultipleFileWithFields(w http.ResponseWriter, r *http.Request) {
	err := getFormFile(r, "file", "test.txt")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = getFormFile(r, "file1", "test1.txt")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if r.FormValue("foo") != "bar" {
		w.WriteHeader(400)
		return
	}
	w.Write([]byte(`OK`))
}

func httpRouterGet(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(`OK`))
}

func httpRouterGetTimeout(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(2 * time.Second)
	w.Write([]byte(`OK`))
}

func runHttp(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func runHttps(addr string) {
	err := http.ListenAndServeTLS(addr, "./test/server.crt", "./test/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func request(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if strings.HasPrefix(url, string(body)) {
		return fmt.Errorf("bad url, get: %s except: %s\n", body, url)
	}
	return nil
}

func manyRequest(addr string) *errgroup.Group {
	eg := &errgroup.Group{}
	eg.Go(func() error {
		time.Sleep(5 * time.Second)
		for count := 0; count < 10; count++ {
			url := fmt.Sprintf("%s/%s?d=%d", addr, "url", count)
			func(url string) {
				eg.Go(func() error {
					return request(url)
				})
			}(url)
		}
		return nil
	})
	return eg
}

func TestApi(t *testing.T) {

	http.HandleFunc("/get", httpRouterGet)
	http.HandleFunc("/getBasicAuth", basicAuth(httpRouterGet, "admin", "123456", "Please enter your username and password for this site"))
	http.HandleFunc("/timeout", httpRouterGetTimeout)
	http.HandleFunc("/checkHeader", httpCheckHeaders)
	http.HandleFunc("/checkUserAgent", httpCheckUserAgent)
	http.HandleFunc("/upload", httpUploadFile)
	http.HandleFunc("/uploadWithFields", httpUploadFileWithFields)
	http.HandleFunc("/uploadMultiple", httpUploadMultipleFile)
	http.HandleFunc("/uploadMultipleWithFields", httpUploadMultipleFileWithFields)

	go runHttp(":1111")
	go runHttps(":1112")
	time.Sleep(time.Second)

	state := lua.NewState()
	defer state.Close()

	luahttp.Preload(state)
	luatime.Preload(state)
	inspect.Preload(state)
	plugin.Preload(state)

	t.Run("test_client", func(t *testing.T) {
		assert.NoError(t, state.DoFile("./test/test_client.lua"))
	})

	t.Run("test_server_accept", func(t *testing.T) {
		eg := manyRequest("http://127.0.0.1:1113")
		assert.NoError(t, state.DoFile("./test/test_server_accept.lua"))
		assert.NoError(t, eg.Wait())
	})

	t.Run("test_server_handle", func(t *testing.T) {
		eg := manyRequest("http://127.0.0.1:2113")
		assert.NoError(t, state.DoFile("./test/test_server_handle.lua"))
		assert.NoError(t, eg.Wait())
	})

	t.Run("test_server_accept_stop", func(t *testing.T) {
		assert.NoError(t, state.DoFile("./test/test_server_accept_stop.lua"))
	})

	t.Run("test_serve_static", func(t *testing.T) {
		assert.NoError(t, state.DoFile("./test/test_serve_static.lua"))
	})
}
