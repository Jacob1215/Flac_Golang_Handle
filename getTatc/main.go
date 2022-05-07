package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strconv"
)

func main() {
	getTact()
}

func getTact() {
	FilePath:= "C:/Users/JACOB/Desktop/seismic_wave/12_21/waves_PSa_final_12_21.xlsx"
	xlsx, err := excelize.OpenFile(FilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("waves_PSa_final_12_21" )
	PSa := make([][]string,400)
	for _,line :=range rows[1:]{
		for j,sa :=range line{
			PSa[j] =append(PSa[j],sa)
		}
	}
	Ts := make([][]string,400)
	for j,wavse :=range PSa{
		var max []float64
		var nomax []float64
		for _,sa :=range wavse{
			saF,_ :=strconv.ParseFloat(sa,64)
			max =append(max,saF)
			nomax =append(max,saF)
		}
		var M float64
		for _,value :=range max{
			if value >= M{
				M = value
			} else {
				continue
			}
		}
		var tsMax float64
		for k := 0; k<96; k++{
			if nomax[k] == M{
				tsMax,_ = strconv.ParseFloat(rows[k][0],64)
			}
		}
		sqrtMax := M/1.4142135623
		//fmt.Println(sqrtMax)
		flag := false
		for i:=0;i<96;i++{
			//fmt.Println(nomax)
			ts,_ := strconv.ParseFloat(rows[i][0],64)
				if sqrtMax >= nomax[i] && flag == false && ts > tsMax {
					//fmt.Println(sqrtMax)
					Ts[j] = append(Ts[j],rows[0][j],rows[i][0])
					flag = true
				}
		}
		fmt.Println(Ts)
	}

	FilePath2:= "C:/Users/JACOB/Desktop/seismic_wave/12_21/waves_number_all_final.xlsx"
	xlsx2, err := excelize.OpenFile(FilePath2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows2 := xlsx2.GetRows("waves_number_all_final" )
	for i,recordNo :=range Ts[1:]{
		for _,line :=range rows2{
			if recordNo[0] == line[2]{
				Ts[i+1] =append(Ts[i+1],line[7],line[8])
			}
		}
	}
	objFile := "Ts_tezhengzhouqi.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(Ts)
}




