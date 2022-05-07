package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main()  {
	//没有写signal_wave和pgv_residual_wave的
	filePath := "C:/Users/JACOB/Desktop/seismic_wave/data_3-15_Final_2/PGa_PI_Calc/"
	//现在进行写入文件夹的操作
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/data_3-15_Final_2/PI_IM_all_301-313.csv"
	var ResData [][]string
	for i := 301;i<=313;i++ {
		for j := 1; j <= 10; j++ {
			i_string := strconv.Itoa(i)
			j_string := strconv.Itoa(j)
			PGV_resid_File := filePath + "PGV_resid_" + i_string +"_"+j_string+ ".csv"
			Pulse_Indicator_File := filePath + "Pulse_Indicator_" + i_string +"_"+j_string+ ".csv"
			Pulse_Wave_File := filePath + "Pulse_Wave_" + i_string +"_"+j_string+ ".csv"
			Record_dt_File := filePath + "Record_dt_" + i_string +"_"+j_string+ ".csv"
			Tp_File := filePath + "Tp_" + i_string +"_"+j_string+ ".csv"
			PGV_resid := getSingalPara(PGV_resid_File)
			Pules_Indicator := getSingalPara(Pulse_Indicator_File)
			Record_dt := getSingalPara(Record_dt_File)
			Tp := getSingalPara(Tp_File)
			CSV_Pulse,PGV_pulse := getPulse_CSV(Record_dt, Pulse_Wave_File)
			ResData = append(ResData, []string{
				PGV_resid, Tp, Pules_Indicator,
				strconv.FormatFloat(CSV_Pulse, 'E', -1, 64),
				strconv.FormatFloat(PGV_pulse,'E',-1,64),
			})
		}
	}

	//NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/data_1-15/"+st_string+"MAXtunnel_disp_DI.csv"
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

	headline := []string{"PGV_Resid","Tp","Pulse_Indicator","CSV_Pulse","PGV_Pusle"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(ResData)
}

func getPulse_CSV(dt string, file string) (float64,float64) {
	x,_ :=ioutil.ReadFile(file)
	xlines :=strings.Split(string(x),"\n")
	PGV := getPGV(xlines)
	var result float64
	for i:=0;i<len(xlines);i++{
		value,_ := strconv.ParseFloat(xlines[i],64)
		dt_s,_ := strconv.ParseFloat(dt,64)
		result += math.Pow(value,2)*dt_s
	}
	return result,PGV
}

func getPGV(xlines []string) float64 {
	max := 0.0
	for _,value := range xlines{
		a,_ := strconv.ParseFloat(value,64)
		valueAbs := math.Abs(a)
		if valueAbs >max{
			max = valueAbs
		}
	}
	return max
}

func getSingalPara(file string) string{
	x,_ := ioutil.ReadFile(file)
	Xlines := strings.Split(string(x),"\n")
	return Xlines[0]
}