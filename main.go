package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const fileName = "prime"

type flags struct {
	path string
}
func initFlag(args *flags) {
	flag.StringVar(&args.path, "path", "", "path that stores primes list.")
}
var args flags

func init() {
	initFlag(&args)
}

type Response struct {
	Output      int `json:"output"`
	ElapsedTime int `json:"elapsedTime"`
}

// CORSMiddleWare defines Access-Control-Allow
func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	var err error
	flag.Parse()
	if args.path == "" {
		if args.path, err = os.Getwd(); err != nil {
			panic(err)
		}
	}
	// load primes
	prime := NewPrime()
	if err = prime.LoadPrimes(filepath.Join(args.path, fileName), math.MaxInt32); err != nil {
		panic(err)
	}

	// start server
	router := gin.New()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}), CORSMiddleWare())
	router.Use(gin.Recovery())

	// v1 API
	v1 := router.Group("/v1")
	{
		v1.GET("/prime", func(c *gin.Context) {
			number, err := strconv.Atoi(c.Query("number"))
			if err != nil {
				c.String(http.StatusBadRequest, "number must be integer")
			} else {
				startTime := time.Now().UnixNano()
				result := prime.BinarySearch(0, len(prime)-1, number)
				endTime := time.Now().UnixNano()
				c.JSON(http.StatusOK, &Response{
					Output:      result,
					ElapsedTime: int(endTime-startTime),
				})
			}
		})
	}
	router.Run(":8080")
}