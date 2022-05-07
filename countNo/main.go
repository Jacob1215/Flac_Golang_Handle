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
	objFile := "waves_number_all_final.csv"
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
	//这个文件是获得反应谱
	filetitle := "waves_all_12-21.xlsx"
	filePath := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+filetitle
	count,waves := CountNum(filePath)
		fmt.Println(count)
	headline := []string{"EarthquakeName","Year", "RecordNo", "Pulse", "Ds5_75","Ds5_95","AI","Mag","Rrup","PGA","Vs_jianqiebosu_m/s"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//划分5_75的持时
	Ds5_75(waves)
	//划分非脉冲
	PulseOrNot(waves)
	//jin远场计数
	Rrup(waves)
	w.Flush()
	w.WriteAll(waves)
}

func Rrup(waves [][]string) {
	var Far [][]string
	var Jin [][]string
	FarCount := 0
	JinCount := 0
	for _,line := range waves{
		distance,_:=strconv.ParseFloat(line[8],64)
		if distance<=30 {
			Far = append(Far,line)
			FarCount += 1
		} else if distance > 30 {
			Jin = append(Jin,line)
			JinCount += 1
		}
	}
	Fc := strconv.Itoa(FarCount)
	Jc:= strconv.Itoa(JinCount)
	objFile := "waves_jin_or_yuan_final.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	headline := []string{"jin_wave",Jc}
	_= w.Write(headline)
	w.Flush()
	_=w.WriteAll(Jin)
	headline2 := []string{"Far_wave",Fc}
	_= w.Write(headline2)
	w.Flush()
	_=w.WriteAll(Far)
	return
}

func PulseOrNot(waves [][]string) {
	var pulse [][]string
	var noPulse [][]string
	pulseCount :=0
	noPulseCount := 0
	var pulse_jin [][]string
	var noPulse_jin [][]string
	pulseCountjin := 0
	noPulseCountjin := 0
	for _,line := range waves {
		TpPulse,_:= strconv.ParseFloat(line[3],64)
		distance,_ := strconv.ParseFloat(line[8],64)
		if TpPulse != 0{
			pulse = append(pulse,line)
			pulseCount+=1
		} else if TpPulse ==0 {
			noPulse = append(noPulse,line)
			noPulseCount +=1
		}
		if TpPulse != 0 && distance <=30{
			pulse_jin=append(pulse_jin,line)
			pulseCountjin +=1
		} else if TpPulse ==0 &&distance <=30{
			noPulse_jin = append(noPulse_jin,line)
			noPulseCountjin +=1
		}
	}
	pc := strconv.Itoa(pulseCount)
	npc := strconv.Itoa(noPulseCount)
	pjc := strconv.Itoa(pulseCountjin)
	npjc := strconv.Itoa(noPulseCountjin)

	objFile := "waves_pulse_or_not_all-final.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	headline := []string{"pulse_wave",pc}
	_= w.Write(headline)
	w.Flush()
	_=w.WriteAll(pulse)
	headline2 := []string{"noPulse_wave",npc}
	_= w.Write(headline2)
	w.Flush()
	_=w.WriteAll(noPulse)
	headline3 := []string{"Pulse_jin_wave",pjc}
	_= w.Write(headline3)
	w.Flush()
	_=w.WriteAll(pulse_jin)
	headline4 := []string{"noPulse_jin_wave",npjc}
	_= w.Write(headline4)
	w.Flush()
	_=w.WriteAll(noPulse_jin)
	return
}

