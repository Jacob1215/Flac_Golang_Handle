package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"log"
	"os"
	"strconv"
)

func main()  {
	//现在进行写入文件夹的操作
	objFile := "waves_PSa_12_21.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21"+objFile
	nfs, err := os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot create file,err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	//这个文件是获得反应谱
	for i:=1;i<=2;i++{
		if i!=3{
			i_string := strconv.Itoa(i)
			filetitle := "0.4gt_p"+i_string
			filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/"+filetitle+"/_SearchResults.xlsx"
			res := getFile(filepath)
			w.Flush()
			w.WriteAll(res)
		}
	}
}

func getFile(filepath string)([][]string){
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("_SearchResults" )
	var resData [][]string
	for i:= 94;i<190;i++{//0到5秒的波
		line := rows[i]
		resData = append(resData,line)
	}
	fmt.Println(resData)
	return resData
}
