package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
)

func main()  {
	filetitle:="waves_PSa_final_12_21.xlsx"
	filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+filetitle
	rows := getAllPsa(filepath)
	getDs_75_Psa(rows)
}

func getDs_75_Psa(rows [][]string) {
	filetitle:="waves_short_or_long_all_final.xlsx"
	filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+filetitle
	xlsx,err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := xlsx.GetRows("waves_short_or_long_all_final")
	var short_75 [][]string
	var long_75 [][]string
	var short_95 [][]string
	var long_95 [][]string
	for i := 804;i<1168;i++{
		short_75 = append(short_75,lines[i])
	}
	for j := 1169;j<1204;j++{
		long_75 = append(long_75,lines[j])
	}
	for i:= 1205;i<1518;i++{
		short_95 = append(short_95,lines[i])
	}
	for j := 1519 ;j<1604;j++{
		long_95 = append(long_95,lines[j])
	}
	records1 := make(map[string]bool)
	short_75Psa := make([][]string,1000)
	counts := 0
	for _,num := range short_75{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records1[num[2]]==false{
				counts +=1
				records1[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					short_75Psa[j]=append(short_75Psa[j],rows[j][k])
				}
			}
		}
	}
	records2 := make(map[string]bool)
	long_75_Psa := make([][]string,99)
	counts2 := 0
	for _,num := range long_75{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records2[num[2]]==false{
				counts2 +=1
				records2[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					long_75_Psa[j]=append(long_75_Psa[j],rows[j][k])
				}
			}
		}
	}
	records3 := make(map[string]bool)
	short_95Psa := make([][]string,100)
	counts3 := 0
	for _,num := range short_95{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records3[num[2]]==false{
				counts3 +=1
				records3[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					short_95Psa[j]=append(short_95Psa[j],rows[j][k])
				}
			}
		}
	}
	records4 := make(map[string]bool)
	long_95_Psa := make([][]string,99)
	counts4 := 0
	for _,num := range long_95{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records4[num[2]]==false{
				counts4 +=1
				records4[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					long_95_Psa[j]=append(long_95_Psa[j],rows[j][k])
				}
			}
		}
	}
	fmt.Println(counts)
	fmt.Println(counts2)
	fmt.Println(counts3)
	fmt.Println(counts4)
	objFile := "short_or_long_all_psa.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(short_75Psa)
	w.Flush()
	_=w.WriteAll(long_75_Psa)
	w.Flush()
	_=w.WriteAll(short_95Psa)
	w.Flush()
	_=w.WriteAll(long_95_Psa)


}



func getAllPsa(filepath string) [][]string{
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("waves_PSa_final_12_21" )
	return rows
}