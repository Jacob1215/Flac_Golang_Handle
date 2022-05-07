package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	//这个文件夹是要用的
	filePath := "G:/res_301-400_2/seismic_"
	//现在进行写入文件夹的操作
	objFile := "IM_all_371-400.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/data_3-15_Final_2/IM_3-15/"+objFile
	var ResData [][]string
	for i:=371;i<=400;i++{
		i_string := strconv.Itoa(i)
		for j:= 1;j<=10;j++{
			j_string := strconv.Itoa(j)
			Tf,arms,SMA,AI,PGA,Accx,TimeStep := getAccx(filePath,i_string,j_string)
			Ds5_95:= Tf.tf95-Tf.tf5
			Ds5_95_string := strconv.FormatFloat(Ds5_95,'E',-1,64)
			PGV,vrms := getVel(filePath,i_string,j_string,Tf)
			PGD,drms := getDid(filePath,i_string,j_string,Tf)
			arms_string := strconv.FormatFloat(arms,'f',4,64)
			vrms_string := strconv.FormatFloat(vrms,'f',4,64)
			drms_string := strconv.FormatFloat(drms,'f',4,64)
			SMA_string := strconv.FormatFloat(SMA,'f',4,64)
			AI_string := strconv.FormatFloat(AI,'f',4,64)
			//getPSA
			TimePre,PSA,PSV,PSD,ASI:=getPSA(Accx,TimeStep,i_string,j_string)
			PSA_string := strconv.FormatFloat(PSA,'f',4,64)
			PSV_string := strconv.FormatFloat(PSV,'f',4,64)
			PSD_string := strconv.FormatFloat(PSD,'f',4,64)
			ASI_string := strconv.FormatFloat(ASI,'f',4,64)
			ResData = append(ResData,[]string{"seismic_all_IM",i_string,j_string,PGA,PGV,PGD,
				arms_string,vrms_string,drms_string,AI_string,SMA_string,TimePre,PSA_string,PSV_string,PSD_string,ASI_string,Ds5_95_string})
		}
	}

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
	objFileTitle := "IM_1"
	var title []string
	title = append(title, objFileTitle)
	err = w.Write(title)
	if err != nil {
		log.Fatalf("cannot write title：%v", err)
	}
	headline := []string{"filename","seismic_NO","No_Tiaofu","PGA/m/s2", "PGV/m/s", "PGD/m",
		"arms/m/s2","vrms/m/sec","drms/m","AI/m/sec","SMA/m/s2","TimePre","PSA","PSV","PSD","ASI","Ds5_95/s"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(ResData)
}

