package main

import (
	"flag"
	"fmt"
	prime2 "github.com/dnk90/find_left_most_prime/prime"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

type flags struct {
	memprofile string
}
func initFlag(args *flags) {
	flag.StringVar(&args.memprofile, "memprofile", "", "write memory profile to `file`")
}
var args flags

func init() {
	log.SetPrefix("API")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	initFlag(&args)
}

type Response struct {
	Output      int32 `json:"output"`
	ElapsedTime int   `json:"elapsedTime"`
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
	flag.Parse()
	// load primes
	log.Println("Loading Primes")
	prime := prime2.SieveOfSundaram(math.MaxInt32)
	log.Println("Finished loading primes, Starting API")
	// write memory profile after load primes
	if args.memprofile != "" {
		f, err := os.Create(args.memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
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
				result := prime.BinarySearch(0, int32(len(prime)-1), int32(number))
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