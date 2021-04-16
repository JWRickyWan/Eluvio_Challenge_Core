package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"
)
//maximum concurrency set as 5
const Max int = 5
func main() {
	args := os.Args
	URL :="https://eluv.io/items/"
	path :=args[1]
	//read the input file
	file,err:= os.Open(path)
	if err!=nil{
		log.Fatal(err)
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	var IDs []string
	//store the id each item as string into the array ID
	for scanner.Scan(){
		IDs = append(IDs, scanner.Text())
	}
	//Initialize the IDStore
	idStore:=NewStore()
	//initialize the concurrent channel
	var ch =make(chan int,len(IDs))
	var tasks sync.WaitGroup
	//If we have less than Max number of tasks, in this case, 5, we don't need Max number of concurrency.
	//therefore, here we take the min of two numbers
	concurrency := int(math.Min(float64(Max), float64(len(IDs))))
	tasks.Add(concurrency)
	//this block was inspired by a stack overflow discussion
	//the basic idea is to implement semaphore to avoid the race condition and achieve concurrency.
	//https://stackoverflow.com/questions/25306073/always-have-x-number-of-goroutines-running-at-any-time
	for i:=0;i<concurrency;i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok {
					tasks.Done()
					return
				}
				Request(URL,IDs[a], idStore)
			}
		}()
	}
	for i:=0;i<len(IDs);i++{
		ch<- i
	}
	close(ch)
	tasks.Wait()
	result, err := os.Create("result.txt")
	if err!=nil {
		log.Fatal("*******Can't create result.csv!*******")
	}
	defer result.Close()
	//store the response from the server into "result.txt" in a format of "ID", "response"
	for id,content:=range idStore.IDMap{
		Entry:= id+string(content)+"\n"
		bytes,err:=io.WriteString(result,Entry)
		if err!=nil{
			log.Fatal(err)
		}
		//keep tracking of number of bytes written in each response for potential reference
		fmt.Println("Wrote "+strconv.Itoa(bytes)+" byte of data")
	}

}

//the function to request with header
func Request (URL string,id string,idStore *IDStore)  {
	req,err:=http.NewRequest("GET",URL+id,nil)
	if err != nil {
		log.Fatal(err)
	}
	//here set the authorization. I assume that authorization for every request is the same
	//if not, it can be easily changed to read from input file.Similar to line 31-33
	req.Header.Set("Authorization", "Y1JGMmR2RFpRc211MzdXR2dLNk1UY0w3WGpI")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Cannot get response!")
	}
	//check if we are sending too many requests at the same time
	if resp.StatusCode==429{
		log.Fatal("Too many requests!")
	}
	//check if the response is code 200
	if resp.StatusCode!=http.StatusOK {
		fmt.Println("Something is wrong with your request! got an "+ strconv.Itoa(resp.StatusCode)+" status code")
	}
	defer resp.Body.Close()
	//read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal("*******Error in accessing the body of the retrieved Info!*******")
	}
	//store the response into IDStore
	idStore.PutBlock(id,body)
}


