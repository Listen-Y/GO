package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main1() {

	var num int64

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer waitGroup.Done()
			start := time.Now()
			// 通过get请求
			req, err := http.NewRequest(http.MethodGet, "http://localhost:9877/api/teacher?pageNum=1&pageSize=10", nil)
			if err != nil {
				panic(err)
			}
			// 设置header
			req.Header.Set("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIxIiwiZXhwIjoxNjUzNDg2MTcwfQ.uL_LwRNMXSd-8bxBU94l-n9Bzo6UObsDSlEqR6lBppw")
			req.Header.Set("Cookie", "Idea-65d8d5c3=26490892-a9bb-4d3d-ae03-9901525193f9; Webstorm-f6d54fcd=592c7868-1e73-4489-aa8e-83d7427526aa; Goland-fb5d00a=833c5de8-e7b3-41cc-85c2-a2cfd58292f1; SECKEY_ABVK=HomcNfVN5hnuUALOgaQp+OEXMorxxOGva0KXu1iSiIU%3D; BMAP_SECKEY=1cKOn2eMNQPSVAxEdhICX_NmyiaTUei3GYTOFKb4nMwALzOM1Pkml0Y7kSDpaDVUsLoNa_x5hPmern2Vw8OLr4mUVINuwOQwA4ubfBunCjngvkwzUd_0pGdWgwmKv1IdkrC4DQy6QMdQrY2U36Wuf2mruQFmnXecx3JlN1lPyR719jCdEbo5AYemqk8vXHt_")

			c := http.Client{}
			resp, err := c.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(body))
			subMs := time.Now().Sub(start).Milliseconds()
			atomic.AddInt64(&num, subMs)
		}()
	}

	waitGroup.Wait()
	fmt.Printf("平均耗时: %d", num/100)
}

func main() {

	var num int64

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer waitGroup.Done()

			for i := 0; i < 100; i++ {

				// 通过get请求
				req, err := http.NewRequest(http.MethodGet, "http://localhost:9877/api/teacher?pageNum=1&pageSize=10", nil)
				if err != nil {
					panic(err)
				}
				// 设置header
				req.Header.Set("token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIxIiwiZXhwIjoxNjUzNDg2MTcwfQ.uL_LwRNMXSd-8bxBU94l-n9Bzo6UObsDSlEqR6lBppw")
				req.Header.Set("Cookie", "Idea-65d8d5c3=26490892-a9bb-4d3d-ae03-9901525193f9; Webstorm-f6d54fcd=592c7868-1e73-4489-aa8e-83d7427526aa; Goland-fb5d00a=833c5de8-e7b3-41cc-85c2-a2cfd58292f1; SECKEY_ABVK=HomcNfVN5hnuUALOgaQp+OEXMorxxOGva0KXu1iSiIU%3D; BMAP_SECKEY=1cKOn2eMNQPSVAxEdhICX_NmyiaTUei3GYTOFKb4nMwALzOM1Pkml0Y7kSDpaDVUsLoNa_x5hPmern2Vw8OLr4mUVINuwOQwA4ubfBunCjngvkwzUd_0pGdWgwmKv1IdkrC4DQy6QMdQrY2U36Wuf2mruQFmnXecx3JlN1lPyR719jCdEbo5AYemqk8vXHt_")

				c := http.Client{}
				resp, err := c.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()
				_, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				if resp.StatusCode != 200 {
					atomic.AddInt64(&num, 1)
				}
			}
		}()
	}

	waitGroup.Wait()
	fmt.Printf("成功率: %d %", ((10000-num)/10000)*100)
}
