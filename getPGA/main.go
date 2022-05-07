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
	parameters := getAllPsa()
	postFile(parameters)

}

func postFile(parameters [][]string) {
	objFile := "parameters_all_12_27.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/new/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(parameters)
}

func getAllPsa()  [][]string{
	parameters :=make([][]string,401)
	parameters[0] = append(parameters[0],"waves_no")
	filepath2 := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/need_waves.xlsx"
	xlsx2, _ := excelize.OpenFile(filepath2)
	wavesRecords := xlsx2.GetRows("Sheet1" )
	for i:=1;i<=2;i++{
		//获取PGA
		i_string := strconv.Itoa(i)
		filepath := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/DONE/"+i_string+".xlsx"
		xlsx, err := excelize.OpenFile(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rows := xlsx.GetRows("Sheet1" )
		//fmt.Println(rows)
		headline := "seismic_waves_"+i_string
		data := make([][]string,len(rows)+1)
		data[0] = append(data[0],headline)
		linesCount := strconv.Itoa(len(rows)-1)
		timestep := rows[2][0]
		line2:= linesCount+" "+timestep
		//fmt.Println(rows[2][0])
		time_step,_ := strconv.ParseFloat(rows[2][0],64)
		a:= strconv.Itoa(len(rows)-1)
		b,_:= strconv.ParseFloat(a,64)
		duration := time_step *b
		fmt.Println(duration)
		data[1] = append(data[1],line2)
		for j,value := range rows{
			if j != 0{
				data[j+1] =append(data[j+1],value[1])
			}
		}
		//fmt.Println(data)
		i_string = strconv.Itoa(i)
		objFile :=i_string+".txt"
		NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/new/"+objFile
		nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
		defer nfs.Close()
		nfs.Seek(0, io.SeekEnd)
		w := csv.NewWriter(nfs)
		//设置属性
		w.Comma = ','
		w.UseCRLF = true
		w.Flush()
		_=w.WriteAll(data)
		//获取各PGA等参数
		parameters[i] =append(parameters[i],i_string)
		parameters[i] = append(parameters[i],wavesRecords[i-1][1])
		duration_string := strconv.FormatFloat(duration,'E',-1,64)
		parameters[i] = append(parameters[i],duration_string)
		parameters[0] =append(parameters[0],"records_no")
		parameters[0] =append(parameters[0],"duration")
		for j,para:= range rows{
			if j>=1&&j<29{
				if i == 1{
					parameters[0] = append(parameters[0],para[2])
				}
				parameters[i] = append(parameters[i],para[3])
			}
		}
	}
	//fmt.Println(parameters)
	return parameters
}
