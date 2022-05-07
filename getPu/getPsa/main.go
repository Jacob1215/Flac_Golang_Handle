package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strings"
)

func main()  {
	//成功了，只需要后面把需要的波全部搞到一起
	filetitle :="waves_PSa_12_21.xlsx"
	filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+filetitle
	getFile(filepath)
}

func getFile(filepath string) {
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("waves_PSa_12_21" )
	//先获取周期
	records := make(map[string]bool)
	PSa := make([][]string,4000)
	wavesCount :=0
	for i:=0;i<25;i++{
		var headline []string
		if i<16{
			headline = rows[i*96][5:30]
		} else if i>=16 &&i<20{
			headline = rows[i*96][5:28]
		} else if i>=20 &&i<21{
			headline = rows[i*96][5:105]
		} else if i>=21 && i<25 {
			headline = rows[i*96][5:55]
		}
		fmt.Println(headline[len(headline)-1:])
		for k,head := range headline{
			RSN := strings.Split(head," ")
			Num := strings.Split(RSN[0],"-")
			if records[Num[1]] == false{
				fmt.Println(Num[1])
				wavesCount +=1
				records[Num[1]] = true
				for j:=0;j<96;j++{//0-95行
					if j==0 {
						PSa[j]=append(PSa[j],Num[1])
					} else {
						PSa[j]=append(PSa[j],rows[i*96+j][k+5])
					}
				}
			}
		}

	}
	objFile := "waves_final_12_21_TODO.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(PSa)
	fmt.Println(records)
	fmt.Println(wavesCount)
}
