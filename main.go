package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/target-sum", targetSumHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error while starting server")
	}
}

type inputParam struct {
	Target  int   `json:"target"`
	Numbers []int `json:"numbers"`
}

type output struct {
	Answer [][]int `json:"Answer"`
}

func targetSumHandler(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("problem with reading body")
		http.Error(writer, "iSE", http.StatusInternalServerError)

	}

	var ip inputParam
	var op output
	err = json.Unmarshal(body, &ip)
	if err != nil {
		fmt.Printf("problem reading post body %v", err)
		http.Error(writer, "iSE", http.StatusInternalServerError)
	}
	op.Answer = returnSumIndex(ip.Numbers, ip.Target)

	jsonResp, err := json.Marshal(op)
	if err != nil {
		fmt.Println("error while marshalling response json")
	}
	fmt.Printf("\n %v", ip)

	fmt.Println("\nexiting")
	writer.Write(jsonResp)

}

func returnSumIndex(numberArray []int, targetSum int) [][]int {
	seenMap := make(map[int]int)
	var Answer [][]int
	for index, num := range numberArray {
		numToSearch := targetSum - num
		if seenIndex, ok := seenMap[numToSearch]; ok {
			Answer = append(Answer, []int{seenIndex, index})
		}
		seenMap[num] = index
	}

	return Answer

}
