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
	filePath := "G:/Data_3-15_Final"
	start := 214
	end:= 238
	//getMaxMoment(filePath,start,end)
	//getAccx(filePath,start,end)
	//getPGv(filePath,start,end)
	//getPGD(filePath,start,end)
	//getAxial(filePath,start,end)
	//getShear(filePath,start,end)
	//getTunnelDisp(filePath,start,end)
	getGamaDisp(filePath,start,end)
}

func getGamaDisp(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"/max_disp/seismic"+i_string+"_tunnel_maxdisp_max.csv"
		X,_ := ioutil.ReadFile(file)
		Xlines := strings.Split(string(X),"\n")
		EachCondition := 18
		for j := 0;j<10;j++{
			headline := j*EachCondition
			header := Xlines[headline]
			start := j*EachCondition +2
			end := (j+1)*EachCondition
			var MaxM[]float64
			for m:= start;m<end;m++{
				Xline2 := strings.Split(Xlines[m],",")
				Max,_ := strconv.ParseFloat(strings.Replace(Xline2[1],"\r","",-1),64)
				MaxM = append(MaxM,Max)
			}
			max := MaxM[0]
			maxElement := 0
			D := 14.5
			DI := max/D
			PGa := strconv.Itoa(j+1)
			var Target int
			var TargetStr string
			var LabelBinary int
			switch {
			case DI<=0.0006:
				Target = 0
				LabelBinary = 0
				TargetStr = "None"
			case DI> 0.0006 &&DI<=0.0025:
				Target = 1
				LabelBinary = 0
				TargetStr = "Minor"
			case DI>0.0025 && DI<=0.004:
				Target = 2
				LabelBinary = 1
				TargetStr = "Moderate"
			case DI>0.004 && DI<= 0.006:
				Target = 3
				LabelBinary = 1
				TargetStr = "Extensive"
			case DI>0.006:
				Target = 4
				LabelBinary = 1
				TargetStr = "Collapse"
			}
			rows = append(rows,[]string{header,
				PGa,
				strconv.Itoa(maxElement),
				strconv.FormatFloat(max,'E',-1,64),
				strconv.FormatFloat(DI,'E',-1,64),
				strconv.Itoa(Target),TargetStr,
				strconv.Itoa(LabelBinary),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := filpath+"/"+st_string+"MAXtunnel_disp_DI_gama_4-9.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxElment","maxDisp","DI","Target","TargetStr","LabelBinary"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}

func getTunnelDisp(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"/max_disp/seismic"+i_string+"_tunnel_maxdisp_max.csv"
		X,_ := ioutil.ReadFile(file)
		Xlines := strings.Split(string(X),"\n")
		EachCondition := 18
		for j := 0;j<10;j++{
			headline := j*EachCondition
			header := Xlines[headline]
			start := j*EachCondition +2
			end := (j+1)*EachCondition
			var MaxM[]float64
			for m:= start;m<end;m++{
				Xline2 := strings.Split(Xlines[m],",")
				Max,_ := strconv.ParseFloat(strings.Replace(Xline2[1],"\r","",-1),64)
				MaxM = append(MaxM,Max)
			}
			max := 0.0
			maxElement := 0
			for k,value :=range MaxM{
				if math.Abs(value) > max {
					max = value
					maxElement = k
				}
			}
			D := 14.5
			DI := max/D
			PGa := strconv.Itoa(j+1)
			var Target int
			var TargetStr string
			var LabelBinary int
			switch {
			case DI<=0.0006:
				Target = 0
				LabelBinary = 0
				TargetStr = "None"
			case DI> 0.0006 &&DI<=0.0025:
				Target = 1
				LabelBinary = 0
				TargetStr = "Minor"
			case DI>0.0025 && DI<=0.004:
				Target = 2
				LabelBinary = 1
				TargetStr = "Moderate"
			case DI>0.004 && DI<= 0.006:
				Target = 3
				LabelBinary = 1
				TargetStr = "Extensive"
			case DI>0.006:
				Target = 4
				LabelBinary = 1
				TargetStr = "Collapse"
			}
			rows = append(rows,[]string{header,
				PGa,
				strconv.Itoa(maxElement),
				strconv.FormatFloat(max,'E',-1,64),
				strconv.FormatFloat(DI,'E',-1,64),
				strconv.Itoa(Target),TargetStr,
				strconv.Itoa(LabelBinary),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := filpath+"/"+st_string+"MAXtunnel_disp_DI.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxElment","maxDisp","DI","Target","TargetStr","LabelBinary"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}

func getShear(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)

		file := filpath+"/shear/seismic"+i_string+"_shear_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 34
		for j := 0;j<10;j++{
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
	NewFileName := filpath+"/"+st_string+"element_shear.csv"
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

		file := filpath+"/axial/seismic"+i_string+"_axial_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 34
		for j := 0;j<10;j++{
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
	NewFileName := filpath+"/"+st_string+"element_axial.csv"
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

func getPGD(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"/soil_xdisp/seismic"+i_string+"_soil_xdisp_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 17
		for j := 0;j<10;j++{
			headline := j*EachCondition
			header := lines[headline]
			Accxline := j*EachCondition+2
			line2 := strings.Split(lines[Accxline],",")
			accx,_ :=strconv.ParseFloat( strings.Replace(line2[3],"\r","",-1),64)
			PGa := strconv.Itoa(j+1)
			rows = append(rows,[]string{header,
				PGa,line2[3],strconv.FormatFloat(math.Abs(accx),'E',-1,64),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := filpath+"/"+st_string+"PGD_soil_xdisp_real.csv"
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
	headline := []string{"seismicNo+PGA","PGa","PGD"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}

func getPGv(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en ;i++{
		i_string := strconv.Itoa(i)

		file := filpath+"/xvel/seismic"+i_string+"_xvel_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 17
		for j := 0;j<10;j++{
			headline := j*EachCondition
			header := lines[headline]
			Accxline := j*EachCondition+2
			line2 := strings.Split(lines[Accxline],",")
			accx,_ :=strconv.ParseFloat( strings.Replace(line2[3],"\r","",-1),64)
			PGa := strconv.Itoa(j+1)
			rows = append(rows,[]string{header,
				PGa,line2[3],strconv.FormatFloat(math.Abs(accx),'E',-1,64),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := filpath+"/"+st_string+"pgv_real.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxVel"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}

func getAccx(filpath string,st,en int) {
	var rows [][]string
	for i := st;i<=en;i++{
		i_string := strconv.Itoa(i)
		file := filpath+"/accx/seismic"+i_string+"_accx_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 17
		for j := 0;j<10;j++{
			headline := j*EachCondition
			header := lines[headline]
			Accxline := j*EachCondition+2
			line2 := strings.Split(lines[Accxline],",")
			accx,_ :=strconv.ParseFloat( strings.Replace(line2[3],"\r","",-1),64)
			PGa := strconv.Itoa(j+1)
			rows = append(rows,[]string{header,
				PGa,line2[3],strconv.FormatFloat(math.Abs(accx),'E',-1,64),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := filpath+"/"+st_string+"PGa_real_6.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxAccx"}
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
		file := filpath+"/moment/seismic"+i_string+"_moment_max.csv"
		b,err := ioutil.ReadFile(file)
		if err!=nil{
			fmt.Println("read file error")
		}
		lines := strings.Split(string(b),"\n")
		EachCondition := 34
		for j := 0;j<10;j++{
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
			var LabelBinary int
			switch {
			case DI<=1.0:
				Target = 0
				LabelBinary = 0
				TargetStr = "None"
			case DI> 1 &&DI<=1.5:
				Target = 1
				LabelBinary = 0
				TargetStr = "Minor"
			case DI>1.5 && DI<=2.5:
				Target = 2
				LabelBinary = 1
				TargetStr = "Moderate"
			case DI>2.5 && DI<= 3.5:
				Target = 3
				LabelBinary = 1
				TargetStr = "Extensive"
			case DI>3.5:
				Target = 4
				LabelBinary = 1
				TargetStr = "Collapse"
			}
			rows = append(rows,[]string{header,
				PGa,
				strconv.Itoa(maxElement),
				strconv.FormatFloat(max,'E',-1,64),
				strconv.FormatFloat(DI,'E',-1,64),
				strconv.Itoa(Target),TargetStr,
				strconv.Itoa(LabelBinary),
			})
		}
	}
	st_string :=strconv.Itoa(st)
	NewFileName := filpath+"/"+st_string+"element_moment_DI_MRD.csv"
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

	headline := []string{"seismicNo+PGA","PGa","maxElment","maxMoment","DI","Target","TargetStr","LabelBinary"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
	return
}
