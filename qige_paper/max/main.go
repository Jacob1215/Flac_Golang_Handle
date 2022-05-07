package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main()  {
	filePath := "C:/Users/JACOB/Desktop/21-10-17-datacode-qige/22-2-18/backfill_grouting/"
	start := 0
	end:= 12
	getMaxMoment(filePath,start,end)
	//getAccx(filePath,start,end)
	//getPGv(filePath,start,end)//这个比较好用，但是PGA不太好用。
	//getPGD(filePath,start,end)
	getAxial(filePath,start,end)
	getShear(filePath,start,end)
	//getTunnelDisp(filePath,start,end)

}

func getShear(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"case"+i_string+"El_shear_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 10
		for j := 0;j<1;j++{
			headline := j*EachCondition
			header := lines[headline]
			start := j*EachCondition +2
			end := (j+1)*EachCondition
			var MaxM[]float64
			for m:= start;m<end;m++{
				line2 := strings.Split(lines[m],",")
				MaxMoment,_ := strconv.ParseFloat(strings.Replace(line2[3],"\r","",-1),64)
				MaxM = append(MaxM,MaxMoment)
			}
			max := 0.0
			maxElement := 0
			for k,value :=range MaxM{
				if math.Abs(value) > max {
					max = value
					maxElement = k
				}
			}
			PGa := strconv.Itoa(j+1)
			rows = append(rows,[]string{header,
				PGa,
				strconv.Itoa(maxElement),
				strconv.FormatFloat(max,'E',-1,64),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/data_1-15/"+st_string+"element_shear.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxElment","maxShear"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}


func getAxial(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"case"+i_string+"El_shear_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 34
		for j := 0;j<1;j++{
			headline := j*EachCondition
			header := lines[headline]
			start := j*EachCondition +2
			end := (j+1)*EachCondition
			var MaxM[]float64
			for m:= start;m<end;m++{
				line2 := strings.Split(lines[m],",")
				MaxMoment,_ := strconv.ParseFloat(strings.Replace(line2[3],"\r","",-1),64)
				MaxM = append(MaxM,MaxMoment)
			}
			max := 0.0
			maxElement := 0
			for k,value :=range MaxM{
				if value > max {
					max = value
					maxElement = k
				}
			}
			PGa := strconv.Itoa(j+1)
			rows = append(rows,[]string{header,
				PGa,
				strconv.Itoa(maxElement),
				strconv.FormatFloat(max,'E',-1,64),
			})
		}
	}
	st_string := strconv.Itoa(st)
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/"+st_string+"element_axial.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxElment","maxAxial","DI"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}
func getMaxMoment(filpath string,st,en int){
	var rows [][]string
	for i := st;i<=en;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"/moment/"+"seismic_"+i_string+"_moment_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 10
		for j := 0;j<1;j++{
			headline := j*EachCondition
			header := lines[headline]
			start := j*EachCondition +2
			end := (j+1)*EachCondition
			var MaxM[]float64
			for m:= start;m<end;m++{
				line2 := strings.Split(lines[m],",")
				MaxMoment,_ := strconv.ParseFloat(strings.Replace(line2[3],"\r","",-1),64)
				MaxM = append(MaxM,MaxMoment)
			}
			max := 0.0
			maxElement := 0
			for k,value :=range MaxM{
				if math.Abs(value) > max {
					max = value
					maxElement = k
				}
			}
			MRD := 413.64*1000
			DI := math.Abs(max)/MRD
			PGa := strconv.Itoa(j+1)
			var Target int
			var TargetStr string
			switch {
			case DI<=1.0:
				Target = 1
				TargetStr = "None"
			case DI> 1 &&DI<=1.5:
				Target = 2
				TargetStr = "Minor"
			case DI>1.5 && DI<=2.5:
				Target = 3
				TargetStr = "Moderate"
			case DI>2.5 && DI<= 3.5:
				Target = 4
				TargetStr = "Extensive"
			case DI>3.5:
				Target = 5
				TargetStr = "Collapse"
			}
			rows = append(rows,[]string{header,
				PGa,
				strconv.Itoa(maxElement),
				strconv.FormatFloat(max,'E',-1,64),
				strconv.FormatFloat(DI,'E',-1,64),
				strconv.Itoa(Target),TargetStr,
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/"+st_string+"element_moment_DI_MRD.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxElment","maxMoment","DI","Target","TargetStr"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}
