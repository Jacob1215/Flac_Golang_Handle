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
	getRrupPsa(rows)
}

func getRrupPsa(rows [][]string) {
	filetitle:="waves_jin_or_yuan_final.xlsx"
	filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+filetitle
	xlsx,err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := xlsx.GetRows("waves_jin_or_yuan_final")
	var jin [][]string
	var far [][]string
	for i := 1;i<161;i++{
		jin = append(jin,lines[i])
	}
	for j:=162;j<400;j++ {
		far = append(far,lines[j])
	}
	records := make(map[string]bool)
	rupPsa := make([][]string,1000)
	counts := 0
	for _,num := range jin{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records[num[2]]==false{
				counts +=1
				records[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					rupPsa[j]=append(rupPsa[j],rows[j][k])
				}
			}
		}
	}
	farrupPsa := make([][]string,99)
	counts2 := 0
	for _,num := range far{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records[num[2]]==false{
				counts2 +=1
				records[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					farrupPsa[j]=append(farrupPsa[j],rows[j][k])
				}
			}
		}
	}
	fmt.Println(counts)
	fmt.Println(counts2)
	objFile := "jin_or_yuan_psa_12_21.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(rupPsa)
	w.Flush()
	_=w.WriteAll(farrupPsa)

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