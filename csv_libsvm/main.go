package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {
	filePath := "C:/Users/JACOB/Desktop/seismic_wave/machine-learning/code/data1.csv"
	data,_ := ioutil.ReadFile(filePath)
	lines := strings.Split(string(data),"\n")
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var trainLines,preLines []string
	for i:= 1;i<=len(lines);i++{
		if i*4%5 == 0{
			preLines = append(preLines,lines[i-1])
		} else {
			trainLines = append(trainLines,lines[i-1])
		}
	}
	train :="train"
	pre := "predict"
	csvToLibsvm(trainLines,train)
	csvToLibsvm(preLines,pre)
}

func csvToLibsvm(lines []string,train string) {
	var TrainData string

	for i,line := range lines{
		//fmt.Println(line)
		data := strings.Split(line,",")
		i+=1
		var tezheng string
		tezheng = data[0]+" "
		for j,val := range data[1:]{
			j_string :=strconv.Itoa(j+1)
			tezheng = tezheng+j_string+ ":"+val+" "
		}
		TrainData = TrainData +tezheng

	}
	write(TrainData,train)
	fmt.Printf(train+"有%d个数据",len(lines))
	return
}

func write(TrainData,train string) {
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/machine-learning/code/"+train+".txt"
	nfs, err := os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot create file,err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := bufio.NewWriter(nfs)

	//这里必须刷新，才能将数据写入文件。
	n,_ :=w.WriteString(TrainData)
	fmt.Println("写入%d个字节",n)
	w.Flush()
}



