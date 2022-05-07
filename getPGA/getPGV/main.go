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
	getPgv()
}

func getPgv() {
	var PGV_PGD [][]string
	headline := []string{"num","PGV","PGD","line","dytime"}
	for i:=1;i<=400;i++{
		i_string := strconv.Itoa(i)
		filepath := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/vel_History/工作簿"+i_string+".xlsx"
		xlsx, err := excelize.OpenFile(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rows := xlsx.GetRows("Sheet1" )
		var vel []float64
		var dis []float64
		for j := 1;j<len(rows);j++{
			vel_float,_ :=strconv.ParseFloat(rows[j][1],64)
			vel = append(vel,vel_float)
			dis_float,_:=strconv.ParseFloat(rows[j][3],64)
			dis = append(dis,dis_float)
		}
		lines := len(vel)
		line_string := strconv.Itoa(lines)
		lines2,_ := strconv.ParseFloat(line_string,64)
		stepTime,_ := strconv.ParseFloat(rows[2][0],64)
		duration := lines2*stepTime
		PGV:=getPeek(vel)
		PGD:=getPeek(dis)
		//fmt.Println(rows)
		PGV_PGD = append(PGV_PGD,[]string{i_string,
			strconv.FormatFloat(PGV,'E',-1,64),
			strconv.FormatFloat(PGD,'E',-1,64),line_string,
			strconv.FormatFloat(duration,'E',-1,64)})
	}
	//fmt.Println(data)
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/vel_History/I_PGD_PGV.csv"
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	_=w.Write(headline)
	w.Flush()
	_=w.WriteAll(PGV_PGD)
	//获取各PGA等参数
}

func getPeek(vel []float64) float64 {
	var max float64
	for _,value := range vel{
		if math.Abs(value)>math.Abs(max){
			max = value
		} else {
			continue
		}
	}
	return max
}