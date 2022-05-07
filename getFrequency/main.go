package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strconv"
)

func main()  {
	//getFrequency()
	fenlei()
}

func fenlei() {
	filepath := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/Freq_FourierAmplitude_all.xlsx"
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("Freq_FourierAmplitude_all")
	rows2 := rows[2:]
	highFreq := 0
	lowFreq := 0
	for _,line :=range rows2{
		value, _ := strconv.ParseFloat(line[1],64)
		if value >= 10.0{
			highFreq +=1
		} else {
			lowFreq +=1
		}
	}
	fmt.Println(highFreq,lowFreq)
}

func getFrequency() {
	maxFreq :=make([][]string,401)
	for i:=1;i<=400;i++ {
		i_string := strconv.Itoa(i)
		filepath := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/DONE/" + i_string + ".xlsx"
		xlsx, err := excelize.OpenFile(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rows := xlsx.GetRows("Sheet2")
		var max float64 = 0
		var maxFrequency string
		for _,line := range rows{
			FourierAmplitude,_ := strconv.ParseFloat(line[2],64)
			if FourierAmplitude >max{
				max = FourierAmplitude
				maxFrequency = line[0]
			}
		}
		maxFA := strconv.FormatFloat(max,'E',-1,64)

		maxFreq[i] = append(maxFreq[i],i_string)
		maxFreq[i] = append(maxFreq[i],maxFrequency)
		maxFreq[i] =append(maxFreq[i],maxFA)
	}
	objFile := "Freq_FourierAmplitude_all.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	headline := []string{"wavse_no","Frequency","MaxFourierAmplitude"}
	_=w.Write(headline)
	w.Flush()
	_=w.WriteAll(maxFreq)
}
