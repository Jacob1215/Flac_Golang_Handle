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
	getPulsePsa(rows)
}

func getPulsePsa(rows [][]string) {
	filetitle:="waves_pulse_or_not_all-final.xlsx"
	filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+filetitle
	xlsx,err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := xlsx.GetRows("waves_pulse_or_not_all-final")
	var pulse [][]string
	var noPulse [][]string
	for i:=402;i<505;i++{
		pulse =append(pulse,lines[i])
	}
	for j:=506;j<644;j++{
		noPulse = append(noPulse,lines[j])
	}
	records := make(map[string]bool)
	pulsePsa := make([][]string,100)
	counts := 0
	for _,num := range pulse{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records[num[2]]==false{
				counts +=1
				records[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					pulsePsa[j]=append(pulsePsa[j],rows[j][k])
				}
			}
		}
	}
	noPulsePsa := make([][]string,99)
	counts2 := 0
	for _,num := range noPulse{
		for k,bianhao :=range rows[0]{
			if num[2] == bianhao && records[num[2]]==false{
				counts2 +=1
				records[num[2]]=true
				for j:=0;j<96;j++{//0-95行
					noPulsePsa[j]=append(noPulsePsa[j],rows[j][k])
				}
			}
		}
	}
	objFile := "pulse_or_nopulse_psa_12_21.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(pulsePsa)
	w.Flush()
	_=w.WriteAll(noPulsePsa)
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