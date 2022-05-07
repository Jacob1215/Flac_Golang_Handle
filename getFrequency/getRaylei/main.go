package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"math"
	"os"
	"strconv"
)

func main()  {
	FilePath:= "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/Freq_FourierAmplitude_all.xlsx"
	xlsx, err := excelize.OpenFile(FilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("Freq_FourierAmplitude_all" )
	raley :=make([][]string,402)
	for i,line := range rows{
		i_string := strconv.Itoa(i)
		zuni,pinlv := getRaylei(line[1])
		linjiezuni  :=strconv.FormatFloat(zuni,'E',-1,64)
		zuixiaopinlv :=strconv.FormatFloat(pinlv,'E',-1,64)
		raley[i] =append(raley[i],i_string,linjiezuni,zuixiaopinlv)
	}
	objFile := "RayleiZuni.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(raley)
}

func getRaylei(Freq string) (float64,float64){
	Vs := 250.0
	H:= 54.5
	Ffund,_ := strconv.ParseFloat(Freq,64)
	zuni := 0.05
	f1 := Vs/4/H
	f2 := Ffund/f1
	w1 := 2*3.14*f1
	w2 := 2*3.14*f2
	alpha := 2*zuni*w1*w2/(w1+w2)
	beta := 2*zuni/(w1+w2)
	linJieZuNi := math.Sqrt(alpha*beta)
	zuiXiaoZhongXinPinlv := math.Sqrt(alpha/beta)
	return linJieZuNi,zuiXiaoZhongXinPinlv
}
