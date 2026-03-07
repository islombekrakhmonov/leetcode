package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// func main() {
// 	ctx := context.Background()

// 	minioClient, err := minio.New("cdn-api.easyto.travel", &minio.Options{
// 		Creds:  credentials.NewStaticV4("minio-admin", "axeiTo1aiebaiPi8Cohn4wei", ""),
// 		Secure: true, // or false if HTTP
// 	})
// 	if err != nil {
// 		fmt.Println("here", err)
// 		return
// 	}

// 	bucket := "dev"
// 	cdnPrefix := "https://cdn-api.easyto.travel/dev/" // e.g. https://cdn.example.com/

// 	var count int
// 	// List all objects
// 	for obj := range minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true}) {

// 		if obj.Err != nil {
// 			fmt.Println("Error listing object:", obj.Err)
// 			continue
// 		}

// 		fileURL := cdnPrefix + obj.Key

// 		// Check file availability via HEAD
// 		client := http.Client{Timeout: 5 * time.Second}
// 		resp, err := client.Head(fileURL)

// 		if err != nil {
// 			fmt.Println("❌ Error accessing:", fileURL)
// 			count++
// 			continue
// 		}

// 		if resp.StatusCode != http.StatusOK {
// 			fmt.Printf("❌ Broken (%d): %s\n", resp.StatusCode, fileURL)
// 			count++
// 			continue
// 		}

// 		// Optional: check for zero-byte files
// 		if resp.Header.Get("Content-Length") == "0" {
// 			fmt.Println("❌ Zero-byte file:", fileURL)
// 			count++
// 			continue
// 		}

// 		fmt.Println("count", count)
// 		// fmt.Println("✔ OK:", fileURL)
// 	}
// 	fmt.Println("count", count)
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"github.com/minio/minio-go/v7"
// 	"github.com/minio/minio-go/v7/pkg/credentials"
// )

func main() {
	ctx := context.Background()

	minioClient, err := minio.New("cdn-api.easyto.travel", &minio.Options{
		Creds:  credentials.NewStaticV4("minio-admin", "axeiTo1aiebaiPi8Cohn4wei", ""),
		Secure: true, // or false if HTTP
	})
	if err != nil {
		fmt.Println("here", err)
		return
	}

	bucket := "prod"
	cdnPrefix := "https://cdn-api.easyto.travel/prod/" // e.g. https://cdn.example.com/

	// STEP 1 — gather all object keys (fast)
	var keys []string
	for obj := range minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true}) {
		if obj.Err == nil {
			keys = append(keys, obj.Key)
		}
	}

	fmt.Println("Total objects:", len(keys))

	// STEP 2 — run concurrent workers
	workerCount := 30 // adjust as needed (20–200)
	jobs := make(chan string, workerCount)
	wg := sync.WaitGroup{}

	client := http.Client{Timeout: 5 * time.Second}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range jobs {
				fileURL := cdnPrefix + key
				resp, err := client.Head(fileURL)

				if err != nil {
					fmt.Println("❌ Error accessing:", fileURL)
					continue
				}
				if resp.StatusCode != http.StatusOK {
					fmt.Printf("❌ Broken (%d): %s\n", resp.StatusCode, fileURL)
					continue
				}
				if resp.Header.Get("Content-Length") == "0" {
					fmt.Println("❌ Zero-byte file:", fileURL)
					continue
				}

				// fmt.Println("✔ OK:", fileURL)
			}
		}()
	}

	// Feed jobs
	for _, key := range keys {
		jobs <- key
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Done")
}

func maximizeExpressionOfThree(nums []int) int {
	max1, max2 := math.MinInt, math.MinInt
	min := math.MaxInt

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max1 {
			max1, max2 = num, max1
		} else if num > max2 {
			max2 = num
		}
	}

	return max1 + max2 - min
}