func getPSA(Accx []float64, TimeStep []float64,i_string,j_string string) (string,float64,float64,float64,float64) {
	//单位统一为m和S
	//fmt.Println(Accx)
	Displace:= make([]float64,len(Accx))
	Velocity:= make([]float64,len(Accx))
	AbsAcce := make([]float64,len(Accx))
	//部分参数
	Damp := 0.05//阻尼比0.02
	//TA := TimeStep[100]-TimeStep[99]
	Dt := TimeStep[100]-TimeStep[99] //地震记录步长
	//初始化
	var MaxD,MaxV,MaxA []float64
	var k_string []string
	for i:=0.00;i<=6.0;i+=0.05{
		NaturalFrequency := 2*3.1415926/i
		DampFrequency := NaturalFrequency*math.Sqrt(1-Damp*Damp)
		e_t := math.Exp(-Damp*NaturalFrequency*Dt)
		s := math.Sin(DampFrequency*Dt)
		c := math.Cos(DampFrequency*Dt)
		//fmt.Println(NaturalFrequency,DampFrequency,e_t,s,c)
		var A [2][2]float64
		A[0][0] = e_t*(s*Damp/math.Sqrt(1-Damp*Damp)+c)
		A[0][1] = e_t*s/DampFrequency
		A[1][0] = -NaturalFrequency*e_t*s/math.Sqrt(1-Damp*Damp)
		A[1][1] = e_t*(-s*Damp/math.Sqrt(1-Damp*Damp)+c)

		d_f := (2*Damp*Damp-1)/(NaturalFrequency*NaturalFrequency*Dt)
		d_3t := Damp/(NaturalFrequency*NaturalFrequency*NaturalFrequency*Dt)

		var B[2][2]float64
		B[0][0] = e_t*((d_f+Damp/NaturalFrequency)*s/DampFrequency+(2*d_3t+1/(NaturalFrequency*NaturalFrequency))*c)-2*d_3t
		B[0][1] = -e_t*(d_f*s/DampFrequency+2*d_3t*c)-1/NaturalFrequency/NaturalFrequency+2*d_3t
		B[1][0] = e_t*((d_f+Damp/NaturalFrequency)*(c-Damp/math.Sqrt(1-Damp*Damp)*s)-(2*d_3t+
			1/NaturalFrequency/NaturalFrequency)*(DampFrequency*s+Damp*NaturalFrequency*c))+1/(NaturalFrequency*NaturalFrequency*Dt)
		B[1][1] = e_t*(1/(NaturalFrequency*NaturalFrequency*Dt)*c+s*Damp/(NaturalFrequency*DampFrequency*Dt))-1/(NaturalFrequency*NaturalFrequency*Dt)
	//fmt.Println(A[0][0])
		for j:=0;j< len(Accx)-1;j++ {
			Displace[j+1] =A[0][0]*Displace[j]+A[0][1]*Velocity[j]+B[0][0]*Accx[j]+B[0][1]*Accx[j+1]
			Velocity[j+1] = A[1][0]*Displace[j]+A[1][1]*Velocity[j]+B[1][0]*Accx[j]+B[1][1]*Accx[j+1]
			AbsAcce[j+1] = 2*Damp*NaturalFrequency*Velocity[j+1]-NaturalFrequency*NaturalFrequency*Displace[j+1]
		}
		MaxD = append(MaxD,getAbsMax(Displace))
		MaxV= append(MaxV,getAbsMax(Velocity))
		if i==0.0{
			MaxA= append(MaxA,getAbsMax(Accx))
		} else {
			MaxA = append(MaxA,getAbsMax(AbsAcce))
		}
		k_string =append(k_string,strconv.FormatFloat(i,'E',-1,64))
	}
	WriteFile(k_string,MaxA,MaxV,MaxD,i_string,j_string)
	TimePre,PSA:= getPS(k_string,MaxA)
	_,PSV :=getPS(k_string,MaxV)
	_,PSD := getPS(k_string,MaxD)
	ASI := getASI(k_string,MaxA)

	return TimePre,PSA,PSV,PSD,ASI
}

func getASI(k_string []string, a []float64) float64{
	var result float64
	for i:=2;i<=10;i++ {
		fmt.Println(k_string[i])
		result += a[i]*0.05
	}
	return result
}

func getPS(k_string []string, a []float64) (string,float64) {
	var max float64
	var TimePre string
	for i,value :=range a{
		if math.Abs(value)>max{
			max=math.Abs(value)
			TimePre = k_string[i]
		}
	}
	return TimePre,max
}



func WriteFile(k_string []string, acce []float64, velocity []float64, displace []float64,i_string,j_string string) {
	var rows [][]string
	for i,_ :=range k_string{
		rows = append(rows,[]string{
			k_string[i],
			strconv.FormatFloat(acce[i],'E',-1,64),
			strconv.FormatFloat(velocity[i],'E',-1,64),
			strconv.FormatFloat(displace[i],'E',-1,64),
		})
	}
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/data_3-15_Final_2/PSA_wave/seismic_"+i_string+"_"+j_string+"_PSA.csv"
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

	headline := []string{"TimeHistory","PSA","PSV","PSD"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
}

func getAbsMax(displace []float64) float64 {
	var max float64
	for _,value := range displace{
		if math.Abs(value)>=max{
			max = math.Abs(value)
		} else {
			continue
		}
	}
    return max
}

func getDid(filePath string, i_string string, j_string string, Tf TimeTf) (string,float64){
	node := 15
	filepath :=filePath+i_string+"_0."+j_string+"_soil_xdisp.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")

	}
	lines := strings.Split(string(b),"\n")
	linesCount := len(lines)/node
	rows := lines[:linesCount]
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	TimeXvdis,Xdis := getData(rows)
	//获取Pd
	Pd := getPa(TimeXvdis,Tf.tf5i,Tf.tf95i,Xdis)
	//获取drms
	drms := math.Sqrt(Pd)
	fmt.Printf("drms: %f\n",drms)
	PGD :=getPG(Xdis)
	fmt.Printf("PGD: %s\n",PGD)
	return PGD,drms
}

