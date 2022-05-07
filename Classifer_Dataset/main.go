package main

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"math/big"
	"os"
	"strconv"
)

func main()  {
	FileName := "C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/Data_classifier_Binary.xlsx"
	xlsx,err := excelize.OpenFile(FileName)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsM := xlsx.GetRows("Data_M_750")
	rowsD := xlsx.GetRows("Data_D_750")
	//fmt.Println(rows)
	headlineM := rowsM[0][0:22]
	headlineD := rowsD[0][0:22]
	//fmt.Println(headline)
	for i:= 1;i<=14;i++{
		for k := 1;k<=10;k++{
			flag := make(map[int64]bool)
			var resDataM [][]string
			var resDataD [][]string
			resDataM = append(resDataM,headlineM)
			resDataD = append(resDataD,headlineD)
			for j:=0;j<i*50;j++{
				line := getRands(flag)
				resDataM = append(resDataM,rowsM[line][0:22])
				resDataD = append(resDataD,rowsD[line][0:22])
			}
			stringM := "M"
			stringD := "D"
			WriteFile(resDataM,i,k,stringM)
			WriteFile(resDataD,i,k,stringD)
		}
	}
}

func WriteFile(m [][]string,i,k int,string2 string) {
	//fmt.Println(data)
	i_string := strconv.Itoa(i*50)
	k_string := strconv.Itoa(k)
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/Dataset_training/Data_"+string2+"_"+i_string+"_"+k_string+"_Binary.csv"
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(m)
	//获取各PGA等参数
}

func getRands(flag map[int64]bool) int64 {
	RandS,_ := rand.Int(rand.Reader,big.NewInt(750))
	Rands := RandS.Int64()
	if flag[Rands] == false && Rands != 0 {
		flag[Rands] = true
		return Rands
	}
	return getRands(flag)
}






