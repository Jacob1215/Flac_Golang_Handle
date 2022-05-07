package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"strconv"
)

func main()  {
	FilePath:= "C:/Users/JACOB/Desktop/seismic_wave/12_21/Ts_tezhengzhouqi.xlsx"
	xlsx, err := excelize.OpenFile(FilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("Ts_tezhengzhouqi" )
	data := make([]int,15)
	for _,line :=range rows{
		Ts,_ := strconv.ParseFloat(line[1],64)
		if Ts <=0.1{
			data[1] +=1
		} else if Ts>0.1 && Ts<=0.2{
			data[2] +=1
		} else if Ts > 0.2&&Ts<=0.3{
			data[3]+=1
		} else if Ts > 0.3&&Ts<=0.4{
			data[4]+=1
		}else if Ts > 0.4&&Ts<=0.5{
			data[5]+=1
		}else if Ts > 0.5&&Ts<=0.6{
			data[6]+=1
		}else if Ts > 0.6&&Ts<=0.7{
			data[7]+=1
		}else if Ts > 0.7&&Ts<=0.8{
			data[8]+=1
		}else if Ts > 0.8&&Ts<=0.9{
			data[9]+=1
		}else if Ts > 0.9&&Ts<=1.0{
			data[10]+=1
		}else if Ts > 1.0&&Ts<=1.2{
			data[11]+=1
		}else if Ts > 1.2 &&Ts<=1.5{
			data[12]+=1
		}else if Ts > 1.5{
			data[13]+=1
		}
	}

	fmt.Println(data)



}