func getVel(filePath string, i_string string, j_string string,Tf TimeTf) (string,float64){
	node := 15
	filepath :=filePath+i_string+"_0."+j_string+"_xvel.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")

	}
	lines := strings.Split(string(b),"\n")
	linesCount := len(lines)/node
	rows := lines[:linesCount]
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	TimeXvel,Xvel := getData(rows)

	//获取Pv
	Pv:= getPv(TimeXvel,Tf.tf5i,Tf.tf95i,Xvel)
	//获取vrms
	vrms :=math.Sqrt(Pv)
	fmt.Printf("vrms:%f\n",vrms)
	PGV:=getPG(Xvel)
	fmt.Printf("PGV: %s\n",PGV)
	return PGV,vrms
}

func WritePGA(TimeStep []float64, xvel2 []float64,i_string,j_string string) {
	var rows [][]string
	Dt := TimeStep[100]-TimeStep[99]
	line := strconv.FormatFloat(Dt,'E',-1,64)
	rows = append(rows,[]string{line})
	for _,value :=range xvel2{
		rows = append(rows,[]string{
			strconv.FormatFloat(value/9.81,'E',-1,64),
		})
	}
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/data_3-15_Final_2/PGa_PI_Calc/seismic_"+i_string+"_"+j_string+"_PGA.acc"
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

	headline := []string{"Vel_TimeHistory"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(rows)
}
func getAccx(filePath string, i_string,j_string string) (TimeTf,float64,float64,float64,string,[]float64,[]float64){
	node := 15
	filepath :=filePath+i_string+"_0."+j_string+"_accx.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
	}
	lines := strings.Split(string(b),"\n")
	fmt.Println(len(lines))
	linesCount := len(lines)/node
	rows := lines[:linesCount]
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	TimeAccx,Accx := getData(rows)
	Tf :=getJifenTF95(TimeAccx,Accx)//这个的目的是获得全局积分并获取tf95的坐标
	//把加速度时程记录下来
	WritePGA(TimeAccx,Accx,i_string,j_string)
	//获取Pa
	Pa :=getPa(TimeAccx,Tf.tf5i,Tf.tf95i,Accx)
	fmt.Printf("Pa: %f\n",Pa)
	//获取arms
	arms:= math.Sqrt(Pa)
	fmt.Printf("arms: %f\n",arms)
	SMA  := getSMA(Accx)
	fmt.Printf("SMA: %f\n",SMA)
	//ASI := getASI()//这个后面单独写一个去提
	//PSA 也是单独去写一个提PSa的就行了。
	AI := getAI(TimeAccx,Accx)
	fmt.Printf("AI: %f\n",AI)
	PGA :=getPG(Accx)
	fmt.Printf("PGA: %s\n",PGA)
	return Tf,arms,SMA,AI,PGA,Accx,TimeAccx
}
func getPG(pg []float64) string {
	var max float64
	for _,value := range pg{
		if math.Abs(value)>=max{
			max = math.Abs(value)
		} else {
			continue
		}
	}
	pgMax_string := strconv.FormatFloat(max,'f',4,64)
	return pgMax_string
}
func getAI(time []float64, accx []float64) float64 {
	var AI float64
	for i:=1;i<len(time);i++{
		result := Function_dingjifen(time[i-1],time[i],accx[i])
		AI +=math.Abs(result) // 这个是总计的积分。
		}
		return 3.14/2*AI/9.8
}
func getSMA(accx []float64) float64{//这个有点难统计，暂时放这里
	//不是这样统计的，不过也还行。
	//data := make(map[string]int,len(accx))
	var up []float64
	for i :=1;i<len(accx)-1;i++{
		if accx[i]>0 && accx[i]> accx[i+1] && accx[i]>accx[i-1]{
			up =append(up,math.Abs(accx[i]))
		} else if accx[i] < 0 && accx[i]< accx[i-1] && accx[i]<accx[i+1] {
			up = append(up,math.Abs(accx[i]))
		}
	}
	sort.Float64s(up)
	//fmt.Println(up)
	SMA := up[len(up)-3]
	return SMA
}
func getPv(tf []float64,tf_5i,tf_95i int, acc []float64) float64{
	var jifen float64
	for i:=tf_5i;i<=tf_95i;i++{
		result := Function_dingjifen(tf[i-1],tf[i],acc[i])
		jifen +=math.Abs(result)
	}
	Pv := jifen/float64(tf_95i-tf_5i)
	return Pv
}

