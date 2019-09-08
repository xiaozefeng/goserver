package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
	"github.com/xiaozefeng/goserver/handler"
	"github.com/xiaozefeng/goserver/pkg/errno"
	"io/ioutil"
	"regexp"
	"time"
)

type logWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w logWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path
		reg := regexp.MustCompile(`(/v1/users|/login)`)
		if ! reg.MatchString(path) {
			return
		}

		// skip for the health check request
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// read the body content
		var bs []byte
		if c.Request.Body != nil {
			bs, _ = ioutil.ReadAll(c.Request.Body)
		}

		// restore the io.ReaderCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bs))

		// the basic info
		method := c.Request.Method
		ip := c.ClientIP()

		lw := &logWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		c.Writer = lw

		// continue
		c.Next()

		//  calculates the latency
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// get code adn message
		var response handler.Response

		if err := json.Unmarshal(lw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response struct, "+
				"body :`%s`", lw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}
		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
	}
}
