package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main()  {
	res_pred := getFile()
	test := getXtest()
	Y_TEST := getYtest()
	filetest :="C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/XGBoost_xtest-D.csv"
	filepred := "C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/XGBoost_ypred-D.csv"
	//fileYtest := "C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/XGBoost_ytest-M.csv"
	writefile(test,filetest)
	var res [][]string
	res = append(res,res_pred,Y_TEST)
	writefile(res,filepred)
}

func getYtest() []string {
	filepath :="C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/XGBoost_REG_YTEST_D.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return nil
	}
	lines := strings.Split(string(b),"\n")
	pred := []string{}
	for _,value := range lines{
		line := strings.Split(value," ")
		for _,v2 := range line{
			v3:= strings.Replace(v2," ","",-1)
			pred = append(pred,v3)
		}

	}
	res_pred2 := pred[0:len(pred)]
	return res_pred2
}

func writefile(test [][]string, file string) {
	nfs, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot create file,err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	//一次写入多行
	w.WriteAll(test)
}

func getXtest() [][]string{
	filepath :="C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/XGBoost_REG_XTEST_D.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return nil
	}
	lines := strings.Split(string(b), "[")

	var test [][]string
	for _,value := range lines{
		line := strings.Split(value,"\n")
		var line_test []string
		for _,v2 := range line{
			v3 := strings.Split(v2," ")
			for _,v4 := range v3{
				v5 := strings.Replace(v4," ","",-1)
				if v5!= ""{
					line_test = append(line_test,v5)
				}
			}

		}
		test = append(test,line_test)
	}
	fmt.Println(test)
	return test
}

func getFile() []string{
	filepath :="C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/XGBoost_REG_YPRED_D.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return nil
	}
	lines := strings.Split(string(b),"\n")
	pred := []string{}
	for _,value := range lines{
		line := strings.Split(value," ")
		for _,v2 := range line{
			v3:= strings.Replace(v2," ","",-1)
			if v3 !=""{
				pred = append(pred,v3)
			}
		}
	}
	res_pred := pred[0:len(pred)]
	return res_pred
}