func getPa(tf []float64,tf_5i,tf_95i int, acc []float64) float64{
	var jifen float64
	for i:=tf_5i;i<=tf_95i;i++{
		result := Function_dingjifen(tf[i-1],tf[i],acc[i])
		jifen +=math.Abs(result)
	}
	Pa := jifen/float64(tf_95i-tf_5i)
	return Pa
}

type TimeTf struct{
	jifen,jifen2 float64
	tf5,tf75,tf95 float64
	tf5i,tf75i,tf95i int
}
//定积分计算
func Function_dingjifen(t1,t2 float64,acc float64) float64 {
	result :=math.Pow(acc,2)*(t2-t1)
	return result
}
func getJifenTF95(tf []float64, acc []float64) (TimeTf){
	var Tf TimeTf
	for i:=1;i<len(tf);i++{
		result := Function_dingjifen(tf[i-1],tf[i],acc[i])
		Tf.jifen +=math.Abs(result) // 这个是总计的积分。
	}
	for i:=1; i<len(tf);i++{
		result2 := Function_dingjifen(tf[i-1],tf[i],acc[i])
		//fmt.Println(result2)
		Tf.jifen2 += math.Abs(result2)
		//这一段代码有问题。
		if Tf.jifen2<=0.05*Tf.jifen{
			Tf.tf5 = tf[i]
			Tf.tf5i = i
		}else if Tf.jifen2 >= 0.7*Tf.jifen && Tf.jifen2 <= 0.75*Tf.jifen{
			Tf.tf75 = tf[i]
			Tf.tf75i = i
		}else if Tf.jifen2 >= 0.9*Tf.jifen && Tf.jifen2 <= 0.95*Tf.jifen{
			Tf.tf95 = tf[i]
			Tf.tf95i= i
		}
	}

	return Tf
}

func getData(lines []string, ) ([]float64,[]float64) {
	var TimeSeries []float64
	var Data []float64
	for i := 4; i < len(lines); i++ {
			//把这txt给读出来。
			data := strings.Replace(lines[i], " ", ",", -1)
			data2 := strings.Split(data, ",,,,,,,,,,,,")
			Time := strings.Replace(data2[0],",","",-1)
			TimeInt := strings.Replace(Time,"\r","",-1)
			TimeFloat,_ := strconv.ParseFloat(TimeInt,64)
			TimeSeries = append(TimeSeries,TimeFloat)
			data3 := strings.Replace(data2[1], ",", "", -1)
			//fmt.Println(data3)
			data4 := strings.Replace(data3, "\r", "", -1)
			dataint, err := strconv.ParseFloat(data4, 64)
			if err != nil {
				log.Fatalf("can not convert, err is %+v", err)
			}
			Data = append(Data, dataint)
	}
	return TimeSeries,Data
}
