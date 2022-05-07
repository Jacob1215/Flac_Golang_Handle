package main


import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"log"
	"os"
	"strconv"
)

func main()  {
	//本文件是通过下载好的波通过xlsx文件获得特性的。但还未对特性做分类
	//现在进行写入文件夹的操作
	objFile := "waves_all_12-21.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, err := os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot create file,err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	for i:=3;i<=3;i++{
		var ResData [][]string
		i_string := strconv.Itoa(i)
		filetitle := "0.4g_P"+i_string
		filepath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/"+filetitle+"/_SearchResults.xlsx"
		//得到波的各种特性
		Wave := getFile(filepath)
		fmt.Println(Wave)
		//分类持时
		short,long := getDs5_75(Wave.Duration5_75)
		fmt.Println(short,long)
		for j:=0;j<23;j++{
			Record_string := strconv.FormatFloat(Wave.RecordNo[j],'E',-1,64)
			Pulse_string := strconv.FormatFloat(Wave.Pulse[j],'E',-1,64)
			Ds5_75_string := strconv.FormatFloat(Wave.Duration5_75[j],'E',-1,64)
			Ds5_95_string := strconv.FormatFloat(Wave.Duration5_95[j],'E',-1,64)
			AI_string := strconv.FormatFloat(Wave.AriasIntensity[j],'E',-1,64)
			Mag_string := strconv.FormatFloat(Wave.Magnitude[j],'E',-1,64)
			Rrup_string := strconv.FormatFloat(Wave.Rrup[j],'E',-1,64)
			PGA_string := strconv.FormatFloat(Wave.PGA[j],'E',-1,64)
			Vs30_string := strconv.FormatFloat(Wave.Vs[j],'E',-1,64)
			ResData = append(ResData,[]string{Wave.EarthquakeName[j],Wave.Year[j],Record_string,
				Pulse_string,Ds5_75_string,Ds5_95_string,AI_string,Mag_string,Rrup_string,PGA_string,
				Vs30_string,
			})
		}

		//var title []string
		//title = append(title, filetitle)
		//err = w.Write(title)
		//if err != nil {
		//	log.Fatalf("cannot write title：%v", err)
		////}
		//headline := []string{"EarthquakeName","Year", "RecordNo", "Pulse", "Ds5_75","Ds5_95","AI","Mag","Rrup","PGA","Vs_jianqiebosu_m/s"}
		//err = w.Write(headline)
		//if err != nil {
		//	log.Fatalf("can not write, err is %+v", err)
		//}
		//这里必须刷新，才能将数据写入文件。
		w.Flush()
		w.WriteAll(ResData)
	}
}

func getDs5_75(duration5_75 []float64) (int,int)  {
	var shortCount int
	var longCount int
	for _,value :=range duration5_75{
		if value <= 25 {
			shortCount += 1
		} else if value >25{
			longCount +=1
		}
	}
	return shortCount,longCount
}

type WavePerts struct {
	RecordNo []float64
	Pulse []float64
	Duration5_75 []float64
	Duration5_95 []float64
	AriasIntensity []float64
	EarthquakeName []string
	Year []string
	Magnitude []float64
	Rrup []float64
	PGA []float64
	Vs []float64

}
func getFile(filepath string) WavePerts {
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows := xlsx.GetRows("_SearchResults" )
	var Wave WavePerts
	for i:=34;i<57;i++{
		line := rows[i]
		recordNo,_ :=strconv.ParseFloat(line[2],64)
		tpPulse,_ := strconv.ParseFloat(line[5],64)
		D5_75,_ := strconv.ParseFloat(line[6],64)
		D5_95,_ := strconv.ParseFloat(line[7],64)
		AI,_ :=strconv.ParseFloat(line[8],64)
		Mag,_:= strconv.ParseFloat(line[12],64)
		Rr,_:= strconv.ParseFloat(line[15],64)
		vs,_:= strconv.ParseFloat(line[16],64)
		pga := 0.1
		Wave.RecordNo = append(Wave.RecordNo,recordNo)
		Wave.Pulse = append(Wave.Pulse,tpPulse)
		Wave.Duration5_75 = append(Wave.Duration5_75,D5_75)
		Wave.Duration5_95 = append(Wave.Duration5_95,D5_95)
		Wave.AriasIntensity = append(Wave.AriasIntensity,AI)
		Wave.EarthquakeName = append(Wave.EarthquakeName,line[9])
		Wave.Year = append(Wave.Year,line[10])
		Wave.Magnitude = append(Wave.Magnitude,Mag)
		Wave.Rrup= append(Wave.Rrup,Rr)
		Wave.PGA = append(Wave.PGA,pga)
		Wave.Vs= append(Wave.Vs,vs)

	}
	fmt.Println(Wave)
	return Wave
}