func Ds5_75(waves [][]string){
	//输出文件夹
	objFile := "waves_short_or_long_all_final.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_wave_12-17/12_21/"+objFile
	nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	//DS5_75划分长短持时，没有划分近远场
	var Ds5_75_short [][]string
	var Ds5_75_long [][]string
	longCount :=0
	shortCount:=0
	for _,line := range waves {
		D5_75,_:= strconv.ParseFloat(line[4],64)
		if D5_75<=25 {
			Ds5_75_short = append(Ds5_75_short,line)
			shortCount+=1
		} else if D5_75 > 25 {
			Ds5_75_long = append(Ds5_75_long,line)
			longCount +=1
		}
	}
	//75划分断层距
	longJinCount :=0
	shortJinCount:=0
	longYuanCount := 0
	shortYuanCount :=0
	var Ds5_75_short_jin [][]string
	var Ds5_75_long_jin [][]string
	var Ds5_75_short_yuan [][]string
	var Ds5_75_long_yuan [][]string
	for _,line := range waves {
		D5_75,_:= strconv.ParseFloat(line[4],64)
		distance,_:= strconv.ParseFloat(line[8],64)
		if D5_75<=25 && distance<=30 {
			Ds5_75_short_jin = append(Ds5_75_short_jin,line)
			shortJinCount+=1
		} else if D5_75 > 25 && distance <=30 {
			Ds5_75_long_jin = append(Ds5_75_long_jin,line)
			longJinCount +=1
		} else if D5_75 <=25 && distance>30{
			Ds5_75_short_yuan = append(Ds5_75_short_yuan,line)
			shortYuanCount +=1
		} else if D5_75 >25 && distance >30 {
			Ds5_75_long_yuan = append(Ds5_75_long_yuan,line)
			longYuanCount +=1
		}
	}
	lJc := strconv.Itoa(longJinCount)
	sjc:=strconv.Itoa(shortJinCount)
	lyc:= strconv.Itoa(longYuanCount)
	syc := strconv.Itoa(shortYuanCount)
	headline5 := []string{"ds5_75_short_jin",sjc}
	_= w.Write(headline5)
	w.Flush()
	_=w.WriteAll(Ds5_75_short_jin)
	headline6 := []string{"ds5_75_long_jin",lJc}
	_= w.Write(headline6)
	w.Flush()
	_=w.WriteAll(Ds5_75_long_jin)
	headline7 := []string{"ds5_75_short_yuan",syc}
	_= w.Write(headline7)
	w.Flush()
	_=w.WriteAll(Ds5_75_short_yuan)
	headline8 := []string{"ds5_75_long_yuan",lyc}
	_= w.Write(headline8)
	w.Flush()
	_=w.WriteAll(Ds5_75_long_yuan)

	//95断层距
	var Ds5_95_short_jin [][]string
	var Ds5_95_long_jin [][]string
	var Ds5_95_short_yuan [][]string
	var Ds5_95_long_yuan [][]string
	longJinCount2 :=0
	longYuanCount2:=0
	shortJinCount2 := 0
	shortYuanCount2 := 0
	for _,line := range waves {
		D5_95,_:= strconv.ParseFloat(line[5],64)
		distance,_:= strconv.ParseFloat(line[8],64)
		if D5_95<=45 &&distance <=30 {
			Ds5_95_short_jin = append(Ds5_95_short_jin,line)
			shortJinCount2+=1
		} else if D5_95 > 45 && distance <=30{
			Ds5_95_long_jin = append(Ds5_95_long_jin,line)
			longJinCount2 +=1
		} else if D5_95 <= 45 &&distance>30{
			Ds5_95_short_yuan =append(Ds5_95_short_yuan,line)
			shortYuanCount2 +=1
		} else if D5_95 >45 && distance >30 {
			Ds5_95_long_yuan = append(Ds5_95_long_yuan,line)
			longYuanCount2 +=1
		}
	}
	sjc2:=strconv.Itoa(shortJinCount2)
	lJc2 := strconv.Itoa(longJinCount2)
	syc2:= strconv.Itoa(shortYuanCount2)
	lyc2 := strconv.Itoa(longYuanCount2)
	headline9 := []string{"ds5_95_short_jin",sjc2}
	_= w.Write(headline9)
	w.Flush()
	_=w.WriteAll(Ds5_95_short_jin)
	headline10 := []string{"ds5_95_long_jin",lJc2}
	_= w.Write(headline10)
	w.Flush()
	_=w.WriteAll(Ds5_95_long_jin)
	headline11 := []string{"ds5_95_short_yuan",syc2}
	_= w.Write(headline11)
	w.Flush()
	_=w.WriteAll(Ds5_95_short_yuan)
	headline12 := []string{"ds5_95_long_yuan",lyc2}
	_= w.Write(headline12)
	w.Flush()
	_=w.WriteAll(Ds5_95_long_yuan)
	//Ds5_95划分长短持时，没有划分近远场。
	var Ds5_95_short [][]string
	var Ds5_95_long [][]string
	longCount2 :=0
	shortCount2:=0
	for _,line := range waves {
		D5_95,_:= strconv.ParseFloat(line[5],64)
		if D5_95<=45  {
			Ds5_95_short = append(Ds5_95_short,line)
			shortCount2+=1
		} else if D5_95 > 45 {
			Ds5_95_long = append(Ds5_95_long,line)
			longCount2 +=1
		}
	}
	lc := strconv.Itoa(longCount)
	lc2:=strconv.Itoa(longCount2)
	sc := strconv.Itoa(shortCount)
	sc2:=strconv.Itoa(shortCount2)

	headline := []string{"ds5_75_short",sc}
	_= w.Write(headline)
	w.Flush()
	_=w.WriteAll(Ds5_75_short)
	headline2 := []string{"ds5_75_long",lc}
	_= w.Write(headline2)
	w.Flush()
	_=w.WriteAll(Ds5_75_long)
	headline3 := []string{"ds5_95_short",sc2}
	_= w.Write(headline3)
	w.Flush()
	_=w.WriteAll(Ds5_95_short)
	headline4 := []string{"ds5_95_long",lc2}
	_= w.Write(headline4)
	w.Flush()
	_=w.WriteAll(Ds5_95_long)

	//DS75按时间进行划分

	var D75_0_10 [][]string
	var D75_10_20 [][]string
	var D75_20_30 [][]string
	var D75_30_40 [][]string
	var D75_40_45 [][]string
	var D75_45_50 [][]string
	var D75_50_60 [][]string
	var D75_60_70 [][]string
	var D75_70_80 [][]string
	var D75_80_90 [][]string
	var D75_90_100 [][]string

	D75_Count_0_10 :=0
	D75_Count_10_20 :=0
	D75_Count_20_30 :=0
	D75_Count_30_40 :=0
	D75_Count_40_45 :=0
	D75_Count_45_50 := 0
	D75_Count_50_60 :=0
	D75_Count_60_70 := 0
	D75_Count_70_80:=0
	D75_Count_80_90 := 0
	D75_Count_90_100 :=0
	for _,line := range waves {
		D5_75,_:= strconv.ParseFloat(line[4],64)
		switch {
		case D5_75<=10:
			D75_0_10 =append(D75_0_10,line)
			D75_Count_0_10 +=1
		case D5_75>10&&D5_75<=20:
			D75_10_20 =append(D75_10_20,line)
			D75_Count_10_20 +=1
		case D5_75>20&&D5_75<=30:
			D75_20_30 =append(D75_20_30,line)
			D75_Count_20_30 +=1
		case D5_75>30&&D5_75<=40:
			D75_30_40 =append(D75_30_40,line)
			D75_Count_30_40 +=1
		case D5_75>40&&D5_75<=45:
			D75_40_45 =append(D75_40_45,line)
			D75_Count_40_45+=1
		case D5_75>45&&D5_75<=50:
			D75_45_50 =append(D75_45_50,line)
			D75_Count_45_50 +=1
		case D5_75>50&&D5_75<=60:
			D75_50_60 =append(D75_50_60,line)
			D75_Count_50_60 +=1
		case D5_75>60&&D5_75<=70:
			D75_60_70 =append(D75_60_70,line)
			D75_Count_60_70 +=1
		case D5_75>70&&D5_75<=80:
			D75_70_80 =append(D75_70_80,line)
			D75_Count_70_80 +=1
		case D5_75>80&&D5_75<=90:
			D75_80_90 =append(D75_80_90,line)
			D75_Count_80_90 +=1
		case D5_75>90&&D5_75<=100:
			D75_90_100 =append(D75_90_100,line)
			D75_Count_90_100 +=1
		}
	}
	d75_0_10 := strconv.Itoa(D75_Count_0_10)
	d75_10_20 := strconv.Itoa(D75_Count_10_20)
	d75_20_30 := strconv.Itoa(D75_Count_20_30)
	d75_30_40 := strconv.Itoa(D75_Count_30_40)
	d75_40_45 := strconv.Itoa(D75_Count_40_45)
	d75_45_50 := strconv.Itoa(D75_Count_45_50)
	d75_50_60 := strconv.Itoa(D75_Count_50_60)
	d75_60_70 := strconv.Itoa(D75_Count_60_70)
	d75_70_80 := strconv.Itoa(D75_Count_70_80)
	d75_80_90 := strconv.Itoa(D75_Count_80_90)
	d75_90_100  := strconv.Itoa(D75_Count_90_100)

	h1 := []string{"ds5_75_0_10",d75_0_10}
	_= w.Write(h1)
	w.Flush()
	_=w.WriteAll(D75_0_10)
	h2 := []string{"ds5_75_10",d75_10_20}
	_= w.Write(h2)
	w.Flush()
	_=w.WriteAll(D75_10_20)
	h3 := []string{"ds5_75_20",d75_20_30}
	_= w.Write(h3)
	w.Flush()
	_=w.WriteAll(D75_20_30)
	h4 := []string{"ds5_75_30",d75_30_40}
	_= w.Write(h4)
	w.Flush()
	_=w.WriteAll(D75_30_40)
	h5 := []string{"ds5_75_40",d75_40_45}
	_= w.Write(h5)
	w.Flush()
	_=w.WriteAll(D75_40_45)
	h6 := []string{"ds5_75_45",d75_45_50}
	_= w.Write(h6)
	w.Flush()
	_=w.WriteAll(D75_45_50)
	h7 := []string{"ds5_75_50",d75_50_60}
	_= w.Write(h7)
	w.Flush()
	_=w.WriteAll(D75_50_60)
	h8 := []string{"ds5_75_60",d75_60_70}
	_= w.Write(h8)
	w.Flush()
	_=w.WriteAll(D75_60_70)
	h9 := []string{"ds5_75_70",d75_70_80}
	_= w.Write(h9)
	w.Flush()
	_=w.WriteAll(D75_70_80)
	h10 := []string{"ds5_75_80",d75_80_90}
	_= w.Write(h10)
	w.Flush()
	_=w.WriteAll(D75_80_90)
	h11 := []string{"ds5_75_90",d75_90_100}
	_= w.Write(h11)
	w.Flush()
	_=w.WriteAll(D75_90_100)

	//DS75按时间进行划分

	var D95_0_10 [][]string
	var D95_10_20 [][]string
	var D95_20_30 [][]string
	var D95_30_40 [][]string
	var D95_40_45 [][]string
	var D95_45_50 [][]string
	var D95_50_60 [][]string
	var D95_60_70 [][]string
	var D95_70_80 [][]string
	var D95_80_90 [][]string
	var D95_90_100 [][]string

	D95_Count_0_10 :=0
	D95_Count_10_20 :=0
	D95_Count_20_30 :=0
	D95_Count_30_40 :=0
	D95_Count_40_45 :=0
	D95_Count_45_50 := 0
	D95_Count_50_60 :=0
	D95_Count_60_70 := 0
	D95_Count_70_80:=0
	D95_Count_80_90 := 0
	D95_Count_90_100 :=0
	for _,line := range waves {
		D5_75,_:= strconv.ParseFloat(line[5],64)
		//这里没有改，只改了后面读取的位置，因为懒得改。
		switch {
		case D5_75<=10:
			D95_0_10 =append(D95_0_10,line)
			D95_Count_0_10 +=1
		case D5_75>10&&D5_75<=20:
			D95_10_20 =append(D95_10_20,line)
			D95_Count_10_20 +=1
		case D5_75>20&&D5_75<=30:
			D95_20_30 =append(D95_20_30,line)
			D95_Count_20_30 +=1
		case D5_75>30&&D5_75<=40:
			D95_30_40 =append(D95_30_40,line)
			D95_Count_30_40 +=1
		case D5_75>40&&D5_75<=45:
			D95_40_45 =append(D95_40_45,line)
			D95_Count_40_45+=1
		case D5_75>45&&D5_75<=50:
			D95_45_50 =append(D95_45_50,line)
			D95_Count_45_50 +=1
		case D5_75>50&&D5_75<=60:
			D95_50_60 =append(D95_50_60,line)
			D95_Count_50_60 +=1
		case D5_75>60&&D5_75<=70:
			D95_60_70 =append(D95_60_70,line)
			D95_Count_60_70 +=1
		case D5_75>70&&D5_75<=80:
			D95_70_80 =append(D95_70_80,line)
			D95_Count_70_80 +=1
		case D5_75>80&&D5_75<=90:
			D95_80_90 =append(D95_80_90,line)
			D95_Count_80_90 +=1
		case D5_75>90&&D5_75<=100:
			D95_90_100 =append(D95_90_100,line)
			D95_Count_90_100 +=1
		}
	}
	d95_0_10 := strconv.Itoa(D95_Count_0_10)
	d95_10_20 := strconv.Itoa(D95_Count_10_20)
	d95_20_30 := strconv.Itoa(D95_Count_20_30)
	d95_30_40 := strconv.Itoa(D95_Count_30_40)
	d95_40_45 := strconv.Itoa(D95_Count_40_45)
	d95_45_50 := strconv.Itoa(D95_Count_45_50)
	d95_50_60 := strconv.Itoa(D95_Count_50_60)
	d95_60_70 := strconv.Itoa(D95_Count_60_70)
	d95_70_80 := strconv.Itoa(D95_Count_70_80)
	d95_80_90 := strconv.Itoa(D95_Count_80_90)
	d95_90_100  := strconv.Itoa(D95_Count_90_100)

	k1 := []string{"ds5_95_0_10",d95_0_10}
	_= w.Write(k1)
	w.Flush()
	_=w.WriteAll(D95_0_10)
	k2 := []string{"ds5_95_10",d95_10_20}
	_= w.Write(k2)
	w.Flush()
	_=w.WriteAll(D95_10_20)
	k3 := []string{"ds5_95_20",d95_20_30}
	_= w.Write(k3)
	w.Flush()
	_=w.WriteAll(D95_20_30)
	k4 := []string{"ds5_95_30",d95_30_40}
	_= w.Write(k4)
	w.Flush()
	_=w.WriteAll(D95_30_40)
	k5 := []string{"ds5_95_40",d95_40_45}
	_= w.Write(k5)
	w.Flush()
	_=w.WriteAll(D95_40_45)
	k6 := []string{"ds5_95_45",d95_45_50}
	_= w.Write(k6)
	w.Flush()
	_=w.WriteAll(D95_45_50)
	k7 := []string{"ds5_95_50",d95_50_60}
	_= w.Write(k7)
	w.Flush()
	_=w.WriteAll(D95_50_60)
	k8 := []string{"ds5_95_60",d95_60_70}
	_= w.Write(k8)
	w.Flush()
	_=w.WriteAll(D95_60_70)
	k9 := []string{"ds5_95_70",d95_70_80}
	_= w.Write(k9)
	w.Flush()
	_=w.WriteAll(D95_70_80)
	k10 := []string{"ds5_95_80",d95_80_90}
	_= w.Write(k10)
	w.Flush()
	_=w.WriteAll(D95_80_90)
	k11 := []string{"ds5_95_90",d95_90_100}
	_= w.Write(k11)
	w.Flush()
	_=w.WriteAll(D95_90_100)
	return
}

func CountNum(filepath string) (int,[][]string) {

	b, e := excelize.OpenFile(filepath)
	if e != nil {
		fmt.Println("read file error")
	}

	lines := b.GetRows("waves_all_12-21" )
	records :=make(map[string]int)
	for i:=0;i<792;i++{
		records[lines[i][2]]+=1
	}
	count := len(records)//得到能用数据数量
	var waves [][]string
	used :=make(map[string]bool)
	for i := range records{
		fmt.Println(used[i])
		for _,line := range lines{
			if i == line[2] && used[i] == false{
				waves = append(waves,line)
				used[i] = true
				break
			} else if i == "0"{
				break
			}
		}
	}
	return count,waves
}
