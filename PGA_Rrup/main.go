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
	PGA_Rrup()
}

func PGA_Rrup() {

	FilePath:= "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/parameters_all.xlsx"
	xlsx, err := excelize.OpenFile(FilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("parameters_all" )
	Filename := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/waves_number_all_final.xlsx"
	xlsx2, _ := excelize.OpenFile(Filename)
	rows2:= xlsx2.GetRows("waves_number_all_final")
	PgaRrup := make([][]string,402)
	for i,line1 := range rows{
		i_string :=strconv.Itoa(i)
		for _,line2 := range rows2{
			if line1[1] == line2[2]{
				PgaRrup[i] = append(PgaRrup[i],i_string)
				PgaRrup[i] = append(PgaRrup[i],line1[1])
				PgaRrup[i] = append(PgaRrup[i],line1[2])
				PgaRrup[i] = append(PgaRrup[i],line2[8])
				break
			}else {
				continue
			}
		}
	}
	objFile := "PGA_Rrup2.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	w.Flush()
	_=w.WriteAll(PgaRrup )



}